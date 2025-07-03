package generate

import (
	"backend_template_personal/repository"
	"backend_template_personal/service"
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type Fields struct {
	Name string
	Type string
}

var entityName string
var fieldString string

var GenerateCRUD = &cobra.Command{
	Use:   "generate-crud",
	Short: "Use for generate CRUD",
	Run: func(cmd *cobra.Command, args []string) {
		ctx = context.Background()
		db := repository.GormSqlClient
		repo := repository.NewProjectRepository(db.GetDB())
		codeGenSvc := service.NewCodeGenService(logger, repo)
		fields, err := parseFields(fieldString)
		if err != nil {
			logger.Fatal("faile to parse fields")
		}
		err = codeGenSvc.GenerateCRUD(ctx, entityName, fields, projectName)
		if err != nil {
			logger.Error("failed to generate", err)
		}
		logger.Info("generate crud successfully")
	},
}

func parseFields(input string) ([]service.Field, error) {
	var fields []service.Field
	pairs := strings.Split(input, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid field format: %s", pair)
		}
		fields = append(fields, service.Field{
			Name: parts[0],
			Type: parts[1],
		})
	}
	return fields, nil
}
