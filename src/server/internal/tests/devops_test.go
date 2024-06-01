package tests

import (
	"deploy-buddy/server/internal/model"
	"deploy-buddy/server/internal/repository"
	utils "deploy-buddy/server/internal/utils/slack"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDevOpsRepository(t *testing.T) {
	godotenv.Load("../../.env")

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao abrir o banco de dados em memória: %v", err)
	}

	err = db.AutoMigrate(&model.Orgs{}, &model.User{})
	if err != nil {
		t.Fatalf("Erro ao migrar as tabelas: %v", err)
	}

	slackService := utils.NewSlackService()

	repo := repository.NewDevOpsRepository(db, slackService)

	testUser := &model.User{
		ID:       uuid.New(),
		Username: "testuser",
		GHP:      "testtoken",
	}
	db.Create(testUser)

	t.Run("CreateOrg", func(t *testing.T) {
		testOrg := &model.Orgs{
			StringConnection: "test_connection_string",
		}
		message, err := repo.CreateOrg(testOrg, testUser.ID, "", "testrepo")

		assert.NoError(t, err)
		assert.Equal(t, "Organization created successfully", message)

		var createdOrg model.Orgs
		err = db.First(&createdOrg, "string_connection = ?", "test_connection_string").Error
		assert.NoError(t, err)
		assert.Equal(t, testOrg.StringConnection, createdOrg.StringConnection)
	})

	t.Run("RetrieveMetadatas", func(t *testing.T) {
		testOrg := &model.Orgs{
			ID:               uuid.New(),
			StringConnection: "test_connection_string_2",
		}
		db.Create(testOrg)

		message, err := repo.RetrieveMetadatas(testUser, testOrg, "testuser", "testtoken", "testuser", "testrepo")

		assert.NoError(t, err)
		assert.Equal(t, "Metadata deployed and pull request opened successfully", message)
	})
}
