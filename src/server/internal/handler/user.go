package handler

import (
	"encoding/json"
	"net/http"

	"deploy-buddy/server/internal/model"
	"deploy-buddy/server/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Create a new user with the provided information.
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.CreateUserRequest true "Create User"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest model.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload", "message": err.Error()})
		return
	}

	if err := createUserRequest.Validate(); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	user := &model.User{
		Name:     createUserRequest.Name,
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
		GHP:      createUserRequest.GHP,
	}

	err = h.repo.Create(user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to create user", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, user)
}

// AuthenticateUser authenticates a user.
// @Summary Authenticate a user
// @Description Authenticate a user using their email and password.
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.AuthenticateUserRequest true "Authenticate User"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/users/auth [post]
func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var authenticateUserRequest model.AuthenticateUserRequest

	err := json.NewDecoder(r.Body).Decode(&authenticateUserRequest)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload", "message": err.Error()})
		return
	}

	user, err := h.repo.Authenticate(authenticateUserRequest.Email, authenticateUserRequest.Password)
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{"error": "Invalid credentials", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

// GetAllUsers returns all users.
// @Summary Get all users
// @Description Get all users.
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [get]
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch users", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, users)
}

// GetUser returns a user by ID.
// @Summary Get a user
// @Description Get a user by ID.
// @Tags users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var userID uuid.UUID
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID", "message": err.Error()})
		return
	}

	user, err := h.repo.FindByID(userID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "User not found", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

// UpdateUser updates a user.
// @Summary Update a user
// @Description Update a user by ID.
// @Tags users
// @Param id path string true "User ID"
// @Accept json
// @Produce json
// @Param user body model.User true "Update User"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userID uuid.UUID
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID", "message": err.Error()})
		return
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload", "message": err.Error()})
		return
	}

	user.ID = userID

	updatingPassword := user.Password != ""
	err = h.repo.Update(&user, updatingPassword)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to update user", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

// DeleteUser deletes a user.
// @Summary Delete a user
// @Description Delete a user by ID.
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userID uuid.UUID
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID", "message": err.Error()})
		return
	}

	err = h.repo.Delete(userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to delete user", "message": err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "User deleted"})
}

// ApproveUser approves a user.
// @Summary Approve a user
// @Description Approve a user by ID.
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/approve/{id} [post]
func (h *UserHandler) ApproveUser(w http.ResponseWriter, r *http.Request) {
	var userID uuid.UUID
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID", "message": err.Error()})
		return
	}

	approved, err := h.repo.Approve(userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to approve user", "message": err.Error()})
		return
	}

	if !approved {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "User approved"})
}

// DeclineUser declines a user.
// @Summary Decline a user
// @Description Decline a user by ID.
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/decline/{id} [post]
func (h *UserHandler) DeclineUser(w http.ResponseWriter, r *http.Request) {
	var userID uuid.UUID
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID", "message": err.Error()})
		return
	}

	declined, err := h.repo.Decline(userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to decline user", "message": err.Error()})
		return
	}

	if !declined {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "User declined"})
}
