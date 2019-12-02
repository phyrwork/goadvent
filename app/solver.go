package app

import (
	"fmt"
	"io"
	"strconv"
)

type Solution interface {
	String() string
	IsError() bool
}


type Solver interface {
	Solve(rd io.Reader) Solution
}


type SolverFunc func (io.Reader) Solution

func (f SolverFunc) Solve(r io.Reader) Solution { return f(r) }


type Int int

func (i Int) String() string { return strconv.Itoa(int(i)) }

func (i Int) IsError() bool { return false }


type String string

func (s String) String() string { return string(s) }

func (s String) IsError() bool { return false }


type Rune rune

func (r Rune) String() string { return string([]rune{rune(r)}) }

func (r Rune) IsError() bool { return false }


type Error string

func (e Error) String() string { return string(e) }

func (e Error) IsError() bool { return true }


func NewError(v interface{}) Error { return Error(fmt.Sprint(v)) }

func Errorf(format string, val ...interface{}) Error {
	return Error(fmt.Sprintf(format, val...))
}