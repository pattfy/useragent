export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE := on
export CGO_ENABLED := 0


GO_TOOLS=\
golang.org/x/tools/cmd/goimports \
github.com/golang/lint/golint \
github.com/jessevdk/go-assets-builder

xargs := $(shell which gxargs xargs | head -n1)

all: fmt imports lint vet

fmt:
	find . -name '*.go' | grep -vE '/(vendor|assets)/' | $(xargs) gofmt -l | $(xargs) -r false

imports:
	find . -name '*.go' | grep -vE '(/(vendor|assets)/)' | $(xargs) goimports -l | $(xargs) -r false

lint:
	golint ./... | grep -vE '^(vendor|assets)/' | $(xargs) -r false

vet:
	go vet ./...

assets:
	go-assets-builder --package assets -o assets/assets.go bots.yml


tools:
	@for pkg in $(GO_TOOLS); do \
		go install $$pkg; \
	done

