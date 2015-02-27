// This module defines a type for mutable, integer-valued gauges.
// Gauges are variable values and can be used to track e.g. the
// current number of concurrent connections. All operations on gauges
// are thread-safe.
package ekg_core

import (
    "sync/atomic"
    "runtime"
)


// A mutable, integer-valued gauge.
type Gauge struct {
    v int64
}


// Create a new, zero initialized, gauge.
func newGauge() (*Gauge) {
    gauge := new(Gauge)
    return gauge
}


// Get the current value of the gauge.
func (gauge *Gauge) Read() int64 {
    v := atomic.LoadInt64(&gauge.v)
    return v
}


func (gauge *Gauge) ReadI(v interface{}) interface{} {
    return gauge.Read()
}


// Increase the gauge by one.
func (gauge *Gauge) Inc() {
    gauge.Add(1)
}


// Decrease the gauge by one.
func (gauge *Gauge) Dec() {
    gauge.Subtract(1)
}


// Increase the gauge by the given amount.
func (gauge *Gauge) Add(i int64) {
    gauge.AddSubtract(i)
}


// Decrease the gauge by the given amount.
func (gauge *Gauge) Subtract(i int64) {
    gauge.AddSubtract(i * (-1))
}


// See Subtract.
func (gauge *Gauge) Sub(i int64) {
    gauge.Subtract(i)
}


// Increase or decrease the gauge by the given amount.
func (gauge *Gauge) AddSubtract(i int64) {
    atomic.AddInt64(&gauge.v, i)
    runtime.Gosched()
}


// Set the gauge to the given value.
func (gauge *Gauge) Set(i int64) {
    atomic.StoreInt64(&gauge.v, i)
    runtime.Gosched()
}
