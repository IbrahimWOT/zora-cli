package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/IbrahimWOT/zora-cli/internal/converter"
	"github.com/IbrahimWOT/zora-cli/internal/validator"
	"github.com/spf13/cobra"
)

var (
	inputPath    string
	outputFormat string
	outputPath   string
)

var rootCmd = &cobra.Command{
	Use:   "zora",
	Short: "Zora is a lightweight CLI tool to convert JSON and YAML files",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validator.ValidateFormat(inputPath, outputFormat)
		if err != nil {
			return err
		}

		if outputPath == "./output" {
			err := os.MkdirAll(outputPath, 0755)
			if err != nil {
				return fmt.Errorf("failed to create output directory: %v", err)
			}
			filename := "output." + outputFormat
			outputPath = filepath.Join(outputPath, filename)
		}

		err = converter.Transform(inputPath, outputFormat, outputPath)
		if err != nil {
			return err
		}

		fmt.Printf("🚀 Success: File successfully converted and saved to %s\n", outputPath)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Path to input file (Required)")
	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Desired output format: json or yaml (Required)")
	rootCmd.Flags().StringVarP(&outputPath, "path", "p", "./output", "Directory or path for the generated file")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
}
