.PHONY: test test-verbose bench-problem test-cover build-bench-runner clean dev-setup new-problem build-problem-creator

# Run all tests
test:
	go test ./problems/...

# Run all tests with verbose output
test-verbose:
	go test -v ./problems/...

# Run tests for a specific problem
bench: build-bench-runner
	@if [ -z "$(PROBLEM)" ]; then \
		echo "Usage: make bench PROBLEM=<problem-name>"; \
		echo "Example: make bench PROBLEM=two-sum"; \
		exit 1; \
	fi
	./bin/run-bench $(PROBLEM)

# Run tests with coverage
test-cover:
	go test -cover ./problems/...

# Run tests with detailed coverage report
test-cover-html:
	go test -coverprofile=coverage.out ./problems/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build the test runner
build-bench-runner:
	go build -o bin/run-bench cmd/run-bench/main.go

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
