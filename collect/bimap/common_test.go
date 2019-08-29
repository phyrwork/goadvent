package bimap

import "github.com/phyrwork/goadvent/collect/set"

var delegateTests = map[string]struct{
	k []interface{}
	v []interface{}
	p []struct{K, V interface{}}
}{
	"empty": {
		[]interface{}{},
		[]interface{}{},
		[]struct{K, V interface{}}{},
	},
	"mirror": {
		[]interface{}{1, 2},
		[]interface{}{2, 1},
		[]struct{K, V interface{}}{{1, 2}, {2, 1}},
	},
	"sequence": {
		[]interface{}{1, 2, 3},
		[]interface{}{4, 5, 6},
		[]struct{K, V interface{}}{{1, 4}, {2, 5},  {3, 6}},
	},
}

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
