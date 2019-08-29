package bimap

import (
	"fmt"
	"github.com/phyrwork/goadvent/collect/set"
)

type Bimap struct {
	v map[interface{}]interface{} // key -> val
	k map[interface{}]interface{} // val -> key
}

func NewBimap() *Bimap {
	v := make(map[interface{}]interface{})
	k := make(map[interface{}]interface{})
	return &Bimap{v, k}
}

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

func (m Bimap) Values() set.Set {
	n := make(set.Set)
	for _, v := range m.v {
		n[v] = struct{}{}
	}
	return n
}

func (m Bimap) Key(v interface{}) (interface{}, bool) {
	k, ok := m.k[v]
	return k, ok

}

func (m Bimap) Keys() set.Set {
	n := make(set.Set)
	for _, k := range m.k {
		n[k] = struct{}{}
	}
	return n
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

func (m Bimap) Pairs() []struct {K, V interface{}} {
	p := make([]struct {K, V interface{}}, 0, len(m.v))
	for k, v := range m.v {
		p = append(p, struct{K, V interface{}}{k, v})
	}
	return p
}

func (m Bimap) Inv() *Bimap { return &Bimap{m.k, m.v} }