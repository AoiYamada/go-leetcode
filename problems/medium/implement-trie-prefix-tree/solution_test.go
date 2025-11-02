package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementTriePrefixTree(t *testing.T) {
	testCases := []struct {
		name     string
		actions  []Action
		words    []string
		expected []*bool
	}{
		{"Example 1", []Action{"insert", "search", "search", "startsWith", "insert", "search"}, []string{"apple", "apple", "app", "app", "app", "app"}, []*bool{nil, boolPtr(true), boolPtr(false), boolPtr(true), nil, boolPtr(true)}},
		// cover all branches
		{"My test 1", []Action{"insert", "search", "startsWith"}, []string{"app", "apt", "abb"}, []*bool{nil, boolPtr(false), boolPtr(false)}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for implement-trie-prefix-tree (medium)
			t.Run(tc.name, func(t *testing.T) {
				obj := Constructor()

				for pos, action := range tc.actions {
					word := tc.words[pos]

					switch action {
					case Insert:
						obj.Insert(word)

					case Search:
						result := obj.Search(word)
						expected := *tc.expected[pos]

						assert.Equal(t, expected, result)

					case StartsWith:
						result := obj.StartsWith(word)
						expected := *tc.expected[pos]

						assert.Equal(t, expected, result)

					default:
					}
				}
			})
		})
	}
}
