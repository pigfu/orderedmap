package orderedmap

import (
	"math/rand"
	"testing"
)

var (
	insertKeys = []int{9, 2, 6, 8, 599, 4, 9, 10, 5, 8, 100}
	deleteKeys = []int{5, 8, 9, 7, 999}
)

func cmpKV(k1, k2 int, v1, v2 int32) int {
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}

	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return 1
	}
	return 0
}

func TestSkipList(t *testing.T) {
	skipList := NewSkipList[int, int32](cmpKV)
	skipMap := make(map[int]NodeI[int, int32])
	for i, k := range insertKeys {
		skipMap[k] = skipList.Insert(k, int32(i))
	}
	t.Log(skipList)

	for iter := skipList.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}

	for _, k := range deleteKeys {
		if skipMap[k] == nil {
			continue
		}
		skipList.Delete(skipMap[k])
	}
	t.Log(skipList)
}

func TestDoubleList(t *testing.T) {
	dlList := NewDoubleList[int, int]()
	dlMap := make(map[int]NodeI[int, int])
	for i, k := range insertKeys {
		dlMap[k] = dlList.Insert(k, i)
	}
	t.Log(dlList)

	for iter := dlList.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}

	for _, k := range deleteKeys {
		if dlMap[k] == nil {
			continue
		}
		dlList.Delete(dlMap[k])
	}
	t.Log(dlList)
}

func TestNew(t *testing.T) {
	m := New[int, int]()
	for i, k := range insertKeys {
		m.Set(k, i)
	}
	for iter := m.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}
	t.Log("------------")
	for _, k := range deleteKeys {
		m.Del(k)
	}

	for iter := m.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}
}

func cmp(k1, k2 int64) int {
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}

	return 0
}
func TestNewCmp(t *testing.T) {
	m := NewCmp[int64, int](cmp)
	for i, k := range insertKeys {
		m.Set(int64(k), i)
	}
	for iter := m.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}
	t.Log("------------")
	for _, k := range deleteKeys {
		m.Del(int64(k))
	}

	for iter := m.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}
}

func cmpVal(v1, v2 int) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return 1
	}
	return 0
}
func TestNewCmpVal(t *testing.T) {
	m := NewCmpVal[int64, int](cmpVal)
	for i, k := range insertKeys {
		m.Set(int64(k), i)
	}
	for iter := m.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}
	t.Log("------------")
	for _, k := range deleteKeys {
		m.Del(int64(k))
	}

	for iter := m.Iter(); iter.Next(); {
		t.Log(iter.KV())
	}
}

var (
	N = int64(100000)
)

func BenchmarkNew(b *testing.B) {
	m := New[int64, int]()
	for i := 0; i < b.N; i++ {
		m.Set(rand.Int63n(N), i)
	}
}

func BenchmarkNewCmp(b *testing.B) {
	m := NewCmp[int64, int](cmp)
	for i := 0; i < b.N; i++ {
		m.Set(rand.Int63n(N), i)
	}
}

func BenchmarkNewCmpVal(b *testing.B) {
	m := NewCmpVal[int64, int](cmpVal)
	for i := 0; i < b.N; i++ {
		m.Set(rand.Int63n(N), i)
	}
}
