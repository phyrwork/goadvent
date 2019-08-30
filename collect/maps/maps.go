package maps // because 'map' is a reserved word

type Map map[interface{}]interface{}

func New(p ...struct{K, V interface{}}) Map {
	m := make(Map, len(p))
	for _, p := range p {
		m[p.K] = p.V
	}
	return m
}

func (m Map) Pair(k interface{}) struct {K, V interface{}} {
	v := m[k]
	return struct{K, V interface{}}{k, v}
}

func (m Map) Inv() Map {
	n := make(Map, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}