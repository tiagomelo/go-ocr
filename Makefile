.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test
## test: runs unit tests
test:
	@ go test -v ./... -count=1

.PHONY: coverage
## coverage: run unit tests and generate coverage report in html format
coverage:
	@ go test -coverprofile=coverage.out ./...  && go tool cover -html=coverage.out