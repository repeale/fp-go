package fp

func Pipe2[T1, T2, R any](fn1 func(T1) T2, fn2 func(T2) R) func(T1) R {
	return func(t1 T1) R {
		return fn2(fn1(t1))
	}
}

func Pipe3[T1, T2, T3, R any](fn1 func(T1) T2, fn2 func(T2) T3, fn3 func(T3) R) func(T1) R {
	return func(t1 T1) R {
		return fn3(fn2(fn1(t1)))
	}
}

func Pipe4[T1, T2, T3, T4, R any](fn1 func(T1) T2, fn2 func(T2) T3, fn3 func(T3) T4, fn4 func(T4) R) func(T1) R {
	return func(t1 T1) R {
		return fn4(fn3(fn2(fn1(t1))))
	}
}
