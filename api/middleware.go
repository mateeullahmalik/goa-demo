package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/mateeullahmalik/goa-demo/internal/errors"
	"github.com/mateeullahmalik/goa-demo/internal/log"
	"github.com/mateeullahmalik/goa-demo/internal/log/hooks"
	"github.com/mateeullahmalik/goa-demo/internal/random"

	httpmiddleware "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// Recovery is a method that tries to recover from panics.
func Recovery() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer errors.Recover(log.Fatal)
			h.ServeHTTP(w, r)
		})
	}
}

// Log logs incoming HTTP requests and outgoing responses.
// It uses the request ID set by the RequestID middleware or creates a short unique request ID if missing for each incoming request
// and logs it with the request and corresponding response details.
func Log(ctx context.Context) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Context().Value(middleware.RequestIDKey)
			if reqID == nil {
				reqID, _ = random.String(8, random.Base62Chars)
			}
			started := time.Now()

			ctx = log.ContextWithPrefix(ctx, fmt.Sprintf("%s-%s", logPrefix, reqID))

			log.WithContext(ctx).
				WithField("from", logFrom(r)).
				WithField("req", r.Method+" "+r.URL.String()).
				Debugf("Request")

			rw := httpmiddleware.CaptureResponse(w)
			h.ServeHTTP(rw, r)

			log.WithContext(ctx).
				WithField("status", rw.StatusCode).
				WithField("bytes", rw.ContentLength).
				WithField("time", time.Since(started).String()).
				Debugf("Response")
		})
	}
}

// logFrom makes a best effort to compute the request client IP.
func logFrom(req *http.Request) string {
	if f := req.Header.Get("X-Forwarded-For"); f != "" {
		return f
	}
	f := req.RemoteAddr
	ip, _, err := net.SplitHostPort(f)
	if err != nil {
		return f
	}
	return ip
}

// ErrorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func ErrorHandler(ctx context.Context, w http.ResponseWriter, err error) {
	id := ctx.Value(middleware.RequestIDKey).(string)
	_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
	log.WithContext(ctx).Errorf("[%s] %s", id, err.Error())
}

func init() {
	log.AddHook(hooks.NewContextHook(middleware.RequestIDKey, func(ctxValue interface{}, msg string, fields hooks.ContextHookFields) (string, hooks.ContextHookFields) {
		fields["prefix"] = fmt.Sprintf("%s-%s", logPrefix, ctxValue)
		return msg, fields
	}))
}
