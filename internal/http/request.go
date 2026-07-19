package http

import (
	"net/http"
	"strings"

	"pwncat/internal/config"
	"pwncat/internal/input"
)

type Request struct {
	config *config.Config
	value  *input.Value
}

func NewRequest(cfg *config.Config) *Request {
	return &Request{config: cfg}
}

func (r *Request) SetValue(val *input.Value) {
	r.value = val
}

func (r *Request) Build() (*http.Request, error) {
	url := r.replace(r.config.URL)

	var body strings.Reader
	if r.config.Method == "POST" {
		body = strings.Reader{}
	}

	req, err := http.NewRequest(r.config.Method, url, &body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "pwncat/1.0")
	return req, nil
}

func (r *Request) replace(text string) string {
	result := text
	for key, val := range r.value.Values {
		result = strings.ReplaceAll(result, key, string(val))
	}
	return result
}
