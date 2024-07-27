package design

type LRUCache struct {
	capacity int
	size     int
	entries  map[int]*Entry
	head     *Entry
	tail     *Entry
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		size:     0,
		entries:  make(map[int]*Entry),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	// If the key is already in the cache
	if entry, exists := this.entries[key]; exists {
		// Delete entry from the current position and move to head
		this.deleteEntry(entry)
		this.updateHead(entry)
		return entry.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// If the key already exists in the cache
	if entry, exists := this.entries[key]; exists {
		entry.value = value
		this.deleteEntry(entry)
		this.updateHead(entry)
	} else {
		// Create new node
		entry := &Entry{key: key, value: value}
		if this.size >= this.capacity {
			delete(this.entries, this.tail.key)
			this.deleteEntry(this.tail)
			this.size--
		}
		// Update head
		this.updateHead(entry)
		this.entries[key] = entry
		this.size++
	}
}

func (this *LRUCache) deleteEntry(entry *Entry) {
	// If the given entry is not the head
	if entry.previous != nil {
		entry.previous.next = entry.next
	} else {
		this.head = entry.next
	}
	// If the given entry is not the tail
	if entry.next != nil {
		entry.next.previous = entry.previous
	} else {
		this.tail = entry.previous
	}
}

func (this *LRUCache) updateHead(entry *Entry) {
	// Make entry as the new head
	entry.next = this.head
	entry.previous = nil
	// If head is not null
	if this.head != nil {
		this.head.previous = entry
	}
	// Update the head
	this.head = entry
	// If there's only single node
	if this.tail == nil {
		this.tail = entry
	}
}

type Entry struct {
	key, value     int
	previous, next *Entry
}
