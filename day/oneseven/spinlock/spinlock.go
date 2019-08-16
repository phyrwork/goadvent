package spinlock

import "container/ring"

func NewSequence(i, n int) <-chan int {
	c := make(chan int)
	go func () {
		for n > 0 {
			c <- i
			i++
			n--
		}
		close(c)
	}()
	return c
}

type Spinlock struct {
	curr *ring.Ring
	step int
}

func NewSpinlock(step int) *Spinlock {
	r := ring.New(1)
	r.Value = 0
	return &Spinlock{r, step}
}

func (l *Spinlock) Peek(o int) int { return l.curr.Move(o).Value.(int) }

func (l *Spinlock) Put(v int) {
	curr := l.curr.Move(l.step)
	next := ring.New(1)
	next.Value = v
	curr.Link(next)
	l.curr = next
}

func (l *Spinlock) Stream(c <-chan int) {
	for v := range c {
		l.Put(v)
	}
}