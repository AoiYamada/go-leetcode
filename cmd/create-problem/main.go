package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run create_problem.go <difficulty> <problem-name>")
		fmt.Println("Example: go run create_problem.go easy two-sum")
		fmt.Println("Example: go run create_problem.go hard median-of-arrays")
		return
	}

	difficulty := os.Args[1]
	problemName := os.Args[2]

	// Validate difficulty
	validDifficulties := map[string]bool{
		"easy":   true,
		"medium": true,
		"hard":   true,
		"extra":  true, // this one is for non leetcode problems
	}
	if !validDifficulties[difficulty] {
		fmt.Printf("Error: Difficulty must be 'easy', 'medium', 'hard, or 'extra', got '%s'\n", difficulty)
		return
	}

	// Create title case converter
	titleCaser := cases.Title(language.English)

	paths := map[string]string{
		"solution.go": `package main

// {{.DifficultyTitle}}: {{.ProblemNameTitle}}
// Solution for {{.ProblemName}} ({{.Difficulty}})
func solve() {
    // Your solution here
}
`,
		"solution_test.go": `package main

import "testing"

func Test{{.ProblemNameTitle}}(t *testing.T) {
    tests := []struct {
        name string
    }{
        {"test case 1"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation for {{.ProblemName}} ({{.Difficulty}})
        })
    }
}
`,
		"README.md": `# {{.ProblemNameTitle}} ({{.DifficultyTitle}})

Problem description goes here.
`,
	}

	data := struct {
		Difficulty       string
		ProblemName      string
		DifficultyTitle  string
		ProblemNameTitle string
	}{
		Difficulty:       difficulty,
		ProblemName:      problemName,
		DifficultyTitle:  titleCaser.String(difficulty),
		ProblemNameTitle: toCamelCase(problemName),
	}

	for filename, content := range paths {
		fullPath := filepath.Join("problems", difficulty, problemName, filename)
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		if err != nil {
			panic(fmt.Sprintf("Failed to create directory: %v", err))
		}

		tmpl, err := template.New(filename).Parse(content)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse template for %s: %v", filename, err))
		}

		file, err := os.Create(fullPath)
		if err != nil {
			panic(fmt.Sprintf("Failed to create file %s: %v", fullPath, err))
		}
		defer file.Close()

		err = tmpl.Execute(file, data)
		if err != nil {
			panic(fmt.Sprintf("Failed to execute template for %s: %v", filename, err))
		}
		fmt.Printf("Created: %s\n", fullPath)
	}
}

// toCamelCase converts kebab-case or snake_case to CamelCase
func toCamelCase(s string) string {
	titleCaser := cases.Title(language.English)
	result := ""
	nextUpper := true

	for _, char := range s {
		if char == '-' || char == '_' {
			nextUpper = true
			continue
		}
		if nextUpper {
			result += string(titleCaser.String(string(char)))
			nextUpper = false
		} else {
			result += string(char)
		}
	}
	return result
}
