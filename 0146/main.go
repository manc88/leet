package main

import (
	"container/list"
	"sync"
)

type Item struct {
	uid   int
	value int
	pos   *list.Element
}

type LRUCache struct {
	mu       sync.RWMutex
	hashmap  map[int]*Item
	list     *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hashmap:  make(map[int]*Item, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	newItem := &Item{
		uid:   key,
		value: value,
		pos:   nil,
	}

	//if item is already in cache
	if item, ok := this.hashmap[newItem.uid]; ok {
		//delete from list and hashmap
		this.list.Remove(item.pos)
		delete(this.hashmap, item.uid)

		//insert into end of list and update into hm
		newItem.pos = this.list.PushBack(newItem.uid)
		this.hashmap[newItem.uid] = newItem
		return
	}

	if len(this.hashmap) >= this.capacity {
		f := this.list.Front()
		this.list.Remove(f)
		delete(this.hashmap, f.Value.(int))
	}

	newItem.pos = this.list.PushBack(newItem.uid)
	this.hashmap[newItem.uid] = newItem
}

func (this *LRUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()

	if item, ok := this.hashmap[key]; ok {
		//update pos
		this.list.Remove(item.pos)
		delete(this.hashmap, item.uid)

		//insert
		item.pos = this.list.PushBack(item.uid)
		this.hashmap[item.uid] = item

		return item.value
	}

	return -1
}
