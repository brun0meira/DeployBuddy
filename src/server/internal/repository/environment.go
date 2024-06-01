package repository

import (
	"deploy-buddy/server/internal/model"
	slack "deploy-buddy/server/internal/utils/slack"
	"errors"

	"gorm.io/gorm"
)

type EnvironmentRepository struct {
	db           *gorm.DB
	slackService *slack.SlackService
}

func NewEnvironmentRepository(db *gorm.DB, slackService *slack.SlackService) *EnvironmentRepository {
	return &EnvironmentRepository{
		db:           db,
		slackService: slackService,
	}
}

func (repo *EnvironmentRepository) CreateEnvironment(env *model.Environment) error {
	if env.Name == "" {
		return errors.New("Name is required")
	}

	if env.StringConnection == "" {
		return errors.New("StringConnection is required")
	}

	repo.db.Create(env)

	return nil
}

func (repo *EnvironmentRepository) GetEnvironments() ([]model.Environment, error) {
	var environments []model.Environment

	repo.db.Find(&environments)

	return environments, nil
}
