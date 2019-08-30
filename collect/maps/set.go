package maps

import "github.com/phyrwork/goadvent/collect/set"

type Set struct {
	M Map
}

func (d Set) Values() set.Set {
	o := make(set.Set, len(d.M))
	for _, v := range d.M {
		o[v] = struct{}{}
	}
	return o
}

func (d Set) Keys() set.Set {
	o := make(set.Set, len(d.M))
	for k := range d.M {
		o[k] = struct{}{}
	}
	return o
}

func (d Set) Pairs() set.Set {
	o := make(set.Set, len(d.M))
	for k, v := range d.M {
		p := struct{K, V interface{}}{k, v}
		o[p] = struct{}{}
	}
	return o
}