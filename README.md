# 🚀 OSV-SCAN Custom Build

### This version is a modification of Google's open-source software, OSV-SCAN, which supports Flutter & Dart and is written in Golang and improved with Python.

---

## ✨ What's New?

- 📊 **HTML Reporting**: Generate visually appealing HTML reports with ease.
- 📝 **Custom Naming**: Personalize your reports with bespoke names.
- ❌ **CVSS Rule Enforcement**: Implement rules to automatically fail scans based on CVSS scores.
- 🏷️ **Project Tagging**: Streamline project identification with custom tags.

---

## 🛠️ Usage

```bash
osv-scanner scan --html --project="project-name" -L $(pwd)/pubspec.lock
```

### **Mandatory**:
- `-L`: Specify the path to the package lockfile.
- `--html`: Enable HTML report generation.

### **Optional**:
- `--project`: Define a project name for the report.
- `--name`: Customize the output report name.
- `--failOnCVSS`: Set a CVSS score threshold for failing.

## 🚩 All Flags Explained

```bash
NAME:
   osv-scanner scan - scans various mediums for dependencies and matches it against the OSV database

USAGE:
   osv-scanner scan command [command options] [directory1 directory2...]

DESCRIPTION:
   scans various mediums for dependencies and matches it against the OSV database

COMMANDS:
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --docker value, -D value [ --docker value, -D value ]            scan docker image with this name
   --lockfile value, -L value [ --lockfile value, -L value ]        scan package lockfile on this path
   --sbom value, -S value [ --sbom value, -S value ]                scan sbom file on this path
   --config value                                                   set/override config file
   --format value, -f value                                         sets the output format; value can be: table, json, markdown, sarif, gh-annotations (default: "table")
   --json                                                           sets output to json (deprecated, use --format json instead) (default: false)
   --output value                                                   saves the result to the given file path
   --skip-git                                                       skip scanning git repositories (default: false)
   --recursive, -r                                                  check subdirectories (default: false)
   --experimental-call-analysis                                     [Deprecated] attempt call analysis on code to detect only active vulnerabilities (default: false)
   --no-ignore                                                      also scan files that would be ignored by .gitignore (default: false)
   --call-analysis value [ --call-analysis value ]                  attempt call analysis on code to detect only active vulnerabilities
   --no-call-analysis value [ --no-call-analysis value ]            disables call graph analysis
   --verbosity value                                                specify the level of information that should be provided during runtime; value can be: error, warn, info, verbose (default: "info")
   --experimental-local-db                                          checks for vulnerabilities using local databases (default: false)
   --experimental-offline                                           checks for vulnerabilities using local databases that are already cached (default: false)
   --experimental-all-packages                                      when json output is selected, prints all packages (default: false)
   --experimental-licenses-summary                                  report a license summary, implying the --experimental-all-packages flag (default: false)
   --experimental-licenses value [ --experimental-licenses value ]  report on licenses based on an allowlist
   --html                                                           generate an HTML report (developed by chahid001) (default: false)
   --failOnCVSS value                                               add CVSS rule for failing (developed by chahid001) (default: -1)
   --project value                                                  add project name (developed by chahid001) (default: "false")
   --name value                                                     add report name (developed by chahid001) (default: "false")
   --help, -h                                                       show help

```
