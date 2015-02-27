// This module defines a type for tracking statistics about a series
// of events. An event could be handling of a request and the value
// associated with the event -- the value you'd pass to 'add' -- could
// be the amount of time spent serving that request (e.g. in
// milliseconds). All operations are thread safe.
package ekg_core

import (
    "sync/atomic"
)


// An metric for tracking events.
type Distribution struct {
    v int64
}


// Add a value to the distribution.
func newDistribution() (*Distribution) {
    distrib := new(Distribution)
    return distrib
}


// Get the current statistical summary for the event being tracked.
func (distrib *Distribution) Read() int64 {
    v := atomic.LoadInt64(&distrib.v)
    return v
}


func (distrib *Distribution) ReadI(v interface{}) interface{} {
    return distrib.Read()
}


// Add a value to the distribution.
func (distrib *Distribution) Add(i int64) {
}



// Add the same value to the distribution N times.
func (distrib *Distribution) AddN(i, n int64) {
}
