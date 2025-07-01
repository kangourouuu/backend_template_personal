package v1

import (
	"net/http"

	"backend_template_personal/common/apperror"
	"backend_template_personal/common/response"
	"backend_template_personal/service"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, apperror.ErrInvalidRequest(err.Error()))
		return
	}

	if req.EntityName == "" || req.ProjectName == "" || len(req.Fields) == 0 {
		c.JSON(http.StatusBadRequest, apperror.ErrInvalidRequest("fields is required"))
		return
	}

	if err := h.codeGenService.GenerateCRUD(c.Request.Context(), req.EntityName, req.Fields, req.ProjectName); err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrInternalServer(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessCreated("CRUD generated successfully"))
}

func (h *CodeGenHandler) GenerateProject(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrInvalidRequest(err.Error()))
		return
	}
	if req.ProjectName == "" {
		c.JSON(http.StatusBadRequest, apperror.ErrInvalidRequest("ProjectName is required"))
		return
	}
	if err := h.codeGenService.GenerateProject(c.Request.Context(), req.ProjectName, req.Port); err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrInternalServer(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessCreated("New Project is generated"))
}
