package explosive

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	MarkerOpen = '('
	MarkerClose = ')'
)

type Marker struct {
	len int
	count int
}

func (m Marker) String() string { return fmt.Sprintf("(%vx%v)", m.len, m.count) }

var MarkerRegexp = regexp.MustCompile(`\((\d+)x(\d+)\)`)

func DecodeMarker(s string) (Marker, error) {
	m := Marker{}
	re := MarkerRegexp.FindStringSubmatch(s)
	if len(re) != 3 {
		return m, fmt.Errorf("not a marker")
	}
	var err error
	if m.len, err = strconv.Atoi(re[1]); err != nil {
		return m, fmt.Errorf("int decode error: %v", err)
	}
	if m.count, err = strconv.Atoi(re[2]); err != nil {
		return m, fmt.Errorf("int decode error: %v", err)
	}
	return m, nil
}

type Tokenizer struct {
	sc *bufio.Scanner
	count int
}

func NewTokenizer(r io.Reader) *Tokenizer {
	sc := bufio.NewScanner(r)
	t := &Tokenizer{sc, 0}
	sc.Split(t.ScanTokens)
	return t
}

func (t *Tokenizer) Next() bool { return t.sc.Scan() }

func (t *Tokenizer) Text() string { return t.sc.Text() }

func (t *Tokenizer) Count() int { return t.count }

func (t *Tokenizer) ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// find start of marker
	if s := bytes.IndexRune(data, MarkerOpen); s > 0 {
		// follows plain text
		t.count = 1
		return s, data[:s], nil
	} else if s == 0 {
		// find end of marker
		if e := bytes.IndexRune(data, MarkerClose); e >= 0 {
			m, err := DecodeMarker(string(data[:e+1]))
			if err != nil {
				return 0, nil, fmt.Errorf("marker decode error: %v", err)
			}
			// consume to end of marked word
			i := e + 1
			j := i + m.len
			if j <= len(data) {
				// marked word is completely present, emit
				t.count = m.count
				return j, data[i:j], nil
			}
			if atEOF {
				return 0, nil, fmt.Errorf("malformed input: marked word extends past eof")
			}
			// fallthrough to request more data
		}
	}
	if atEOF && len(data) > 0 {
		// unmarked final data
		t.count = 1
		// ignore trailing newline
		// TODO: this isn't very robust
		if data[len(data)-1] == '\n' {
			return len(data), data[:len(data)-1], nil
		}
		return len(data), data, nil
	}
	// request more data
	return 0, nil, nil
}

func Inflate(r io.Reader, w io.Writer, rec bool) (int, error)  {
	e := true
	if w == nil {
		e = false
	} else {
		e = !reflect.ValueOf(w).IsNil()
	}
	l := 0
	t := NewTokenizer(r)
	for t.Next() {
		n := t.Count()
		s := t.Text()
		var c []byte
		if w != nil {
			c = []byte(s)
		}
		m := len(s)
		// recursively inflate
		if rec && n > 1 {
			r := strings.NewReader(s)
			var b *bytes.Buffer
			if e {
				b = &bytes.Buffer{}
			}
			var err error
			m, err = Inflate(r, b, true)
			if err != nil {
				return 0, fmt.Errorf("subinflate error: %v", err)
			}
			if b != nil {
				c = b.Bytes()
			}
		}
		// output
		for i := 0; i < n; i++ {
			l += m
			if e {
				if _, err := w.Write(c); err != nil {
					return 0, fmt.Errorf("write error: %v", err)
				}
			}
		}
	}
	return l, nil
}

func NewSolver(rec bool) app.SolverFunc {
	return func (r io.Reader) app.Solution {
		l, err := Inflate(r, nil, rec)
		if err != nil {
			return app.Errorf("inflate error: %v", err)
		}
		return app.Int(l)
	}
}

