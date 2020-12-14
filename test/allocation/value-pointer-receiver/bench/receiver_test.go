package receiver_test

import "testing"

type some struct {
	par0  int
	par1  int
	par2  int
	par3  int
	par4  int
	par5  int
	par6  int
	par7  int
	par8  int
	par9  int
	par10 int
	par11 int
	par12 int
	par13 int
	par14 int
	par15 int
	par16 int
	par17 int
	par18 int
	par19 int
	par20 int
	par21 int
	par22 int
	par23 int
	par24 int
	par25 int
	par26 int
	par27 int
	par28 int
	par29 int
	// par30 int
	// par31 int
	// par32 int
	// par33 int
	// par34 int
	// par35 int
	// par36 int
	// par37 int
	// par38 int
	// par39 int
}

func (s *some) sum(i int) int {
	out := s.par0 +
		s.par1 +
		s.par2 +
		s.par3 +
		s.par4 +
		s.par5 +
		s.par6 +
		s.par7 +
		s.par8 +
		s.par9 +
		s.par10 +
		s.par11 +
		s.par12 +
		s.par13 +
		s.par14 +
		s.par15 +
		s.par16 +
		s.par17 +
		s.par18 +
		s.par19 +
		s.par20 +
		s.par21 +
		s.par22 +
		s.par23 +
		s.par24 +
		s.par25 +
		s.par26 +
		s.par27 +
		s.par28 +
		s.par29 +
		// s.par30 +
		// s.par31 +
		// s.par32 +
		// s.par33 +
		// s.par34 +
		// s.par35 +
		// s.par36 +
		// s.par37 +
		// s.par38 +
		// s.par39 +
		i
	return out
}

func BenchmarkReceiverTest(b *testing.B) {
	s := some{}
	for i := 0; i < b.N; i++ {
		s.sum(i)
	}
}
