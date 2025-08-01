package xe

import "sync"

type list[K comparable, V any] struct {
	data sync.Map
}

func newList[K comparable, V any]() *list[K, V] {
	return &list[K, V]{}
}

func (t *list[K, V]) Set(key K, val V) {
	t.data.Store(key, val)
}

func (t *list[K, V]) Get(key K) (v V, ok bool) {
	var a any
	a, ok = t.data.Load(key)
	if ok {
		v = a.(V)
	}
	return
}

func (t *list[K, V]) Del(key K) {
	t.data.Delete(key)
}

func (t *list[K, V]) LoadAndDel(key K) (v V, ok bool) {
	var a any
	a, ok = t.data.LoadAndDelete(key)
	if ok {
		v = a.(V)
	}
	return
}
