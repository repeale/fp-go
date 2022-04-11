package fp

import (
	"testing"
)

func TestCurry2_Example(t *testing.T) {
	res := Curry2(func(a int, b int) int { return a + b })(1)(2)

	if res != 3 {
		t.Error("Should allow to curry a sum. Received:", res)
	}
}

func TestCurry3_Example(t *testing.T) {
	res := Curry3(func(a int, b int, c int) int { return a + b + c })(1)(2)(3)

	if res != 6 {
		t.Error("Should allow to curry a sum. Received:", res)
	}
}

func TestCurry4_Example(t *testing.T) {
	res := Curry4(func(a int, b int, c int, d int) int { return a + b + c + d })(1)(2)(3)(4)

	if res != 10 {
		t.Error("Should allow to curry a sum. Received:", res)
	}
}
