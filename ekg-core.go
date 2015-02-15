// A module for defining metrics that can be monitored.
//
// Metrics are used to monitor program behavior and performance. All
// metrics have
//
//  * a name, and
//
//  * a way to get the metric's current value.
//
// This module provides a way to register metrics in a global \"metric
// store\". The store can then be used to get a snapshot of all
// metrics. The store also serves as a central place to keep track of
// all the program's metrics, both user and library defined.
//
// Here's an example of creating a single counter, used to count the
// number of request served by a web server:
//
// > import System.Metrics
// > import qualified System.Metrics.Counter as Counter
// >
// > main = do
// >     store <- newStore
// >     requests <- createCounter "myapp.request_count" store
// >     -- Every time we receive a request:
// >     Counter.inc requests
//
// This module also provides a way to register a number of predefined
// metrics that are useful in most applications. See e.g.
// 'registerGcMetrics'.
package ekg_core

import (
)


// A mutable metric store.
type Store struct {
    metrics map[string]interface{}
}


// Create a new, empty metric store.
func New () *Store {
    store := Store{
        metrics: make(map[string]interface{}),
    }
    return &store
}


// Create and register a zero-initialized counter.
func (store *Store) CreateCounter(name string) *Counter {
    counter := new(Counter)
    store.metrics[name] = counter
    return counter
}


// Create and register a zero-initialized gauge.
func (store *Store) CreateGauge(name string) *Gauge {
    gauge := new(Gauge)
    store.metrics[name] = gauge
    return gauge
}


// Create and register an empty label.
func (store *Store) CreateLabel(name string) *Label {
    label := new(Label)
    store.metrics[name] = label
    return label
}


// Create and register an event tracker.
func (store *Store) CreateDistribution(name string) *Distribution {
    distrib := new(Distribution)
    store.metrics[name] = distrib
    return distrib
}
