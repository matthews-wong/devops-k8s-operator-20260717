IMG ?= ghcr.io/matthews-wong/devops-k8s-operator:latest

.PHONY: build
build: ## Build the manager binary.
	go build -o bin/manager main.go

.PHONY: fmt
fmt: ## Run go fmt.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet.
	go vet ./...

.PHONY: validate
validate: ## Build, vet, and schema-validate manifests.
	go build ./...
	go vet ./...
	kubeconform -strict -summary \
		config/crd/ping.yaml \
		config/rbac/role.yaml \
		config/manager/deployment.yaml \
		config/samples/ping-sample.yaml

.PHONY: docker-build
docker-build: ## Build the container image.
	docker build -t $(IMG) .

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
