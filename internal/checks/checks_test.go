package checks

import (
	"strings"
	"testing"
	"time"
)

func TestNewRunnerUsesProvidedTimeout(t *testing.T) {
	runner := NewRunner(7 * time.Second)

	if runner.Timeout != 7*time.Second {
		t.Fatalf("expected timeout %s, got %s", 7*time.Second, runner.Timeout)
	}
}

func TestNewRunnerUsesDefaultTimeout(t *testing.T) {
	runner := NewRunner(0)

	if runner.Timeout != 20*time.Second {
		t.Fatalf("expected default timeout %s, got %s", 20*time.Second, runner.Timeout)
	}
}

func TestCommandAvailableFindsKnownCommand(t *testing.T) {
	runner := NewRunner(time.Second)
	result := runner.CommandAvailable("go")

	if result.Status != StatusPass {
		t.Fatalf("expected pass, got %s: %s", result.Status, result.Summary)
	}

	if !strings.Contains(result.Name, "command:go") {
		t.Fatalf("unexpected result name: %s", result.Name)
	}
}

func TestCommandAvailableRejectsMissingCommand(t *testing.T) {
	runner := NewRunner(time.Second)
	result := runner.CommandAvailable(
		"edgeproof-command-that-should-not-exist-123456",
	)

	if result.Status != StatusFail {
		t.Fatalf("expected fail, got %s", result.Status)
	}
}

func TestRunCapturesSuccessfulOutput(t *testing.T) {
	runner := NewRunner(2 * time.Second)
	result := runner.Run("printf", "hello-edgeproof")

	if result.Status != StatusPass {
		t.Fatalf("expected pass, got %s: %s", result.Status, result.Summary)
	}

	if result.Output != "hello-edgeproof" {
		t.Fatalf("unexpected output: %q", result.Output)
	}
}

func TestRunReturnsFailureForMissingCommand(t *testing.T) {
	runner := NewRunner(time.Second)
	result := runner.Run(
		"edgeproof-command-that-should-not-exist-654321",
	)

	if result.Status != StatusFail {
		t.Fatalf("expected fail, got %s", result.Status)
	}
}

func TestRunTimesOut(t *testing.T) {
	runner := NewRunner(20 * time.Millisecond)
	result := runner.Run("sleep", "1")

	if result.Status != StatusFail {
		t.Fatalf("expected fail, got %s", result.Status)
	}

	if result.Summary != "command timed out" {
		t.Fatalf("unexpected summary: %q", result.Summary)
	}
}

func TestTruncateLeavesShortTextUnchanged(t *testing.T) {
	input := "short text"

	if output := truncate(input, 100); output != input {
		t.Fatalf("expected %q, got %q", input, output)
	}
}

func TestTruncateLimitsLongText(t *testing.T) {
	input := strings.Repeat("x", 30)
	output := truncate(input, 10)

	if !strings.HasPrefix(output, strings.Repeat("x", 10)) {
		t.Fatalf("unexpected truncated output: %q", output)
	}

	if !strings.Contains(output, "output truncated") {
		t.Fatalf("missing truncation marker: %q", output)
	}
}
