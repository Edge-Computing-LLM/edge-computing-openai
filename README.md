# EdgeProof

**Codex-operated verification for private LLM infrastructure on Linux, k3s, and constrained NVIDIA GPUs.**

EdgeProof is the OpenAI Build Week proof and verification repository for the
[Edge-Computing-LLM](https://github.com/Edge-Computing-LLM) platform.

It provides:

- a read-only Go verification CLI,
- a repository-level Codex Agent Skill,
- sanitized evidence,
- Build Week documentation,
- and competition-ready demonstration material.

## Why EdgeProof exists

A local LLM platform is more than a model server. A working deployment depends
on Linux, Kubernetes, GPU runtime support, scheduling, model availability,
networking, telemetry, and operational verification.

A command can succeed while the underlying platform is unhealthy. EdgeProof
creates a clear, reproducible verification entry point and documents the
boundaries of what is proven.

## What is verified

The `edge-openai` CLI checks selected infrastructure contracts:

- required host commands,
- NVIDIA GPU visibility,
- Kubernetes API access,
- k3s node access,
- NVIDIA RuntimeClass,
- allocatable Kubernetes GPU resources,
- Helm releases,
- GPU infrastructure namespace access,
- observability namespace access,
- and selected deployment, StatefulSet, and DaemonSet listings.

The current version verifies that these commands and contracts are reachable.
It does not yet semantically parse every Kubernetes object or automatically
repair failures.

## Verified reference result

The reference environment produced:

```text
Passed:   13
Warnings: 0
Failed:   0
Verdict:  VERIFIED
```

Reference hardware and software:

- Lenovo ThinkPad T450s
- Intel Core i7-5600U
- 12 GiB RAM
- NVIDIA GeForce 940M
- 1 GiB VRAM
- Ubuntu 24.04
- k3s 1.36
- NVIDIA container runtime
- Ollama
- Open WebUI
- OpenTelemetry Collector
- Prometheus
- Grafana

This is a demonstrated environment, not a minimum requirement.

## Quick start

Clone:

```bash
git clone https://github.com/Edge-Computing-LLM/edge-computing-openai.git
cd edge-computing-openai
```

Run local validation:

```bash
make check
```

Run live infrastructure verification:

```bash
make prove
```

Or build and run directly:

```bash
go build -o bin/edge-openai ./cmd/edge-openai

./bin/edge-openai version
./bin/edge-openai inspect
./bin/edge-openai prove
./bin/edge-openai prove --format json
./bin/edge-openai prove --verbose
```

`make prove` requires access to the expected local Linux, k3s, Helm, and NVIDIA
environment. GitHub Actions runs unit tests, formatting, vetting, race tests, and
build validation, but does not run the live cluster proof.

## Commands

### `version`

```bash
./bin/edge-openai version
```

Prints the CLI version.

### `inspect`

```bash
./bin/edge-openai inspect
```

Checks required commands and selected host and Kubernetes contracts.

### `prove`

```bash
./bin/edge-openai prove
```

Extends inspection with GPU and observability namespace checks.

Output formats:

```bash
./bin/edge-openai prove --format human
./bin/edge-openai prove --format json
```

Bounded command output:

```bash
./bin/edge-openai prove --verbose
```

Custom timeout:

```bash
./bin/edge-openai prove --timeout 30s
```

## Platform relationship

EdgeProof does not replace the platform repositories.

| Repository | Responsibility |
|---|---|
| [`edge-cli`](https://github.com/Edge-Computing-LLM/edge-cli) | Unified installation and operational CLI |
| [`k3s-nvidia-edge`](https://github.com/Edge-Computing-LLM/k3s-nvidia-edge) | Linux, k3s, and NVIDIA infrastructure |
| [`llm-observability-stack`](https://github.com/Edge-Computing-LLM/llm-observability-stack) | Local inference and observability workloads |
| [`gguf-observability`](https://github.com/Edge-Computing-LLM/gguf-observability) | Prompt-free GGUF runtime evidence |
| [`edge-llm-tests`](https://github.com/Edge-Computing-LLM/edge-llm-tests) | Cross-repository validation |
| [`Edge-Computing-LLM-Documentation`](https://github.com/Edge-Computing-LLM/Edge-Computing-LLM-Documentation) | End-user documentation |
| **edge-computing-openai** | Codex Skill, proof CLI, evidence, and Build Week material |

## OpenAI and local-model boundary

- GPT-5.6 through Codex CLI was used to design, generate, refactor, test, and
  document the platform.
- A local GGUF model performs local inference.
- The local model is not presented as GPT-5.6.
- The local runtime does not require an OpenAI API key.
- EdgeProof does not send local prompts or model responses to OpenAI.

## Codex Agent Skill

The repository includes:

```text
.agents/skills/edge-llm-operator/
```

The skill instructs Codex to:

- start with read-only inspection,
- use existing platform tools,
- treat successful `Completed` jobs correctly,
- avoid reading or displaying secrets,
- avoid storing prompts and model responses,
- distinguish historical warnings from current health,
- and request approval before a mutating repair.

## Privacy and evidence policy

EdgeProof must not intentionally collect or publish:

- prompt text,
- model response text,
- API keys,
- GitHub tokens,
- Kubernetes Secrets,
- kubeconfig credentials,
- private keys,
- browser session data,
- model weights,
- or unrestricted raw logs.

Generated live reports are ignored by Git. A sanitized example is available at:

```text
evidence/examples/verified-reference.json
```

## Build Week work

Substantial Build Week work across the organization includes:

- automatic CPU and NVIDIA orchestration,
- Go-first validation,
- low-VRAM model profiles,
- prompt-free GGUF evidence,
- k3s networking checks,
- observability validation,
- repository documentation,
- the EdgeProof Go CLI,
- and the Codex Agent Skill.

The submission does not claim that every historical line of organization code
was created during Build Week.

See [Build Week Evidence](docs/BUILD-WEEK-EVIDENCE.md).

## Development

```bash
make check
```

Equivalent checks:

```bash
gofmt -w cmd internal
go test ./...
go test -race ./...
go vet ./...
go build ./...
```

Run live proof:

```bash
make prove
```

## Security

Version `0.1.0` is read-only.

Future mutating functionality must use:

- explicit approval,
- allow-listed actions,
- bounded execution,
- clear rollback instructions,
- and post-change verification.

See [Security Policy](SECURITY.md).

## Documentation

- [Architecture](docs/ARCHITECTURE.md)
- [Build Week Evidence](docs/BUILD-WEEK-EVIDENCE.md)
- [Demo Guide](docs/DEMO-GUIDE.md)
- [Known Limitations](docs/KNOWN-LIMITATIONS.md)
- [Privacy](docs/PRIVACY.md)
- [Devpost Description](competition/devpost-description.md)
- [Video Script](competition/video-script.md)
- [Submission Checklist](competition/submission-checklist.md)
- [Contributing](CONTRIBUTING.md)
- [Support](SUPPORT.md)
- [Changelog](CHANGELOG.md)

## License

MIT License. See [LICENSE](LICENSE).
