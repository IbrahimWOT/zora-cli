package converter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IbrahimWOT/zora-cli/internal/models"
	"gopkg.in/yaml.v3"
)

func Transform(inputPath string, outputFormat string, outputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("could not read input file: %v", err)
	}

	var config models.Config

	if json.Unmarshal(data, &config) != nil {
		if err := yaml.Unmarshal(data, &config); err != nil {
			return fmt.Errorf("failed to parse input file as JSON or YAML: %v", err)
		}
	}

	var finalData []byte
	var marshalErr error

	if outputFormat == "json" {
		finalData, marshalErr = json.MarshalIndent(config, "", "  ")
	} else {
		finalData, marshalErr = yaml.Marshal(config)
	}

	if marshalErr != nil {
		return fmt.Errorf("failed to convert data to %s: %v", outputFormat, marshalErr)
	}

	err = os.WriteFile(outputPath, finalData, 0644)
	if err != nil {
		return fmt.Errorf("failed to save output file: %v", err)
	}

	return nil
}
