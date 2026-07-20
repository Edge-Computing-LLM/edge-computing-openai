# Privacy

EdgeProof is designed to verify infrastructure without collecting application
content.

The project must not intentionally record prompt text, model response text, API
keys, GitHub tokens, Kubernetes Secrets, kubeconfig credentials, private keys,
browser session data, model files, or unrestricted raw logs.

The current Go CLI uses read-only commands and limits retained command output.
