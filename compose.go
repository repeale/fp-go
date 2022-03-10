package fp

func Compose2[T1, T2, R any](fn2 func(T2) R, fn1 func(T1) T2) func(T1) R {
	return func(t1 T1) R {
		return fn2(fn1(t1))
	}
}

func Compose3[T1, T2, T3, R any](fn3 func(T3) R, fn2 func(T2) T3, fn1 func(T1) T2) func(T1) R {
	return func(t1 T1) R {
		return fn3(fn2(fn1(t1)))
	}
}

func Compose4[T1, T2, T3, T4, R any](fn4 func(T4) R, fn3 func(T3) T4, fn2 func(T2) T3, fn1 func(T1) T2) func(T1) R {
	return func(t1 T1) R {
		return fn4(fn3(fn2(fn1(t1))))
	}
}
