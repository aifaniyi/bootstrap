package filter

// header filters
const (
	Headers = `
	package filter

import "net/http"

// AddHeaders : setup request context
func AddHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
`
)
