package fp

func Curry2[T1, T2, R any](fn func(T1, T2) R) func(T1) func(T2) R {
	return func(t1 T1) func(T2) R {
		return func(t2 T2) R {
			return fn(t1, t2)
		}
	}
}

func Curry3[T1, T2, T3, R any](fn func(T1, T2, T3) R) func(T1) func(T2) func(T3) R {
	return func(t1 T1) func(T2) func(T3) R {
		return func(t2 T2) func(T3) R {
			return func(t3 T3) R {
				return fn(t1, t2, t3)
			}
		}
	}
}

func Curry4[T1, T2, T3, T4, R any](fn func(T1, T2, T3, T4) R) func(T1) func(T2) func(T3) func(T4) R {
	return func(t1 T1) func(T2) func(T3) func(T4) R {
		return func(t2 T2) func(T3) func(T4) R {
			return func(t3 T3) func(T4) R {
				return func(t4 T4) R {
					return fn(t1, t2, t3, t4)
				}
			}
		}
	}
}
