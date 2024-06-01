package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesforceCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Metadata struct {
	ID          uuid.UUID `json:"id"`
	OrgID       uuid.UUID `json:"org_id"`
	Content     string    `json:"content"`
	RetrievedAt time.Time `json:"retrieved_at"`
}

type GithubBranchDetails struct {
	OrgName      string `json:"org_name"`
	RepoName     string `json:"repo_name"`
	SourceBranch string `json:"source_branch"`
	NewBranch    string `json:"new_branch"`
}

type PullRequestDetails struct {
	OrgName      string `json:"org_name"`
	RepoName     string `json:"repo_name"`
	Title        string `json:"title"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
}

type Package struct {
	ID               uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	OwnerID          uuid.UUID      `gorm:"type:uuid;" json:"owner_id"`
	EnvironmentID    *uuid.UUID     `gorm:"type:uuid;" json:"environment_id,omitempty"`
	SalesforceAuthID uuid.UUID      `json:"salesforce_auth_id"`
	MetadataID       uuid.UUID      `json:"metadata_id"`
	GithubBranchID   uuid.UUID      `json:"github_branch_id"`
	PullRequestID    uuid.UUID      `json:"pull_request_id"`
	Status           string         `json:"status"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type CreatePackageRequest struct {
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description"`
	StringConnection string `json:"string_connection"`
	RepoName         string `json:"repo_name"`
	Owner            string `json:"owner"`
}
