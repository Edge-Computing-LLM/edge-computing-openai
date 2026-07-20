# Security Policy

## Supported version

The current supported development version is `0.1.x`.

## Reporting a vulnerability

Do not open a public issue containing credentials, private cluster output, or
an exploitable security detail.

Report security concerns privately to the repository owner through GitHub.

## Sensitive data

Never submit:

- API keys,
- GitHub tokens,
- Kubernetes Secrets,
- kubeconfig files,
- private keys,
- model prompts,
- model responses,
- model weights,
- or private infrastructure addresses that are not necessary to reproduce an
  issue.

## Operational safety

EdgeProof version 0.1.0 is read-only. Future mutating features must use
explicit approval, allow-listed actions, timeouts, and rollback instructions.
