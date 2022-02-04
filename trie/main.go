package main

import "fmt"

// AlphabetSize
const AlphabetSize = 26

// Node
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert
func (t *Trie) Insert(w string) {
	wl := len(w)
	currentNode := t.root
	for i := 0; i < wl; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			// create node and put it
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	wl := len(w)
	currentNode := t.root
	for i := 0; i < wl; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd {
		return true
	}

	return false
}

func main() {
	testTrie := InitTrie()
	//fmt.Println(testTrie.root)

	testTrie.Insert("aragon")
	fmt.Println(testTrie.Search("aragon"))

	fmt.Println(testTrie.root)

}
