package reporter

import (
    "os"
    "os/exec"
    "fmt"
    "strconv"
)

func parse_result_npm(inputFile string, project string, path string, failOn int) {

    if failOn == -1 {
        failOn = 10
    }

    if project == "" {
        project = "-"
    }

    if path == "" {
        path = "osv-scan-report.html"
    }

    pythonScript := "/osv-scanner/python-script/parser_npm.py"

    jsonData, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        os.Exit(1)
    }

    tmpFile, err := os.CreateTemp("", "osv-scan-data")
    if err != nil {
        fmt.Println("Error creating temporary file:", err)
        os.Exit(1)
    }
    defer os.Remove(tmpFile.Name()) 

    if _, err := tmpFile.Write(jsonData); err != nil {
        fmt.Println("Error writing JSON data to temporary file:", err)
        os.Exit(1)
    }

    if err := tmpFile.Close(); err != nil {
        fmt.Println("Error closing temporary file:", err)
        os.Exit(1)
    }

    args := []string{tmpFile.Name(), project, path, strconv.Itoa(failOn)}

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


func parse_result_pub(inputFile string, project string, path string, failOn int) {

    if failOn == -1 {
        failOn = 10
    }

    if project == "" {
        project = "-"
    }

    if path == "" {
        path = "osv-scan-report.html"
    }

    pythonScript := "/osv-scanner/python-script/parser_pub.py"

    jsonData, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        os.Exit(1)
    }

    tmpFile, err := os.CreateTemp("", "osv-scan-data")
    if err != nil {
        fmt.Println("Error creating temporary file:", err)
        os.Exit(1)
    }
    defer os.Remove(tmpFile.Name()) 

    if _, err := tmpFile.Write(jsonData); err != nil {
        fmt.Println("Error writing JSON data to temporary file:", err)
        os.Exit(1)
    }

    if err := tmpFile.Close(); err != nil {
        fmt.Println("Error closing temporary file:", err)
        os.Exit(1)
    }

    args := []string{tmpFile.Name(), project, path, strconv.Itoa(failOn)}

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

