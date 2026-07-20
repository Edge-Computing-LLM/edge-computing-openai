package report

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/Edge-Computing-LLM/edge-computing-openai/internal/checks"
)

type Report struct {
	Project     string          `json:"project"`
	GeneratedAt time.Time       `json:"generated_at"`
	Results     []checks.Result `json:"results"`
	Passed      int             `json:"passed"`
	Failed      int             `json:"failed"`
	Warnings    int             `json:"warnings"`
	Verdict     string          `json:"verdict"`
}

func New(results []checks.Result) Report {
	report := Report{
		Project:     "EdgeProof",
		GeneratedAt: time.Now(),
		Results:     results,
		Verdict:     "VERIFIED",
	}

	for _, result := range results {
		switch result.Status {
		case checks.StatusPass:
			report.Passed++
		case checks.StatusWarning:
			report.Warnings++
		case checks.StatusFail:
			report.Failed++
			report.Verdict = "NOT VERIFIED"
		}
	}

	return report
}

func WriteHuman(writer io.Writer, report Report, includeOutput bool) {
	fmt.Fprintln(writer, "============================================================")
	fmt.Fprintln(writer, "EdgeProof verification")
	fmt.Fprintln(writer, "============================================================")

	for _, result := range report.Results {
		fmt.Fprintf(
			writer,
			"[%s] %s — %s\n",
			result.Status,
			result.Name,
			result.Summary,
		)

		if includeOutput && result.Output != "" {
			fmt.Fprintln(writer, result.Output)
			fmt.Fprintln(writer)
		}
	}

	fmt.Fprintln(writer, "============================================================")
	fmt.Fprintf(writer, "Passed:   %d\n", report.Passed)
	fmt.Fprintf(writer, "Warnings: %d\n", report.Warnings)
	fmt.Fprintf(writer, "Failed:   %d\n", report.Failed)
	fmt.Fprintf(writer, "Verdict:  %s\n", report.Verdict)
}

func WriteJSON(writer io.Writer, report Report) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	return encoder.Encode(report)
}
