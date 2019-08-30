package maps

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
