package main

import (
	"log"
	"net/http"
	"os"
)

type Middleware func(http.Handler) http.Handler

// --- Creates the stack of middlewares ---

func CreateMiddlewareStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

// --- Middlewares ---

func LogIncommingRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ENV") == "DEV" {
			isAuthenticated := r.Header.Get("Authenticated")
			if isAuthenticated == "" || isAuthenticated == "false" {
				log.Printf("📤 %v %v 📤\n", r.Method, r.URL.Path)
			} else {
				log.Printf("📤 %v %v 📤\n", r.Method, r.URL.Path)
			}
		}
		next.ServeHTTP(w, r)
	})
}
