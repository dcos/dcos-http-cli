PLATFORM?=$(shell uname | tr [A-Z] [a-z])
export GO111MODULE = on

.PHONY: default
default:
	@make $(PLATFORM)

.PHONY: install
install:
	@make plugin
	dcos plugin add -u ./build/plugins/dcos-http-cli.$(PLATFORM).zip

.PHONY: plugin
plugin: $(PLATFORM)
	mkdir -p build/plugins
	(cd build/$(PLATFORM); zip -r ../plugins/dcos-http-cli.$(PLATFORM).zip .)

.PHONY: darwin linux windows
darwin linux windows:
	GOOS=$(@) go build -mod=vendor -o build/$(@)/bin/dcos-http ./cmd/dcos-http
	cp plugin.toml build/$(@)/
	cp -R completion build/$(@)

.PHONY: vet
vet: lint
	go vet -mod=vendor ./...

.PHONY: lint
lint:
	# Can be simplified once https://github.com/golang/lint/issues/320 is fixed.
	golint -set_exit_status $(go list -mod=vendor ./...)

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: clean
clean:
	rm -rf build
