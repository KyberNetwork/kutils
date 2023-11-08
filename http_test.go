package kutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHttpCfg_NewRestyClient(t *testing.T) {
	tests := []struct {
		name   string
		cfg    *HttpCfg
		assert func(*testing.T, *resty.Client)
	}{
		{
			"happy",
			&HttpCfg{
				HttpClient:       http.DefaultClient,
				BaseUrl:          "base",
				Headers:          map[string][]string{"Header": {"value"}},
				RetryCount:       1,
				RetryWaitTime:    time.Millisecond,
				RetryMaxWaitTime: time.Second,
				Debug:            true,
			},
			func(t *testing.T, c *resty.Client) {
				assert.NotNil(t, c)
				assert.Equal(t, http.DefaultClient, c.GetClient())
				assert.Equal(t, "base", c.BaseURL)
				assert.Equal(t, []string{"value"}, c.Header["Header"])
				assert.Equal(t, 1, c.RetryCount)
				assert.Equal(t, time.Millisecond, c.RetryWaitTime)
				assert.Equal(t, time.Second, c.RetryMaxWaitTime)
				assert.Equal(t, true, c.Debug)
			},
		},
		{
			"nil config",
			nil,
			func(t *testing.T, c *resty.Client) {
				assert.NotNil(t, c)
				assert.NotNil(t, c.GetClient())
				assert.Empty(t, c.BaseURL)
				assert.Empty(t, c.Header)
				assert.Empty(t, c.RetryCount)
				assert.Empty(t, c.Debug)
			},
		},
		{
			"only base",
			&HttpCfg{
				BaseUrl: "base",
			},
			func(t *testing.T, c *resty.Client) {
				assert.NotNil(t, c)
				assert.NotNil(t, c.GetClient())
				assert.Equal(t, "base", c.BaseURL)
				assert.Empty(t, c.Header)
				assert.Empty(t, c.RetryCount)
				assert.Empty(t, c.Debug)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.cfg
			tt.assert(t, h.NewRestyClient())
		})
	}
}

func TestJSONUnmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    any
	}
	tests := []struct {
		name    string
		args    args
		wantV   any
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"happy",
			args{
				data: []byte(`{"i": 1, "s": {"j": 2}}`),
				v: &struct {
					I int
					S any
				}{},
			},
			&struct {
				I int
				S any
			}{1, map[string]any{"j": json.Number("2")}},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, JSONUnmarshal(tt.args.data, tt.args.v),
				fmt.Sprintf("JSONUnmarshal(%v, %v)", tt.args.data, tt.args.v))
			assert.Equal(t, tt.wantV, tt.args.v)
		})
	}
}

func Test_retryableHttpError(t *testing.T) {
	type args struct {
		r   *resty.Response
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"nil",
			args{
				r:   nil,
				err: nil,
			},
			false,
		},
		{
			"429",
			args{
				r: &resty.Response{
					RawResponse: &http.Response{
						StatusCode: http.StatusTooManyRequests,
					},
				},
				err: nil,
			},
			true,
		},
		{
			"too many redirects",
			args{
				r: &resty.Response{
					RawResponse: &http.Response{},
				},
				err: &url.Error{Err: errors.New("stopped after 10 redirects")},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, retryableHttpError(tt.args.r, tt.args.err), "retryableHttpError(%v, %v)",
				tt.args.r, tt.args.err)
		})
	}
}
