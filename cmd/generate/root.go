/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package generate

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "codegen",
	Short: "A CLI tool for generating Go backend code",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	GenerateProject.Flags().StringVarP(&projectName, "name", "n", "", "ProjectName")
	GenerateProject.Flags().IntVarP(&port, "port", "p", 9003, "Port")
	GenerateProject.MarkFlagRequired("name")

	GenerateCRUD.Flags().StringVarP(&entityName, "entity", "E", "", "EntityName")
	GenerateCRUD.Flags().StringVarP(&projectName, "name", "n", "", "ProjectName")
	GenerateCRUD.Flags().StringVarP(&fieldString, "fields", "f", "", "Field format: name:type")

	rootCmd.AddCommand(GenerateCRUD)
	rootCmd.AddCommand(GenerateProject)
}
