#!/usr/bin/env bash
# Validation gate for the devops-k8s-operator repo.
# Used on every change before commit: build, vet, and schema-validate manifests.
set -euo pipefail

cd "$(dirname "$0")/.."

echo "[1/3] go build ./..."
go build ./...

echo "[2/3] go vet ./..."
go vet ./...

echo "[3/3] kubeconform schema validation"
kubeconform -strict -summary \
  config/crd/ping.yaml \
  config/rbac/role.yaml \
  config/manager/deployment.yaml \
  config/manager/service.yaml \
  config/samples/ping-sample.yaml

echo "OK: all validations passed"
