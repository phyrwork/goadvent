package abba

type Address string

type Sequence [4]rune

func (s Sequence) Abba() bool {
	return s[0] == s[3] && s[1] == s[2]
}

type Sequencer struct {
	addr Address
	cur  int
	seq  Sequence
}

func NewSequencer(addr Address) *Sequencer {
	return &Sequencer{
		addr: addr,
		cur:  0,
		seq:  Sequence{},
	}
}

func (s *Sequencer) Next() bool {
	for s.cur + len(s.seq) <= len(s.addr) {
		// extract sequence
		w := []rune(string(s.addr))[s.cur:s.cur+4]
		s.cur++
		// can't span hypernet boundary
		span := func (w []rune) bool {
			for _, c := range w {
				switch c {
				case '[', ']':
					return true
				}
			}
			return false
		}
		if span(w) {
			continue
		}
		// yield
		copy(s.seq[:], w)
		return true
	}
	return false
}

func (s *Sequencer) Sequence() Sequence { return s.seq }