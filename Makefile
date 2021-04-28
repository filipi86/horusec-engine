GO ?= go
GOFMT ?= gofmt
GO_FILES ?= $$(find . -name '*.go' | grep -v vendor)
GOLANG_CI_LINT ?= ./bin/golangci-lint
GO_IMPORTS ?= goimports
GO_IMPORTS_LOCAL ?= github.com/ZupIT/horusec-engine
HORUSEC ?= horusec
DOCKER_COMPOSE ?= docker-compose

compose:
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.yaml up -d --build --force-recreate

fmt:
	$(GOFMT) -w $(GO_FILES)

lint:
    ifeq ($(wildcard $(GOLANG_CI_LINT)), $(GOLANG_CI_LINT))
		$(GOLANG_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...
    else
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
		$(GOLANG_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...
    endif

coverage:
	curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec-devkit/develop/scripts/coverage.sh | bash -s 65 .

test:
	$(GO) clean -testcache && $(GO) test -v ./... -timeout=2m -parallel=1 -failfast -short

fix-imports:
    ifeq (, $(shell which $(GO_IMPORTS)))
		$(GO) get -u golang.org/x/tools/cmd/goimports
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GO_FILES)
    else
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GO_FILES)
    endif

security:
    ifeq (, $(shell which $(HORUSEC)))
		curl -fsSL https://horusec.io/bin/install.sh | bash
		$(HORUSEC) start -p="./" -e="true"
    else
		$(HORUSEC) start -p="./" -e="true"
    endif

pipeline: fmt fix-imports lint test coverage build security

