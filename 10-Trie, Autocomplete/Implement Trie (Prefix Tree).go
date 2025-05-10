type Trie struct {
	startNode *Node
}

type Node struct {
	isWord   bool
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
	}
	node.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.startNode
	for _, r := range word {
		next, ok := node.children[r]
		if !ok {
			return false
		}
		node = next
	}
	return node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.startNode
	for _, r := range prefix {
		next, ok := node.children[r]
		if !ok {
			return false
		}
		node = next
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
