.PHONY: default
default:
	@make $(shell uname | tr [A-Z] [a-z])

.PHONY: zip
zip: darwin linux
	mkdir -p build/releases
	(cd build/linux; zip -r ../releases/dcos-http-cli.linux.zip .)
	(cd build/darwin; zip -r ../releases/dcos-http-cli.darwin.zip .)

.PHONY: darwin linux
darwin linux:
	GOOS=$(@) go build -o build/$(@)/bin/dcos-http ./main.go
	cp plugin.toml build/$(@)/

.PHONY: vet
vet: lint
	go vet ./main.go

.PHONY: lint
lint:
	# Can be simplified once https://github.com/golang/lint/issues/320 is fixed.
	golint -set_exit_status ./main.go

.PHONY: vendor
vendor:
	dep ensure

.PHONY: clean
clean:
	rm -rf build
