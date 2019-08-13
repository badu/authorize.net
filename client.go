package authorize_net

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	SandboxEndpoint = "https://apitest.authorize.net/xml/v1/request.api"
	Endpoint        = "https://api.authorize.net/xml/v1/request.api"

	DefaultTimeOut               = 30 * time.Second
	DefaultKeepAlive             = 30 * time.Second
	DefaultMaxIdleConns          = 100
	DefaultMaxIdleConnsPerHost   = 100
	DefaultIdleConnTimeout       = 90 * time.Second
	DefaultTLSHandshakeTimeout   = 10 * time.Second
	DefaultExpectContinueTimeout = 1 * time.Second
	DefaultMaxConnsPerHost       = 50

	ContentType     = "Content-Type"
	Accept          = "Accept"
	ApplicationJSON = "application/json; charset=utf-8"
	UserAgent       = "User-Agent"
	Authorization   = "Authorization"
	Bearer          = "Bearer "
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type HttpClientCfg struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

type Client struct {
	cfg *HttpClientCfg
	log Logger
}

func NewAPIClient(cfg *HttpClientCfg, logger Logger) *Client {
	if cfg == nil {
		cfg = &HttpClientCfg{}
	}
	if cfg.HTTPClient == nil {
		// build a "default" http client
		cfg.HTTPClient = &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
				Proxy:               http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   DefaultTimeOut,
					KeepAlive: DefaultKeepAlive,
				}).DialContext,
				MaxIdleConns:          DefaultMaxIdleConns,
				IdleConnTimeout:       DefaultIdleConnTimeout,
				TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
				ExpectContinueTimeout: DefaultExpectContinueTimeout,
				MaxConnsPerHost:       DefaultMaxConnsPerHost,
			},
		}
	}

	result := Client{
		cfg: cfg,
		log: logger,
	}

	return &result
}

// prepareRequest build the request
func (c Client) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{}) (*http.Request, error) {

	var (
		body *bytes.Buffer
		err  error
	)
	// Detect postBody type and post.
	if postBody != nil {
		body = &bytes.Buffer{}
		if err = json.NewEncoder(body).Encode(postBody); err != nil {
			return nil, err
		}
	}

	// Setup path and query parameters
	uri, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Generate a new request
	var request *http.Request
	if body != nil {
		if request, err = http.NewRequest(method, uri.String(), body); err != nil {
			return nil, err
		}
	} else {
		if request, err = http.NewRequest(method, uri.String(), nil); err != nil {
			return nil, err
		}
	}

	// set Content-Type header
	request.Header.Set(ContentType, ApplicationJSON)
	// set Accept header
	request.Header.Set(Accept, ApplicationJSON)

	// Override request host, if applicable
	if c.cfg.Host != "" {
		request.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	request.Header.Add(UserAgent, c.cfg.UserAgent)

	// yeah, context should never be nil, but we never know
	if ctx != nil {
		// add context to the request, for cancellation
		request = request.WithContext(ctx)
	}
	// add rest of the "default" headers (if any)
	for header, value := range c.cfg.DefaultHeader {
		request.Header.Add(header, value)
	}
	return request, nil
}

// do do the request.
func (c Client) do(request *http.Request, result interface{}, printRaw bool) error {
	response, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		c.log.Printf("Error : %v\n", err)
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()

	if response.StatusCode != http.StatusOK {

	}

	if printRaw {
		fmt.Println("---")
		fmt.Println(string(body))
		fmt.Println("---")
	}
	// The BOM identifies that the text is UTF-8 encoded, but it should be removed before decoding.
	body = bytes.TrimPrefix(body, []byte{239, 187, 191})

	if err := json.Unmarshal(body, &result); err != nil {
		c.log.Printf("Error : %v\nresponse : %q", err, string(body))
		return err
	}

	return nil
}
