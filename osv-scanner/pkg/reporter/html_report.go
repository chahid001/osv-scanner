package reporter

import (
	"os/exec"
	"fmt"
	"strconv"
	"os"
)



func parse_result(input string, project string, path string, failOn int) {

	if failOn == -1 {
		failOn = 10	
	}

	if project == "" {
		project = "-"
	}

	if path == "" {
		path = "osv-scan-report.html"
	}

    pythonScript := "/workspace/python-script/parser.py"

	args := []string{input, project, path, strconv.Itoa(failOn)}

	cmd := exec.Command(pythonScript, args...)
    output, err := cmd.CombinedOutput()

    if err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok {
            fmt.Println("Error running Python script (exit code:", exitErr.ExitCode(), ")")
            if len(output) > 0 {
                fmt.Println(string(output))
            }
            os.Exit(1)
        } else {
            fmt.Println("Error executing Python script:", err)
            os.Exit(1)
        }
    }
    fmt.Println(string(output))
}