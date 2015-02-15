package main

import (
    "github.com/adarqui/ekg-core-go"
    "fmt"
)

func main() {
    store := ekg_core.New()
    counter := store.CreateCounter("test_counter")
    counter.Inc()
    fmt.Println(counter.Read())
    counter.Add(10)
    fmt.Println(counter.Read())
    fmt.Println("done.")
}
