.POSIX:

bin/:
	mkdir -p $@

bin/golangci-lint: | bin/
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.61.0

clean:
	rm -rf bin test-coverage test-coverage.html

lint: bin/golangci-lint
	./bin/golangci-lint run --fix ./... ./pkg/*

test:
	go test -coverprofile test-coverage ./... ./pkg/*

test-coverage: test
	go tool cover -html=$@ -o $@.html
	xdg-open $@.html

.PHONY: lint test test-coverage
