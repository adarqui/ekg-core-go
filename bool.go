package ekg_core

import (
	"sync"
)

type Bool struct {
	mtx *sync.Mutex
	v   bool
}

// Create a new, false initialized, boolean.
func newBool() *Bool {
	boolean := new(Bool)
	boolean.mtx = &sync.Mutex{}
	return boolean
}

// Get the current value of the boolean.
func (boolean *Bool) Read() bool {
	boolean.mtx.Lock()
	v := boolean.v
	boolean.mtx.Unlock()
	return v
}

func (boolean *Bool) ReadI(v interface{}) interface{} {
	return boolean.Read()
}

// Set the boolean to true.
func (boolean *Bool) True() {
	boolean.Set(true)
}

// Set the boolean to false.
func (boolean *Bool) False() {
	boolean.Set(false)
}

// Set the boolean to <truth>.
func (boolean *Bool) Set(truth bool) {
	boolean.mtx.Lock()
	boolean.v = truth
	boolean.mtx.Unlock()
}

// Toggle the boolean value.
func (boolean *Bool) Toggle() {
	v := boolean.Read()
	if v == true {
		boolean.Set(false)
	} else {
		boolean.Set(true)
	}
}
