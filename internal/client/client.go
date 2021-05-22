package client

import (
	"github.com/henvic/httpretty"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type ClientOption = func(tripper http.RoundTripper) http.RoundTripper

type funcTripper struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (tr funcTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return tr.roundTrip(req)
}

func NewHTTPClient(opts ...ClientOption) *http.Client {
	tr := http.DefaultTransport
	for _, opt := range opts {
		tr = opt(tr)
	}

	return &http.Client{
		Transport: tr,
	}
}

func VerboseLog(out io.Writer, logTraffic bool, colorize bool) ClientOption {
	logger := &httpretty.Logger{
		Time:            true,
		TLS:             false,
		Colors:          colorize,
		RequestHeader:   logTraffic,
		RequestBody:     logTraffic,
		ResponseHeader:  logTraffic,
		ResponseBody:    logTraffic,
		Formatters:      []httpretty.Formatter{&httpretty.JSONFormatter{}},
		MaxResponseBody: 10000,
	}
	logger.SetOutput(out)
	logger.SetBodyFilter(func(h http.Header) (skip bool, err error) {
		return true, nil
	})
	return logger.RoundTripper
}

var jsonTypeRE = regexp.MustCompile(`[/+]json($|;)`)

func inspectableMIMEType(t string) bool {
	return strings.HasPrefix(t, "text/") || jsonTypeRE.MatchString(t)
}

// AddHeaderFunc is an AddHeader that gets the string value from a function
func AddHeaderFunc(name string, getValue func(*http.Request) (string, error)) ClientOption {
	return func(tr http.RoundTripper) http.RoundTripper {
		return &funcTripper{
			roundTrip: func(req *http.Request) (*http.Response, error) {
				if req.Header.Get(name) != "" {
					return tr.RoundTrip(req)
				}
				value, err := getValue(req)
				if err != nil {
					return nil, err
				}
				if value != "" {
					req.Header.Add(name, value)
				}
				return tr.RoundTrip(req)
			}}
	}
}

// AddHeader turns a RoundTripper into one that adds a request header
func AddHeader(name, value string) ClientOption {
	return func(tr http.RoundTripper) http.RoundTripper {
		return &funcTripper{roundTrip: func(req *http.Request) (*http.Response, error) {
			if req.Header.Get(name) == "" {
				req.Header.Add(name, value)
			}
			return tr.RoundTrip(req)
		}}
	}
}
