package pair

type Tuple[A, B any] struct {
	a A
	b B
}

func Pair[A, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{a, b}
}

func Fst[A, B any](t Tuple[A, B]) A {
	return t.a
}

func Snd[A, B any](t Tuple[A, B]) B {
	return t.b
}

func Get[A, B any](t Tuple[A, B]) (A, B) {
	return t.a, t.b
}
