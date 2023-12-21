package orderedmap

type NodeI[K any, V any] interface {
	GetValue() V
	SetValue(V)
	String() string
}
type IterI[K any, V any] interface {
	Next() bool
	KV() (K, V)
}
type ListI[K any, V any] interface {
	Insert(K, V) NodeI[K, V]
	Delete(NodeI[K, V])
	Iter() IterI[K, V]
	String() string
}

type Compare[K any] func(k1, k2 K) int
type CmpValue[V any] func(v1, v2 V) int

type OrderedMap[K comparable, V any] struct {
	kv     map[K]NodeI[K, V]
	l      ListI[K, V]
	cmpVal bool
}

func New[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		kv: make(map[K]NodeI[K, V]),
		l:  NewDoubleList[K, V](),
	}
}

// NewCmp compare key
func NewCmp[K comparable, V any](cmp Compare[K]) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		kv: make(map[K]NodeI[K, V]),
		l:  NewSkipList[K, V](func(k1, k2 K, v1, v2 V) int { return cmp(k1, k2) }),
	}
}

// NewCmpVal compare value
func NewCmpVal[K comparable, V any](cmp CmpValue[V]) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		kv:     make(map[K]NodeI[K, V]),
		l:      NewSkipList[K, V](func(k1, k2 K, v1, v2 V) int { return cmp(v1, v2) }),
		cmpVal: true,
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be the zero value of V.
func (m *OrderedMap[K, V]) Get(key K) (value V, ok bool) {
	var v NodeI[K, V]
	v, ok = m.kv[key]
	if ok {
		value = v.GetValue()
	}
	return
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (m *OrderedMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if value, ok := m.kv[key]; ok {
		return value.GetValue()
	}

	return defaultValue
}

// Set will set (or update) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced.
func (m *OrderedMap[K, V]) Set(key K, value V) {
	node, ok := m.kv[key]
	if ok && !m.cmpVal {
		node.SetValue(value)
		return
	}
	if ok && m.cmpVal {
		m.l.Delete(node)
	}

	m.kv[key] = m.l.Insert(key, value)
	return
}

// Len returns the number of elements in the map.
func (m *OrderedMap[K, V]) Len() int {
	return len(m.kv)
}

// Del will remove a key from the map. It will return true if the key was
// exist,otherwise return false
func (m *OrderedMap[K, V]) Del(key K) {
	node, ok := m.kv[key]
	if ok {
		m.l.Delete(node)
		delete(m.kv, key)
	}
}

// Iter iterate all keys in the map.
func (m *OrderedMap[K, V]) Iter() IterI[K, V] {
	return m.l.Iter()
}

// Keys returns all the keys in the map.
func (m *OrderedMap[K, V]) Keys() (keys []K) {
	keys = make([]K, 0, m.Len())
	for iter := m.l.Iter(); iter.Next(); {
		k, _ := iter.KV()
		keys = append(keys, k)
	}
	return keys
}
