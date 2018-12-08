package pack

// one -
type one struct {
	A string
	b string
}

var o = &one{A: "a", b: "bb"}

// O -
func O() *one {
	return o
}
