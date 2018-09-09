package main

import (
	"errors"
	"fmt"
	"github.com/ericr/solanalyzer/analyzers"
	"github.com/ericr/solanalyzer/reports"
	"github.com/ericr/solanalyzer/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/x-cray/logrus-prefixed-formatter"
	"time"
)

var version = "0.1-beta"

type options struct {
	Verbose bool
}

func main() {
	options := &options{}
	rootCmd := &cobra.Command{
		Use: "solanalyzer path",
		Short: "SolAnalyzer is a static analyzer for the Solidity programming " +
			"language, with a focus on finding security bugs.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("Missing target path")
			}
			return nil
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&prefixed.TextFormatter{FullTimestamp: true})
			setLoggingLevel(options)
		},
		Run: func(cmd *cobra.Command, args []string) {
			printHeader()

			logrus.Info("Starting new session")
			session := sessions.NewSession()

			logrus.Info("Parsing sources")
			session.ParsePath(args)

			logrus.Info("Scanning sources")
			session.VisitSources()

			session.AddAnalyzer(&analyzers.CompilerVersionAnalyzer{})
			session.AddAnalyzer(&analyzers.FunctionVisibilityAnalyzer{})

			logrus.Info("Analyzing sources")
			session.Analyze()

			logrus.Info("Generating report")
			session.GenerateReport(&reports.ConsoleReport{})
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&options.Verbose, "verbose", "v", false,
		"verbose output")
	rootCmd.Execute()
}

func setLoggingLevel(options *options) {
	if options.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func printHeader() {
	fmt.Printf("\nSolAnalyzer v%s\n", version)
	fmt.Printf("Copyright %d Eric Rafaloff\n\n", time.Now().Year())
	fmt.Printf("This is beta software. Please report issues at " +
		"https://github.com/EricR/solanalyzer/issues/.\n\n")
}
