# LeetCode Solutions in Go

This repository contains my solutions to LeetCode problems, organized by difficulty and implemented in Go with comprehensive testing.

## Structure

- `problems/`: Solutions organized by difficulty (easy/medium/hard)
- `cmd/test-runner/`: Command-line tool to run specific tests
- `pkg/common/`: Common data structures and utilities
- `.devcontainer/`: Docker development environment setup

## Usage

### Running Tests

```bash
# Run all tests
make test

# Run tests for a specific problem
make test-problem PROBLEM=two-sum

# Run with verbose output
make test-verbose

# Run with coverage
make test-cover
```

### Using the Test Runner

```bash
# Run all tests for two-sum problem
go run cmd/test-runner/main.go two-sum

# Run specific test case
go run cmd/test-runner/main.go two-sum TestTwoSumHashMap
```

### Debugging in VSCode

1. Open the problem file you want to debug
2. Set breakpoints in your solution code
3. Use `Ctrl+Shift+P` → "Debug: Start Debugging"
4. Choose "Debug Specific Problem Test"
5. Enter the problem path and test name when prompted

## Development Environment

This project uses a devcontainer for consistent development environment:

1. Install Docker and VSCode with Dev Containers extension
2. Open the project in VSCode
3. When prompted, click "Reopen in Container"

## Adding New Problems

```bash
make new-problem DIFFICULTY=<easy|medium|hard|extra> PROBLEM=<problem name>
```

where `extra` is for non leetcode problems.
