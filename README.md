# fp-go - Functional programming helpers

## Install

Requires go 1.18+

    go get github.com/lib/pq

## Features

- [Currying](#currying)
- [Variations to get only what you really need](#variations)

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
fp.Map[int64, string](func(x int64) string {
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
