package kutils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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

func Test_h2c(t *testing.T) {
	body := "hello"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(body))
		assert.Nil(t, r.TLS, nil)
	})
	h2s := &http2.Server{}
	h2cHandler := h2c.NewHandler(handler, h2s)
	tests := []struct {
		name          string
		serverHandler http.Handler
		clientCfg     *HttpCfg
		wantErr       assert.ErrorAssertionFunc
		expects       func(*testing.T, *resty.Response)
	}{
		{
			"h1 server, h1 client",
			handler,
			&HttpCfg{},
			assert.NoError,
			func(t *testing.T, resp *resty.Response) {
				assert.Equal(t, "HTTP/1.1", resp.Proto())
				assert.Equal(t, http.StatusOK, resp.StatusCode())
				assert.Equal(t, body, resp.String())
			},
		},
		{
			"h1 server, h2c client",
			handler,
			&HttpCfg{UseH2c: true},
			assert.Error,
			nil,
		},
		{
			"h2 server, h1 client",
			h2cHandler,
			&HttpCfg{},
			assert.NoError,
			func(t *testing.T, resp *resty.Response) {
				assert.Equal(t, "HTTP/1.1", resp.Proto())
				assert.Equal(t, http.StatusOK, resp.StatusCode())
				assert.Equal(t, body, resp.String())
			},
		},
		{
			"h2 server, h2c client",
			h2cHandler,
			&HttpCfg{UseH2c: true},
			assert.NoError,
			func(t *testing.T, resp *resty.Response) {
				assert.Equal(t, "HTTP/2.0", resp.Proto())
				assert.Equal(t, http.StatusOK, resp.StatusCode())
				assert.Equal(t, body, resp.String())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ln, err := net.Listen("tcp", "127.0.0.1:")
			assert.NoError(t, err)
			srv := &http.Server{
				Addr:    ln.Addr().String(),
				Handler: tt.serverHandler,
			}

			go func() {
				_ = srv.Serve(ln)
			}()
			defer func(srv *http.Server, ctx context.Context) {
				_ = srv.Shutdown(ctx)
			}(srv, context.Background())

			client := tt.clientCfg.NewRestyClient()
			resp, err := client.R().EnableTrace().Get("http://" + srv.Addr)

			tt.wantErr(t, err)
			if tt.expects != nil {
				tt.expects(t, resp)
			}
		})
	}
}
