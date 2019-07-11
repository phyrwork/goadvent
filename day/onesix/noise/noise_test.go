package noise

import "testing"

func TestDecodeMode(t *testing.T) {
	in := []string{
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
	r := NewSlice(in...)
	want := "easter"
	got, err := DecodeMode(r)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("unexpected value: want %v, got %v", want, got)
	}
}

