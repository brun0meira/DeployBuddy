package routes

import (
	"deploy-buddy/server/internal/handler"

	"github.com/go-chi/chi/v5"
)

type DevOpsRoutes struct {
	Handler *handler.DevOpsHandler
}

func (ur *DevOpsRoutes) NewRouter(r chi.Router) {
	// r.Use(middlewares.JWTAuthMiddleware)
	r.Post("/", ur.Handler.CreateOrg)
}
