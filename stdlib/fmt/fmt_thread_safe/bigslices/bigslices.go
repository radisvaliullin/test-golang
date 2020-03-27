package bigslices

var Size = 1024 * 128
var WriteCycles = 1024 * 128

// var WriteCycles = 1024 * 1024 * (1024 / Size)
var BigSlice0, BigSlice1 = SetBigSlice0and1()

var CharA byte = byte('a')
var CharB byte = byte('b')

func SetBigSlice0and1() ([]byte, []byte) {

	bs0 := make([]byte, Size)
	bs1 := make([]byte, Size)
	for i := 0; i < (Size - 1); i++ {
		bs0[i] = CharA
		bs1[i] = CharB
	}
	bs0[Size-1] = byte('\n')
	bs1[Size-1] = byte('\n')
	return bs0, bs1
}
