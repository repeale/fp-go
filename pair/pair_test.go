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

func TestMapFst(t *testing.T) {
	res := MapFst[string](func(x int) int { return x + 1 })(Pair(41, "42"))
	if Fst(res) != 42 {
		t.Error("MapFirst should return 42 as a first value. Received:", Fst(res))
	}
	if Snd(res) != "42" {
		t.Error("MapFst should leave the second value \"42\" alone. Received:", Snd(res))
	}
}

func TestMapSnd(t *testing.T) {
	res := MapSnd[string](func(x int) int { return x + 1 })(Pair("42", 41))
	if Fst(res) != "42" {
		t.Error("MapSnd should leave the first value \"42\" alone. Received:", Fst(res))
	}
	if Snd(res) != 42 {
		t.Error("MapFirst should return 42 as a second value. Received:", Snd(res))
	}
}

func TestMapBoth(t *testing.T) {
	res := MapBoth(
		func(x bool) bool { return !x },
		func(x int) int { return x + 1 },
	)(Pair(false, 41))

	if Fst(res) != true {
		t.Error("MapBoth should return true as the first value. Received:", Fst(res))
	}
	if Snd(res) != 42 {
		t.Error("MapBoth should return 42 as the second value. Received:", Snd(res))
	}
}

func TestCheckFst_True(t *testing.T) {
	res := CheckFst[string](func(x int) bool { return x > 10 })(Pair(42, "42"))
	if res != true {
		t.Error("CheckFst should return true. Received:", res)
	}
}

func TestCheckFst_False(t *testing.T) {
	res := CheckFst[string](func(x int) bool { return x < 10 })(Pair(42, "42"))
	if res != false {
		t.Error("CheckFst should return false. Received:", res)
	}
}

func TestCheckSnd_True(t *testing.T) {
	res := CheckSnd[int](func(x string) bool { return x == "42" })(Pair(42, "42"))
	if res != true {
		t.Error("CheckFst should return true. Received:", res)
	}
}

func TestCheckSnd_False(t *testing.T) {
	res := CheckSnd[int](func(x string) bool { return x == "1" })(Pair(42, "42"))
	if res != false {
		t.Error("CheckFst should return false. Received:", res)
	}
}

func TestCheckBoth_True_True(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x > 10 },
		func(x string) bool { return x == "42" },
	)(Pair(42, "42"))

	if res != true {
		t.Error("CheckBoth should return true. Received:", res)
	}
}

func TestCheckBoth_True_False(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x > 10 },
		func(x string) bool { return x != "42" },
	)(Pair(42, "42"))

	if res != false {
		t.Error("CheckBoth should return false. Received:", res)
	}
}

func TestCheckBoth_False_True(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x < 10 },
		func(x string) bool { return x == "42" },
	)(Pair(42, "42"))

	if res != false {
		t.Error("CheckBoth should return false. Received:", res)
	}
}

func TestCheckBoth_False_False(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x < 10 },
		func(x string) bool { return x != "42" },
	)(Pair(42, "42"))

	if res != false {
		t.Error("CheckBoth should return false. Received:", res)
	}
}
