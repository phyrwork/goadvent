package deduct

const SequenceLen = 5

type Sequence [SequenceLen]int

func NewSequence(i int) Sequence {
	return Sequence{
		(i / 1) % 10,
		(i / 10) % 10,
		(i / 100) % 10,
		(i / 1000) % 10,
		(i / 10000) % 10,
	}
}

func (b Sequence) String() string {
	var s [SequenceLen]rune
	for p := range s {
		s[p] = '0' + rune(b[p])
	}
	return string(s[:])
}

func (b Sequence) c(a Sequence) (int, [SequenceLen]bool) {
	var m [SequenceLen]bool
	c := 0
	for p := 0; p < len(a); p++ {
		if b[p] == a[p] {
			c++
			m[p] = true
		}
	}
	return c, m
}

func (b Sequence) C(a Sequence) int {
	c, _ := b.c(a)
	return c
}

func (b Sequence) CD(a Sequence) (int, int) {
	c, f := b.c(a)
	an, bn := make(map[int]int), make(map[int]int)
	for p, m := range f {
		if m {
			// C, don't count
			continue
		}
		ai, bi := a[p], b[p]
		an[ai], bn[bi] = an[ai] + 1, bn[bi] + 1
	}
	d := 0
	for i := range bn {
		ain, bin := an[i], bn[i]
		if ain < bin {
			d += ain
		} else {
			d += bin
		}
	}
	return c, d
}

func (b Sequence) D(a Sequence) int {
	_, d := b.CD(a)
	return d
}

type SequenceSet map[Sequence]struct{}

func DefaultSequenceSet() SequenceSet {
	m := make(SequenceSet, 9000)
	for a := 10000; a <= 99999; a++ {
		m[NewSequence(a)] = struct{}{}
	}
	return m
}

type Entry struct {
	B Sequence // Symbol sequence
	C int      // Position matching symbols
	D int      // Non-position matching symbols (after position matched are c)
}

func Solve(s SequenceSet, g []Entry) SequenceSet {
	o := make(SequenceSet)
	for a := range s {
		ok := true
		for _, e := range g {
			c, d := a.CD(e.B)
			if c != e.C || d != e.D {
				ok = false
				break
			}
		}
		if ok {
			o[a] = struct{}{}
		}
	}
	return o
}

