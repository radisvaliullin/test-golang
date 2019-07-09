package main

import "log"

// There is list of stations like []int{3,2,1,0,4},
// where index it is station and
// value it is how many station can you pass from current index station to next station by increasing station number.
// For example above: from station 0 (index) you can pass to station 1, 2 or 3 (indexes);
// from station 1 (index) to station 2 and 3 (index);
// from station 2 (index) to station 3 (index) because value is 1;
// etc.
// You should implement function answerting true false
// if you can pass from station 0 to out of end station.
// For example above will return false, because you can not pass far than station 3.
func main() {

	ins := [][]int{
		[]int{3, 2, 1, 1, 4},
		[]int{3, 2, 1, 0, 4},
		[]int{3, 2, 1, 1, 4, 0, 0, 0, 0, 5},
	}

	for _, in := range ins {
		res := pass(in)

		log.Printf("for in %v you result is %v", in, res)
	}
}

func pass(in []int) bool {
	if len(in) == 0 {
		return false
	}
	for i := in[0]; i > 0; i-- {
		log.Printf("i - %v", i)

		if i >= len(in) {
			return true
		}
		res := pass(in[i:])
		if res {
			return true
		}
	}
	return false
}
