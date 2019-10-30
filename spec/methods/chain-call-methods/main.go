package main

import "log"

type s0 struct {
	s1 *s1
}

func (s *s0) getS1() *s1 {
	return s.s1
}

type s1 struct {
	one int
}

func (s *s1) getOne() int {
	return s.one
}

func main() {

	os0 := s0{}

	// should be panic
	log.Printf("one of s1 of s0 - %v", os0.getS1().getOne())
}
