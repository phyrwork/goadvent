package noise

import "testing"

var example = []string{
	"eedadn",
	"drvtee",
	"eandsr",
	"raavrd",
	"atevrs",
	"tsrnev",
	"sdttsa",
	"rasrtv",
	"nssdts",
	"ntnada",
	"svetve",
	"tesnvt",
	"vntsnd",
	"vrdear",
	"dvrsen",
	"enarar",
}

func TestDecoders(t *testing.T) {
	tests := map[string]struct {
		in     []string
		decode WordDecoder
		want   string
	}{
		"mode": {example, NewColumnDecoder(Mode), "easter"},
		"inv mode": {example, NewColumnDecoder(InvMode), "advent"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := NewSlice(test.in...)
			ans := test.decode(r)
			if ans.IsError() {
				t.Errorf("unexpected error: %v", ans)
			}
			if got := ans.String(); got != test.want {
				t.Errorf("unexpected value: want %v, got %v", test.want, got)
			}
		})
	}
}

