package sif

type Layer [][]rune

func (l Layer) EachDigit(f func (d rune) error) error {
	for _, r := range l {
		for _, c := range r {
			if err := f(c); err != nil {
				return err
			}
		}
	}
	return nil
}

func (l Layer) String() string {
	s := make([]rune, 0)
	for _, r := range l {
		for _, c := range r {
			s = append(s, '0' + c)
		}
		s = append(s, '\n')
	}
	return string(s[:len(s)-1])
}

type DigitIterator interface {
	EachDigit(func (d rune) error) error
}

func Count(d DigitIterator, f func (d rune) bool) int {
	s := 0
	d.EachDigit(func (d rune) error {
		if f != nil {
			if !f(d) {
				return nil
			}
		}
		s++
		return nil
	})
	return s
}

const (
	Black = 0
	White = 1
	Clear = 2
)

type Composite []Layer

func (i Composite) Flatten() Layer {
	if len(i) == 0 {
		return nil
	}
	o := make(Layer, len(i[0]))
	for z := range i[0] {
		l := make([]rune, len(i[0][z]))
		copy(l, i[0][z])
		o[z] = l
	}
	for y, r := range o {
		for x := range r {
			for z := 1; z < len(i); z++ {
				l := i[z]
				if y > len(l) {
					continue
				}
				if x > len(l[y]) {
					continue
				}
				if o[y][x] != Clear {
					continue
				}
				o[y][x] = l[y][x]
			}
		}
	}
	return o
}