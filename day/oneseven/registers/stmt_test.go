package registers

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		in string
		want []Stmt
	}{
		"inc gt": {
			"smi inc 781 if epx > -2",
			[]Stmt{
				{
					Op:OpMod{
						Reg: "smi",
						Op:  OpInc,
						Arg: 781,
					},
					Cond:OpCmp{
						Reg: "epx",
						Op:  OpGt,
						Arg: -2,
					},
				},
			},
		},
		"dec ne": {
			"yrf dec -813 if jzm != 6",
			[]Stmt{
				{
					Op:OpMod{
						Reg: "yrf",
						Op:  OpDec,
						Arg: -813,
					},
					Cond:OpCmp{
						Reg: "jzm",
						Op:  OpNeq,
						Arg: 6,
					},
				},
			},
		},
		"dec gte": {
			"c dec -10 if a >= 1",
			[]Stmt{
				{
					Op:OpMod{
						Reg: "c",
						Op:  OpDec,
						Arg: -10,
					},
					Cond:OpCmp{
						Reg: "a",
						Op:  OpGte,
						Arg: 1,
					},
				},
			},
		},
		"inc gt, dec ne": {
			"smi inc 781 if epx > -2\nyrf dec -813 if jzm != 6",
			[]Stmt{
				{
					Op:OpMod{
						Reg: "smi",
						Op:  OpInc,
						Arg: 781,
					},
					Cond:OpCmp{
						Reg: "epx",
						Op:  OpGt,
						Arg: -2,
					},
				},
				{
					Op:OpMod{
						Reg: "yrf",
						Op:  OpDec,
						Arg: -813,
					},
					Cond:OpCmp{
						Reg: "jzm",
						Op:  OpNeq,
						Arg: 6,
					},
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Parse(strings.NewReader(test.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected stmts: want %v, got %v", test.want, got)
			}
		})
	}
}

