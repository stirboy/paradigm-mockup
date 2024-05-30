package requestid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

var RequestIDHeader = "X-Request-Id"

type RequestId string

const key RequestId = "requestId"

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, key, uuid.New())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetReqID(ctx context.Context) uuid.UUID {
	if ctx == nil {
		return uuid.Nil
	}
	if reqID, ok := ctx.Value(key).(uuid.UUID); ok {
		return reqID
	}
	return uuid.Nil
}
