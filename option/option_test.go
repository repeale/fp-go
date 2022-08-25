package opt

import (
	"strconv"
	"testing"
)

func TestSome(t *testing.T) {
	res := Some("val")
	if res.hasValue != true {
		t.Error("Some should return a struct with hasValue set to true. Received:", res.hasValue)
	}
}

func TestNone(t *testing.T) {
	res := None[string]()
	if res.hasValue != false {
		t.Error("None should return a struct with hasValue set to false. Received:", res.hasValue)
	}
}

func TestIsSome_Some(t *testing.T) {
	res := IsSome(Some("value"))
	if res != true {
		t.Error("IsSome should return true. Received:", res)
	}
}
func TestIsSome_None(t *testing.T) {
	res := IsSome(None[string]())
	if res != false {
		t.Error("IsSome should return false. Received:", res)
	}
}

func TestIsNone_Some(t *testing.T) {
	res := IsNone(None[string]())
	if res != true {
		t.Error("IsNone should return true. Received:", res)
	}
}
func TestIsNone_None(t *testing.T) {
	res := IsNone(Some("value"))
	if res != false {
		t.Error("IsNone should return false. Received:", res)
	}
}

func TestGetOrElse_Some(t *testing.T) {
	res := GetOrElse(func() string { return "fail" })(Some("val"))
	if res != "val" {
		t.Error("GetOrElse should return the Some value. Received:", res)
	}
}

func TestGetOrElse_None(t *testing.T) {
	res := GetOrElse(func() string { return "elseValue" })(None[string]())
	if res != "elseValue" {
		t.Error("GetOrElse should return the onNone() value. Received:", res)
	}
}

func TestMatch_onSome(t *testing.T) {
	res := Match(
		func() int { return 0 },
		func(x string) int {
			intVar, _ := strconv.Atoi(x)
			return 1 + intVar
		})(Some("10"))
	if res != 11 {
		t.Error("Match should return the onSome() value, through the callback function. Received:", res)
	}
}

func TestMatch_onNone(t *testing.T) {
	res := Match(func() string { return "onNone" }, func(x string) string { return x + x })(None[string]())
	if res != "onNone" {
		t.Error("Match should return the onNone() return value. Received:", res)
	}
}

func TestMap_Some(t *testing.T) {
	res := Map(
		func(x string) int {
			intVar, _ := strconv.Atoi(x)
			return 1 + intVar
		})(Some("10"))
	if res.Value != 11 {
		t.Error("Map should return the a Some with the result of the callback function over the value. Received:", res.Value)
	}
}

func TestMap_None(t *testing.T) {
	res := Map(func(x string) string { return x + x })(None[string]())
	if res.hasValue != false {
		t.Error("Map should return a None value. Received:", res.Value)
	}
}

func TestChain_Some(t *testing.T) {
	res := Chain(func(x string) Option[string] { return Some(x + x) })(Some("val"))
	if res.hasValue != true {
		t.Error("Chain should return a Some of string. Received:", res.Value)
	}
}
func TestChain_None(t *testing.T) {
	res := Chain(func(x string) Option[string] { return Some(x + x) })(None[string]())
	if res.hasValue != false {
		t.Error("Chain should return a None value. Received:", res.Value)
	}
}
