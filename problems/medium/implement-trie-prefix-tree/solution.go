package main

type Trie struct {
	// use array here to avoid dynamic malloc make this faster than map
	// fixed size wasted some memory though
	chars [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		[26]*Trie{},
		false,
	}
}

func (trie *Trie) Insert(word string) {
	current := trie

	for pos, char := range word {
		subTriePtr := current.chars[char-'a']
		if subTriePtr == nil {
			subTrie := Constructor()
			current.chars[char-'a'] = &subTrie
			current = &subTrie
		} else {
			current = subTriePtr
		}

		if pos == len(word)-1 {
			current.isEnd = true
		}
	}
}

func (trie *Trie) Search(word string) bool {
	current := trie

	for _, char := range word {
		subTriePtr := current.chars[char-'a']
		if subTriePtr == nil {
			return false
		} else {
			current = subTriePtr
		}
	}

	return current.isEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	current := trie

	for _, char := range prefix {
		subTriePtr := current.chars[char-'a']
		if subTriePtr == nil {
			return false
		} else {
			current = subTriePtr
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// helpers
type Action string

const (
	Insert     Action = "insert"
	Search     Action = "search"
	StartsWith Action = "startsWith"
)

func boolPtr(val bool) *bool {
	return &val
}
