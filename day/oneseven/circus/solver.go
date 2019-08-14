package circus

import "io"

func Base(r io.Reader) (string, error) {
	descs, err := Parse(r)
	if err != nil {
		return "", err
	}
	c, err := New(descs...)
	if err != nil {
		return "", err
	}
	bt, err := c.Base()
	if err != nil {
		return "", err
	}
	return bt.Name, nil
}
