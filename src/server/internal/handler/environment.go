package handler

import (
	"encoding/json"
	"net/http"

	"deploy-buddy/server/internal/model"
	"deploy-buddy/server/internal/repository"

	"github.com/go-chi/render"
)

type EnvironmentHandler struct {
	repo *repository.EnvironmentRepository
}

func NewEnvironmentHandler(repo *repository.EnvironmentRepository) *EnvironmentHandler {
	return &EnvironmentHandler{repo: repo}
}

// CreateEnv creates a new environment.
// @Summary Create a new environment
// @Description Create a new environment with the provided information.
// @Tags environments
// @Accept json
// @Produce json
// @Param environment body model.CreateEnvRequest true "Create Environment"
// @Success 201 {object} model.Environment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/environments [post]
func (h *EnvironmentHandler) CreateEnv(w http.ResponseWriter, r *http.Request) {
	var req model.CreateEnvRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload", "message": err.Error()})
		return
	}

	environment := &model.Environment{
		Name:             req.Name,
		Description:      req.Description,
		StringConnection: req.StringConnection,
		Packages:         []model.Package{},
	}

	err = h.repo.CreateEnvironment(environment)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to create environment", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, environment)
}

// GetEnvironments returns all environments.
// @Summary Get all environments
// @Description Get all environments.
// @Tags environments
// @Produce json
// @Success 200 {array} model.Environment
// @Failure 500 {object} map[string]string
// @Router /api/v1/environments [get]
func (h *EnvironmentHandler) GetEnvironments(w http.ResponseWriter, r *http.Request) {
	environments, err := h.repo.GetEnvironments()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to get environments", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, environments)
}