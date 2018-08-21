package main

import "log"

func main() {

	s := newStore()
	s.add("one", 1)
	s.add("two", 2)

	sm := s.getM()
	log.Printf("sm is %+v", sm)

	s.add("three", 3)
	s.add("four", 4)

	log.Printf("sm is %+v", sm)

	sm["44"] = 44

	s.add("end", 42)
}

type store struct {
	m map[string]int ``
}

func newStore() *store {
	s := &store{
		m: map[string]int{},
	}
	log.Printf("init store, m is %+v", s.m)
	return s
}

func (s *store) add(key string, val int) {
	s.m[key] = val
	log.Printf("add key, m is %+v", s.m)
}

func (s *store) getM() map[string]int {
	return s.m
}
