package reporter

import (
	"fmt"
	"io"
	"github.com/google/osv-scanner/internal/output"
	"github.com/google/osv-scanner/pkg/models"
)

// JSONReporter prints vulnerability results in JSON format to stdout. Runtime information
// will be written to stderr.
type JSONReporter struct {
	hasErrored bool
	stdout     io.Writer
	stderr     io.Writer
	level      VerbosityLevel
	failRule   int
	project    string
	name_rp    string
}

func NewJSONReporter(stdout io.Writer, stderr io.Writer, level VerbosityLevel, failRule int, project string, name_rp string) *JSONReporter {
	return &JSONReporter{
		stdout:     stdout,
		stderr:     stderr,
		level:      level,
		hasErrored: false,
		failRule: failRule,
		project: project,
		name_rp: name_rp,
	}
}

func (r *JSONReporter) Errorf(format string, a ...any) {
	fmt.Fprintf(r.stderr, format, a...)
	r.hasErrored = true
}

func (r *JSONReporter) HasErrored() bool {
	return r.hasErrored
}

func (r *JSONReporter) Warnf(format string, a ...any) {
	if WarnLevel <= r.level {
		fmt.Fprintf(r.stderr, format, a...)
	}
}

func (r *JSONReporter) Infof(format string, a ...any) {
	if InfoLevel <= r.level {
		fmt.Fprintf(r.stderr, format, a...)
	}
}

func (r *JSONReporter) Verbosef(format string, a ...any) {
	if VerboseLevel <= r.level {
		fmt.Fprintf(r.stderr, format, a...)
	}
}

func (r *JSONReporter) PrintResult(vulnResult *models.VulnerabilityResults) (error) {
	

	in, nil := output.PrintJSONResults(vulnResult, r.stdout)

	parse_result(in, r.project, r.name_rp, r.failRule)

	return nil
}

