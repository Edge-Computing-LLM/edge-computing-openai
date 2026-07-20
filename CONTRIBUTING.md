# Contributing

## Development requirements

- Go 1.26 or a compatible current Go release
- Git
- Optional live environment: Linux, k3s, Helm, and NVIDIA tooling

## Local checks

```bash
make check
```

## Live proof

```bash
make prove
```

Live checks require access to the expected local k3s environment.

## Contribution rules

- Keep the default workflow read-only.
- Add tests for behavior changes.
- Do not add credentials or generated reports.
- Do not commit model binaries.
- Document limitations honestly.
- Keep OpenAI-hosted and local-model responsibilities clearly separated.
