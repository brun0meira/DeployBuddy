package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Environment struct {
	EnvironmentID    uuid.UUID  `gorm:"type:uuid;primaryKey;" json:"environment_id"`
	Name             string     `gorm:"type:varchar;not null;" json:"name"`
	Description      string     `gorm:"type:varchar;" json:"description"`
	Packages         []Package  `gorm:"foreignKey:EnvironmentID" json:"packages"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	StringConnection string     `gorm:"type:varchar;" json:"string_connection"`
}

func (env *Environment) BeforeCreate(tx *gorm.DB) (err error) {
	env.EnvironmentID = uuid.New()
	return nil
}

type CreateEnvRequest struct {
	Name             string `json:"name" validate:"required,min=3,max=20"`
	Description      string `json:"description"`
	StringConnection string `json:"string_connection" validate:"required"`
}

func init() {
	validate = validator.New()
}
