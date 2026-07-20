# Demo Guide

## Build and validate

```bash
make check
make prove
```

Expected final output:

```text
Passed:   13
Warnings: 0
Failed:   0
Verdict:  VERIFIED
```

## Suggested recording sequence

```bash
nvidia-smi
kubectl get nodes -o wide
kubectl get runtimeclass nvidia
kubectl get pods -n gpu-operator
kubectl get pods -n llm-observability
helm list -A
make prove
```

Do not display credentials, Kubernetes Secrets, private prompts, or private
model responses in the recording.
