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

It might take few minutes:

```bash
cd "$GOPATH/src/github.com/keni7385/rp10-custom-metrics"
glide install -v
```

(Option `-v` ESSENTIAL, `glide install -h` for help)

## Build

```bash
go build
```

## Test & Deploy

WIP
