package pack

// One -
type one struct {
	A string
	B string
}

var o = &one{A: "a", B: "b"}

// O -
func O() *one {
	return o
}
