package main

import "slices"

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := Constructor()
	for _, word := range products {
		trie.Insert(word)
	}
	return trie.Search(searchWord)
}

type Trie struct {
	startNode *Node
}

type Node struct {
	words    []string
	children map[rune]*Node
}

func Constructor() Trie {
	return Trie{
		startNode: &Node{
			children: make(map[rune]*Node),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.startNode
	for _, r := range word {
		if next, ok := node.children[r]; ok {
			node = next
		} else {
			newNode := &Node{children: make(map[rune]*Node)}
			node.children[r] = newNode
			node = newNode
		}
		node.words = append(node.words, word)
	}
}

func (t *Trie) Search(word string) [][]string {
	res := make([][]string, len(word))

	node := t.startNode
	for i, r := range word {
		next, ok := node.children[r]
		if !ok {
			return res
		}

		slices.Sort(next.words)
		if len(next.words) < 3 {
			res[i] = next.words
		} else {
			res[i] = next.words[:3]
		}

		node = next
	}

	return res
}
