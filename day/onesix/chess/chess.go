package chess

import (
	"crypto/md5"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Generator interface {
	Next() bool
	String() string
	// TODO: should probably add Err() error rather than failing silently with Next() == false
}

type HashFn func (string) string

func HashMd5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s) // fingers crossed it doesn't error
	return fmt.Sprintf("%x", h.Sum(nil))
}

// with very specific requirements...
type HashGenerator struct {
	hash HashFn
	key string
	cur int
	out string
}

func NewHashGenerator(fn HashFn, key string, cur int) *HashGenerator {
	return &HashGenerator{fn, key, cur, ""}
}

func (g *HashGenerator) Next() bool {
	g.cur++
	g.out = g.hash(g.key + strconv.Itoa(g.cur))
	return true
}

func (g *HashGenerator) String() string { return g.out }

type ZeroesGenerator struct {
	sub   Generator
	count int
}

func NewZeroesGenerator(sub Generator, count int) *ZeroesGenerator {
	return &ZeroesGenerator{sub, count}
}

func (g *ZeroesGenerator) Next() bool {
	lz := func (s string) int {
		n := 0
		for _, c := range s {
			if c != '0' {
				break
			}
			n++
		}
		return n
	}
	for g.sub.Next() {
		out := g.sub.String()
		if lz(out) >= g.count {
			return true
		}
	}
	return false
}

func (g *ZeroesGenerator) String() string { return g.sub.String() }

// TODO: rename AppendGenerator
type PasswordGenerator struct {
	zero *ZeroesGenerator
	len  int
	out  string
}

func NewPasswordGenerator(len int, sub Generator, diff int) *PasswordGenerator {
	zero := NewZeroesGenerator(sub, diff)
	return &PasswordGenerator{zero, len, ""}
}

func (g *PasswordGenerator) Next() bool {
	out := make([]rune, g.len)
	for i := range out {
		if !g.zero.Next() {
			// sub generator empty before password complete
			return false
		}
		h := g.zero.String()
		if len(h) < g.zero.count {
			// string not long enough to extract rune from
			return false
		}
		// take first rune after difficulty length
		out[i] = []rune(h)[g.zero.count]
	}
	g.out = string(out)
	return true
}

func (g *PasswordGenerator) String() string { return g.out }

// for lack of a better name. sue me...
// TODO: could probably conflate this with PasswordGenerator by specifying an
//  IndexFn(n int, s string) string
type FillerGenerator struct {
	zero *ZeroesGenerator
	len  int
	out  string
}

func NewFillerGenerator(len int, sub Generator, diff int) *FillerGenerator {
	zero := NewZeroesGenerator(sub, diff)
	return &FillerGenerator{zero, len, ""}
}

func (g *FillerGenerator) Next() bool {
	out := make([]rune, g.len)
	rem := make(map[int]struct{}) // positions left to fill
	for i := range out {
		rem[i] = struct{}{}
	}
	for len(rem) > 0 {
		if !g.zero.Next() {
			// TODO: error?
			return false
		}
		h := g.zero.String()
		if len(h) < g.zero.count + 1 {
			// string not long enough to extract runes from
			return false
		}
		// first rune after zeros now position
		pos := int([]rune(h)[g.zero.count] - '0')
		if _, ok := rem[pos]; !ok {
			// already filled - ignore
			continue
		}
		// fill and mark
		out[pos] = []rune(h)[g.zero.count + 1]
		delete(rem, pos)
	}
	g.out = string(out)
	return true
}

func (g *FillerGenerator) String() string { return g.out }

func SolveAppend(r io.Reader) app.Solution {
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return app.Errorf("reader error: %v", err)
	}
	key := strings.TrimSpace(string(in))
	sub := NewHashGenerator(HashMd5, key, 0)
	gen := NewPasswordGenerator(8, sub, 5)
	if !gen.Next() {
		return app.Errorf("password not generated")
	}
	return app.String(gen.String())
}

func SolveFiller(r io.Reader) app.Solution {
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return app.Errorf("reader error: %v", err)
	}
	key := strings.TrimSpace(string(in))
	sub := NewHashGenerator(HashMd5, key, 0)
	gen := NewFillerGenerator(8, sub, 5)
	if !gen.Next() {
		return app.Errorf("password not generated")
	}
	return app.String(gen.String())
}