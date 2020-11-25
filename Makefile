# Image URL to use all building/pushing image targets
IMG ?= node-down-webhook:latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: node-down-webhook

# Run tests
test: fmt vet
	go test ./... -coverprofile cover.out

# Build manager binary
node-down-webhook: fmt vet
	go build -o bin/node-down-webhook

# Deploy the Node Down Handler in the configured Kubernetes cluster in ~/.kube/config
deploy:
	helm upgrade -i --wait node-down-webhook ./helm/chart --set image=${IMG}

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

docker-build: test
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

# Load the docker image to a local kind cluster
kind-load:
	kind load docker-image ${IMG}
