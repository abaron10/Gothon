package gothonMaps

import "github.com/abaron10/Gothon/gothonSlice"

type OrderedMap[K comparable, V any] struct {
	elementsCH chan Pair[K, V]
	sequence   []Pair[K, V]
	raw        map[K]int
	invoked    bool
}

type Pair[K, V any] struct {
	Key   K
	Value V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{sequence: make([]Pair[K, V], 0), raw: make(map[K]int)}
}

func (om *OrderedMap[K, V]) Add(key K, value V) {
	om.raw[key] = len(om.sequence)
	om.sequence = append(om.sequence, Pair[K, V]{Key: key, Value: value})
}

func (om *OrderedMap[K, V]) Get(key K) (val V, exist bool) {
	if index, ok := om.raw[key]; ok {
		return om.sequence[index].Value, true
	}

	return
}

func (om *OrderedMap[K, V]) Delete(key K) (val V, exist bool) {
	if index, ok := om.raw[key]; ok {
		del := om.deleteTrace(index)
		return del.Value, true
	}

	return
}

func (om *OrderedMap[K, V]) PopItem(last bool) (K, V) {
	var index int
	if last {
		index = -1
	}

	del := om.deleteTrace(index)
	return del.Key, del.Value
}

func (om *OrderedMap[K, V]) deleteTrace(index int) Pair[K, V] {
	pair := gothonSlice.Pop(&om.sequence, index)
	delete(om.raw, pair.Key)
	return pair
}

func (om *OrderedMap[K, V]) Items() (Pair[K, V], bool) {

	if !om.invoked {
		om.generator()
		om.invoked = true
	}

	item, ok := <-om.elementsCH
	if !ok {
		om.invoked = false
		return Pair[K, V]{}, false
	}

	return item, true
}

func (om *OrderedMap[K, V]) generator() {
	ch := make(chan Pair[K, V], len(om.sequence))
	for _, item := range om.sequence {
		ch <- item
	}
	close(ch)
	om.elementsCH = ch
}
