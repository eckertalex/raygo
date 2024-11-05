# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n "s/^##//p" ${MAKEFILE_LIST} | column -t -s ":" |  sed -e "s/^/ /"

.PHONY: confirm
confirm:
	@echo "Are you sure? (y/n) \c"
	@read answer; \
	if [ "$$answer" != "y" ]; then \
		echo "Aborting."; \
		exit 1; \
	fi

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy and format all .go files
.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...

## build/raygo: build the cmd/raygo application
.PHONY: build/raygo
build/raygo:
	@go build -o=./tmp/raygo ./cmd/raygo

## run/raygo: run the cmd/raygo application
.PHONY: run/raygo
run/raygo: build/raygo
	@./tmp/raygo > ./tmp/image.ppm

# vim: set tabstop=4 shiftwidth=4 noexpandtab
