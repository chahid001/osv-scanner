package output

import (
	"encoding/json"
	"io"
	"bytes"

	"github.com/google/osv-scanner/pkg/models"
)

// PrintJSONResults writes results to the provided writer in JSON format
func PrintJSONResults(vulnResult *models.VulnerabilityResults, outputWriter io.Writer) (string, error) {
	var buf bytes.Buffer


	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(vulnResult); err != nil {
		return "", err
	}
	jsonString := buf.String()
	return jsonString, nil
}
