// Perform 100,000 atomic increments using 100 concurrent writers.
package main

import (
    "github.com/adarqui/ekg-core-go"
    "fmt"
    "time"
)

func main() {
    store := ekg_core.New()
    counter := store.CreateCounter("test.counter")

    for n := 1; n <= 100; n++ {
        go func() {
            for iters := 1 ; iters <= 100000 ; iters++ {
                counter.Inc()
            }
        }()
    }

    time.Sleep(2 * time.Second)
    fmt.Println(counter.Read())
    fmt.Println("done.")
}
