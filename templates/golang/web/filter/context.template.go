package filter

// contexts filter
const (
	Context = `
	package filter

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type key int

const (
	requestID key = iota
)

// AddContext : setup request context
func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create and add unique request uuid
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestID, id)

		// ... add other needed params
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID :
func GetRequestID(ctx context.Context) string {
	id := "-"
	if reqID := ctx.Value(requestID); reqID != nil {
		id = reqID.(string)
	}
	return id
}`
)
