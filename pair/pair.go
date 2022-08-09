package pair

// Base Pair struct
type Pair[A, B any] struct {
	a A
	b B
}

// Constructor for a Pair
func New[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{a, b}
}

// Getter for the first value of the Pair
func Fst[A, B any](t Pair[A, B]) A {
	return t.a
}

// Getter for the second value of the Pair
func Snd[A, B any](t Pair[A, B]) B {
	return t.b
}

// Getter for both values of the Pair
func Get[A, B any](t Pair[A, B]) (A, B) {
	return t.a, t.b
}

// Execute the function on the first value of the Pair
func MapFst[B, A, R any](fn func(A) R) func(Pair[A, B]) Pair[R, B] {
	return func(t Pair[A, B]) Pair[R, B] {
		return Pair[R, B]{fn(t.a), t.b}
	}
}

// Execute the function on the second value of the Pair
func MapSnd[A, B, R any](fn func(B) R) func(Pair[A, B]) Pair[A, R] {
	return func(t Pair[A, B]) Pair[A, R] {
		return Pair[A, R]{t.a, fn(t.b)}
	}
}

// Execute the functions on both the first and second values of the Pair
func MapBoth[A, B, R, S any](fnF func(A) R, fnS func(B) S) func(Pair[A, B]) Pair[R, S] {
	return func(t Pair[A, B]) Pair[R, S] {
		return Pair[R, S]{fnF(t.a), fnS(t.b)}
	}
}

// Helper to chcek if the first value of the Pair satisfies a predicate
func CheckFst[B, A any](fn func(A) bool) func(Pair[A, B]) bool {
	return func(t Pair[A, B]) bool {
		return fn(t.a)
	}
}

// Helper to chcek if the second value of the Pair satisfies a predicate
func CheckSnd[A, B any](fn func(B) bool) func(Pair[A, B]) bool {
	return func(t Pair[A, B]) bool {
		return fn(t.b)
	}
}

// Helper to chcek if both the first and the second values of the Pair satisfies their respective predicate
func CheckBoth[A, B any](fnF func(A) bool, fnS func(B) bool) func(Pair[A, B]) bool {
	return func(t Pair[A, B]) bool {
		return fnF(t.a) && fnS(t.b)
	}
}

// Merge the elements of a Pair with a Curried function
func MergeC[A, B, C any](fn func(A) func(B) C) func(Pair[A, B]) C {
	return func(t Pair[A, B]) C {
		return fn(t.a)(t.b)
	}
}

// Merge the elements of a Pair with a non-Curried function
func Merge[A, B, C any](fn func(A, B) C) func(Pair[A, B]) C {
	return func(t Pair[A, B]) C {
		return fn(t.a, t.b)
	}
}
