package main

import "fmt"

func main() {
	hashMap := Constructor()
	hashMap.Put(5, 10)
	hashMap.Put(6, 11)
	hashMap.Put(5, 12)
	fmt.Println(hashMap.Get(6))
	hashMap.Remove(6)
	fmt.Println(hashMap.Get(6))
	fmt.Println(hashMap.Get(5))
	hashMap.Put(4, 13)
	hashMap.Put(2, 111)
	fmt.Println(hashMap.Get(2))
}

type MyHashMap struct {
	n       int
	buckets []*LinkedList
}

func Constructor() MyHashMap {
	n := 1
	m := MyHashMap{
		n:       n,
		buckets: make([]*LinkedList, n),
	}
	for i := range m.buckets {
		m.buckets[i] = &LinkedList{}
	}
	return m
}

func (this *MyHashMap) Put(key int, value int) {
	this.buckets[this.hash(key)].put(key, value)
}

func (this *MyHashMap) Get(key int) int {
	return this.buckets[this.hash(key)].get(key)
}

func (this *MyHashMap) Remove(key int) {
	this.buckets[this.hash(key)].remove(key)
}

func (this *MyHashMap) hash(key int) int {
	return key % this.n
}

type node struct {
	key   int
	value int
	next  *node
}

type LinkedList struct {
	head *node
}

func (this *LinkedList) get(key int) int {
	curr := this.head
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.next
	}
	return -1
}

func (this *LinkedList) put(key, value int) {
	curr := this.head
	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		curr = curr.next
	}
	this.head = &node{key, value, this.head}
}

func (this *LinkedList) remove(key int) {
	fakehead := &node{next: this.head}
	prev := fakehead
	curr := this.head
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			break
		}
		prev = prev.next
		curr = curr.next
	}
	this.head = fakehead.next
}
