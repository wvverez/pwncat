package http

import (
	"crypto/tls"
	"net/http"
	"time"

	"pwncat/internal/config"
)

type Client struct {
	httpClient *http.Client
	config     *config.Config
}

func NewClient(cfg *config.Config) *Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.Insecure,
		},
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
	}

	if cfg.Cert != "" && cfg.Key != "" {
		if cert, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key); err == nil {
			transport.TLSClientConfig.Certificates = []tls.Certificate{cert}
		}
	}

	return &Client{
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   cfg.Timeout,
		},
		config: cfg,
	}
}

func (c *Client) Do(req *Request) (*Response, error) {
	httpReq, err := req.Build()
	if err != nil {
		return nil, err
	}

	start := time.Now()
	resp, err := c.httpClient.Do(httpReq)
	duration := time.Since(start)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return NewResponse(resp, duration), nil
}
