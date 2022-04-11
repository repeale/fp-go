package fp

// Allow to transfrom a function that receives two params in single functions that take one single param each
func Curry2[T1, T2, R any](fn func(T1, T2) R) func(T1) func(T2) R {
	return func(t1 T1) func(T2) R {
		return func(t2 T2) R {
			return fn(t1, t2)
		}
	}
}

// Like Curry2 but with three params
func Curry3[T1, T2, T3, R any](fn func(T1, T2, T3) R) func(T1) func(T2) func(T3) R {
	return func(t1 T1) func(T2) func(T3) R {
		return func(t2 T2) func(T3) R {
			return func(t3 T3) R {
				return fn(t1, t2, t3)
			}
		}
	}
}

// Like Curry2 but with four params
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
