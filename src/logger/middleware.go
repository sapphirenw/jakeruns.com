package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

const (
	Purple = "\033[35m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		latency := time.Since(start)
		status := ww.Status()

		// print with colors
		logString := fmt.Sprintf(
			"\"%s%s%s %s%s %s%s\" from %s%s%s - %s%d%s %s%dB%s in %s%s%s",
			Purple, r.Method, Reset,
			Blue, r.RequestURI, r.Proto, Reset,
			Reset, r.RemoteAddr, Reset,
			Yellow, status, Reset,
			Blue, ww.BytesWritten(), Reset,
			Green, latency, Reset,
		)

		// without colors
		// logString := fmt.Sprintf(
		// 	"\"%s %s %s\" from %s - %d %dB in %s",
		// 	r.Method,
		// 	r.RequestURI, r.Proto,
		// 	r.RemoteAddr,
		// 	status,
		// 	ww.BytesWritten(),
		// 	latency,
		// )

		Info.Print(logString)
	})
}
