# OpenAI Build Week Evidence

## Existing foundation

The Edge-Computing-LLM organization contained Linux, Kubernetes, NVIDIA,
local-inference, and observability work before the final competition submission.

This submission does not claim that every historical line of code was created
during Build Week.

## Work completed during Build Week

### edge-cli

- Added automatic CPU and NVIDIA orchestration.
- Added Qwen validation for the GeForce 940M profile.
- Added Go validation workflows.
- Added stronger infrastructure-health validation.
- Documented the generic GGUF evidence plane.

### edge-llm-tests

- Added a Go-first cross-repository validation harness.
- Recorded Ubuntu, k3s, and NVIDIA validation evidence.
- Improved GitHub Actions.
- Added open-source secret scanning.
- Added generic GGUF observer validation.

### gguf-observability

- Added read-only Qwen runtime observability.
- Replaced earlier runtime tooling with Go.
- Kept inference evidence prompt-free.
- Generalized GGUF runtime evidence.

### k3s-nvidia-edge

- Added conditional NVIDIA validation.
- Added a Go validation workflow.
- Added stale k3s-network detection.
- Updated the local project inventory.

### llm-observability-stack

- Added CPU fallback support.
- Added a low-VRAM Qwen 1.8B profile.
- Added Go-native runtime tooling.
- Added k3s network-health gates.
- Added low-VRAM multi-model profiles.

### Competition integration

- Created the EdgeProof Go CLI.
- Created the repository-level Codex Agent Skill.
- Created the OpenAI Build Week proof workflow.
- Created Devpost and demonstration material.

## Reference environment

- Ubuntu 24.04.3
- Linux kernel 6.17
- Go 1.26.5
- k3s 1.36.2
- Helm 4.2.3
- NVIDIA driver 580.95.05
- NVIDIA GeForce 940M
- 1 GiB VRAM
- One Kubernetes-allocatable NVIDIA GPU

## Verified live platform

The recorded EdgeProof result contains:

- 13 passed checks
- 0 warnings
- 0 failed checks
- final verdict `VERIFIED`

The verified platform includes a Ready k3s node, NVIDIA RuntimeClass, one
allocatable GPU, GPU Operator components, Ollama, Open WebUI, OpenTelemetry
Collector, Prometheus, and Grafana.
