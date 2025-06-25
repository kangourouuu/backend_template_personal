package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/sirupsen/logrus"
	"backend_template_personal/model"
	"backend_template_personal/repository"
)

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TemplateData struct {
	EntityName      string  `json:"entity_name"`
	EntityNameLower string  `json:"entity_name_lower"`
	Fields          []Field `json:"fields"`
	ProjectName     string  `json:"project_name"`
	Port            int     `json:"port"`

	// Thêm các field mới cho Swagger
	Version     string `json:"version"`
	Host        string `json:"host"`
	BasePath    string `json:"base_path"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var typeValid = []string{
	"int",
	"string",
	"float64",
}

type CodeGenService interface {
	GenerateCRUD(ctx context.Context, entityName string, fields []Field, projectName string) error
	GenerateProject(ctx context.Context, projectName string, port int) error
	Create(*model.ProjectInfo) error
}

type codeGenService struct {
	logger *logrus.Logger
	repo   repository.ProjectRepository
}

func NewCodeGenService(logger *logrus.Logger, repo repository.ProjectRepository) CodeGenService {
	return &codeGenService{
		logger: logger,
		repo:   repo,
	}
}

func (s *codeGenService) Create(prj *model.ProjectInfo) error {
	return s.repo.Create(&model.ProjectInfo{})
}

// Generate CRUD
func (s *codeGenService) GenerateCRUD(ctx context.Context, entityName string, fields []Field, projectName string) error {

	for _, field := range fields {
		valid := false
		for _, typeInput := range typeValid {
			if field.Type == typeInput {
				valid = true
				break
			}
		}
		if !valid {
			s.logger.Infof("invalid input type: %s", field.Type)
			return fmt.Errorf("invalid input type: %s", field.Type)
		}
	}

	data := TemplateData{
		EntityName:      entityName,
		EntityNameLower: strings.ToLower(entityName),
		Fields:          fields,
		ProjectName:     projectName,
	}

	basePath := filepath.Clean(filepath.Join("C:\\", projectName))

	tmplFiles := []struct {
		tmplPath string
		outPath  string
	}{
		{"templates/model.tmpl", filepath.Join(basePath, "model", strings.ToLower(entityName)+".go")},
		{"templates/repository.tmpl", filepath.Join(basePath, "repository", strings.ToLower(entityName)+"_repository.go")},
		{"templates/service.tmpl", filepath.Join(basePath, "service", strings.ToLower(entityName)+"_service.go")},
		{"templates/handler.tmpl", filepath.Join(basePath, "api", "v2", strings.ToLower(entityName)+"_handler.go")},
		{"templates/setup_routes.tmpl", filepath.Join(basePath, "api", strings.ToLower(entityName)+"_routes.go")},
		{"templates/dto.tmpl", filepath.Join(basePath, "dto", strings.ToLower(entityName)+"_dto.go")},
	}

	for _, tf := range tmplFiles {
		if err := s.generateFile(tf.tmplPath, tf.outPath, data); err != nil {
			s.logger.Error("Failed to generate file: ", err)
			return err
		}
	}

	return nil
}

// Generate Project
func (s *codeGenService) GenerateProject(ctx context.Context, projectName string, port int) error {
	// Auto fallback entity name from project name if not provided
	entityName := fallbackEntityFromProject(projectName)

	data := TemplateData{
		ProjectName:     projectName,
		Port:            port,
		Version:         "1.0",
		Host:            fmt.Sprintf("localhost:%d", port),
		BasePath:        "/api/v2",
		Title:           fmt.Sprintf("%s API", projectName),
		Description:     fmt.Sprintf("API Documentation for %s", projectName),
		EntityName:      entityName,
		EntityNameLower: strings.ToLower(entityName),
		Fields:          []Field{},
	}

	basePath := filepath.Clean(filepath.Join("C:\\", projectName))

	dirs := []string{
		basePath,
		filepath.Join(basePath, "api", "v2"),
		filepath.Join(basePath, "build"),
		filepath.Join(basePath, "common", "api_response"),
		filepath.Join(basePath, "common", "err_response"),
		filepath.Join(basePath, "common", "log"),
		filepath.Join(basePath, "common", "limiter"),
		filepath.Join(basePath, "configs"),
		filepath.Join(basePath, "constant"),
		filepath.Join(basePath, "internal", "sqlclient"),
		filepath.Join(basePath, "internal", "redis"),
		filepath.Join(basePath, "middleware"),
		filepath.Join(basePath, "model"),
		filepath.Join(basePath, "repository"),
		filepath.Join(basePath, "server", "http"),
		filepath.Join(basePath, "service"),
		filepath.Join(basePath, "schema"),
		filepath.Join(basePath, "tmp"),
		filepath.Join(basePath, "dto"),
		filepath.Join(basePath, "docs"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			s.logger.Error("Failed to create directory: ", err)
			return err
		}
	}

	tmplFiles := []struct {
		tmplPath string
		outPath  string
	}{
		{"templates/main.tmpl", filepath.Join(basePath, "main.go")},
		{"templates/go.mod.tmpl", filepath.Join(basePath, "go.mod")},
		{"templates/config.tmpl", filepath.Join(basePath, "configs", "config.json")},
		{"templates/http_server.tmpl", filepath.Join(basePath, "server", "http", "http_server.go")},
		{"templates/gorm_sql.tmpl", filepath.Join(basePath, "internal", "sqlclient", "gorm_sql.go")},
		{"templates/init_repo.tmpl", filepath.Join(basePath, "internal", "repository.go")},
		{"templates/log.tmpl", filepath.Join(basePath, "common", "log", "log.go")},
		{"templates/auth.tmpl", filepath.Join(basePath, "middleware", "auth.go")},
		{"templates/env.tmpl", filepath.Join(basePath, ".env")},
		{"templates/err_response.tmpl", filepath.Join(basePath, "common", "err_response", "error.go")},
		{"templates/result_response.tmpl", filepath.Join(basePath, "common", "api_response", "response.go")},
		{"templates/constant.tmpl", filepath.Join(basePath, "constant", "constant.go")},
		{"templates/readme.tmpl", filepath.Join(basePath, "README.md")},
		{"templates/limiter.tmpl", filepath.Join(basePath, "common", "limiter", "limiter.go")},
		{"templates/redis.tmpl", filepath.Join(basePath, "internal", "redis", "redis.go")},
		// {"templates/docs.tmpl", filepath.Join(basePath, "docs", "docs.go")},
		// {"templates/swagger.json.tmpl", filepath.Join(basePath, "docs", "swagger.json")},
		// {"templates/swagger.yaml.tmpl", filepath.Join(basePath, "docs", "swagger.yaml")},
	}

	for _, tf := range tmplFiles {
		if err := s.generateFile(tf.tmplPath, tf.outPath, data); err != nil {
			s.logger.Error("Failed to generate file: ", err)
			return err
		}
	}

	info := &model.ProjectInfo{
		ProjectName: projectName,
		Port:        port,
		CreatedAt:   time.Now(),
	}

	err := s.repo.Create(info)
	if err != nil {
		fmt.Print("fail to insert")
	}

	return nil
}

// fallbackEntityFromProject creates a default entity name from the project name (e.g., "sword_management" -> "Sword")
func fallbackEntityFromProject(project string) string {
	project = strings.TrimSpace(project)
	parts := strings.Split(project, "_")
	if len(parts) == 0 {
		return "Entity"
	}
	return strings.Title(parts[0]) // e.g. "Sword"
}

func (s *codeGenService) generateFile(tmplPath, outPath string, data interface{}) error {
	// Create template with functions
	tmpl := template.New(filepath.Base(tmplPath)).Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(s string) string {
			return template.HTMLEscapeString(s)
		},
		"lowerFirst": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToLower(s[:1]) + s[1:]
		},
		"ToLower": strings.ToLower,
		"sub1": func(i int) int {
			return i - 1
		},
		"SwaggerType": func(goType string) string {
			switch goType {
			case "int":
				return "integer"
			case "float64":
				return "number"
			default:
				return "string"
			}
		},
	})

	// Parse template file
	tmpl, err := tmpl.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Create output directory if not exists
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Create output file
	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer f.Close()

	// Execute template
	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}
