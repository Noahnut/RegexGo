package regexgo

import "testing"

func Test_infix2Post(t *testing.T) {
	r := regexgo{}
	result := r.infix2Post("a(bb)+a")

	if result != "abb.+.a." {
		t.Error(result, "abb.+.a.")
	}

	result = r.infix2Post("ab|c")

	if result != "ab.c|" {
		t.Error(result, "ab.c|")
	}
}
