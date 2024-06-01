package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"deploy-buddy/server/internal/config"
	"deploy-buddy/server/internal/routes"
	"deploy-buddy/server/internal/handler"
	"deploy-buddy/server/internal/repository"

	_ "deploy-buddy/server/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	port := os.Getenv("PORT")
	log.Printf("Port: %v\n", port)

	userRepo := repository.NewUserRepository(c.DB, c.SlackService)
	userHandler := handler.NewUserHandler(userRepo)

	packageRepo := repository.NewPackageRepository(c.DB)
	packageHandler := handler.NewPackageHandler(packageRepo)

	environmentRepo := repository.NewEnvironmentRepository(c.DB, c.SlackService)
	environmentHandler := handler.NewEnvironmentHandler(environmentRepo)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	ur := &routes.UsersRoutes{
		Handler: userHandler,
	}
	usersRouter := ur.NewRouter

	sr := &routes.SlackRoutes{}
	slackRouter := sr.NewRouter

	pr := &routes.PackageRoutes{
		Handler: packageHandler,
	}
	packageRouter := pr.NewRouter

	er := &routes.EnvironmentRoutes{
		Handler: environmentHandler,
	}
	environmentRouter := er.NewRouter

	r.Use(middleware.Logger)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", usersRouter)
		r.Route("/slack", slackRouter)
		r.Route("/packages", packageRouter)
		r.Route("/environments", environmentRouter)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	host := "http://localhost"
	log.Printf("Server running at %v:%v\n", host, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
