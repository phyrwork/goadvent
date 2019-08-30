package set

type Set map[interface{}]struct{}

func New(v ...interface{}) Set {
	s := make(Set, len(v))
	for _, v := range v {
		s[v] = struct{}{}
	}
	return s
}

func (s Set) Array() []interface{} {
	a := make([]interface{}, 0, len(s))
	for v := range s {
		a = append(a, v)
	}
	return a
}
