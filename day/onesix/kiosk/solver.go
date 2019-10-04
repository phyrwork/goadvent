package kiosk

import (
	"fmt"
	"io"
	"strconv"
)

func SolveSumReal(r io.Reader) (string, error) {
	all, err := Read(r)
	if err != nil {
		return "", fmt.Errorf("input error: %v", err)
	}
	val := Validator(AlphaCountHash)
	ok, err := Filter(all, func (r Room) (bool, error) { return val.Valid(r) })
	sum := 0
	for _, r := range ok {
		sum += r.Sector
	}
	return strconv.Itoa(sum), nil
}

func SolveNorthPoleRoom(r io.Reader) (string, error) {
	all, err := Read(r)
	if err != nil {
		return "", fmt.Errorf("input error: %v", err)
	}
	val := Validator(AlphaCountHash)
	ok, err := Filter(all, func (r Room) (bool, error) { return val.Valid(r) })
	for _, r := range ok {
		d := Decrypt(r)
		// yep, this was found by inspection
		if d.Name == "northpole object storage" {
			return strconv.Itoa(d.Sector), nil
		}
	}
	return "", fmt.Errorf("storage room not found")
}

