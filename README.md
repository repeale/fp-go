# fp-go - Functional programming helpers

## Install

Requires go 1.18+

    go get github.com/repeale/fp-go

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

[Every](#every) \
[Filter](#filter) \
[Flat](#flat) \
[FlatMap](#flatmap) \
[Map](#map) \
[Reduce](#reduce) \
[Some](#some)

---

[Compose](#compose) \
[Pipe](#pipe) \
[Curry](#curry)

#### Every

```go
fp.Every(func(x int) bool { return x > 0 })([]int{1, 2, 3})

// => true
```

#### Filter

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

```go
fp.FlatMap(func(x int) []int { return []int{x, x} })([]int{1, 2})

// => []int{1, 1, 2, 2}
```

#### Map

```go
fp.Map(func(x int64) string {
    return strconv.FormatInt(x, 10)
})([]int64{1, 2, 3})

// => []string{"1", "2", "3", "4"}
```

#### Reduce

```go
fp.Reduce(func(acc int, curr int) int { return acc + curr }, 0)([]int{1, 2, 3})

// => 6
```

#### Some

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
