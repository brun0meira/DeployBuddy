package routes

import (
	"deploy-buddy/server/internal/handler"

	"github.com/go-chi/chi/v5"
)

type PackageRoutes struct {
	Handler *handler.PackageHandler
}

func (pr *PackageRoutes) NewRouter(r chi.Router) {
	// r.Use(middlewares.JWTAuthMiddleware)
	r.Post("/", pr.Handler.CreatePackage)
	r.Post("/metadata", pr.Handler.ListMetadatas)
	r.Post("/{packageID}/auth", pr.Handler.AuthenticateSalesforce)
	r.Post("/{packageID}/retrieve", pr.Handler.RetrieveMetadata)
	r.Post("/{packageID}/unpack", pr.Handler.UnpackFiles)
	r.Post("/{packageID}/github", pr.Handler.CreateGithubBranch)
	r.Post("/{packageID}/pr", pr.Handler.OpenPullRequest)
}
