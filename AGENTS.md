# AGENTS.md

## Project purpose

EdgeProof is the OpenAI Build Week verification and presentation layer for the
Edge-Computing-LLM organization.

## Language policy

Use Go for application logic, command execution, reporting, validation, and
tests.

Python helper files are local project-construction utilities and are not part
of the product runtime.

## Safety requirements

- Use read-only inspection by default.
- Never print Kubernetes Secrets.
- Never read or expose API keys, GitHub tokens, private keys, or kubeconfig
  credential data.
- Never store prompt text or model response text in evidence.
- Never commit GGUF models or generated runtime reports.
- Do not perform cluster mutations without explicit user approval.
- Do not claim semantic health when only command execution was checked.

## Required checks

Before committing Go changes, run:

```bash
make check
make prove
```
