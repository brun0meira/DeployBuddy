package routes

import (
	"deploy-buddy/server/internal/handler"

	"github.com/go-chi/chi/v5"
)

type EnvironmentRoutes struct {
	Handler *handler.EnvironmentHandler
}

func (ur *EnvironmentRoutes) NewRouter(r chi.Router) {
	// r.Use(middlewares.JWTAuthMiddleware)
	r.Post("/", ur.Handler.CreateEnv)
	r.Get("/", ur.Handler.GetEnvironments)
}
