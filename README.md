# SolAnalyzer

SolAnalyzer is a static analyzer for the Solidity programming language, with a focus on finding security bugs.

**Warning**: This software is in beta and false negatives may be likely. Please keep that in mind when using this tool.

## Compile

Compiling SolAnalyzer requires Go and GoDep. Once those dependencies are satisfied, `make` can be run to compile the source.

## Run

Running the tool looks like this:

```
$ solanalyzer --help
SolAnalyzer is a static analyzer for the Solidity programming language, with a focus on finding security bugs.

Usage:
  solanalyzer path [flags]

Flags:
  -h, --help      help for solanalyzer
  -v, --verbose   verbose output

$ solanalyzer examples/

SolAnalyzer v0.1-beta
Copyright 2018 Eric Rafaloff

This is beta software. Please report issues at https://github.com/EricR/solanalyzer/issues/.

[2018-09-01T09:28:10-04:00]  INFO Starting new session
[2018-09-01T09:28:10-04:00]  INFO Parsing sources
[2018-09-01T09:28:10-04:00]  INFO Scanning sources
[2018-09-01T09:28:10-04:00]  INFO Analyzing sources
[2018-09-01T09:28:10-04:00]  INFO Generating report

=== Start SolAnalyzer Report ===

Report Date:   Sat Sep  1 9:28 AM 2018
Analyzers Run: compiler-version, function-visibility

High Severity Issues
--------------------
No issues

Medium Severity Issues
----------------------
No issues

Low Severity Issues
-------------------
Title:       Compiler Bug - EventStructWrongData
Description: The version pragma, >0.4.22, can be satisfied by a version of the Solidity compiler that contains a known bug. If a struct is used in an event, the address of the struct is logged instead of the actual data. This bug is reported to be fixed in version 0.5.0.
Source:      examples/reentrancy.sol:1:23
Analyzer ID: compiler-version
Instance ID: c4afc52c128cbd79b10ffeee91b937beeac479ecb864752a8dd0dcf787bdebb8

Informational Severity Issues
-----------------------------
Title:       Default Function Visibility
Description: No visibility is specified for function transfer(address to, uint amount) in contract Vulnerable. The default is public. It should be confirmed that this is desired, and the visibility of the function should be explicitly set.
Source:      examples/reentrancy.sol:6:1
Analyzer ID: function-visibility
Instance ID: 747dbaef6def6d744a1898c482f1325ec704d56ceb1f1fa1b099d7a03008f8e1

Title:       Default Function Visibility
Description: No visibility is specified for function withdraw() in contract Vulnerable. The default is public. It should be confirmed that this is desired, and the visibility of the function should be explicitly set.
Source:      examples/reentrancy.sol:13:1
Analyzer ID: function-visibility
Instance ID: 6d6158a6074c6d05639d365b861615cb5648736afd27ae0a78b6633844c2f317

=== End SolAnalyzer Report ===
```

## Supported Checks

See [here](https://github.com/EricR/solanalyzer/wiki/Supported-Checks) for a list of issues SolAnalyzer is capable of checking for.