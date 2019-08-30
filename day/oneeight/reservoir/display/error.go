package display

import (
	"fmt"
	"github.com/phyrwork/goadvent/vector"
)

type OutOfBoundsError struct {
	R vector.Range
	V vector.Vector
}

func (e OutOfBoundsError) Error() string {
	return fmt.Sprintf("out of bounds: %v in %v", e.V, e.R)
}
