# Safety Policy

The operator is read-only by default.

It must never automatically:

- Print Kubernetes Secrets
- Print authentication tokens
- Delete persistent volumes
- Remove model files
- Uninstall k3s
- Remove NVIDIA drivers
- Reset the cluster
- Delete arbitrary namespaces
- Send local prompts or responses to third parties
- Apply mutating repairs without explicit approval
