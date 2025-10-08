package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/test-runner/main.go <problem-name> [test-case]")
		fmt.Println("Example: go run cmd/test-runner/main.go two-sum")
		fmt.Println("Example: go run cmd/test-runner/main.go two-sum TestTwoSum")
		os.Exit(1)
	}

	problemName := os.Args[1]
	var testCase string
	if len(os.Args) > 2 {
		testCase = os.Args[2]
	}

	// Find the problem directory
	problemPath, err := findProblemPath(problemName)
	if err != nil {
		fmt.Printf("Error finding problem: %v\n", err)
		os.Exit(1)
	}

	// Run benchmark
	err = runBench(problemPath, testCase)
	if err != nil {
		fmt.Printf("Error running test: %v\n", err)
		os.Exit(1)
	}
}

func findProblemPath(problemName string) (string, error) {
	difficulties := []string{
		"easy",
		"medium",
		"hard",
		"extra", // this one is for non leetcode problems
	}

	for _, difficulty := range difficulties {
		path := filepath.Join("problems", difficulty, problemName)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("problem '%s' not found", problemName)
}

func runBench(problemPath, testCase string) error {
	// Use benchmark flags instead of test flags
	args := []string{"test", "-bench=.", "-benchmem", "./" + problemPath}

	if testCase != "" {
		args = append(args, "-run", "^$") // Don't run regular tests
		args = append(args, "-bench", testCase)
	}

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
