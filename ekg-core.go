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

import ()

// Type Information
type Type int

const (
	COUNTER Type = iota
	GAUGE
	LABEL
	DISTRIBUTION
	TIMESTAMP
)

// A representation of a value
type Value struct {
	Typ Type
	Val interface{}
}

// A metric entry
type Metric struct {
	getter func(interface{}) interface{}
	typ    Type
}

// A group entry
type Group struct {
	sampleAction   func() interface{}
	samplerMetrics map[string]Metric
}

// A mutable metric store.
type Store struct {
	metrics     map[string]Metric
	groups      map[int]Group
	stateNextId int
}

// Create a new, empty metric store.
func New() *Store {
	store := Store{
		metrics: make(map[string]Metric),
		groups:  make(map[int]Group),
	}
	return &store
}

// | Register a non-negative, monotonically increasing, integer-valued
// metric. The provided action to read the value must be thread-safe.
// Also see 'CreateCounter'.
func (store *Store) RegisterCounter(name string, cb func(interface{}) interface{}) {
	store.Register(name, cb, COUNTER)
}

// | Register an integer-valued metric. The provided action to read
// the value must be thread-safe. Also see 'CreateGauge'.
func (store *Store) RegisterGauge(name string, cb func(interface{}) interface{}) {
	store.Register(name, cb, GAUGE)
}

// | Register a text metric. The provided action to read the value
// must be thread-safe. Also see 'CreateLabel'.
func (store *Store) RegisterLabel(name string, cb func(interface{}) interface{}) {
	store.Register(name, cb, LABEL)
}

// | Register a distribution metric. The provided action to read the
// value must be thread-safe. Also see 'CreateDistribution'.
func (store *Store) RegisterDistribution(name string, cb func(interface{}) interface{}) {
	store.Register(name, cb, DISTRIBUTION)
}

// | Register a timestampn metric. The provided action to read the
// value must be thread-safe. Also see 'CreateTimestamp'.
func (store *Store) RegisterTimestamp(name string, cb func(interface{}) interface{}) {
	store.Register(name, cb, TIMESTAMP)
}

/*
-- | The value of a sampled metric.
data Value = Counter {-# UNPACK #-} !Int64
           | Gauge {-# UNPACK #-} !Int64
           | Label {-# UNPACK #-} !T.Text
           | Distribution !Distribution.Stats
           deriving Show
*/

// * Registering metrics

// $registering
// Before metrics can be sampled they need to be registered with the
// metric store. The same metric name can only be used once. Passing a
// metric name that has already been used to one of the register
// function is an 'error'.

// register
func (store *Store) Register(name string, sample func(interface{}) interface{}, typ Type) {
	m := Metric{}
	m.getter = sample
	m.typ = typ
	store.metrics[name] = m
}

// | Register an action that will be executed any time one of the
// metrics computed from the value it returns needs to be sampled.
//
// When one or more of the metrics listed in the first argument needs
// to be sampled, the action is executed and the provided getter
// functions will be used to extract the metric(s) from the action's
// return value.
//
// The registered action might be called from a different thread and
// therefore needs to be thread-safe.
//
// This function allows you to sample groups of metrics together. This
// is useful if
//
// * you need a consistent view of several metric or
//
// * sampling the metrics together is more efficient.
//
// For example, sampling GC statistics needs to be done atomically or
// a GC might strike in the middle of sampling, rendering the values
// incoherent. Sampling GC statistics is also more efficient if done
// in \"bulk\", as the run-time system provides a function to sample all
// GC statistics at once.
//
// Note that sampling of the metrics is only atomic if the provided
// action computes @a@ atomically (e.g. if @a@ is a record, the action
// needs to compute its fields atomically if the sampling is to be
// atomic.)
//
// Example usage:
//
// > {-# LANGUAGE OverloadedStrings #-}
// > import qualified Data.HashMap.Strict as M
// > import GHC.Stats
// > import System.Metrics
// >
// > main = do
// >     store <- newStore
// >     let metrics =
// >             [ ("num_gcs", Counter . numGcs)
// >             , ("max_bytes_used", Gauge . maxBytesUsed)
// >             ]
// >     registerGroup (M.fromList metrics) getGCStats store
func (store *Store) RegisterGroup(getters map[string]Metric, cb func() interface{}) {
	g := Group{
		sampleAction:   cb,
		samplerMetrics: getters,
	}
	store.groups[store.stateNextId] = g
	store.stateNextId += 1
}

// | Sample all metrics. Sampling is /not/ atomic in the sense that
// some metrics might have been mutated before they're sampled but
// after some other metrics have already been sampled.
func (store *Store) SampleAll() map[string]Value {
	sample_metrics := store.readAllRefs()
	sample_groups := store.SampleGroups()
	return merge(sample_metrics, sample_groups)
}

func (store *Store) SampleGroups() map[string]Value {
	res := make(map[string]Value)
	for _, v := range store.groups {
		r := v.sampleAction()
		for j, w := range v.samplerMetrics {
			res[j] = Value{Typ: w.typ, Val: w.getter(r)}
		}
	}
	return res
}

func (store *Store) SampleOne() {
}

func (store *Store) readAllRefs() map[string]Value {
	res := make(map[string]Value)
	for k, v := range store.metrics {
		res[k] = Value{Typ: v.typ, Val: v.getter(nil)}
	}
	return res
}

func merge(a, b map[string]Value) map[string]Value {
	for k, v := range b {
		a[k] = v
	}
	return a
}
