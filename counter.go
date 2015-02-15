// This module defines a type for mutable, integer-valued counters.
// Counters are non-negative, monotonically increasing values and can
// be used to track e.g. the number of requests served since program
// start.  All operations on counters are thread-safe.
package ekg_core

import (
    "github.com/adarqui/math"
    "sync/atomic"
    "runtime"
)


// A mutable, integer-valued counter.
type Counter struct {
    v int64
}


// Create a new, zero initialized, counter.
func newCounter() (*Counter) {
    counter := new(Counter)
    return counter
}


// Get the current value of the counter.
func (counter *Counter) Read() int64 {
    v := atomic.LoadInt64(&counter.v)
    return v
}


// Increase the counter by one.
func (counter *Counter) Inc() {
    counter.Add(1)
}


// Add the argument to the counter.
func (counter *Counter) Add(i int64) {
    atomic.AddInt64(&counter.v, math.AbsInt64(i))
    runtime.Gosched()
}
