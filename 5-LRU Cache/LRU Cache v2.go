package main

type LRUCache struct {
	capacity int
	cache    map[int]*listNode
	head     *listNode
	tail     *listNode
}

type listNode struct {
	key   int
	value int
	prev  *listNode
	next  *listNode
}

// Constructor инициализирует LRUCache
func Constructor(capacity int) LRUCache {
	head := &listNode{}
	tail := &listNode{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*listNode),
		head:     head,
		tail:     tail,
	}
}

// Get возвращает значение по ключу и обновляет порядок
func (c *LRUCache) Get(key int) int {
	if node, exists := c.cache[key]; exists {
		c.moveToFront(node)
		return node.value
	}
	return -1
}

// Put добавляет или обновляет ключ, и перемещает его в начало
func (c *LRUCache) Put(key int, value int) {
	if node, exists := c.cache[key]; exists {
		node.value = value
		c.moveToFront(node)
	} else {
		if len(c.cache) == c.capacity {
			// удаление least recently used
			lru := c.tail.prev
			c.remove(lru)
			delete(c.cache, lru.key)
		}
		newNode := &listNode{key: key, value: value}
		c.addToFront(newNode)
		c.cache[key] = newNode
	}
}

// moveToFront перемещает узел в начало списка
func (c *LRUCache) moveToFront(node *listNode) {
	c.remove(node)
	c.addToFront(node)
}

// addToFront вставляет узел сразу после head
func (c *LRUCache) addToFront(node *listNode) {
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}

// remove удаляет узел из списка
func (c *LRUCache) remove(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
