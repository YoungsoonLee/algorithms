package main

import "fmt"

// AlphabetSize ...
const AlphabetSize = 26

// Node ...
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert ...
func (t *Trie) Insert(w string) {
	wl := len(w)
	currentNode := t.root
	for i := 0; i < wl; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// search
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

	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	t := InitTrie()
	fmt.Println(t)

	t.Insert("aragorn")
	fmt.Println(t.Search("aragorn"))

}
