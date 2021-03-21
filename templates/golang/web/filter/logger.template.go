package filter

// logger filter
const (
	Logger = `
	package filter

import (
	"net/http"
	"time"

	"gitlab.com/aifaniyi/go-libs/logger"
)

// LoggingMiddleware : creates an access log entry for each request received
// using the log format:
//
// http-method  requestURI  http-protocol  useragent  remote-addr
// referer  request-length  time-taken response-status response-length
func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		// Call the next handler, which can be another
		// middleware in the chain, or the final handler.
		writer := apiResponseWriter{ResponseWriter: w}
		next.ServeHTTP(&writer, r)

		duration := time.Since(start)

		referer := r.Referer()
		if referer == "" {
			referer = "-"
		}

		ua := r.UserAgent()
		if ua == "" {
			ua = "-"
		}

		id := "-"
		if reqID := r.Context().Value(requestID); reqID != nil {
			id = reqID.(string)
		}

		// format: req-id, http-method  requestURI  http-protocol  useragent  remote-addr
		// referer  request-length  time-taken response-status response-length
		logger.Info.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%v\t%d\t%d",
			id, r.Method, r.RequestURI, r.Proto, ua,
			r.RemoteAddr, referer, r.ContentLength, duration,
			writer.Status, writer.Length,
		)
	})
}
`
)
