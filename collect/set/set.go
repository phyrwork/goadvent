package set

type Set map[interface{}]struct{}

func (s Set) Array() []interface{} {
	a := make([]interface{}, 0, len(s))
	for v := range s {
		a = append(a, v)
	}
	return a
}
