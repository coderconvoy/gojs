package gojs

import "testing"

func Test_line(t *testing.T) {
	a := []struct {
		in string
		a  string
		b  string
	}{
		{"nnn\np", "nnn", "p"},
		{"nr\n\rrr", "nr", "\rrr"},
		{"rn1\r\npop", "rn1", "pop"},
		{"rrr\rpop", "rrr", "pop"},
		{"hello", "hello", ""},
	}

	for k, v := range a {
		aa, bb := line(v.in)
		cc := v.in[bb:]
		if aa != v.a || cc != v.b {
			t.Logf("t%d - Expected: %s,%s\n--\nGot: %s,%s", k, v.a, v.b, aa, cc)
			t.Fail()
		}
	}
}
