package bimap

import (
	"fmt"
)

type Bimap struct {
	v map[interface{}]interface{} // key -> val
	k map[interface{}]interface{} // val -> key
}

func New() *Bimap {
	v := make(map[interface{}]interface{})
	k := make(map[interface{}]interface{})
	return &Bimap{v, k}
}

func (m Bimap) Len() int { return len(m.v) }

func (m Bimap) Set(k, v interface{}) error {
	if e, ok := m.v[k]; ok && e != v {
		return fmt.Errorf("key %#v exists: %#v", k, v)
	}
	m.v[k], m.k[v] = v, k
	return nil
}

func (m Bimap) Clear(k interface{}) {
	if v, ok := m.v[k]; ok {
		delete(m.v, k)
		delete(m.k, v)
	}
}

func (m Bimap) Value(k interface{}) (interface{}, bool) {
	v, ok := m.v[k]
	return v, ok
}

func (m Bimap) Key(v interface{}) (interface{}, bool) {
	k, ok := m.k[v]
	return k, ok

}

func (m Bimap) Map() map[interface{}]interface{} {
	n := make(map[interface{}]interface{})
	for k, v := range m.v {
		n[k] = v
	}
	return n
}

func (m Bimap) Pair(k interface{}) (struct {K, V interface{}}, bool) {
	v, ok := m.v[k]
	return struct {K, V interface{}}{k, v}, ok
}

func (m Bimap) Inv() *Bimap { return &Bimap{m.k, m.v} }