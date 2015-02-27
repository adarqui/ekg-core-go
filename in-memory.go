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


// Create and register a zero-initialized counter.
func (store *Store) CreateCounter(name string) *Counter {
    counter := newCounter()
    store.Register(name, counter.ReadI)
    return counter
}


// Create and register a zero-initialized gauge.
func (store *Store) CreateGauge(name string) *Gauge {
    gauge := newGauge()
    store.Register(name, gauge.ReadI)
    return gauge
}


// Create and register an empty label.
func (store *Store) CreateLabel(name string) *Label {
    label := newLabel()
    store.Register(name, label.ReadI)
    return label
}


// Create and register an event tracker.
func (store *Store) CreateDistribution(name string) *Distribution {
    distrib := newDistribution()
    store.Register(name, distrib.ReadI)
    return distrib
}


// Create and register a timestamper.
func (store *Store) CreateTimestamp(name string) *Timestamp {
    timestamp := newTimestamp()
    store.Register(name, timestamp.ReadI)
    return timestamp
}
