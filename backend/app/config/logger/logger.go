package logger

import (
	"backend/app/config/requestid"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func New(devMode bool) (*zap.Logger, error) {
	var config zap.Config

	if devMode {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}

	logger, err := config.Build()
	if err != nil {
		fmt.Println("failed to create logger: ", err)
		return nil, err
	}
	zap.ReplaceGlobals(logger)

	return logger, nil
}

// Middleware provides logger middleware
func Middleware(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				t1 := time.Now()
				fields := []zap.Field{
					zap.String("request", r.Proto+" "+r.Method+" "+r.URL.Path),
					zap.Duration("lat", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
				}

				if reqID := requestid.GetReqID(r.Context()); reqID != uuid.Nil {
					fields = append(fields, zap.Any("reqId", reqID))
				}
				l.Info("Served", fields...)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
