package kiosk

import (
	"fmt"
	"sort"
)

type Hash [5]rune

func StringToHash(s string) (Hash, error) {
	var chk Hash
	if len(s) != len(chk) {
		return chk, UnhashableError(fmt.Sprintf("invalid hash: %v: must be len %v", s, len(chk)))
	}
	copy(chk[:], []rune(s))
	return chk, nil
}

type HashFn func (string) (Hash, error)

type UnhashableError string

func (e UnhashableError) Error() string { return string(e) }

type alphaCountItem struct {
	alpha rune
	count int
}

type alphaCountOrder []alphaCountItem

func (o alphaCountOrder) Len() int { return len(o) }

func (o alphaCountOrder) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

func (o alphaCountOrder) Less(i, j int) bool {
	a, b := o[i], o[j]
	if a.count > b.count {
		return true
	}
	if a.count < b.count {
		return false
	}
	// a.count == b.count
	return a.alpha < b.alpha
}

func AlphaCountHash(s string) (Hash, error) {
	m := make(map[rune]int)
	for _, c := range s {
		m[c] = m[c] + 1
	}
	o := make(alphaCountOrder, 0, len(m))
	for c, n := range m {
		if c >= 'a' && c <= 'z' {
			o = append(o, alphaCountItem{c, n})
		}
	}
	var chk Hash
	if len(o) < len(chk) {
		return chk, UnhashableError(fmt.Sprintf("input not hashable: %v: not enough unique runes: %v", s, len(o)))
	}
	sort.Sort(o)
	for i := range chk {
		chk[i] = o[i].alpha
	}
	return chk, nil
}
