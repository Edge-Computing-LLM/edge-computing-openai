# Known Limitations

## Read-only initial release

Version 0.1.0 performs read-only inspection and verification. Automated repair
is not claimed as complete.

## Command-success limitation

The current CLI confirms that selected commands complete successfully. It does
not yet parse and evaluate every field inside every Kubernetes resource.

## Legacy GPU constraints

The reference NVIDIA GeForce 940M has only 1 GiB VRAM and does not support MIG.
MIG-related components with zero desired replicas are not applicable.

## Historical Kubernetes warnings

Old startup or readiness warnings do not necessarily indicate a current
failure. Current resource state should be evaluated separately.

## OpenAI API

The demonstrated local runtime does not require an OpenAI API key. GPT-5.6 was
used through Codex CLI during development.
