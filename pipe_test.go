package fp

import (
	"reflect"
	"testing"
)

func TestPipe2_Example(t *testing.T) {
	res := Pipe2(
		Filter(func(x int) bool { return x > 0 }),
		Map(func(x int) int { return x + 2 }))([]int{1, 2, 3, -1})
	expected := []int{3, 4, 5}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("Should perform left-to-right function composition of two functions. Received:", res)
	}
}

func TestPipe3_Example(t *testing.T) {
	res := Pipe3(
		Filter(func(x int) bool { return x > 0 }),
		Map(func(x int) int { return x + 2 }),
		Map(func(x int) int { return x + 2 }))([]int{1, 2, 3, -1})
	expected := []int{5, 6, 7}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("Should perform left-to-right function composition of two functions. Received:", res)
	}
}

func TestPipe4_Example(t *testing.T) {
	res := Pipe4(
		Filter(func(x int) bool { return x > 0 }),
		Map(func(x int) int { return x + 2 }),
		Map(func(x int) int { return x + 2 }),
		Map(func(x int) int { return x + 2 }))([]int{1, 2, 3, -1})
	expected := []int{7, 8, 9}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("Should perform left-to-right function composition of two functions. Received:", res)
	}
}

/* Generated with the following python code
def pipe_test(n: int) -> str:
	fns = ", ".join(f"fn{i % 2}" for i in range(n))

	expect = 0
	for i in range(n):
		if i % 2 == 0:
			expect += 1
		else:
			expect *= 2

    return f"""func TestPipe{n}_Example(t *testing.T) {{
    fn0 := func(x int) int {{ return x + 1 }}
	fn1 := func(x int) int {{ return x * 2 }}
    res := Pipe{n}({fns})(0)
    if res != {expect} {{
        t.Error("Should perform left-to-right function composition of {n} functions. Received:", res)
    }}
}}"""
*/
func TestPipe5_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe5(fn0, fn1, fn0, fn1, fn0)(0)
	if res != 7 {
		t.Error("Should perform left-to-right function composition of 5 functions. Received:", res)
	}
}

func TestPipe6_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe6(fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 14 {
		t.Error("Should perform left-to-right function composition of 6 functions. Received:", res)
	}
}

func TestPipe7_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe7(fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 15 {
		t.Error("Should perform left-to-right function composition of 7 functions. Received:", res)
	}
}

func TestPipe8_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe8(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 30 {
		t.Error("Should perform left-to-right function composition of 8 functions. Received:", res)
	}
}

func TestPipe9_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe9(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 31 {
		t.Error("Should perform left-to-right function composition of 9 functions. Received:", res)
	}
}

func TestPipe10_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe10(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 62 {
		t.Error("Should perform left-to-right function composition of 10 functions. Received:", res)
	}
}

func TestPipe11_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe11(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 63 {
		t.Error("Should perform left-to-right function composition of 11 functions. Received:", res)
	}
}

func TestPipe12_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe12(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 126 {
		t.Error("Should perform left-to-right function composition of 12 functions. Received:", res)
	}
}

func TestPipe13_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe13(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 127 {
		t.Error("Should perform left-to-right function composition of 13 functions. Received:", res)
	}
}

func TestPipe14_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe14(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 254 {
		t.Error("Should perform left-to-right function composition of 14 functions. Received:", res)
	}
}

func TestPipe15_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe15(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0)(0)
	if res != 255 {
		t.Error("Should perform left-to-right function composition of 15 functions. Received:", res)
	}
}

func TestPipe16_Example(t *testing.T) {
	fn0 := func(x int) int { return x + 1 }
	fn1 := func(x int) int { return x * 2 }
	res := Pipe16(fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1, fn0, fn1)(0)
	if res != 510 {
		t.Error("Should perform left-to-right function composition of 16 functions. Received:", res)
	}
}
