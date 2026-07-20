---
name: edge-llm-operator
description: Inspect and validate the Edge-Computing-LLM Linux, k3s, NVIDIA GPU and local LLM platform. Use this skill for read-only health checks, privacy-safe evidence and approval-based repair planning.
---

# Edge LLM Operator

Use this skill when working with the Edge-Computing-LLM platform.

## Safety rules

1. Begin with read-only inspection.
2. Never print Kubernetes Secrets.
3. Never read or display API keys, GitHub tokens, private keys or
   kubeconfig certificate data.
4. Never retain prompt or model-response text in evidence.
5. Never remove model files or persistent storage.
6. Never uninstall k3s or NVIDIA drivers.
7. Never perform a mutating repair without showing the exact plan
   and receiving explicit approval.
8. Prefer existing Edge-Computing-LLM commands over reimplementing
   functionality.

## Repository responsibilities

- `edge-cli`: platform installation and orchestration
- `k3s-nvidia-edge`: k3s and NVIDIA infrastructure
- `llm-observability-stack`: local inference and observability
- `gguf-observability`: prompt-free GGUF runtime evidence
- `edge-llm-tests`: cross-repository validation
- `edge-computing-openai`: Codex integration and competition proof

## Standard workflow

1. Inspect the Linux host.
2. Inspect the repository workspace.
3. Inspect the k3s node.
4. Check NVIDIA RuntimeClass and allocatable GPU resources.
5. Check NVIDIA device-plugin and DCGM workloads.
6. Check Ollama, Open WebUI, OpenTelemetry, Prometheus and Grafana.
7. Run existing validation commands.
8. Classify each result as pass, warning, fail, skipped or not
   applicable.
9. Generate a proposed repair plan for real failures.
10. Request approval before any mutation.
11. Validate recovery.
12. Generate privacy-safe evidence.

## Kubernetes classification rules

- A pod with phase `Succeeded` or status `Completed` is successful.
- A completed NVIDIA CUDA validator pod is not unhealthy.
- MIG and MPS daemonsets with desired count zero may be not
  applicable on unsupported hardware.
- Historical warning events do not prove that a currently Ready
  workload is unhealthy.
- Current workload state has priority over old startup warnings.

## Evidence restrictions

Do not store:

- Prompt content
- Model-response content
- Secrets
- Authentication headers
- Tokens
- Private keys
- Kubeconfig credentials
- Model weights

Record only safe operational status, timings, counts and sanitized
error categories.
