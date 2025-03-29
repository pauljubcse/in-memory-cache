package in_memory_cache

type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

type List[K comparable, V any] struct {
	head, tail *Node[K, V]
}

func NewList[K comparable, V any]() *List[K, V] {
	return &List[K, V]{}
}

// Insert at front
func (l *List[K, V]) InsertFront(node *Node[K, V]) {
	node.next = l.head
	node.prev = nil
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

func (l *List[K, V]) RemoveNode(node *Node[K, V]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}

}

func (l *List[K, V]) MoveToFront(node *Node[K, V]) {
	l.RemoveNode(node)
	l.InsertFront(node)
}
