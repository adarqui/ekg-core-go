package ekg_core

import (
    "runtime/debug"
)

func (store *Store) RegisterGCMetrics() {
    m := make(map[string]Metric)
    m["gc.last_gc"] = Metric {
        getter: getLastGC,
    }
    m["gc.num_gc"] = Metric {
        getter: getNumGC,
    }
    store.RegisterGroup(m, _getGCStats)
}


func getGCStats() debug.GCStats {
    gc := new(debug.GCStats)
    debug.ReadGCStats(gc)
    return *gc
}


func _getGCStats() interface{} {
    return getGCStats()
}


func getNumGC(gc interface{}) interface{} {
    _gc := gc.(debug.GCStats)
    return _gc.NumGC
}


func getLastGC(gc interface{}) interface{} {
    _gc := gc.(debug.GCStats)
    return _gc.LastGC
}
