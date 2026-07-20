package report

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/Edge-Computing-LLM/edge-computing-openai/internal/checks"
)

func TestNewCreatesVerifiedReport(t *testing.T) {
	results := []checks.Result{
		{
			Name:    "first",
			Status:  checks.StatusPass,
			Summary: "passed",
		},
		{
			Name:    "second",
			Status:  checks.StatusWarning,
			Summary: "warning",
		},
	}

	generated := New(results)

	if generated.Passed != 1 {
		t.Fatalf("expected one pass, got %d", generated.Passed)
	}

	if generated.Warnings != 1 {
		t.Fatalf("expected one warning, got %d", generated.Warnings)
	}

	if generated.Failed != 0 {
		t.Fatalf("expected zero failures, got %d", generated.Failed)
	}

	if generated.Verdict != "VERIFIED" {
		t.Fatalf("unexpected verdict: %s", generated.Verdict)
	}
}

func TestNewCreatesNotVerifiedReport(t *testing.T) {
	results := []checks.Result{
		{
			Name:    "failed-check",
			Status:  checks.StatusFail,
			Summary: "failed",
		},
	}

	generated := New(results)

	if generated.Failed != 1 {
		t.Fatalf("expected one failure, got %d", generated.Failed)
	}

	if generated.Verdict != "NOT VERIFIED" {
		t.Fatalf("unexpected verdict: %s", generated.Verdict)
	}
}

func TestWriteHumanIncludesSummary(t *testing.T) {
	generated := New([]checks.Result{
		{
			Name:    "example",
			Status:  checks.StatusPass,
			Summary: "worked",
		},
	})

	var output bytes.Buffer
	WriteHuman(&output, generated, false)

	text := output.String()

	for _, expected := range []string{
		"EdgeProof verification",
		"[pass] example",
		"Passed:",
		"Verdict:",
		"VERIFIED",
	} {
		if !strings.Contains(text, expected) {
			t.Fatalf("expected output to contain %q:\n%s", expected, text)
		}
	}
}

func TestWriteJSONProducesValidJSON(t *testing.T) {
	generated := New([]checks.Result{
		{
			Name:    "example",
			Status:  checks.StatusPass,
			Summary: "worked",
		},
	})

	var output bytes.Buffer

	if err := WriteJSON(&output, generated); err != nil {
		t.Fatalf("WriteJSON returned error: %v", err)
	}

	var decoded Report

	if err := json.Unmarshal(output.Bytes(), &decoded); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if decoded.Project != "EdgeProof" {
		t.Fatalf("unexpected project: %s", decoded.Project)
	}

	if decoded.Verdict != "VERIFIED" {
		t.Fatalf("unexpected verdict: %s", decoded.Verdict)
	}
}
