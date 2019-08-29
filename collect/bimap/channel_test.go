package bimap

import (
	"github.com/phyrwork/goadvent/collect/set"
	"reflect"
	"testing"
)

func TestChannel_Values(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			want := setFromSlice(test.v...)
			m := bimapFromSlice(test.p...)
			d := Channel{M: m}
			got := make(set.Set, len(test.v))
			for v := range d.Values() {
				got[v] = struct{}{}
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected values: want %v, got %v", want, got)
			}
		})
	}
}

