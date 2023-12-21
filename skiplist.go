package orderedmap

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	maxLevel = 32   //From redis
	score    = 0.25 //From redis
)

type CmpKV[K any, V any] func(k1, k2 K, v1, v2 V) int

type SkipList[K any, V any] struct {
	head        *SkipNode[K, V]
	tail        *SkipNode[K, V]
	updateCache []*SkipNode[K, V]
	//
	cmp   CmpKV[K, V]
	r     *rand.Rand
	level int
}

type SkipNode[K any, V any] struct {
	key   K
	value V
	pre   []*SkipNode[K, V]
	next  []*SkipNode[K, V]
	//
	head bool
	tail bool
}

type SkipListIter[K any, V any] struct {
	head *SkipNode[K, V]
}

func NewSkipList[K any, V any](cmp CmpKV[K, V]) ListI[K, V] {
	sl := &SkipList[K, V]{
		head:        &SkipNode[K, V]{head: true, next: make([]*SkipNode[K, V], maxLevel)},
		tail:        &SkipNode[K, V]{tail: true, pre: make([]*SkipNode[K, V], maxLevel)},
		updateCache: make([]*SkipNode[K, V], maxLevel),
		cmp:         cmp,
		r:           rand.New(rand.NewSource(time.Now().UnixNano())),
		level:       0,
	}
	for i := 0; i < maxLevel; i++ {
		sl.head.next[i] = sl.tail
		sl.tail.pre[i] = sl.head
	}
	return sl
}
func (sl *SkipList[K, V]) randomLevel() int {
	level := 1
	for sl.r.Float64() < score && level <= maxLevel {
		level++
	}
	return level
}
func (sl *SkipList[K, V]) Insert(key K, value V) NodeI[K, V] {
	temp := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for !temp.next[i].tail && sl.cmp(temp.next[i].key, key, temp.next[i].value, value) < 0 {
			temp = temp.next[i]
		}
		sl.updateCache[i] = temp
	}

	//insert new node
	level := sl.randomLevel()

	node := sl.newSkipNode(level, key, value)

	for i := level - 1; i >= 0; i-- {
		if i >= sl.level { //beyond sl.level,set the node to sl.head.next[i]
			sl.head.next[i] = node
			node.pre[i] = sl.head
			node.next[i] = sl.tail
			sl.tail.pre[i] = node
			continue
		}
		//adjust pointer
		sl.updateCache[i].next[i].pre[i] = node
		node.pre[i] = sl.updateCache[i]
		node.next[i] = sl.updateCache[i].next[i]
		sl.updateCache[i].next[i] = node
	}
	if sl.level < level {
		sl.level = level
	}
	return node
}

