# orderedmap
a go ordered map that supports custom sorting rule

## install
```sh
go get github.com/pigfu/orderedmap
```

## Benchmarks

```go
var (
	N = int64(100000)
)

func BenchmarkNew(b *testing.B) {
	m := New[int64, int]()
	for i := 0; i < b.N; i++ {
		m.Set(rand.Int63n(N), i)
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
func BenchmarkNewCmp(b *testing.B) {
	m := NewCmp[int64, int](cmp)
	for i := 0; i < b.N; i++ {
		m.Set(rand.Int63n(N), i)
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
func BenchmarkNewCmpVal(b *testing.B) {
	m := NewCmpVal[int64, int](cmpVal)
	for i := 0; i < b.N; i++ {
		m.Set(rand.Int63n(N), i)
	}
}
```
```bash
goos: windows
goarch: amd64
pkg: github.com/pigfu/orderedmap
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkNew-12                 17014821                71.68 ns/op
BenchmarkNewCmp-12              14796138                76.79 ns/op
BenchmarkNewCmpVal-12            2105310               576.1 ns/op
PASS
ok      github.com/pigfu/orderedmap     6.172s
```
### explain

1. the New function use double linked list,so Set and Del is O(1).
2. the NewCmp function use skip list,so Set is O(logN) (but update is O(1)) and Del is O(1),because every level is double linked list.
3. the NewCmpValue function also use skip list,Set is O(logN) (but update is O(logN),because compare value need firstly delete key and then insert) and Del is O(1).

## example
```go
package main

import (
	. "github.com/pigfu/orderedmap"
	"fmt"
)

var (
	insertKeys = []int{9, 2, 6, 8, 599, 4, 9, 10, 5, 8, 100}
	deleteKeys = []int{5, 8, 9, 7, 999}
)
//key order by insert 
func testNew() {
	m := New[int, int]()
	for i, k := range insertKeys {
		m.Set(k, i)
	}
	for iter := m.Iter(); iter.Next(); {
		fmt.Println(iter.KV())
	}
	fmt.Println("------testNew------")
	for _, k := range deleteKeys {
		m.Del(k)
	}

	for iter := m.Iter(); iter.Next(); {
		fmt.Println(iter.KV())
	}
}

func cmp(k1, k2 int) int {
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}

	return 0
}

// compare with key
func testNewCmp() {
	m := NewCmp[int, int32](cmp)
	for i, k := range insertKeys {
		m.Set(k, int32(i))
	}
	for iter := m.Iter(); iter.Next(); {
		fmt.Println(iter.KV())
	}
	fmt.Println("------testNewCmp------")
	for _, k := range deleteKeys {
		m.Del(k)
	}

	for iter := m.Iter(); iter.Next(); {
		fmt.Println(iter.KV())
	}
}

func cmpVal(v1, v2 int32) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return 1
	}
	return 0
}
//compare with value
func testNewCmpVal() {
	m := NewCmpVal[int, int32](cmpVal)
	for i, k := range insertKeys {
		m.Set(k, int32(i))
	}
	for iter := m.Iter(); iter.Next(); {
		fmt.Println(iter.KV())
	}
	fmt.Println("------testNewCmpVal------")
	for _, k := range deleteKeys {
		m.Del(k)
	}

	for iter := m.Iter(); iter.Next(); {
		fmt.Println(iter.KV())
	}
}

func main() {
	testNew()
	testNewCmp()
	testNewCmpVal()
}
```

