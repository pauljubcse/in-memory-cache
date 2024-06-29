package cache

import (
	"testing"
)

func TestListOperations(t *testing.T) {
	list := NewList[int, int]()

	node1 := &Node[int, int]{key: 1, value: 100}
	node2 := &Node[int, int]{key: 2, value: 200}
	node3 := &Node[int, int]{key: 3, value: 300}

	// Test InsertFront
	list.InsertFront(node1)
	if list.head != node1 || list.tail != node1 {
		t.Errorf("InsertFront failed: head = %v, tail = %v", list.head, list.tail)
	}

	list.InsertFront(node2)
	if list.head != node2 || list.tail != node1 || node2.next != node1 {
		t.Errorf("InsertFront failed: head = %v, tail = %v", list.head, list.tail)
	}

	// Test MoveToFront
	list.MoveToFront(node1)
	if list.head != node1 || node1.next != node2 {
		t.Errorf("MoveToFront failed: head = %v", list.head)
	}

	// Test RemoveNode
	list.InsertFront(node3)
	list.RemoveNode(node3)
	if list.head == node3 || list.tail == node3 {
		t.Errorf("RemoveNode failed: head = %v, tail = %v", list.head, list.tail)
	}
	if node1.next != node2 || node1.prev != nil {
		t.Errorf("RemoveNode failed: node2.next = %v, node1.prev = %v", node2.next, node1.prev)
	}
}
func TestLRUCache(t *testing.T) {
	cache := NewLRUCache[string, string](2)

	cache.Put("key1", "value1")
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected to get 'value1', got %v", val)
	}

	cache.Put("key2", "value2")
	if val, ok := cache.Get("key2"); !ok || val != "value2" {
		t.Errorf("Expected to get 'value2', got %v", val)
	}

	cache.Put("key3", "value3") // This should evict "key1"
	if _, ok := cache.Get("key1"); ok {
		t.Errorf("Expected 'key1' to be evicted")
	}

	if val, ok := cache.Get("key3"); !ok || val != "value3" {
		t.Errorf("Expected to get 'value3', got %v", val)
	}

	cache.Put("key4", "value4") // This should evict "key2"
	if _, ok := cache.Get("key2"); ok {
		t.Errorf("Expected 'key2' to be evicted")
	}

	if val, ok := cache.Get("key4"); !ok || val != "value4" {
		t.Errorf("Expected to get 'value4', got %v", val)
	}
}
