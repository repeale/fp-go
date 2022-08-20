package fp

// Callback function that returns a specific value type
type Lazy[T any] func() T

// Callback function that takes an argument and return a value of the same type
type LazyVal[T any] func(x T) T

// Callback function that takes an argument of a specific type and returns a boolean
type Pred[T any] func(T) bool
