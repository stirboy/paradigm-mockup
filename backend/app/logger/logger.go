package logger

import (
	"backend/app/requestid"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func Logger(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			fields := []zap.Field{
				zap.String("proto", r.Proto),
				zap.String("path", r.URL.Path),
				zap.Duration("lat", time.Since(t1)),
				zap.Int("status", ww.Status()),
				zap.Int("size", ww.BytesWritten()),
			}

			if reqID := requestid.GetReqID(r.Context()); reqID != uuid.Nil {
				fields = append(fields, zap.Any("reqId", reqID))
			}

			defer func() {
				l.Info("Served", fields...)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
