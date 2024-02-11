package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
	logger  *log.Logger
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	l.logger.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(logger *log.Logger, handlerToWrap http.Handler) *Logger {
	return &Logger{
		handler: handlerToWrap,
		logger:  logger,
	}
}
