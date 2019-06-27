package checksum

type OrderedValues struct {
	v []int
	i int
}

func NewOrderedValues(v ...int) *OrderedValues { return &OrderedValues{v, -1} }

func (it *OrderedValues) Reset() error {
	it.i = -1
	return nil
}

func (it *OrderedValues) Next() bool {
	it.i += 1
	return it.i < len(it.v)
}

func (it *OrderedValues) Int() int { return it.v[it.i] }

func (it *OrderedValues) Err() error { return nil }

type OrderedRows struct {
	r []Values
	i int
}

func NewOrderedRows(r ...Values) *OrderedRows { return &OrderedRows{r, -1} }

func (it *OrderedRows) Reset() error {
	it.i = -1
	return nil
}

func (it *OrderedRows) Next() bool {
	it.i += 1
	return it.i < len(it.r)
}

func (it *OrderedRows) Row() Values { return it.r[it.i] }

func (it *OrderedRows) Err() error { return nil }
