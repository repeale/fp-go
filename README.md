# fp-go

[![Go Reference](https://pkg.go.dev/badge/github.com/repeale/fp-go.svg)](https://pkg.go.dev/github.com/repeale/fp-go)
[![Go Report](https://goreportcard.com/badge/github.com/repeale/fp-go)](https://goreportcard.com/badge/github.com/repeale/fp-go)

Fp-go is a collection of Functional Programming helpers powered by Golang [1.18](https://tip.golang.org/doc/go1.18)+ [generics](https://tip.golang.org/doc/go1.18#generics).

<p align="center">
  <img 
    width="500"
    height="313"
    src="https://user-images.githubusercontent.com/9580458/162070974-4367f4b8-bb3d-451c-8114-dd578bad4e46.png"
  >
</p>

## Contents

- [Install](#install)
- [Features](#features)
  - [Currying](#currying)
  - [Variations](#variations)
- [Helpers](#helpers)
  - [Every](#every)
  - [Filter](#filter)
  - [Flat](#flat)
  - [FlatMap](#flatmap)
  - [Map](#map)
  - [Reduce](#reduce)
  - [Some](#some)
  - [Compose](#compose)
  - [Pipe](#pipe)
  - [Curry](#curry)

## Install

Requires go 1.18+

```sh
go get github.com/repeale/fp-go
```

## Features

- [Currying](#currying)
- [Variations](#variations)

### Currying

By default! Data last!

```go
func isPositive(x int) bool {
	return x > 0
}

func main() {
    filterPositive := fp.Filter(isPositive)
    numbers := []int{1, 2, 3, 4, 5}

    filterPositive(numbers)
}
```

### Variations

Variations allows you to get different parameters in the callback function so that you get only only what is really needed.

[Default](#default) \
[WithIndex](#withindex) \
[WithSlice](#withslice)

#### Default

Only the current item is available:

```go
fp.Map[int, string](func(x int) { ... })
```

#### WithIndex

Current element and index are available:

```go
fp.MapWithIndex[int, string](func(x int, i int) { ... })
```

#### WithSlice

Current element, index and the whole slice are available:

```go
fp.MapWithSlice[int, string](func(x int, i int, xs: []int) { ... })
```

## Helpers

#### Every

Variations `EveryWithIndex` and `EveryWithSlice`

```go
fp.Every(func(x int) bool { return x > 0 })([]int{1, 2, 3})

// => true
```

#### Filter

Variations `FilterWithIndex` and `FilterWithSlice`

```go
fp.Filter(func(x int) bool { return x > 0 })([]int{-1, 2, -3, 4})

// => []int{2, 4}
```

#### Flat

```go
fp.Flat([][]int{{1, 2}, {3, 4}})

// => []int{1, 2, 3, 4}
```

#### FlatMap

Variations `FlatMapWithIndex` and `FlatMapWithSlice`

```go
fp.FlatMap(func(x int) []int { return []int{x, x} })([]int{1, 2})

// => []int{1, 1, 2, 2}
```

#### Map

Variations `MapWithIndex` and `MapWithSlice`

```go
fp.Map(func(x int64) string {
    return strconv.FormatInt(x, 10)
})([]int64{1, 2, 3})

// => []string{"1", "2", "3", "4"}
```

#### Reduce

Variations `ReduceWithIndex` and `ReduceWithSlice`

```go
fp.Reduce(func(acc int, curr int) int { return acc + curr }, 0)([]int{1, 2, 3})

// => 6
```

#### Some

Variations `SomeWithIndex` and `SomeWithSlice`

```go
fp.Some(func(x int) bool { return x < 0 })([]int{1, 2, 3})

// => false
```

---

#### Compose

Performs right-to-left function composition.

Variations `Compose2`, `Compose3` and `Compose4` stating the number of functions you are going to compose.

```go
func isPositive(x int) bool {
	return x > 0
}

func sumTwo(x int) int {
	return x + 2
}

Pipe2(fp.Filter(isPositive), fp.Map(sumTwo))([]int{1, 2, 3, -1})

// => []int{3,4,5,1}
```

#### Pipe

Performs left-to-right function composition.

Variations `Pipe2`, `Pipe3` and `Pipe4` stating the number of functions you are going to compose.

```go
func isPositive(x int) bool {
	return x > 0
}

func sumTwo(x int) int {
	return x + 2
}

Pipe2(fp.Filter(isPositive), fp.Map(sumTwo))([]int{1, 2, 3, -1})

// => []int{3,4,5}
```

#### Curry

Variations `Curry2`, `Curry3` and `Curry4` stating the number of params will be curried individually.

```go
curryedSum := Curry2(func(a int, b int) int { return a + b })
curryedSum(1)(2)
```
