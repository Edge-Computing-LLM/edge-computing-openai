# Three-Minute OpenAI Build Week Video

## 0:00–0:20 — The problem

Private LLM infrastructure normally assumes modern GPUs and cloud budgets. I
built this platform on a ten-year-old ThinkPad with a one-gigabyte NVIDIA
GeForce 940M.

## 0:20–0:40 — How Codex helped

I used GPT-5.6 through Codex CLI to build and refactor a Go-first platform for
Linux, k3s, NVIDIA GPU scheduling, local inference, observability, and
validation.

## 0:40–1:05 — Hardware and Kubernetes

Show `nvidia-smi`, the Ready k3s node, NVIDIA RuntimeClass, and one allocatable
GPU. Explain that the local runtime does not require an OpenAI API key.

## 1:05–1:35 — Working platform

Show the NVIDIA device plugin, DCGM exporter, Ollama, Open WebUI, OpenTelemetry
Collector, Prometheus, and Grafana.

## 1:35–2:00 — Local inference

Show one short local inference request. Explain that the local GGUF model is not
GPT-5.6; GPT-5.6 through Codex was used to build the infrastructure.

## 2:00–2:25 — Observability

Show one clear Grafana or Prometheus view containing GPU, Kubernetes, or
inference telemetry.

## 2:25–2:45 — Proof

Run `make prove` and show 13 passed checks, 0 failures, and `VERIFIED`.

## 2:45–3:00 — Impact

EdgeProof shows that Codex can help one independent developer build and verify
serious private AI infrastructure using hardware they already own.
