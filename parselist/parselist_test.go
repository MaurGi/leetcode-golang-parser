package parselist

import "testing"

func TestAbs(t *testing.T) {
	test := []string{"[]", "[1,2,3]", "[3,2,1]"}

	for _, v := range test {
		res := SerializeList(ParseList(v))
		if res != v {
			t.Errorf("wanted %s, got %s", v, res)
		}
	}
}
