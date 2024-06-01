package handler

import (
	"deploy-buddy/server/internal/model"
	"deploy-buddy/server/internal/repository"
	github "deploy-buddy/server/internal/utils/github"
	utils "deploy-buddy/server/internal/utils/slack"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type PackageHandler struct {
	Repo *repository.PackageRepository
}

func NewPackageHandler(repo *repository.PackageRepository) *PackageHandler {
	return &PackageHandler{Repo: repo}
}

// CreatePackage creates a new package.
// @Summary Create a new package
// @Description Create a new package with the provided information.
// @Tags packages
// @Accept json
// @Produce json
// @Param package body model.CreatePackageRequest true "Create Package"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/packages [post]
func (h *PackageHandler) CreatePackage(w http.ResponseWriter, r *http.Request) {
	var packageDetails model.CreatePackageRequest
	err := json.NewDecoder(r.Body).Decode(&packageDetails)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload", "message": err.Error()})
		return
	}

	// Aqui, implemente qualquer validação adicional se necessário
	message, err := h.Repo.CreatePackage(packageDetails)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to create package", "message": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"message": message})
}

type ListMetadatasRequest struct {
	Connection string `json:"string_connection"`
}

type ModifiedItem struct {
	OldValue string
	NewValue string
}

// ListMetadatas lists metadata.
// @Summary List metadata
// @Description List metadata with the provided connection.
// @Tags packages
// @Accept json
// @Produce json
// @Param metadata body ListMetadatasRequest true "List Metadata"
// @Success 200 {object} []repository.MetadataResult
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/packages/metadata [post]
func (h *PackageHandler) ListMetadatas(w http.ResponseWriter, r *http.Request) {
	prodID := "63fa879e-487c-464d-926c-3019ee5da9f4"
	env, err := h.Repo.GetEnvironmentByID(prodID)
	if err != nil {
		log.Printf("Error retrieving environment: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve production environment", "message": err.Error()})
		return
	}

	var req ListMetadatasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload", "message": err.Error()})
		return
	}

	log.Println("Retrieving metadata from string_connection: ", req.Connection)
	currentMetadata, err := h.Repo.ListMetadatas(req.Connection)
	if err != nil {
		log.Printf("Error retrieving current metadata: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve metadata", "message": err.Error()})
		return
	}

	log.Println("Retrieving metadata from production environment: ", env.StringConnection)
	prodMetadata, err := h.Repo.ListMetadatas(env.StringConnection)
	if err != nil {
		log.Printf("Error retrieving production metadata: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve metadata", "message": err.Error()})
		return
	}

	added, removed := compareResults(currentMetadata, prodMetadata)

	log.Printf("Added: %v, Removed: %v", added, removed)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"added":   added,
		"removed": removed,
	})
}

