# devops-k8s-operator

A minimal Kubernetes operator scaffold.

## Purpose

This operator introduces a custom resource `Ping` (`operators.example.com/v1`).
A `Ping` instructs the controller to emit a periodic "ping" log message at a
configurable interval and record the last ping time plus a running count in the
resource's status. It is intentionally tiny — a working, deployable skeleton you
can extend over time (webhooks, tests, metrics, multi-resource, etc.).

## Layout

```
.
├── api/v1/                 # Ping CRD Go types + group registration
├── controllers/            # PingReconciler (the control loop)
├── config/
│   ├── crd/ping.yaml       # CustomResourceDefinition manifest
│   ├── rbac/role.yaml       # ServiceAccount + ClusterRole/Binding
│   ├── manager/deployment.yaml  # controller-manager Deployment + Namespace
│   └── samples/ping-sample.yaml  # example Ping CR
├── main.go                 # manager entrypoint
├── Dockerfile              # distroless build
├── Makefile                # build / validate targets
└── scripts/validate.sh     # validation gate (build + vet + kubeconform)
```

## Prerequisites

- Go 1.23+
- [kubeconform](https://github.com/yannh/kubeconform) on PATH (for manifest validation)
- A Kubernetes cluster (kind/minikube/real) to deploy into

## Usage

Build and run locally against your current kubeconfig context:

```bash
make build
./bin/manager --leader-elect=false
```

Apply the CRD and a sample Ping:

```bash
kubectl apply -f config/crd/ping.yaml
kubectl apply -f config/samples/ping-sample.yaml
```

Watch the operator ping:

```bash
kubectl logs -l app.kubernetes.io/name=devops-k8s-operator -f
kubectl get ping ping-sample -o wide
```

Deploy the manager:

```bash
make docker-build IMG=<your-registry>/devops-k8s-operator:latest
kubectl apply -f config/crd/ping.yaml
kubectl apply -f config/rbac/role.yaml
kubectl apply -f config/manager/deployment.yaml
```

## Validation

This repo has no live cluster, so it is validated by **build + vet + kubeconform
schema checks**, never by deploying:

```bash
./scripts/validate.sh
```

## Continuous integration

Every push to `main` and every pull request runs `.github/workflows/ci.yaml`,
which executes the same gate as a local checkout: `go build ./...`,
`go vet ./...`, `go test ./...`, and `./scripts/validate.sh` (kubeconform
schema checks). Go modules are cached between runs and `kubeconform` is
installed from a pinned release, so a green check on a PR means the manifests
and code validated together.

## Next steps (refinement backlog)

- Add unit/integration tests (`envtest`).
- Add a conversion/mutating webhook example.
- Add `Makefile` deploy targets via kustomize.
