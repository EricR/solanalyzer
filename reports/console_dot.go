package reports

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func processDot(analyzer_id string, dot []byte) string {
	outputDir := filepath.Join(".", "output")

	dotOutputFile := fmt.Sprintf("solanalyzer_%s_%s.dot",
		analyzer_id, time.Now().Format("20060102150405"))

	dotOutputPath := filepath.Join(outputDir, dotOutputFile)

	svgOutputFile := fmt.Sprintf("solanalyzer_%s_%s.svg",
		analyzer_id, time.Now().Format("20060102150405"))

	svgOutputPath := filepath.Join(outputDir, svgOutputFile)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Sprintf("Unable to create directory %s: %s", outputDir, err)
	}

	if err := ioutil.WriteFile(dotOutputPath, dot, 0644); err != nil {
		return fmt.Sprintf("Unable to write to file %s: %s", dotOutputPath, err)
	}

	dotPath, err := exec.LookPath("dot")
	if err != nil {
		return fmt.Sprintf("Got error: %s", err)
	}

	var stdErr bytes.Buffer

	cmd := exec.Command(dotPath, "-Tsvg", "-o", svgOutputPath)
	cmd.Stdin = bytes.NewReader(dot)
	cmd.Stderr = &stdErr

	if err := cmd.Run(); err != nil {
		return fmt.Sprintf("Error running dot: %s: %s", err, stdErr.String())
	}

	return fmt.Sprintf("Wrote SVG file to %s", svgOutputPath)
}
