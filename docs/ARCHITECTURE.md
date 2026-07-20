# Architecture

EdgeProof is the competition-facing verification layer for the
Edge-Computing-LLM organization.

```text
Developer or Codex
        |
        v
edge-computing-openai
        |
        +--> nvidia-smi
        +--> kubectl
        +--> Helm
        +--> edge-cli
        +--> edge-llm-tests
        +--> gguf-observability
        |
        v
k3s-nvidia-edge
        |
        v
llm-observability-stack
        |
        +--> Ollama
        +--> Open WebUI
        +--> OpenTelemetry
        +--> Prometheus
        +--> Grafana
```

The initial Go CLI uses bounded, read-only command execution and returns human
or JSON verification results.
