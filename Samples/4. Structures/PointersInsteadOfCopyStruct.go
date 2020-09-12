package main

//S struct is used to demostarte the use of pointers
type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byCopy() S {

	return S{
		a: 1, b: 2, c: 3,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {

	return &S{
		a: 1, b: 2, c: 3,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}
