.PHONY: test test-verbose test-problem test-cover build-test-runner clean dev-setup new-problem build-problem-creator

# Run all tests
test:
	go test ./problems/...

# Run all tests with verbose output
test-verbose:
	go test -v ./problems/...

# Run tests for a specific problem
test-problem: build-test-runner
	@if [ -z "$(PROBLEM)" ]; then \
		echo "Usage: make test-problem PROBLEM=<problem-name>"; \
		echo "Example: make test-problem PROBLEM=two-sum"; \
		exit 1; \
	fi
	./bin/run-test $(PROBLEM)

# Run tests with coverage
test-cover:
	go test -cover ./problems/...

# Run tests with detailed coverage report
test-cover-html:
	go test -coverprofile=coverage.out ./problems/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build the test runner
build-test-runner:
	go build -o bin/run-test cmd/run-test/main.go

# Build problem creator
build-problem-creator:
	go build -o bin/create-problem cmd/create-problem/main.go

# Create new problem structure
new-problem: build-problem-creator
	@if [ -z "$(DIFFICULTY)" ] || [ -z "$(PROBLEM)" ]; then \
		echo "Usage: make new-problem DIFFICULTY=<easy|medium|hard|extra> PROBLEM=<problem-name>"; \
		echo "Example: make new-problem DIFFICULTY=easy PROBLEM=two-sum"; \
		exit 1; \
	fi
	./bin/create-problem $(DIFFICULTY) $(PROBLEM)

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

clean-coverage:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Setup development environment
dev-setup:
	go mod init go-leetcode
	go mod tidy
	go install -v golang.org/x/tools/gopls@latest

# Format code
fmt:
	go fmt ./problems/...

# Run linter
lint:
	golangci-lint run

# Run benchmarks
bench:
	go test -bench=. ./problems/...