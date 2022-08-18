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
