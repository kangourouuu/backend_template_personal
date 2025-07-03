package generate

import (
	"backend_template_personal/repository"
	"backend_template_personal/service"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ctx context.Context
var projectName string
var port int
var logger = logrus.New()

var GenerateProject = &cobra.Command{
	Use:   "generate-project",
	Short: "ProjectName and Port is required",
	Run: func(cmd *cobra.Command, args []string) {
		ctx = context.Background()
		db := repository.GormSqlClient
		repo := repository.NewProjectRepository(db.GetDB())
		codeGenSvc := service.NewCodeGenService(logger, repo)
		err := codeGenSvc.GenerateProject(ctx, projectName, port)
		if err != nil {
			logger.Error("failed to generate", err)
		}
		logger.Info("generate successfully")
	},
}


