package api

import (
	v1 "backend_template_personal/api/v1"
	"backend_template_personal/repository"
	"backend_template_personal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetUpRoutes(r *gin.RouterGroup) {

	codeGenRepo := repository.NewProjectRepository(repository.GormSqlClient.GetDB())
	codeGenService := service.NewCodeGenService(logrus.New(), codeGenRepo)
	codeGenHandler := v1.NewCodeGenHandler(codeGenService)
	r.POST("/generate-crud", codeGenHandler.GenerateCRUD)
	r.POST("/generate-project", codeGenHandler.GenerateProject)
}
