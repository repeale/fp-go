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
