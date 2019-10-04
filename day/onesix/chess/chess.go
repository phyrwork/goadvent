package chess

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Generator interface {
	Next() bool
	String() string
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

func Solve(r io.Reader) (string, error) {
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("reader error: %v, err")
	}
	key := strings.TrimSpace(string(in))
	sub := NewHashGenerator(HashMd5, key, 0)
	gen := NewPasswordGenerator(8, sub, 5)
	if !gen.Next() {
		return "", fmt.Errorf("password not generated")
	}
	return gen.String(), nil
}