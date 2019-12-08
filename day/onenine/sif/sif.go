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

type Layers []Layer

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