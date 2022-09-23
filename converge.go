package fp

/*
 * Converge is a function that Accepts a converging function and variadic parameters of function  and returns a new function.
 * When invoked, this new function is applied to some arguments,
 * and each branching function is applied to those same arguments.
 * The results of each branching function are passed as arguments to the converging function to produce the return value.
 */

// Apply converge of 2 branch functions
func Converge2[R1, R2, R, T any](fc func(R1, R2) R, f1 func(T) R1, f2 func(T) R2) func(T) R {
	return func(t T) R {
		return fc(f1(t), f2(t))
	}
}
