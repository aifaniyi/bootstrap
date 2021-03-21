package filter

// response writer filter
const (
	ResponseWriter = `
	package filter

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

// apiResponseWriter : response writer with
// response status and response size information
type apiResponseWriter struct {
	http.ResponseWriter
	Status int
	Length int
}

// WriteHeader : writes response status
func (w *apiResponseWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

// Write : writes response data to writer
func (w *apiResponseWriter) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.Length += n
	return n, err
}

func (w *apiResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

	`
)
