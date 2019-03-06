# RP10 Custom Metrics for k8s

It should include:

    - Response time
    - Number of active requests

## Download the project

Firstly, make sure your `GOPATH` is set. Then:

```bash
go get github.com/keni7385/rp10-custom-metrics
```

## Download dependencies

On the first time, it might take few minutes:

```bash
cd "$GOPATH/src/github.com/keni7385/rp10-custom-metrics"
make vendor
```

This passage is optional, because will be automatically performed by the [build step](#build)

## Build

```bash
make build-adapter
```

## Deploy

```bash
make adapter-container
```

It creates a Docker image and push it to the target registry (set var `$REGISTRY`, see `Makefile`). Then will populate the file `deploy/adapter.yaml` (originally a template, where `REGISTRY` and `IMAGE` is filled in).
