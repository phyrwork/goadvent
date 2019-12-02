package kiosk

import (
	"github.com/phyrwork/goadvent/app"
	"io"
)

func SolveSumReal(r io.Reader) app.Solution {
	all, err := Read(r)
	if err != nil {
		return app.Errorf("input error: %v", err)
	}
	val := Validator(AlphaCountHash)
	ok, err := Filter(all, func (r Room) (bool, error) { return val.Valid(r) })
	sum := 0
	for _, r := range ok {
		sum += r.Sector
	}
	return app.Int(sum)
}

func SolveNorthPoleRoom(r io.Reader) app.Solution {
	all, err := Read(r)
	if err != nil {
		return app.Errorf("input error: %v", err)
	}
	val := Validator(AlphaCountHash)
	ok, err := Filter(all, func (r Room) (bool, error) { return val.Valid(r) })
	for _, r := range ok {
		d := Decrypt(r)
		// yep, this was found by inspection
		if d.Name == "northpole object storage" {
			return app.Int(d.Sector)
		}
	}
	return app.Errorf("storage room not found")
}

