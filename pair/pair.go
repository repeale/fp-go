package pair

// Base Tuple struct
type Tuple[A, B any] struct {
	a A
	b B
}

// Constructor for a Tuple
func Pair[A, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{a, b}
}

// Getter for the first value of the Tuple
func Fst[A, B any](t Tuple[A, B]) A {
	return t.a
}

// Getter for the second value of the Tuple
func Snd[A, B any](t Tuple[A, B]) B {
	return t.b
}

// Getter for both values of the Tuple
func Get[A, B any](t Tuple[A, B]) (A, B) {
	return t.a, t.b
}

// Execute the function on the first value of the Tuple
func MapFst[B, A, R any](fn func(A) R) func(Tuple[A, B]) Tuple[R, B] {
	return func(t Tuple[A, B]) Tuple[R, B] {
		return Tuple[R, B]{fn(t.a), t.b}
	}
}

// Execute the function on the second value of the Tuple
func MapSnd[A, B, R any](fn func(B) R) func(Tuple[A, B]) Tuple[A, R] {
	return func(t Tuple[A, B]) Tuple[A, R] {
		return Tuple[A, R]{t.a, fn(t.b)}
	}
}

// Execute the functions on both the first and second values of the Tuple
func MapBoth[A, B, R, S any](fnF func(A) R, fnS func(B) S) func(Tuple[A, B]) Tuple[R, S] {
	return func(t Tuple[A, B]) Tuple[R, S] {
		return Tuple[R, S]{fnF(t.a), fnS(t.b)}
	}
}

// Helper to chcek if the first value of the Tuple satisfies a predicate
func CheckFst[B, A any](fn func(A) bool) func(Tuple[A, B]) bool {
	return func(t Tuple[A, B]) bool {
		return fn(t.a)
	}
}

// Helper to chcek if the second value of the Tuple satisfies a predicate
func CheckSnd[A, B any](fn func(B) bool) func(Tuple[A, B]) bool {
	return func(t Tuple[A, B]) bool {
		return fn(t.b)
	}
}

// Helper to chcek if both the first and the second values of the Tuple satisfies their respective predicate
func CheckBoth[A, B any](fnF func(A) bool, fnS func(B) bool) func(Tuple[A, B]) bool {
	return func(t Tuple[A, B]) bool {
		return fnF(t.a) && fnS(t.b)
	}
}
