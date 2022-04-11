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
