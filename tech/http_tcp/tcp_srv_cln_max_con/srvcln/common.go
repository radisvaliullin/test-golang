package srvcln

import "sync"

type safeCounter struct {
	cnt int
	mx  sync.Mutex
}

func (c *safeCounter) Inc() {
	c.mx.Lock()
	c.cnt++
	c.mx.Unlock()
}

func (c *safeCounter) Dec() {
	c.mx.Lock()
	c.cnt--
	c.mx.Unlock()
}

func (c *safeCounter) Cnt() int {
	var out int
	c.mx.Lock()
	out = c.cnt
	c.mx.Unlock()
	return out
}
