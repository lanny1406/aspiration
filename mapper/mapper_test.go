package mapper

import (
	"testing"
)

func TestUpperEveryN(t *testing.T) {
	type test struct {
		input string
		pos   int
		want  string
		err   string
	}

	tests := []test{
		{"Aspiration.com", 100, "aspiration.com", "not expecting an error"},
		{"Aspiration.com", 3, "asPirAtiOn.cOm", "not expecting an error"},
		{"Aspiration.com", 1, "ASPIRATION.COM", "not expecting an error"},
		{"a..a", 1, "A..A", "not expecting an error"},
		{"Aspiration.com", 0, "aspiration.com", "invalid index"},
		{"Aspiration.com", -1, "aspiration.com", "invalid index"},
		{"......a", 1, "......A", "not expecting an error"},
		{"......A", 1, "......A", "not expecting an error"},
		{"......A", 2, "......a", "not expecting an error"},
		{".#Ea", 2, ".#eA", "not expecting an error"},
	}

	for i, tc := range tests {
		r, e := upperEveryNChars(tc.pos, tc.input)

		if e != nil && e.Error() != tc.err {
			t.Fatalf("input %v post %v error %v", tc.input, tc.pos, e)
		}

		if string(r) != tc.want {
			t.Fatalf("expected %v got %v", tc.want, string(r))
		}

		t.Logf("test table index %v input %v pos %v result %v", i, tc.input, tc.pos, string(r))
	}
}
