package bimap

type Channel struct {
	M    Bimap
	// Channel size
	Size int
}

func (d Channel) Values() <-chan interface{} {
	o := make(chan interface{}, d.Size)
	go func() {
		for v := range d.M.k {
			o <- v
		}
		close(o)
	}()
	return o
}

func (d Channel) Keys() <-chan interface{} {
	o := make(chan interface{}, d.Size)
	go func() {
		for k := range d.M.v {
			o <- k
		}
		close(o)
	}()
	return o
}

func (d Channel) Pairs() <-chan struct {K, V interface{}} {
	o := make(chan struct {K, V interface{}}, d.Size)
	go func() {
		for k, v := range d.M.v {
			o <- struct {K, V interface{}}{k, v}
		}
		close(o)
	}()
	return o
}