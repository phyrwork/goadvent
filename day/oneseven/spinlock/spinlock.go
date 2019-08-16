package spinlock

import (
	"container/ring"
	"log"
)

func NewSequence(i, n int) <-chan int {
	c := make(chan int, 20)
	go func () {
		for n > 0 {
			c <- i
			i++
			n--
			if n % 50000 == 0 {
				log.Printf("sequence remaining: %v", n)
			}
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

func (l *Spinlock) Find(v int) (o int, s *Spinlock) {
	if l.curr == nil {
		return -1, nil
	}
	if l.curr.Value.(int) == v {
		return 0, l
	}
	curr, o := l.curr.Next(), 1
	for curr != l.curr {
		if curr.Value.(int) == v {
			return o, &Spinlock{curr, l.step}
		}
		curr, o = curr.Next(), o + 1
	}
	return -1, nil
}

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