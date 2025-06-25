package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend_template_personal/service"
)

type GenerateRequest struct {
	EntityName  string          `json:"entity_name"`
	Fields      []service.Field `json:"fields"`
	ProjectName string          `json:"project_name"`
	Port        int             `json:"port"`
}

type CodeGenHandler struct {
	codeGenService service.CodeGenService
}

func NewCodeGenHandler(codeGenService service.CodeGenService) *CodeGenHandler {
	return &CodeGenHandler{codeGenService: codeGenService}
}

func (h *CodeGenHandler) GenerateCRUD(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.EntityName == "" || req.ProjectName == "" || len(req.Fields) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.codeGenService.GenerateCRUD(c.Request.Context(), req.EntityName, req.Fields, req.ProjectName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CRUD code generated successfully"})
}

func (h *CodeGenHandler) GenerateProject(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.ProjectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectName is required"})
		return
	}
	if err := h.codeGenService.GenerateProject(c.Request.Context(), req.ProjectName, req.Port); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project code generated successfully"})
}
