.PHONY: setup check vet lint errcheck fmt fmt.check fmt.diff test pre-dep dep build help

GO_FILES_CMD := find . -name 'vendor' -prune -o -name '*.go' -print
GO_PATHS_CMD := $(GO_FILES_CMD) | awk -F/ '{ print $$1 "/" $$2 }' | uniq
PACKAGES_CMD := $(GO_PATHS_CMD) | grep -v '\.go$$' | awk -F/ '{ print $$0 "/..." }'

GO_PATHS := $(shell $(GO_PATHS_CMD))
PACKAGES := . $(shell $(PACKAGES_CMD))

# Show usage if execute "make" without option
.DEFAULT_GOAL := help

pre-dep:  ## Prepare dep command for resolve dependencies
	@if [ -z `which dep 2> /dev/null` ]; then \
		go get -u github.com/golang/dep/cmd/dep;\
	fi

setup: pre-dep ## Get tools
	@if [ -z `which golint 2> /dev/null` ]; then \
		go get -u golang.org/x/lint/golint;\
	fi
	@if [ -z `which errcheck 2> /dev/null` ]; then \
		go get -u github.com/kisielk/errcheck;\
	fi

check: vet lint fmt.check errcheck ## Execute all static analysis

vet: ## Execute go vet
	go vet $(PACKAGES)

lint: ## Execute golint
	echo $(PACKAGES) | xargs -n 1 golint -set_exit_status

errcheck: ## Execute github.com/kisielk/errcheck
	errcheck -ignoretests $(PACKAGES)

fmt: ## Format codes
	gofmt -l -w $(GO_PATHS)

fmt.check: ## Return error status if exist any format diff
	test -z "$(shell gofmt -l $(GO_PATHS))"

fmt.diff: ## Display diffs instead of rewriting files
	gofmt -d $(GO_PATHS)

dep: pre-dep ## Resolve dependencies
	dep ensure

build: dep ## Build binary
	go build -o ./bin/godecov ./cmd

clean: ## Cleanup destination directories
	@rm -rf bin/*

# Calculate coverage on CircleCI
GO_TEST := go test -v -race -p 1 $(GO_TEST_FLAGS)

test: ## Execute all tests
	$(GO_TEST) $(PACKAGES)

# Show help for each commands
help: ## Show options
	@grep -E '^[a-zA-Z_-{\.}]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
