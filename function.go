package fp

type Lazy[T any] func() T

type LazyVal[T any] func(x T) T
