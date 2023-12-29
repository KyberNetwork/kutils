package kutils

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/KyberNetwork/kutils/internal/json"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-retryablehttp"
)

// HttpCfg is the resty http client configs
type HttpCfg struct {
	HttpClient       *http.Client  `json:"-"`
	BaseUrl          string        // client's base url for all methods
	Headers          http.Header   // default headers
	Timeout          time.Duration // request timeout, see http.Client's Timeout
	RetryCount       int           // retry count (exponential backoff), default 0
	RetryWaitTime    time.Duration // first exponential backoff, default 100ms
	RetryMaxWaitTime time.Duration // max exponential backoff, default 2s
	Debug            bool          // whether to log requests and responses
}

// NewRestyClient creates a new resty client with the given configs
func (h *HttpCfg) NewRestyClient() (client *resty.Client) {
	if h == nil {
		return resty.New()
	}

	hc := h.HttpClient
	if hc == nil {
		hc = &http.Client{Timeout: h.Timeout}
	} else if hc.Timeout == 0 {
		hc.Timeout = h.Timeout
	}
	client = resty.NewWithClient(hc)

	client.SetBaseURL(h.BaseUrl).
		SetRetryCount(h.RetryCount).
		AddRetryCondition(retryableHttpError).
		SetDebug(h.Debug)
	client.Header = h.Headers
	if waitTime := h.RetryWaitTime; waitTime != 0 {
		client.SetRetryWaitTime(waitTime)
	}
	if maxWaitTime := h.RetryMaxWaitTime; maxWaitTime != 0 {
		client.SetRetryMaxWaitTime(maxWaitTime)
	}
	client.JSONMarshal = JSONMarshal
	client.JSONUnmarshal = JSONUnmarshal
	return client
}

func retryableHttpError(r *resty.Response, err error) bool {
	if r == nil {
		return false
	}
	switch r.StatusCode() {
	case http.StatusRequestTimeout, http.StatusMisdirectedRequest, http.StatusLocked, http.StatusTooManyRequests:
		return true
	default:
		retry, _ := retryablehttp.DefaultRetryPolicy(context.Background(), r.RawResponse, err)
		return retry
	}
}

// JSONMarshal allows choosing the JSON marshalling implementation with build tag with the same logic as used by gin
func JSONMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// JSONUnmarshal allows choosing the JSON unmarshalling implementation with build tag with the same logic as used by gin
func JSONUnmarshal(data []byte, v any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}
