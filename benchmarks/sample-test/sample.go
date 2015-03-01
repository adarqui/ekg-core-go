// Perform 100,000 atomic increments using 100 concurrent writers.
package main

import (
	"fmt"
	"github.com/adarqui/ekg-core-go"
)

func main() {
	store := ekg_core.New()
	_ = store.CreateCounter("test.counter")
	_ = store.CreateGauge("test.gauge")
	store.CreateLabel("test.label").Set("Hello.")
	_ = store.CreateDistribution("test.distribution")
	_ = store.CreateTimestamp("test.timestamp")
    _ = store.CreateBool("test.bool")

	store.RegisterGCMetrics()

	samples := store.SampleAll()
	fmt.Println(samples)
}
