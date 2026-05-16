package validator

import (
	"errors"
	"path/filepath"
	"strings"
)

func ValidateFormat(inputPath string, outputFormat string) error {

	ext := filepath.Ext(inputPath)
	currentFormat := strings.ToLower(strings.TrimPrefix(ext, "."))
	targetFormat := strings.ToLower(outputFormat)

	if currentFormat == targetFormat {
		return errors.New("requested output must be in different format than input file")
	}

	return nil
}
