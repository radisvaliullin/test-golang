package main

import (
	"fmt"
	"time"
)

func main() {

	// tk := time.NewTicker(time.Second * 2)
	tk := NewInstantTicker(time.Second * 3)

	go func() {
		for t := range tk.C {
			fmt.Println("tick ", t)
		}
		fmt.Println("tick loop end")
	}()

	for {

	}
}

// InstantTicker -
// problem you must use stop to kill C upd goroutine
type InstantTicker struct {
	C       <-chan time.Time
	origTk  *time.Ticker
	stopSig chan struct{}
}

// Stop -
func (t *InstantTicker) Stop() {
	t.origTk.Stop()
	t.stopSig <- struct{}{}
}

// NewInstantTicker -
func NewInstantTicker(d time.Duration) *InstantTicker {

	c := make(chan time.Time)
	iTk := &InstantTicker{
		C:       c,
		origTk:  time.NewTicker(d),
		stopSig: make(chan struct{}, 1),
	}
	go func() {
		c <- time.Now()
		for {
			select {
			case tk := <-iTk.origTk.C:
				c <- tk
			case <-iTk.stopSig:
				return
			}
		}
	}()
	return iTk
}
