package tests

import (
	"deploy-buddy/server/internal/model"
	"deploy-buddy/server/internal/repository"
	utils "deploy-buddy/server/internal/utils/slack"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepository(t *testing.T) {
	godotenv.Load("../../.env")

	// Configure a database in memory for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao abrir o banco de dados em memória: %v", err)
	}

	// Migrate tables to the database
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatalf("Erro ao migrar as tabelas: %v", err)
	}

	// Create a fake Slack service instance
	slackService := utils.NewSlackService()

	// Create user repository
	repo := repository.NewUserRepository(db, slackService)

	// Test to Create function
	t.Run("Create", func(t *testing.T) {
		user := &model.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}
		err := repo.Create(user)
		assert.NoError(t, err)

		var createdUser model.User
		err = db.First(&createdUser, "email = ?", "john@example.com").Error
		assert.NoError(t, err)
		assert.Equal(t, user.Name, createdUser.Name)
		assert.NotEmpty(t, createdUser.Password)
	})

	// Test to FindAll function
	t.Run("FindAll", func(t *testing.T) {
		users, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Greater(t, len(users), 0)
	})

	// Test to FindByID function
	t.Run("FindByID", func(t *testing.T) {
		user := &model.User{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
		}

		db.Create(user)

		foundUser, err := repo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.Name, foundUser.Name)
	})

	// Test to Update function
	t.Run("Update", func(t *testing.T) {
		user := &model.User{
			Name: "John Updated",
		}
		err := repo.Update(user, false)
		assert.NoError(t, err)
	})

	// Test to Delete function
	t.Run("Delete", func(t *testing.T) {
		user := &model.User{
			Email: "user_to_delete@example.com",
		}
		db.Create(user)
		err := repo.Delete(user.ID)
		assert.NoError(t, err)
	})

	// Test to Approve function
	t.Run("Approve", func(t *testing.T) {
		user := &model.User{
			Name: "User To Approve",
		}
		db.Create(user)
		success, err := repo.Approve(user.ID)
		assert.NoError(t, err)
		assert.True(t, success)
	})

	// Test to Decline function
	t.Run("Decline", func(t *testing.T) {
		user := &model.User{
			Name: "User To Decline",
		}
		db.Create(user)
		success, err := repo.Decline(user.ID)
		assert.NoError(t, err)
		assert.True(t, success)
	})
}
