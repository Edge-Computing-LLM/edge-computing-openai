package checks

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Status string

const (
	StatusPass    Status = "pass"
	StatusFail    Status = "fail"
	StatusWarning Status = "warning"
	StatusSkipped Status = "skipped"
)

type Result struct {
	Name    string `json:"name"`
	Status  Status `json:"status"`
	Summary string `json:"summary"`
	Output  string `json:"output,omitempty"`
}

type Runner struct {
	Timeout time.Duration
}

func NewRunner(timeout time.Duration) Runner {
	if timeout <= 0 {
		timeout = 20 * time.Second
	}

	return Runner{Timeout: timeout}
}

func (r Runner) CommandAvailable(name string) Result {
	path, err := exec.LookPath(name)
	if err != nil {
		return Result{
			Name:    "command:" + name,
			Status:  StatusFail,
			Summary: fmt.Sprintf("%s is not installed or is not in PATH", name),
		}
	}

	return Result{
		Name:    "command:" + name,
		Status:  StatusPass,
		Summary: fmt.Sprintf("%s is available at %s", name, path),
	}
}

func (r Runner) Run(name string, args ...string) Result {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	command := exec.CommandContext(ctx, name, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()

	output := strings.TrimSpace(stdout.String())
	errorOutput := strings.TrimSpace(stderr.String())

	if ctx.Err() == context.DeadlineExceeded {
		return Result{
			Name:    strings.Join(append([]string{name}, args...), " "),
			Status:  StatusFail,
			Summary: "command timed out",
			Output:  truncate(output+"\n"+errorOutput, 4000),
		}
	}

	if err != nil {
		return Result{
			Name:    strings.Join(append([]string{name}, args...), " "),
			Status:  StatusFail,
			Summary: err.Error(),
			Output:  truncate(output+"\n"+errorOutput, 4000),
		}
	}

	return Result{
		Name:    strings.Join(append([]string{name}, args...), " "),
		Status:  StatusPass,
		Summary: "command completed successfully",
		Output:  truncate(output, 4000),
	}
}

func truncate(value string, maximum int) string {
	value = strings.TrimSpace(value)

	if len(value) <= maximum {
		return value
	}

	return value[:maximum] + "\n... output truncated ..."
}
