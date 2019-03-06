REGISTRY?=keni7385
TEMP_DIR:=$(shell mktemp -d)
OUT_DIR=./_output

ARCH?=amd64
VERSION?=latest
IMAGE?=rp10_metrics_adapter-$(ARCH):$(VERSION)

.PHONY: all build-adapter adapter-container

all: build-adapter
build-adapter: vendor
	CGO_ENABLE=0 GOARCH=$(ARCH) go build -o $(OUT_DIR)/$(ARCH)/adapter github.com/keni7385/rp10-custom-metrics

vendor: glide.lock
		glide install -v

adapter-container: build-adapter
	cp deploy/Dockerfile_template $(TEMP_DIR)/Dockerfile
	cp $(OUT_DIR)/$(ARCH)/adapter $(TEMP_DIR)/adapter
	cd $(TEMP_DIR) && sed -i "s|BASEIMAGE|scratch|g" Dockerfile
	sed -i 's|REGISTRY|'$(REGISTRY)'|g' deploy/adapter.yaml
	sed -i 's|IMAGE|'$(IMAGE)'|g' deploy/adapter.yaml
	docker build -t $(REGISTRY)/$(IMAGE) $(TEMP_DIR)
	rm -rf $(TEMP_DIR)
