package bimap

import "github.com/phyrwork/goadvent/collect/set"

type Set struct {
	M Bimap
}

func (d Set) Values() set.Set {
	o := make(set.Set, d.M.Len())
	for _, v := range d.M.v {
		o[v] = struct{}{}
	}
	return o
}

func (d Set) Keys() set.Set {
	o := make(set.Set, d.M.Len())
	for _, k := range d.M.k {
		o[k] = struct{}{}
	}
	return o
}

func (d Set) Pairs() set.Set {
	o := make(set.Set, d.M.Len())
	for k, v := range d.M.v {
		p := struct{K, V interface{}}{k, v}
		o[p] = struct{}{}
	}
	return o
}
