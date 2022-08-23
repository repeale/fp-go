def compose_test(n: int) -> str:
	funcs = ["add1", "double"]
	fns = ", ".join(funcs[i % 2] for i in range(n))

	expect = 0
	for i in range(n-1, -1, -1):
		if i % 2 == 0:
			expect += 1
		else:
			expect *= 2

	return f"""
func TestCompose{n}_Example(t *testing.T) {{
    res := Compose{n}({fns})(0)
    if res != {expect} {{
        t.Error("Should perform right-to-left function composition of {n} functions. Received:", res)
    }}
}}"""


def pipe_test(n: int) -> str:
	funcs = ["add1", "double"]
	fns = ", ".join(funcs[i % 2] for i in range(n))

	expect = 0
	for i in range(n):
		if i % 2 == 0:
			expect += 1
		else:
			expect *= 2

	return f"""
func TestPipe{n}_Example(t *testing.T) {{
	res := Pipe{n}({fns})(0)
	if res != {expect} {{
		t.Error("Should perform left-to-right function composition of {n} functions. Received:", res)
	}}
}}"""


def pipe_func(n: int) -> str:
    gen = "[" + ", ".join(f"T{i+1}" for i in range(n)) + ", R any]"
    fns = ", ".join(f"fn{i} func(T{i}) T{i+1}" for i in range(1, n)) + f", fn{n} func(T{n}) R"
    cal = "".join(f"fn{i}(" for i in range(n, 0, -1)) + "t1" + ")" * n

    return f"""// Performs left-to-right function composition of {n} functions
func Pipe{n}{gen}({fns}) func(T1) R  {{
    return func(t1 T1) R {{
        return {cal}
    }}
}}"""


def compose_func(n: int) -> str:
    gen = "[" + ", ".join(f"T{i+1}" for i in range(n)) + ", R any]"
    fns = f"fn{n} func(T{n}) R, " + ", ".join(f"fn{i} func(T{i}) T{i+1}" for i in range(n-1, 0, -1))
    cal = "".join(f"fn{i}(" for i in range(n, 0, -1)) + "t1" + ")" * n

    return f"""// Performs right-to-left function composition of {n} functions
func Compose{n}{gen}({fns}) func(T1) R  {{
    return func(t1 T1) R {{
        return {cal}
    }}
}}"""