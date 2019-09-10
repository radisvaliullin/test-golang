package mocklock

// Do ...
type Do interface {
	Do2(int) int
	Do3(int) int
}

// Run ...
type Run interface {
	Run(int) int
}

// Some ...
type Some struct {
	doer   Do
	runner Run
}

// Do2 ...
func (s *Some) Do2(i int) int {
	return s.doer.Do2(i)
}

// Do3 ...
func (s *Some) Do3(i int) int {
	return s.doer.Do3(i)
}

// Run ...
func (s *Some) Run(i int) int {
	return s.runner.Run(i)
}

// Doer ...
type Doer struct {
}

// Do2 ...
func (d *Doer) Do2(i int) int {
	return i * i
}

// Do3 ...
func (d *Doer) Do3(i int) int {
	return i * i * i
}

// Runner ...
type Runner struct{}

// Run ...
func (r *Runner) Run(i int) int {
	return i * i * i * i
}
