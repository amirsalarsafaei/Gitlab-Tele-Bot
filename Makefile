REVISER_VERSION := v3.3.1

build:
	go build \
		-o notifier-cli


revise-imports: $(GOPATH)/bin/goimports-reviser
	@goimports-reviser -company-prefixes "github.com/amirsalarsafaei/" ./...

$(GOPATH)/bin/goimports-reviser:
	@go install -v github.com/incu6us/goimports-reviser/v3@$(REVISER_VERSION)
