package bimap

import "github.com/phyrwork/goadvent/collect/set"

func bimapFromSlice(p ...struct{K, V interface{}}) Bimap {
	v := make(map[interface{}]interface{}, len(p))
	k := make(map[interface{}]interface{}, len(p))
	for _, p := range p {
		v[p.K], k[p.V] = p.V, p.K
	}
	return Bimap{v, k}
}

func setFromSlice(k ...interface{}) set.Set {
	s := make(set.Set, len(k))
	for _, k := range k {
		s[k] = struct{}{}
	}
	return s
}
