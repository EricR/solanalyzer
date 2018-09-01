package analyzers

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/ericr/solanalyzer/sources"
	"github.com/sirupsen/logrus"
	"regexp"
)

const (
	solVerURL       = "https://raw.githubusercontent.com/ethereum/solidity/develop/CMakeLists.txt"
	solBugsURL      = "https://raw.githubusercontent.com/ethereum/solidity/develop/docs/bugs.json"
	solBugsByVerURL = "https://raw.githubusercontent.com/ethereum/solidity/develop/docs/bugs_by_version.json"
	solSrcVerRegexp = "PROJECT_VERSION \"(.*)\""
)

// CompilerVersionAnalyzer is an analyzer that reports issues related to
// the version of the Solidity compiler declared in a source's solidity pragma.
type CompilerVersionAnalyzer struct {
	LatestVersion    string
	KnownBugs        []*CompilerBug
	KnownBugVersions map[string]*KnownBugVersion
	AffectedBugs     map[string]bool
}

// CompilerBug represents a known bug affecting a version of the Solidity
// compiler.
type CompilerBug struct {
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Introduced  string `json:"introduced"`
	Fixed       string `json:"fixed"`
	Severity    string `json:"severity"`
	URL         string `json:"url"`
}

// KnownBugVersion represents a version of the Solidity compiler with one or
// more known bugs.
type KnownBugVersion struct {
	Bugs []string `json:"bugs"`
}

// Name returns the name of the analyzer.
func (cva *CompilerVersionAnalyzer) Name() string {
	return "compiler version"
}

// ID returns the unique ID of the analyzer.
func (cva *CompilerVersionAnalyzer) ID() string {
	return "compiler-version"
}

// Execute runs the analyzer on a given source.
func (cva *CompilerVersionAnalyzer) Execute(source *sources.Source) ([]*Issue, error) {
	issues := []*Issue{}

	cva.AffectedBugs = map[string]bool{}

	// Try getting the version pragma, otherwise report an issue and return.
	if source.Pragma == nil || source.Pragma.Name != "solidity" {
		issues = append(issues, &Issue{
			Severity:    SeverityInfo,
			Title:       "Missing Version Pragma",
			MsgFormat:   "txt",
			Message:     "No version pragma is declared.",
			analyzer:    cva,
			sourcePath:  source.FilePath,
			sourceStart: source.Pragma.Start,
			sourceStop:  source.Pragma.Stop,
		})
		return issues, nil
	}

	// Fetch the latest version of the Solidity compiler.
	if err := cva.getLatestVersion(); err != nil {
		return issues, err
	}

	// Compare the latest version of the Solidity compiler to the version
	// constraint expressed in the pragma, and report if outdated.
	equal, err := cva.compareToLatest(source.Pragma.Value)
	if err != nil {
		return issues, err
	}

	if !equal {
		msg := fmt.Sprintf("The version pragma, %s, can only be satisfied by "+
			"an outdated version of the Solidity compiler. The source's version "+
			"pragma should be updated to ^%s.",
			source.Pragma.Value, cva.LatestVersion,
		)

		issues = append(issues, &Issue{
			Severity:    SeverityInfo,
			Title:       "Outdated Solidity Compiler",
			MsgFormat:   "txt",
			Message:     msg,
			analyzer:    cva,
			sourcePath:  source.FilePath,
			sourceStart: source.Pragma.Start,
			sourceStop:  source.Pragma.Stop,
		})
	}

	// Fetch a list of known compiler bugs.
	if err := cva.getBugs(); err != nil {
		return issues, err
	}

	// Fetch a list of versions with known compiler bugs.
	if err := cva.getBugVersions(); err != nil {
		return issues, err
	}

	// Match and report on any known compiler bugs.
	matchedBugs, err := cva.matchBugs(source.Pragma.Value)
	if err != nil {
		return issues, err
	}

	for _, bug := range matchedBugs {
		issues = append(issues, &Issue{
			Severity:  cva.bugSeverity(bug.Severity),
			Title:     fmt.Sprintf("Compiler Bug - %s", bug.Name),
			MsgFormat: "txt",
			Message: fmt.Sprintf("The version pragma, %s, can be satisfied "+
				"by a version of the Solidity compiler that contains a known bug. %s "+
				"This bug is reported to be fixed in version %s.",
				source.Pragma.Value, bug.Description, bug.Fixed),
			analyzer:    cva,
			sourcePath:  source.FilePath,
			sourceStart: source.Pragma.Start,
			sourceStop:  source.Pragma.Stop,
		})
	}

	return issues, nil
}

func (cva *CompilerVersionAnalyzer) getLatestVersion() error {
	resp, err := getUrl(solVerURL)
	if err != nil {
		return fmt.Errorf("HTTP error: %s", err)
	}

	verRegexp, err := regexp.Compile(solSrcVerRegexp)
	if err != nil {
		return fmt.Errorf("Regexp error: %s", err)
	}

	versionMatch := verRegexp.FindStringSubmatch(string(resp))
	if len(versionMatch) != 2 {
		return fmt.Errorf("Could not regexp capture latest solc version")
	}

	cva.LatestVersion = versionMatch[1]

	return nil
}

func (cva *CompilerVersionAnalyzer) getBugs() error {
	resp, err := getUrl(solBugsURL)
	if err != nil {
		return fmt.Errorf("HTTP error: %s", err)
	}

	if err := json.Unmarshal(resp, &cva.KnownBugs); err != nil {
		return fmt.Errorf("Error parsing JSON %s", err)
	}

	return nil
}

func (cva *CompilerVersionAnalyzer) getBugVersions() error {
	resp, err := getUrl(solBugsByVerURL)
	if err != nil {
		return fmt.Errorf("HTTP error: %s", err)
	}

	if err := json.Unmarshal(resp, &cva.KnownBugVersions); err != nil {
		return fmt.Errorf("Error parsing JSON %s", err)
	}

	return nil
}

func (cva *CompilerVersionAnalyzer) compareToLatest(version string) (bool, error) {
	constraint, err := semver.NewConstraint(version)
	if err != nil {
		return false, err
	}

	latest, err := semver.NewVersion(cva.LatestVersion)
	if err != nil {
		return false, err
	}

	equal, _ := constraint.Validate(latest)

	return equal, nil
}

func (cva *CompilerVersionAnalyzer) matchBugs(pragmaVer string) ([]*CompilerBug, error) {
	matchedBugNames := []string{}
	matchedBugs := []*CompilerBug{}

	constraint, err := semver.NewConstraint(pragmaVer)
	if err != nil {
		return matchedBugs, err
	}

	for versionStr, knownBugVersion := range cva.KnownBugVersions {
		version, err := semver.NewVersion(versionStr)
		if err != nil {
			logrus.Warnf("Invalid semver version, skipping: %s", err)
			continue
		}

		if equal, _ := constraint.Validate(version); equal {
			for _, bugName := range knownBugVersion.Bugs {
				if !cva.AffectedBugs[bugName] {
					cva.AffectedBugs[bugName] = true
					matchedBugNames = append(matchedBugNames, bugName)
				}
			}
		}
	}

	for _, bugName := range matchedBugNames {
		for _, bug := range cva.KnownBugs {
			if bugName == bug.Name {
				matchedBugs = append(matchedBugs, bug)
			}
		}
	}

	return matchedBugs, nil
}

func (cva *CompilerVersionAnalyzer) bugSeverity(severity string) int {
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
