# SolAnalyzer

SolAnalyzer is a static analyzer for the Solidity programming language, with a focus on finding security bugs.

**Warning**: This software is in beta and false negatives may be likely. Please keep that in mind when using this tool.

## Compile

Compiling solanalyzer requires Go and GoDep. Once those dependencies are satisfied, `make` can be run to build the project.

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

[2018-08-26T19:00:32-04:00]  INFO Starting new session
[2018-08-26T19:00:32-04:00]  INFO Parsing sources
[2018-08-26T19:00:32-04:00]  INFO Scanning sources
[2018-08-26T19:00:32-04:00]  INFO Running analysis
[2018-08-26T19:00:32-04:00]  INFO Generating report

=== Start SolAnalyzer Report ===

Report Date:   Sun Aug 26 19:00 2018
Analyzers Run: Solidity Compiler Version, Function Visibility

High Severity Issues
--------------------
No issues

Medium Severity Issues
----------------------
No issues

Low Severity Issues
-------------------
Title:       Compiler Bug - EventStructWrongData
Description: If a struct is used in an event, the address of the struct is logged instead of the actual data.
Source:      examples/reentrancy.sol:1:16
Analyzer ID: compiler-version
Instance ID: 9634c1c4d4c66a71eb8dff6b1bbd2bb311304e9985bd2a6714844c41403002b3

Informational Severity Issues
-----------------------------
Title:       Outdated Solidity Compiler
Description: The version constraint, =0.4.24, can only be satisfied by an outdated version of the Solidity compiler. The latest version is 0.4.25. The source's pragma declaration should be updated accordingly.
Source:      examples/reentrancy.sol:1:16
Analyzer ID: compiler-version
Instance ID: 06e02d5a72cd15b55d759811ccaad7747fc6b1dc1d68e9dd70650e6389483157

Title:       Public Function
Description: The function transfer(address,uint) in the contract Vulnerable was found to be public. It should be confirmed that this function is intended to be publicly callable.
Source:      examples/reentrancy.sol:6:1
Analyzer ID: function-visibility
Instance ID: 50a1be44962e05fd00c1b2bfc651d545375ea6c627ea945af49d0ea87e6d8d63

Title:       Public Function
Description: The function withdraw() in the contract Vulnerable was found to be public. It should be confirmed that this function is intended to be publicly callable.
Source:      examples/reentrancy.sol:13:1
Analyzer ID: function-visibility
Instance ID: fa50ccb652b67143f96251666641a2bcd509d493c04e1b30913cef3bb2ed3d99

=== End SolAnalyzer Report ===
```

## Supported Checks

See [here](https://github.com/EricR/solanalyzer/wiki/Supported-Checks) for a list of issues SolAnalyzer is capable of checking for.