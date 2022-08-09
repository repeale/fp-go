package pair

import "testing"

func TestPair(t *testing.T) {
	res := Pair(42, "42")
	if res.a != 42 {
		t.Error("Pair should return a struct with 42 as the first value. Received:", res.a)
	}
	if res.b != "42" {
		t.Error("Pair should return a struct with \"42\" as the second value. Received:", res.b)
	}
}

func TestFst(t *testing.T) {
	res := Fst(Pair(42, "42"))
	if res != 42 {
		t.Error("Fst should return 42. Received:", res)
	}
}

func TestSnd(t *testing.T) {
	res := Snd(Pair(42, "42"))
	if res != "42" {
		t.Error("Fst should return \"42\". Received:", res)
	}
}

func TestGet(t *testing.T) {
	res1, res2 := Get(Pair(42, "42"))
	if res1 != 42 {
		t.Error("Get should return 42 as the first value. Received:", res1)
	}
	if res2 != "42" {
		t.Error("Get should return \"42\" as the second value. Received:", res2)
	}
}
