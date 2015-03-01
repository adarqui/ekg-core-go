package ekg_core

import (
)

type Bool struct {
	v bool
}

// Create a new, false initialized, boolean.
func newBool() *Bool {
	boolean := new(Bool)
	return boolean
}

// Get the current value of the boolean.
func (boolean *Bool) Read() bool {
	return boolean.v
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
    boolean.v = truth
}
