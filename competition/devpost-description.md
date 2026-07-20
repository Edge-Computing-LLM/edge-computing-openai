# Devpost Description

## Project name

EdgeProof

## Tagline

Codex-operated verification for private LLM infrastructure on Linux, k3s, and
constrained NVIDIA GPUs.

## Inspiration

Private LLM infrastructure is difficult to reproduce and often assumes modern
GPUs or paid cloud services. I wanted to prove that one developer, using GPT-5.6
through Codex CLI, could build a real local AI platform on an older ThinkPad
with a 1 GiB NVIDIA GPU.

## What it does

EdgeProof verifies Linux tooling, k3s, NVIDIA runtime and GPU scheduling,
Ollama, Open WebUI, OpenTelemetry, Prometheus, Grafana, and the organization
repositories that implement the platform. It produces a human-readable or JSON
verdict without requiring an OpenAI API key for local runtime.

## How it was built

The implementation is Go-first. GPT-5.6 through Codex CLI was used to design,
generate, refactor, test, and document the platform.

## Challenges

The largest challenge was running a useful local model and observability stack
on a legacy GeForce 940M with only 1 GiB VRAM while keeping Kubernetes and GPU
validation reliable.

## Accomplishments

- Working k3s and NVIDIA deployment on legacy hardware.
- Local GGUF inference.
- OpenTelemetry, Prometheus, Grafana, and DCGM observability.
- Go-first orchestration and validation.
- Prompt-free and privacy-conscious proof workflow.
- Reusable Codex Agent Skill.

## What is next

The next stage is approval-based repair planning, richer evidence bundles, and
optional provider adapters while preserving an API-key-free local runtime.