// can reduce one memory allocation and improve performance
func (sl *SkipList[K, V]) newSkipNode(level int, key K, value V) *SkipNode[K, V] {
	switch level {
	case 1:
		n := struct {
			head    SkipNode[K, V]
			indexes [2]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 2:
		n := struct {
			head    SkipNode[K, V]
			indexes [4]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 3:
		n := struct {
			head    SkipNode[K, V]
			indexes [6]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 4:
		n := struct {
			head    SkipNode[K, V]
			indexes [8]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 5:
		n := struct {
			head    SkipNode[K, V]
			indexes [10]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 6:
		n := struct {
			head    SkipNode[K, V]
			indexes [12]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 7:
		n := struct {
			head    SkipNode[K, V]
			indexes [14]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 8:
		n := struct {
			head    SkipNode[K, V]
			indexes [16]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 9:
		n := struct {
			head    SkipNode[K, V]
			indexes [18]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 10:
		n := struct {
			head    SkipNode[K, V]
			indexes [20]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 11:
		n := struct {
			head    SkipNode[K, V]
			indexes [22]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 12:
		n := struct {
			head    SkipNode[K, V]
			indexes [24]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 13:
		n := struct {
			head    SkipNode[K, V]
			indexes [26]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 14:
		n := struct {
			head    SkipNode[K, V]
			indexes [28]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 15:
		n := struct {
			head    SkipNode[K, V]
			indexes [30]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 16:
		n := struct {
			head    SkipNode[K, V]
			indexes [32]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 17:
		n := struct {
			head    SkipNode[K, V]
			indexes [34]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 18:
		n := struct {
			head    SkipNode[K, V]
			indexes [36]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 19:
		n := struct {
			head    SkipNode[K, V]
			indexes [38]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 20:
		n := struct {
			head    SkipNode[K, V]
			indexes [40]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 21:
		n := struct {
			head    SkipNode[K, V]
			indexes [42]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 22:
		n := struct {
			head    SkipNode[K, V]
			indexes [44]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 23:
		n := struct {
			head    SkipNode[K, V]
			indexes [46]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 24:
		n := struct {
			head    SkipNode[K, V]
			indexes [48]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 25:
		n := struct {
			head    SkipNode[K, V]
			indexes [50]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 26:
		n := struct {
			head    SkipNode[K, V]
			indexes [52]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 27:
		n := struct {
			head    SkipNode[K, V]
			indexes [54]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 28:
		n := struct {
			head    SkipNode[K, V]
			indexes [56]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 29:
		n := struct {
			head    SkipNode[K, V]
			indexes [58]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 30:
		n := struct {
			head    SkipNode[K, V]
			indexes [60]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 31:
		n := struct {
			head    SkipNode[K, V]
			indexes [62]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	case 32:
		n := struct {
			head    SkipNode[K, V]
			indexes [64]*SkipNode[K, V]
		}{head: SkipNode[K, V]{key: key, value: value}}
		n.head.pre = n.indexes[:level]
		n.head.next = n.indexes[level:]
		return &n.head
	}

	panic("into the void")
}

func (sl *SkipList[K, V]) Delete(nodeI NodeI[K, V]) {
	// delete node
	node := nodeI.(*SkipNode[K, V])
	for i := 0; i < len(node.next); i++ {
		node.pre[i].next[i] = node.next[i]
		node.next[i].pre[i] = node.pre[i]
		node.pre[i] = nil
		node.next[i] = nil
		if sl.head.next[i].tail { //adjust level
			sl.level--
		}
	}
}

func (sl *SkipList[K, V]) Iter() IterI[K, V] {
	return &SkipListIter[K, V]{
		head: sl.head,
	}
}

func (sl *SkipList[K, V]) String() string {
	temp := sl.head
	str := &strings.Builder{}
	str.WriteString("skip list:")
	str.WriteString(strconv.Itoa(sl.level))
	for i := sl.level - 1; i >= 0; i-- {
		str.WriteString("\n[level:")
		str.WriteString(strconv.Itoa(i + 1))
		str.WriteString("]=>")
		for !temp.next[i].tail {
			str.WriteString("[")
			str.WriteString(fmt.Sprintf("Key:%v,", temp.next[i].key))
			str.WriteString(fmt.Sprintf("Value:%v", temp.next[i].value))
			str.WriteString("],")
			temp = temp.next[i]
		}
		temp = sl.head
	}
	return str.String()
}

func (node *SkipNode[K, V]) GetValue() V {
	return node.value
}

func (node *SkipNode[K, V]) SetValue(value V) {
	node.value = value
}

func (node *SkipNode[K, V]) String() string {
	if node == nil {
		return "node is nil"
	}
	str := &strings.Builder{}
	str.WriteString("[")
	str.WriteString(fmt.Sprintf("Key:%v,", node.key))
	str.WriteString(fmt.Sprintf("Value:%v", node.value))
	str.WriteString("]:\n")
	str.WriteString("pre:")
	for i, n := range node.pre {
		if n == nil {
			continue
		}
		str.WriteString("[level:")
		str.WriteString(strconv.Itoa(i + 1))
		str.WriteString("]=>[")
		if n.head {
			str.WriteString("head")
		} else {
			str.WriteString(fmt.Sprintf("Key:%v,", n.key))
			str.WriteString(fmt.Sprintf("Value:%v", n.value))
		}
		str.WriteString("],")
	}
	str.WriteString("\nnext:")
	for i, n := range node.next {
		if n == nil {
			continue
		}
		str.WriteString("[level:")
		str.WriteString(strconv.Itoa(i + 1))
		str.WriteString("]=>[")
		if n.tail {
			str.WriteString("tail")
		} else {
			str.WriteString(fmt.Sprintf("Key:%v,", n.key))
			str.WriteString(fmt.Sprintf("Value:%v", n.value))
		}
		str.WriteString("],")
	}
	return str.String()
}

func (iter *SkipListIter[K, V]) Next() bool {
	if iter.head.next[0].tail {
		return false
	}
	iter.head = iter.head.next[0]
	return true
}

func (iter *SkipListIter[K, V]) KV() (K, V) {
	return iter.head.key, iter.head.value
}
