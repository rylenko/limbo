.POSIX:

bin/:
	mkdir -p $@

bin/golangci-lint: | bin/
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.61.0

lint: bin/golangci-lint
	go list -f '{{.Dir}}' -m | xargs ./bin/golangci-lint run --fix

.PHONY: lint
