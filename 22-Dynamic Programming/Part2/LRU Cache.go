package main

import "fmt"

func main() {
	lru := Constructor(2)
	fmt.Println(lru)
	fmt.Println(lru.Get(4))
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
}

type Node struct {
	key  int
	prev *Node
	next *Node
}

func (n *Node) del() {
	n.prev.next = n.next
	n.next.prev = n.prev
}

type LRUCache struct {
	cap int
	m   map[int]struct {
		value int
		node  *Node
	}
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{cap: capacity, m: make(map[int]struct {
		value int
		node  *Node
	}), head: &Node{}, tail: &Node{}}
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (this *LRUCache) addToHead(node *Node) {
	this.head.next.prev = node
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		if v.node.prev != this.head {
			v.node.del()
			this.addToHead(v.node)
		}
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.value = value
		this.m[key] = v
		if v.node != this.head.next {
			v.node.del()
			this.addToHead(v.node)
		}
	} else {

		if len(this.m) == this.cap {
			delKey := this.tail.prev.key
			this.tail.prev.del()
			delete(this.m, delKey)
		}

		node := &Node{key: key}
		this.addToHead(node)
		this.m[key] = struct {
			value int
			node  *Node
		}{value: value, node: node}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
