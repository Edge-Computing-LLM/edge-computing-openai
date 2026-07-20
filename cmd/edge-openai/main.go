package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Edge-Computing-LLM/edge-computing-openai/internal/checks"
	"github.com/Edge-Computing-LLM/edge-computing-openai/internal/report"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "version":
		fmt.Printf("edge-openai %s\n", version)
	case "inspect":
		os.Exit(runInspect(os.Args[2:], false))
	case "prove":
		os.Exit(runInspect(os.Args[2:], true))
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(2)
	}
}

func printUsage() {
	fmt.Println(`EdgeProof — Codex-operated private LLM infrastructure verification

Usage:
  edge-openai version
  edge-openai inspect [--format human|json] [--verbose]
  edge-openai prove   [--format human|json] [--verbose]

Commands:
  version   Print the CLI version.
  inspect   Inspect required tools and the Kubernetes environment.
  prove     Run the competition verification workflow.

The CLI performs read-only checks. It does not print Kubernetes
Secrets, prompts, model responses, API keys or model weights.`)
}

func runInspect(arguments []string, prove bool) int {
	commandName := "inspect"
	if prove {
		commandName = "prove"
	}

	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	outputFormat := flags.String("format", "human", "output format: human or json")
	verbose := flags.Bool("verbose", false, "include bounded command output")
	timeout := flags.Duration("timeout", 20*time.Second, "timeout per command")

	if err := flags.Parse(arguments); err != nil {
		return 2
	}

	runner := checks.NewRunner(*timeout)
	results := collectResults(runner, prove)
	generatedReport := report.New(results)

	switch strings.ToLower(*outputFormat) {
	case "human":
		report.WriteHuman(os.Stdout, generatedReport, *verbose)
	case "json":
		if err := report.WriteJSON(os.Stdout, generatedReport); err != nil {
			fmt.Fprintf(os.Stderr, "write JSON report: %v\n", err)
			return 1
		}
	default:
		fmt.Fprintf(os.Stderr, "unsupported format: %s\n", *outputFormat)
		return 2
	}

	if generatedReport.Failed > 0 {
		return 1
	}

	return 0
}

func collectResults(runner checks.Runner, prove bool) []checks.Result {
	results := make([]checks.Result, 0, 20)

	for _, command := range []string{
		"git",
		"go",
		"kubectl",
		"helm",
		"nvidia-smi",
	} {
		results = append(results, runner.CommandAvailable(command))
	}

	results = append(results,
		runner.Run(
			"nvidia-smi",
			"--query-gpu=name,memory.total,driver_version",
			"--format=csv,noheader",
		),
		runner.Run("kubectl", "get", "nodes", "-o", "wide"),
		runner.Run("kubectl", "get", "runtimeclass", "nvidia"),
		runner.Run(
			"kubectl",
			"get",
			"nodes",
			"-o",
			"custom-columns=NAME:.metadata.name,GPU:.status.allocatable.nvidia\\.com/gpu",
		),
		runner.Run("helm", "list", "-A"),
	)

	if prove {
		results = append(results,
			runner.Run(
				"kubectl",
				"get",
				"pods",
				"-n",
				"gpu-operator",
			),
			runner.Run(
				"kubectl",
				"get",
				"pods",
				"-n",
				"llm-observability",
			),
			runner.Run(
				"kubectl",
				"get",
				"deployments,statefulsets,daemonsets",
				"-n",
				"llm-observability",
			),
		)
	}

	return results
}
