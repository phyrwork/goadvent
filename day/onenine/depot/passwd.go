package depot

type Rule func (string) bool

func NewLengthRule(n int) Rule {
	return func (s string) bool {
		return len(s) == n
	}
}

func NewAdjacentLeastRule(n int) Rule {
	return func (s string) bool {
		hits := 0
		var last rune
		for _, cur := range []rune(s) {
			if cur == last {
				hits++
			} else {
				last = cur
				hits = 1
			}
			if hits >= n {
				return true
			}
		}
		return false
	}
}

func NewAdjacentExactRule(n int) Rule {
	return func (s string) bool {
		hits := 0
		var last rune
		for _, cur := range []rune(s) {
			if cur == last {
				hits++
			} else {
				last = cur
				if hits == n {
					return true
				}
				hits = 1
			}
		}
		if hits == n {
			return true
		}
		return false
	}
}

func DecreaseRule(s string) bool {
	w := []rune(s)
	if len(w) == 0 {
		return false
	}
	last := w[0]
	for _, cur := range w {
		if cur < last {
			return false
		}
		last = cur
	}
	return true
}

func Filter(c <-chan string, b int, rules ...Rule) <-chan string {
	ok := func (s string) bool {
		for _, rule := range rules {
			if !rule(s) {
				return false
			}
		}
		return true
	}
	o := make(chan string, b)
	go func () {
		for s := range c {
			if ok(s) {
				o <- s
			}
		}
		close(o)
	}()
	return o
}

func NewRange(from, to int, b int) <-chan int {
	o := make(chan int, b)
	go func () {
		for i := from; i <= to; i++ {
			o <- i
		}
		close(o)
	}()
	return o
}