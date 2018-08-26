package main

import (
	"errors"
	"fmt"
	"github.com/ericr/solanalyzer/analyzers"
	"github.com/ericr/solanalyzer/reports"
	"github.com/ericr/solanalyzer/scanner"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/x-cray/logrus-prefixed-formatter"
	"time"
)

var Version = "0.1-beta"

type Options struct {
	Verbose bool
}

func main() {
	options := &Options{}
	rootCmd := &cobra.Command{
		Use: "solanalyzer path",
		Short: "SolAnalyzer is a static analyzer for the Solidity programming language, " +
			"with a focus on finding security bugs.",
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
			session := scanner.NewSession()

			logrus.Info("Parsing sources")
			session.ParsePath(args)

			logrus.Info("Scanning sources")
			session.Scan()

			group := analyzers.NewGroup("all")
			group.AddAnalyzer(&analyzers.CompilerVersionAnalyzer{})
			group.AddAnalyzer(&analyzers.FunctionVisibilityAnalyzer{})

			logrus.Info("Running analysis")
			group.AnalyzeMany(session.Sources)

			logrus.Info("Generating report")
			report := reports.NewConsoleReport()
			report.Generate(group)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&options.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Execute()
}

func setLoggingLevel(options *Options) {
	if options.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func printHeader() {
	fmt.Printf("\nSolAnalyzer v%s\n", Version)
	fmt.Printf("Copyright %d Eric Rafaloff\n\n", time.Now().Year())
	fmt.Printf("This is beta software. Please report issues at " +
		"https://github.com/EricR/solanalyzer/issues/.\n\n")
}
