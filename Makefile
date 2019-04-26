.PHONY: default
default:
	@make $(shell uname | tr [A-Z] [a-z])

.PHONY: zip
zip: darwin linux windows
	mkdir -p build/releases
	(cd build/linux; zip -r ../releases/dcos-http-cli.linux.zip .)
	(cd build/darwin; zip -r ../releases/dcos-http-cli.darwin.zip .)
	(cd build/windows; zip -r ../releases/dcos-http-cli.windows.zip .)

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
