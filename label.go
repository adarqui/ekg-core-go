// This module defines a type for mutable, string-valued labels.
// Labels are variable values and can be used to track e.g. the
// command line arguments or other free-form values. All operations on
// labels are thread-safe.
package ekg_core

import (
    "sync"
)


// A mutable, text-valued label.
type Label struct {
    mtx *sync.Mutex
    v string
}


// Callback for modify: string -> string.
type modify_hof func(string) string


// Create a new empty label.
func newLabel() (*Label) {
    label := new(Label)
    label.mtx = &sync.Mutex{}
    return label
}


// Get the current value of the label.
func (label *Label) Read() string {
    label.mtx.Lock()
    v := label.v
    label.mtx.Unlock()
    return v
}


// Set the label to the given value.
func (label *Label) Set(nv string) {
    label.mtx.Lock()
    label.v = nv
    label.mtx.Unlock()
}


// Set the label to the result of applying the given function to the
// value.
func (label *Label) Modify(cb modify_hof) {
    label.mtx.Lock()
    label.v = cb(label.v)
    label.mtx.Unlock()
}
