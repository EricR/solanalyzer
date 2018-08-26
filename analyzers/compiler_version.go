package analyzers

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/ericr/solanalyzer/scanner"
	"github.com/sirupsen/logrus"
	"regexp"
)

const (
	SolVerURL       = "https://raw.githubusercontent.com/ethereum/solidity/develop/CMakeLists.txt"
	SolBugsURL      = "https://raw.githubusercontent.com/ethereum/solidity/develop/docs/bugs.json"
	SolSrcVerRegexp = "PROJECT_VERSION \"?[0-9.]+(?=\")"
)

// CompilerVersionAnalyzer is an analyzer that reports issues related to
// the version of the Solidity compiler declared in a source's pragma.
type CompilerVersionAnalyzer struct {
	LatestVersion string
}

// CompilerBug represents a known bug affecting the Solidity compiler.
type CompilerBug struct {
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Introduced  string `json:"introduced"`
	Fixed       string `json:"fixed"`
	Severity    string `json:"severity"`
}

// GetName returns the name of the analyzer.
func (osp *CompilerVersionAnalyzer) GetName() string {
	return "Solidity Compiler Version"
}

// GetID returns the unique ID of the analyer.
func (osp *CompilerVersionAnalyzer) GetID() string {
	return "compiler-version"
}

// Execute runs the analyzer on a given source.
func (cva *CompilerVersionAnalyzer) Execute(source *scanner.Source) ([]*Issue, error) {
	issues := []*Issue{}

	if source.Pragma == nil || source.Pragma.Version == "" {
		issues = append(issues, &Issue{
			Severity:   SeverityInfo,
			Title:      "Missing Pragma Declaration",
			MsgFormat:  "txt",
			Message:    "No pragma is declared.",
			analyzer:   cva,
			sourcePath: source.FilePath,
		})

		return issues, nil
	}

	version, err := cva.getLatestVersion()
	if err != nil {
		return issues, err
	}

	equal, _, err := cva.compareVersions(source.Pragma.Version, version)
	if err != nil {
		return issues, err
	}

	if !equal {
		msg := fmt.Sprintf("The version constraint, %s, can only be satisfied by "+
			"an outdated version of the Solidity compiler. The latest version is %s. "+
			"The source's pragma declaration should be updated accordingly.",
			source.Pragma.Version, version,
		)

		issues = append(issues, &Issue{
			Severity:   SeverityInfo,
			Title:      "Outdated Solidity Compiler",
			MsgFormat:  "txt",
			Message:    msg,
			analyzer:   cva,
			sourcePath: source.FilePath,
			tokens:     source.Pragma.Tokens,
		})
	}

	bugs, err := cva.getCompilerBugs()
	if err != nil {
		return issues, err
	}

	matchedBugs, err := cva.matchCompilerBugs(source.Pragma.Version, bugs)
	if err != nil {
		return issues, err
	}

	for _, bug := range matchedBugs {
		issues = append(issues, &Issue{
			Severity:   cva.compilerBugSeverity(bug.Severity),
			Title:      fmt.Sprintf("Compiler Bug - %s", bug.Name),
			MsgFormat:  "txt",
			Message:    bug.Description,
			analyzer:   cva,
			sourcePath: source.FilePath,
			tokens:     source.Pragma.Tokens,
		})
	}

	return issues, nil
}

func (cva *CompilerVersionAnalyzer) getLatestVersion() (string, error) {
	resp, err := getUrl(SolVerURL)
	if err != nil {
		return "", fmt.Errorf("HTTP error: %s", err)
	}

	verRegexp, err := regexp.Compile("PROJECT_VERSION \"(.*)\"")
	if err != nil {
		return "", fmt.Errorf("Regexp error: %s", err)
	}

	versionMatch := verRegexp.FindStringSubmatch(string(resp))
	if len(versionMatch) != 2 {
		return "", fmt.Errorf("Could not regexp capture latest solc version")
	}

	return versionMatch[1], nil
}

func (cva *CompilerVersionAnalyzer) compareVersions(usedVersion string, latestVersion string) (bool, []error, error) {
	constraint, err := semver.NewConstraint(usedVersion)
	if err != nil {
		return false, []error{}, err
	}

	latest, err := semver.NewVersion(latestVersion)
	if err != nil {
		return false, []error{}, err
	}

	equal, errors := constraint.Validate(latest)

	return equal, errors, nil
}

func (cva *CompilerVersionAnalyzer) getCompilerBugs() ([]*CompilerBug, error) {
	cbugs := []*CompilerBug{}

	resp, err := getUrl(SolBugsURL)
	if err != nil {
		return cbugs, fmt.Errorf("HTTP error: %s", err)
	}

	if err := json.Unmarshal(resp, &cbugs); err != nil {
		return cbugs, fmt.Errorf("Error parsing JSON %s", err)
	}

	return cbugs, nil
}

func (cva *CompilerVersionAnalyzer) matchCompilerBugs(pragmaVer string, bugs []*CompilerBug) ([]*CompilerBug, error) {
	matchedBugs := []*CompilerBug{}

	pragmaConstraint, err := semver.NewVersion(cva.constraintToLowestVersion(pragmaVer))
	if err != nil {
		return matchedBugs, err
	}

	for _, bug := range bugs {
		if bug.Introduced == "" {
			bug.Introduced = "0.0.0"
		}

		var constrStr string

		if bug.Fixed == "" {
			constrStr = fmt.Sprintf(">= %s", bug.Introduced)
		} else {
			constrStr = fmt.Sprintf(">= %s, < %s", bug.Introduced, bug.Fixed)
		}

		bugConstraint, err := semver.NewConstraint(constrStr)
		if err != nil {
			logrus.Warnf("Invalid semver constraint, skipping: %s", err)
			continue
		}

		if affected, _ := bugConstraint.Validate(pragmaConstraint); affected {
			matchedBugs = append(matchedBugs, bug)
		}
	}

	return matchedBugs, nil
}

func (cva *CompilerVersionAnalyzer) constraintToLowestVersion(pragmaVer string) string {
	switch pragmaVer[0] {
	case '=', '^', '>', '<', '~':
		return pragmaVer[1:len(pragmaVer)]
	default:
		return pragmaVer
	}
}

func (cva *CompilerVersionAnalyzer) compilerBugSeverity(severity string) int {
	switch severity {
	case "very low":
		return SeverityLow
	case "low":
		return SeverityLow
	case "medium":
		return SeverityMed
	case "medium/high":
		return SeverityHigh
	case "high":
		return SeverityHigh
	default:
		return SeverityHigh // would rather aim high than low here
	}
}