func compareResults(oldResults, newResults []map[string]interface{}) (added []string, removed []string) {
	// oldMap := make(map[string]bool)
	// newMap := make(map[string]bool)

	// fmt.Println("Old results: ", oldResults)
	// fmt.Println("New results: ", newResults)

	// // Fill the map for old results
	// for _, item := range oldResults {
	// 	if id, ok := item["id"].(string); ok {
	// 		oldMap[id] = true
	// 	}
	// }

	// // Check for additions and build map for new results
	// for _, item := range newResults {
	// 	if id, ok := item["id"].(string); ok {
	// 		newMap[id] = true
	// 		if !oldMap[id] {
	// 			added = append(added, id)
	// 		}
	// 	}
	// }

	// // Check for removals
	// for id := range oldMap {
	// 	if !newMap[id] {
	// 		removed = append(removed, id)
	// 	}
	// }

	// read ./metadatas/2e86b113-1bb8-4768-abf2-98d7bf3d31b6.json
	file1, err := os.Open("./metadatas/2e86b113-1bb8-4768-abf2-98d7bf3d31b6.json")
	if err != nil {
		log.Println("Error opening file:", err)
	}
	defer file1.Close()

	// read ./metadatas/bc4718e2-28ce-45db-ab49-ec7d6a635aa2.json
	file2, err := os.Open("./metadatas/bc4718e2-28ce-45db-ab49-ec7d6a635aa2.json")
	if err != nil {
		log.Println("Error opening file:", err)
	}
	defer file2.Close()

	var oldMetadataList, newMetadataList []map[string]interface{}

	// Unmarshal the JSON data
	err = json.NewDecoder(file1).Decode(&oldMetadataList)
	if err != nil {
		log.Println("Error unmarshalling old.json:", err)
	}
	err = json.NewDecoder(file2).Decode(&newMetadataList)
	if err != nil {
		log.Println("Error unmarshalling new.json:", err)
	}

	fmt.Println("Old metadata list: ", oldMetadataList)
	fmt.Println("New metadata list: ", newMetadataList)

	for key, value := range oldMetadataList[0] {
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			if !compareSlices(value.([]interface{}), newMetadataList[0][key].([]interface{})) {
				added = append(added, key)
			}
		} else if newMetadataList[0][key] != value {
			added = append(added, key)
		}
	}

	for key, value := range newMetadataList[0] {
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			if !compareSlices(value.([]interface{}), oldMetadataList[0][key].([]interface{})) {
				removed = append(removed, key)
			}
		} else if oldMetadataList[0][key] != value {
			removed = append(removed, key)
		}
	}

	return added, removed
}

func compareSlices(slice1, slice2 []interface{}) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func (h *PackageHandler) AuthenticateSalesforce(w http.ResponseWriter, r *http.Request) {
	// Extract the Salesforce authentication URL from the request context or payload
	// Here, you need to ensure you have this URL available, either passed in a secure way or configured securely
	sfAuthURL := r.Context().Value("salesforceAuthURL").(string)

	// Write the URL to a temp file needed by the sfdx command
	file, err := os.Create("temp.txt")
	if err != nil {
		log.Printf("Error creating file for Salesforce auth: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to create authentication file"})
		return
	}
	defer file.Close()

	_, err = file.WriteString(sfAuthURL)
	if err != nil {
		log.Printf("Error writing to Salesforce auth file: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to write authentication URL"})
		return
	}

	// Execute the sfdx command to authenticate
	cmd := exec.Command("sfdx", "force:auth:sfdxurl:store", "-f", "temp.txt", "-a", "aliasNameHere")
	if err := cmd.Run(); err != nil {
		log.Printf("Salesforce authentication failed: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Salesforce authentication failed"})
		return
	}

	// Cleanup the temp file
	if err := os.Remove("temp.txt"); err != nil {
		log.Printf("Error removing temp file: %v", err)
		// Don't fail the request if only file removal fails
	}

	// Respond success
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{"message": "Authenticated with Salesforce successfully"})
}

func (h *PackageHandler) RetrieveMetadata(w http.ResponseWriter, r *http.Request) {
	// Definir o ID da organização para criar um nome de diretório único para os metadados
	orgID := r.Context().Value("orgID").(string)
	folderName := "./metadatas/" + orgID

	// Criar diretório se não existir
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, os.ModePerm)
	}

	// Comando para recuperar os metadados usando sfdx
	cmd := exec.Command("sfdx", "force:mdapi:retrieve", "-r", folderName, "-u", "aliasNameHere", "-k", "./metadatas/package.xml")
	if err := cmd.Run(); err != nil {
		log.Printf("Error retrieving Salesforce metadata: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve Salesforce metadata"})
		return
	}

	// Descompactar os metadados no diretório especificado
	cmdUnpack := exec.Command("unzip", folderName+"/unpackaged.zip", "-d", folderName)
	if err := cmdUnpack.Run(); err != nil {
		log.Printf("Error unpacking Salesforce metadata: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to unpack Salesforce metadata"})
		return
	}

	// Remover o arquivo zip após descompactar
	cmdDeleteZip := exec.Command("rm", "-rf", folderName+"/unpackaged.zip")
	if err := cmdDeleteZip.Run(); err != nil {
		log.Printf("Error deleting zip file: %v", err)
		// Não falhar a requisição se apenas a remoção do arquivo zip falhar
	}

	// Resposta de sucesso
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "Salesforce metadata retrieved and unpacked successfully",
		"path":    folderName,
	})
}

