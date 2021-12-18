package utilities

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

var HTTPTimeout = 10 * time.Second
var AllowInsecureCert = false

type Options struct {
	Method  string
	Body    []byte
	URL     string
	Headers map[string]string

	// AltClient is an alternative HTTP client interface, mostly useful only for testing.
	AltClient HTTPClient
}

func GetHTTPClientInterface() HTTPClient {
	return GetHTTPClient()
}


func SetAllowInsecureCert(allowInsecureCert bool) {
	AllowInsecureCert = allowInsecureCert
}

func GetHTTPClient() *http.Client {
	var client *http.Client
	if AllowInsecureCert {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	client.Timeout = HTTPTimeout
	return client
}


type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
}

func RequestJSON(opts *Options) (*http.Response, error) {
	if opts.Headers == nil {
		opts.Headers = make(map[string]string)
	}
	opts.Headers["Content-Type"] = "application/json"
	return Request(opts)
}

func Request(o *Options) (*http.Response, error) {
	req, err := http.NewRequest(o.Method, o.URL, bytes.NewBuffer(o.Body))
	if err != nil {
		fmt.Printf("Error Request: %v", err)
		return nil, err
	}

	if len(o.Headers) > 0 {
		for k, v := range o.Headers {
			req.Header.Set(k, v)
		}
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	var client HTTPClient
	if o.AltClient != nil {
		client = o.AltClient
	} else {
		client = GetHTTPClientInterface()
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return res, nil
}
