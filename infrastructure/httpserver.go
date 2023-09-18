package infrastructure

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/httplog"

	"github.com/Renos-id/go-starter-template/lib/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitChiRouter() *chi.Mux {
	r := chi.NewRouter()
	// Logger
	logger := httplog.NewLogger(os.Getenv("APP_NAME"), httplog.Options{
		JSON: true,
	})

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/health")) //AWS Health Check

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(10 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := response.WriteSuccess(os.Getenv("APP_NAME"), map[string]any{})
		response.ToJSON(w, r)
	})
	return r
}