func (h *PackageHandler) UnpackFiles(w http.ResponseWriter, r *http.Request) {
	// Obter o caminho do diretório dos metadados do contexto da requisição
	folderName := r.Context().Value("metadataFolderPath").(string)
	zipPath := folderName + "/unpackaged.zip"
	unpackedPath := folderName + "/unpackaged"

	// Comando para descompactar os metadados no diretório especificado
	cmdUnpack := exec.Command("unzip", zipPath, "-d", folderName)
	if err := cmdUnpack.Run(); err != nil {
		log.Printf("Error unpacking Salesforce metadata: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to unpack Salesforce metadata"})
		return
	}

	// Remover o arquivo zip após descompactar
	cmdDeleteZip := exec.Command("rm", "-rf", zipPath)
	if err := cmdDeleteZip.Run(); err != nil {
		log.Printf("Error deleting zip file: %v", err)
		// Não falhar a requisição se apenas a remoção do arquivo zip falhar
	}

	// Organizar os arquivos descompactados (Opcional)
	// Este passo pode envolver renomear diretórios, mover arquivos, etc.
	cmdOrganize := exec.Command("mv", unpackedPath, folderName)
	if err := cmdOrganize.Run(); err != nil {
		log.Printf("Error organizing unpacked files: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to organize unpacked files"})
		return
	}

	// Resposta de sucesso
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "Files unpacked and organized successfully",
		"path":    folderName,
	})
}

func (h *PackageHandler) CreateGithubBranch(w http.ResponseWriter, r *http.Request) {
	// Recuperar as informações necessárias do contexto ou corpo da requisição
	owner := r.Context().Value("githubOwner").(string)
	repoName := r.Context().Value("githubRepoName").(string)
	baseBranch := "main" // Branch base, pode ser dinâmico conforme a necessidade

	// Criar um identificador único para a nova branch
	branchID := uuid.New().String()

	// Instanciar cliente GitHub, isso poderia ser injetado ou recuperado de algum serviço centralizado
	gc := github.NewGithubClient("yourGitHubUsername", "yourGitHubToken")

	// Criar a branch no GitHub
	createdBranch, err := gc.CreateBranch(owner, repoName, baseBranch, branchID)
	if err != nil {
		log.Printf("Failed to create GitHub branch: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to create GitHub branch"})
		return
	}

	// Resposta de sucesso
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "GitHub branch created successfully",
		"branch":  createdBranch,
	})
}

func (h *PackageHandler) OpenPullRequest(w http.ResponseWriter, r *http.Request) {
	// Recuperar as informações necessárias do contexto ou corpo da requisição
	userID := r.Context().Value("githubOwner").(string)
	repoName := r.Context().Value("githubRepoName").(string)
	sourceBranch := r.Context().Value("sourceBranch").(string) // A branch de onde o PR será aberto
	targetBranch := "main"                                     // A branch para onde o PR será direcionado, pode ser configurável

	// Instanciar cliente GitHub, isso poderia ser injetado ou recuperado de algum serviço centralizado
	gc := github.NewGithubClient("yourGitHubUsername", "yourGitHubToken")

	// Criar título e descrição para o PR
	prTitle := "Update from " + sourceBranch
	prDescription := "Pull request opened at " + time.Now().Format("2006-01-02 15:04:05")

	// Abrir o pull request no GitHub
	user, err := repository.NewUserRepository(h.Repo.DB, utils.NewSlackService()).FindByID(uuid.MustParse(userID))
	prID, err := gc.OpenPullRequest(user, repoName, prTitle, sourceBranch, targetBranch, prDescription)
	if err != nil {
		log.Printf("Failed to open GitHub pull request: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to open GitHub pull request"})
		return
	}

	// Resposta de sucesso
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message":       "GitHub pull request opened successfully",
		"pullRequestID": prID,
	})
}
