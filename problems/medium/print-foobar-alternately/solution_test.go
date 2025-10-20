package main

import (
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintFoobarAlternately(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected string
	}{
		{"Example 1", 1, "foobar"},
		{"Example 2", 2, "foobarfoobar"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var output strings.Builder
			// Add mutex to protect concurrent writes
			var mu sync.Mutex

			// Create custom print functions that capture output
			printFoo := func() {
				mu.Lock()
				defer mu.Unlock()
				output.WriteString("foo")
			}

			printBar := func() {
				mu.Lock()
				defer mu.Unlock()
				output.WriteString("bar")
			}

			foobar := NewFooBar(tc.n)

			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				foobar.Foo(printFoo)
			}()

			go func() {
				defer wg.Done()
				foobar.Bar(printBar)
			}()

			wg.Wait()

			assert.Equal(t, tc.expected, output.String())
		})
	}
}
