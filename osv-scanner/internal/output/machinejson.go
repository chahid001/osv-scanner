package output

import (
	"encoding/json"
	"io"
	"bytes"
	"os"
	"path/filepath"	

	"github.com/google/osv-scanner/pkg/models"
)

// PrintJSONResults writes results to the provided writer in JSON format

func PrintJSONResults(vulnResult *models.VulnerabilityResults, outputWriter io.Writer) error {
	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(vulnResult); err != nil {
		return err
	}

	// Write JSON string to file
	if err := os.MkdirAll(filepath.Dir("/osv-scanner/reports/osv-scan-report.json"), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create("/osv-scanner/reports/osv-scan-report.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	return nil
}
