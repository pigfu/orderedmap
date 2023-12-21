package orderedmap

import (
	"fmt"
	"strings"
)

type DoubleList[K, V any] struct {
	head *DoubleNode[K, V]
	tail *DoubleNode[K, V]
	//length int
}

type DoubleNode[K any, V any] struct {
	key   K
	value V
	pre   *DoubleNode[K, V]
	next  *DoubleNode[K, V]
	//
	head bool
	tail bool
}

// DoubleListIter to iteration
type DoubleListIter[K, V any] struct {
	head *DoubleNode[K, V]
}

func NewDoubleList[K, V any]() ListI[K, V] {
	dl := &DoubleList[K, V]{
		head: &DoubleNode[K, V]{head: true},
		tail: &DoubleNode[K, V]{tail: true},
	}
	dl.head.next = dl.tail
	dl.tail.pre = dl.head
	return dl
}

// Insert push back list
func (dl *DoubleList[K, V]) Insert(key K, value V) NodeI[K, V] {
	node := &DoubleNode[K, V]{key: key, value: value}

	dl.tail.pre.next = node
	node.pre = dl.tail.pre
	node.next = dl.tail
	dl.tail.pre = node

	return node
}

// Delete del appoint node
func (dl *DoubleList[K, V]) Delete(nodeI NodeI[K, V]) {
	node := nodeI.(*DoubleNode[K, V])
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (dl *DoubleList[K, V]) Iter() IterI[K, V] {
	return &DoubleListIter[K, V]{dl.head}
}

func (dl *DoubleList[K, V]) String() string {
	str := &strings.Builder{}
	str.WriteString("double linked list:")
	head := dl.head.next
	for !head.tail {
		str.WriteString("[")
		str.WriteString(fmt.Sprintf("Key:%v,", head.key))
		str.WriteString(fmt.Sprintf("Value:%v", head.value))
		str.WriteString("],")
		head = head.next
	}
	return str.String()
}

func (node *DoubleNode[K, V]) GetValue() V {
	return node.value
}

func (node *DoubleNode[K, V]) SetValue(value V) {
	node.value = value
}

func (node *DoubleNode[K, V]) String() string {
	if node == nil {
		return "node is nil"
	}
	str := &strings.Builder{}
	str.WriteString("[")
	str.WriteString(fmt.Sprintf("Key:%v,", node.key))
	str.WriteString(fmt.Sprintf("Value:%v", node.value))
	str.WriteString("]")
	return str.String()
}

func (iter *DoubleListIter[K, V]) Next() bool {
	if iter.head.next.tail {
		return false
	}
	iter.head = iter.head.next
	return true
}

func (iter *DoubleListIter[K, V]) KV() (K, V) {
	return iter.head.key, iter.head.value
}
