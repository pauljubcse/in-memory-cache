package cache

type LRUCache[K comparable, V any] struct {
	capacity int
	list     *List[K, V]
	cache    map[K]*Node[K, V]
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		list:     NewList[K, V](),
		cache:    make(map[K]*Node[K, V]),
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	var zero V
	if node, found := c.cache[key]; found {
		c.list.MoveToFront(node)
		return node.value, true
	}
	return zero, false
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if node, found := c.cache[key]; found {
		node.value = value
		c.list.MoveToFront(node)
		return
	}
	if len(c.cache) == c.capacity {
		delete(c.cache, c.list.tail.key)
		c.list.RemoveNode(c.list.tail)
	}
	newNode := &Node[K, V]{
		key:   key,
		value: value,
	}
	c.list.InsertFront(newNode)
	c.cache[key] = newNode
}
