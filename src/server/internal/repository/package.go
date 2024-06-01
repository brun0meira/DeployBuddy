package repository

import (
	"bufio"
	"deploy-buddy/server/internal/model"
	"fmt"
	"log"
	"os/exec"

	"bytes"
	"encoding/json"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PackageRepository struct {
	DB *gorm.DB
}

// MetadataResult holds the type and result of metadata
type MetadataResult struct {
	FullName      string          `json:"fullName"`
	CreatedByName string          `json:"createdByName"`
	Type          string          `json:"type"`
	Result        json.RawMessage `json:"result"`
}

func NewPackageRepository(db *gorm.DB) *PackageRepository {
	return &PackageRepository{DB: db}
}

// Criar pacote
func (repo *PackageRepository) CreatePackage(packageDetails model.CreatePackageRequest) (string, error) {
	result := repo.DB.Create(&packageDetails)
	if result.Error != nil {
		return "", result.Error
	}
	return "Package created successfully", nil
}

// Autenticar com Salesforce
func (repo *PackageRepository) AuthenticateSalesforce(strConnection string) (string, error) {

	uuidForConnection := uuid.New()

	log.Println(strConnection)

	file, err := os.Create("./metadatas/temp.txt")
	if err != nil {
		log.Println("Error creating file:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(strConnection)
	if err != nil {
		log.Println("Error writing to file:", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Println("Error flushing file:", err)
	}

	cmdLogin := exec.Command("sfdx", "force:auth:sfdxurl:store", "-f", "./metadatas/temp.txt", "-a", uuidForConnection.String())
	if err := cmdLogin.Run(); err != nil {
		return "", fmt.Errorf("failed to authenticate with Salesforce: %w", err)
	}

	return uuidForConnection.String(), nil
}

func (repo *PackageRepository) ListMetadatas(strConnection string) ([]map[string]interface{}, error) {
	uuid, err := repo.AuthenticateSalesforce(strConnection)
	if err != nil {
		return nil, err
	}

	// List all metadata types
	metadataTypes := []string{
		// "ApexClass",
		// "ApexTrigger",
		// "VisualforcePage",
		// "LightningComponentBundle",
		// "Workflow",
		// "ValidationRule",
		// "Profile",
		// "PermissionSet",
		// "Layout",
		"CustomObject",
		// "EmailTemplate",
		// "CustomTab",
		// "StaticResource",
		// "CustomLabel",
		"CustomField",
	}

	// Save to file with the name of the connection
	file, err := os.Create(fmt.Sprintf("./metadatas/%s.json", uuid))
	if err != nil {
		log.Println("Error creating file:", err)
	}
	defer file.Close()

	var metadataResults []map[string]interface{}

	for _, metadataType := range metadataTypes {
		fmt.Printf("Metadata type: %s\n", metadataType)
		cmdList := exec.Command("sfdx", "force:mdapi:listmetadata", "--metadata-type", metadataType, "-o", uuid)
		var out bytes.Buffer
		cmdList.Stdout = &out
		if err := cmdList.Run(); err != nil {
			return nil, fmt.Errorf("failed to list metadata for type %s: %w", metadataType, err)
		}

		// Capture the JSON output
		var result json.RawMessage
		if err := json.Unmarshal(out.Bytes(), &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal metadata for type %s: %w", metadataType, err)
		}

		metadataResults = append(metadataResults, map[string]interface{}{
			"Type":   metadataType,
			"Result": result,
		})
	}

	// Convert metadataResults to JSON
	metadataJSON, err := json.Marshal(metadataResults)
	if err != nil {
		return nil, fmt.Errorf("failed to convert metadata to JSON: %w", err)
	}

	// Write to file
	writer := bufio.NewWriter(file)
	_, err = writer.Write(metadataJSON)
	if err != nil {
		log.Println("Error writing to file:", err)
	}

	return metadataResults, nil
}

// Recuperar metadados do Salesforce
func (repo *PackageRepository) RetrieveMetadata(packageID string) (model.Metadata, error) {
	// Implementar lógica de recuperação de metadados
	// Get the package details
	var packageDetails model.CreatePackageRequest

	result := repo.DB.Where("package_id = ?", packageID).First(&packageDetails)
	if result.Error != nil {
		return model.Metadata{}, result.Error
	}

	// Get the connection details
	return model.Metadata{}, nil
}

// Descompactar e organizar arquivos
func (repo *PackageRepository) UnpackFiles(packageID string, filePath string) error {
	// Implementar lógica para descompactar e organizar arquivos
	return nil
}

// Criar branch no GitHub
func (repo *PackageRepository) CreateGithubBranch(packageID string, branchDetails model.GithubBranchDetails) error {
	// Implementar lógica para criar um branch no GitHub
	return nil
}

// Abrir pull request no GitHub
func (repo *PackageRepository) OpenPullRequest(packageID string, prDetails model.PullRequestDetails) (string, error) {
	// Implementar lógica para abrir um pull request
	return "", nil
}

func (repo *PackageRepository) GetEnvironmentByID(id string) (*model.Environment, error) {
	var env model.Environment
	err := repo.DB.Where("environment_id = ?", id).First(&env).Error
	return &env, err
}
