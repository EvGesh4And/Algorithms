package main

type LRUCache struct {
	cap      int
	m        map[int]*listNode
	fakeHead *listNode
	fakeTail *listNode
}

func (this *LRUCache) moveHead(node *listNode) {
	node.prev.next, node.next.prev = node.next, node.prev
	this.fakeHead.next, this.fakeHead.next.prev, node.next = node, node, this.fakeHead.next
	node.prev = this.fakeHead
}

func (this *LRUCache) add(key, value int) {
	node := &listNode{
		key:   key,
		value: value,
	}
	this.fakeHead.next, this.fakeHead.next.prev, node.next = node, node, this.fakeHead.next
	node.prev = this.fakeHead

	this.m[key] = node

	if len(this.m) > this.cap {
		delNode := this.fakeTail.prev
		delNode.prev.next, delNode.next.prev = delNode.next, delNode.prev
		delete(this.m, delNode.key)
	}
}

func Constructor(capacity int) LRUCache {
	fakeHead := &listNode{}
	fakeTail := &listNode{}
	fakeHead.next, fakeTail.prev = fakeTail, fakeHead
	return LRUCache{
		cap:      capacity,
		m:        make(map[int]*listNode),
		fakeHead: fakeHead,
		fakeTail: fakeTail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value
		this.moveHead(node)
	} else {
		this.add(key, value)
	}
}

type listNode struct {
	key   int
	value int
	next  *listNode
	prev  *listNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
