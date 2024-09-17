package server

import (
	"log"
	"net/http"
	"time"
)

func (s *ApiServer) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("started %s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("completed in %v\n", time.Since(start))
	})
}
