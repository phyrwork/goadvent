package kiosk

type RuneShifter interface {
	Shift(w []rune) []rune
	// TODO: should have a 'Idem() bool' method to signal that the
	//  shift is copy-on-write
}

type LinearShifter struct {
	Count int
}

func (r LinearShifter) Shift(w []rune) []rune {
	for i := range w {
		w[i] += rune(r.Count)
	}
	return w
}

type ModShifter struct {
	Count int
	Mod int
}

func (r ModShifter) Shift(w []rune) []rune {
	w = LinearShifter{r.Count}.Shift(w)
	for i := range w {
		w[i] %= rune(r.Mod)
	}
	return w
}

type WrapShifter struct {
	Count int
	Base rune
	Lim rune
}

func (r WrapShifter) Shift(w []rune) []rune {
	// shift back so that we can mod between base and lim
	w = LinearShifter{int(-r.Base)}.Shift(w)
	// shift and mod
	w = ModShifter{r.Count, int(r.Lim - r.Base + 1)}.Shift(w)
	// shift forwards to restore original base
	return LinearShifter{int(r.Base)}.Shift(w)
}

type CopyShifter struct {
	Shifter RuneShifter
}

func (r CopyShifter) Shift(w []rune) []rune {
	x := make([]rune, len(w))
	copy(x, w)
	x = r.Shifter.Shift(x)
	return x
}

type MaskShifter struct {
	Shifter RuneShifter
	Mask MaskFn
}

func (r MaskShifter) Shift(w []rune) []rune {
	x := CopyShifter{r.Shifter}.Shift(w)
	// Replace runes in x based on mask of w
	mask := func (i int, _ rune) bool { return r.Mask(i, w[i]) }
	ReplaceRunes(x, mask, func (i int, _ rune) rune { return w[i] })
	return x
}