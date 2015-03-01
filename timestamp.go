package ekg_core

import (
	"sync"
	"time"
)

type Timestamp struct {
	mtx *sync.Mutex
	v   time.Time
}

// Create a new, time initialized, timestamp.
func newTimestamp() *Timestamp {
	timestamp := new(Timestamp)
	timestamp.mtx = &sync.Mutex{}
	timestamp.Stamp()
	return timestamp
}

// Get the current value of the timestamp.
func (timestamp *Timestamp) Read() time.Time {
	timestamp.mtx.Lock()
	v := timestamp.v
	timestamp.mtx.Unlock()
	return v
}

func (timestamp *Timestamp) ReadI(v interface{}) interface{} {
	return timestamp.Read()
}

// Stamp the time.
func (timestamp *Timestamp) Stamp() {
	timestamp.mtx.Lock()
	timestamp.v = time.Now()
	timestamp.mtx.Unlock()
}
