package http

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	StatusCode int
	Headers    map[string][]string
	Body       string
	Size       int
	Lines      int
	Words      int
	Time       time.Duration
}

func NewResponse(resp *http.Response, duration time.Duration) *Response {
	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)
	size := len(bodyBytes)

	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
		Size:       size,
		Lines:      strings.Count(body, "\n"),
		Words:      len(strings.Fields(body)),
		Time:       duration,
	}
}
