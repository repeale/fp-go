package fp

import (
	"reflect"
	"testing"
)

func TestCompose2_Example(t *testing.T) {
	res := Compose2(
		Map(func(x int) int { return x + 2 }),
		Filter(func(x int) bool { return x > 0 }))([]int{1, 2, 3, -1})
	expected := []int{3, 4, 5}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("Should perform right-to-left function composition of two functions. Received:", res)
	}
}

func TestCompose3_Example(t *testing.T) {
	res := Compose3(
		Map(func(x int) int { return x + 2 }),
		Map(func(x int) int { return x + 2 }),
		Filter(func(x int) bool { return x > 0 }))([]int{1, 2, 3, -1})
	expected := []int{5, 6, 7}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("Should perform right-to-left function composition of two functions. Received:", res)
	}
}

func TestCompose4_Example(t *testing.T) {
	res := Compose4(
		Map(func(x int) int { return x + 2 }),
		Map(func(x int) int { return x + 2 }),
		Map(func(x int) int { return x + 2 }),
		Filter(func(x int) bool { return x > 0 }))([]int{1, 2, 3, -1})
	expected := []int{7, 8, 9}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("Should perform right-to-left function composition of two functions. Received:", res)
	}
}

/* Generated with the following python code
def compose_test(n: int) -> str:
    fns = ", ".join(f"fn{i % 2}" for i in range(n))

	expect = 0
	for i in range(n-1, -1, -1):
		if i % 2 == 0:
			expect += 1
		else:
			expect *= 2

    return f"""func TestCompose{n}_Example(t *testing.T) {{
    fn0 := func(x int) int {{ return x + 1 }}
	fn1 := func(x int) int {{ return x * 2 }}
    res := Compose{n}({fns})(0)
    if res != {expect} {{
        t.Error("Should perform right-to-left function composition of {n} functions. Received:", res)
    }}
}}"""
*/

func TestCompose5_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose5(fn0, fn1, fn0, fn1, fn0)(0)
	if res != 7 {
		t.Error("Should perform right-to-left function composition of 5 functions. Received:", res)
	}
}

func TestCompose6_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose6(fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 7 {
		t.Error("Should perform right-to-left function composition of 6 functions. Received:", res)
	}
}

func TestCompose7_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose7(fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 15 {
		t.Error("Should perform right-to-left function composition of 7 functions. Received:", res)
	}
}

func TestCompose8_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose8(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 15 {
		t.Error("Should perform right-to-left function composition of 8 functions. Received:", res)
	}
}

func TestCompose9_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose9(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 31 {
		t.Error("Should perform right-to-left function composition of 9 functions. Received:", res)
	}
}

func TestCompose10_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose10(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 31 {
		t.Error("Should perform right-to-left function composition of 10 functions. Received:", res)
	}
}

func TestCompose11_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose11(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 63 {
		t.Error("Should perform right-to-left function composition of 11 functions. Received:", res)
	}
}

func TestCompose12_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose12(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 63 {
		t.Error("Should perform right-to-left function composition of 12 functions. Received:", res)
	}
}

func TestCompose13_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose13(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 127 {
		t.Error("Should perform right-to-left function composition of 13 functions. Received:", res)
	}
}

func TestCompose14_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose14(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 127 {
		t.Error("Should perform right-to-left function composition of 14 functions. Received:", res)
	}
}

func TestCompose15_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose15(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 255 {
		t.Error("Should perform right-to-left function composition of 15 functions. Received:", res)
	}
}

func TestCompose16_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Compose16(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 255 {
		t.Error("Should perform right-to-left function composition of 16 functions. Received:", res)
	}
}
