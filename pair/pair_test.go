package pair

import "testing"

func eq[A, B comparable](ps1, ps2 []Pair[A, B]) bool {
	if len(ps1) != len(ps2) {
		return false
	}

	for i := 0; i < len(ps1); i++ {
		if !Eq(ps1[i])(ps2[i]) {
			return false
		}
	}

	return true
}

func TestNew(t *testing.T) {
	res := New(42, "42")
	if res.a != 42 {
		t.Error("New should return a struct with 42 as the first value. Received:", res.a)
	}
	if res.b != "42" {
		t.Error("New should return a struct with \"42\" as the second value. Received:", res.b)
	}
}

func TestFst(t *testing.T) {
	res := Fst(New(42, "42"))
	if res != 42 {
		t.Error("Fst should return 42. Received:", res)
	}
}

func TestSnd(t *testing.T) {
	res := Snd(New(42, "42"))
	if res != "42" {
		t.Error("Fst should return \"42\". Received:", res)
	}
}

func TestGet(t *testing.T) {
	res1, res2 := Get(New(42, "42"))
	if res1 != 42 {
		t.Error("Get should return 42 as the first value. Received:", res1)
	}
	if res2 != "42" {
		t.Error("Get should return \"42\" as the second value. Received:", res2)
	}
}

func TestMapFst(t *testing.T) {
	res := MapFst[string](func(x int) int { return x + 1 })(New(41, "42"))
	if Fst(res) != 42 {
		t.Error("MapFirst should return 42 as a first value. Received:", Fst(res))
	}
	if Snd(res) != "42" {
		t.Error("MapFst should leave the second value \"42\" alone. Received:", Snd(res))
	}
}

func TestMapSnd(t *testing.T) {
	res := MapSnd[string](func(x int) int { return x + 1 })(New("42", 41))
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
	)(New(false, 41))

	if Fst(res) != true {
		t.Error("MapBoth should return true as the first value. Received:", Fst(res))
	}
	if Snd(res) != 42 {
		t.Error("MapBoth should return 42 as the second value. Received:", Snd(res))
	}
}

func TestCheckFst_True(t *testing.T) {
	res := CheckFst[string](func(x int) bool { return x > 10 })(New(42, "42"))
	if res != true {
		t.Error("CheckFst should return true. Received:", res)
	}
}

func TestCheckFst_False(t *testing.T) {
	res := CheckFst[string](func(x int) bool { return x < 10 })(New(42, "42"))
	if res != false {
		t.Error("CheckFst should return false. Received:", res)
	}
}

func TestCheckSnd_True(t *testing.T) {
	res := CheckSnd[int](func(x string) bool { return x == "42" })(New(42, "42"))
	if res != true {
		t.Error("CheckFst should return true. Received:", res)
	}
}

func TestCheckSnd_False(t *testing.T) {
	res := CheckSnd[int](func(x string) bool { return x == "1" })(New(42, "42"))
	if res != false {
		t.Error("CheckFst should return false. Received:", res)
	}
}

func TestCheckBoth_True_True(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x > 10 },
		func(x string) bool { return x == "42" },
	)(New(42, "42"))

	if res != true {
		t.Error("CheckBoth should return true. Received:", res)
	}
}

func TestCheckBoth_True_False(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x > 10 },
		func(x string) bool { return x != "42" },
	)(New(42, "42"))

	if res != false {
		t.Error("CheckBoth should return false. Received:", res)
	}
}

func TestCheckBoth_False_True(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x < 10 },
		func(x string) bool { return x == "42" },
	)(New(42, "42"))

	if res != false {
		t.Error("CheckBoth should return false. Received:", res)
	}
}

func TestCheckBoth_False_False(t *testing.T) {
	res := CheckBoth(
		func(x int) bool { return x < 10 },
		func(x string) bool { return x != "42" },
	)(New(42, "42"))

	if res != false {
		t.Error("CheckBoth should return false. Received:", res)
	}
}

func TestMergeC(t *testing.T) {
	mul := func(a int) func(int) int { return func(b int) int { return a * b } }

	res := MergeC(mul)(New(21, 2))
	if res != 42 {
		t.Error("MergeC should return 42. Received:", res)
	}
}

func TestMerge(t *testing.T) {
	mul := func(a, b int) int { return a * b }

	res := Merge(mul)(New(21, 2))
	if res != 42 {
		t.Error("Merge should return 42. Received:", res)
	}
}

func TestZip_EqLen(t *testing.T) {
	res := Zip[bool]([]int{1, 2, 3, 4, 5})([]bool{true, true, false, false, true})
	want := []Pair[int, bool]{New(1, true), New(2, true), New(3, false), New(4, false), New(5, true)}
	if !eq(res, want) {
		t.Error("Zip should have returned", want, ". Received:", res)
	}
}

func TestZip_Large_Fst(t *testing.T) {
	res := Zip[bool]([]int{1, 2, 3, 4, 5})([]bool{true, true, false})
	want := []Pair[int, bool]{New(1, true), New(2, true), New(3, false)}
	if !eq(res, want) {
		t.Error("Zip should have returned", want, ". Received:", res)
	}
}

func TestZip_Large_Snd(t *testing.T) {
	res := Zip[bool]([]int{1, 2, 3})([]bool{true, true, false, false, true})
	want := []Pair[int, bool]{New(1, true), New(2, true), New(3, false)}
	if !eq(res, want) {
		t.Error("Zip should have returned", want, ". Received:", res)
	}
}

func TestEq_True_True(t *testing.T) {
	res := Eq(New(42, 21))(New(42, 21))
	if res != true {
		t.Error("Eq should return true. Received:", res)
	}
}

func TestEq_True_False(t *testing.T) {
	res := Eq(New(42, 21))(New(42, 100))
	if res != false {
		t.Error("Eq should return false. Received:", res)
	}
}

func TestEq_False_True(t *testing.T) {
	res := Eq(New(21, 42))(New(100, 42))
	if res != false {
		t.Error("Eq should return false. Received:", res)
	}
}

func TestEq_False_False(t *testing.T) {
	res := Eq(New(42, 21))(New(84, 42))
	if res != false {
		t.Error("Eq should return false. Received:", res)
	}
}
