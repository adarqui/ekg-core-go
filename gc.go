package ekg_core

import (
    "runtime/debug"
    "runtime"
)


type GCMemStats struct {
    gcs *debug.GCStats
    mem *runtime.MemStats
}


func (store *Store) RegisterGCMetrics() {
    m := make(map[string]Metric)
    m["gcs.last_gc"] = Metric { getter: getGcsLastGC }
    m["gcs.next_gc"] = Metric { getter: getMemNextGC }
    m["gcs.num_gc"] = Metric { getter: getGcsNumGC }
    m["gcs.alloc"] = Metric { getter: getMemAlloc }
    m["gcs.total_alloc"] = Metric { getter: getMemTotalAlloc }
    m["gcs.sys"] = Metric { getter: getMemSys }
    m["gcs.lookups"] = Metric { getter: getMemLookups }
    m["gcs.mallocs"] = Metric { getter: getMemMallocs }
    m["gcs.frees"] = Metric { getter: getMemFrees }
    m["gcs.heap_alloc"] = Metric { getter: getMemHeapAlloc }
    m["gcs.heap_sys"] = Metric { getter: getMemHeapSys }
    m["gcs.heap_idle"] = Metric { getter: getMemHeapIdle }
    m["gcs.heap_inuse"] = Metric { getter: getMemHeapInuse }
    m["gcs.heap_released"] = Metric { getter: getMemHeapReleased }
    m["gcs.heap_objects"] = Metric { getter: getMemHeapObjects }
    m["gcs.pause_total_ns"] = Metric { getter: getMemPauseTotalNs }
    m["gcs.enable_gc"] = Metric { getter: getMemEnableGC }
    m["gcs.debug_gc"] = Metric { getter: getMemDebugGC }
    store.RegisterGroup(m, _getGCStats)
}


func getGCStats() GCMemStats {
    gcm := GCMemStats{}
    gcm.gcs = new(debug.GCStats)
    gcm.mem = new(runtime.MemStats)
    debug.ReadGCStats(gcm.gcs)
    runtime.ReadMemStats(gcm.mem)
    return gcm
}


func _getGCStats() interface{} {
    return getGCStats()
}


func getGcsNumGC(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.gcs.NumGC
}


func getGcsLastGC(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.gcs.LastGC
}


func getMemNextGC(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.NextGC
}


func getMemAlloc(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.Alloc
}


func getMemTotalAlloc(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.TotalAlloc
}


func getMemSys(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.Sys
}


func getMemLookups(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.Lookups
}


func getMemMallocs(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.Mallocs
}


func getMemFrees(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.Frees
}


func getMemHeapAlloc(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.HeapAlloc
}


func getMemHeapSys(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.HeapSys
}


func getMemHeapIdle(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.HeapIdle
}


func getMemHeapInuse(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.HeapInuse
}


func getMemHeapReleased(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.HeapReleased
}


func getMemHeapObjects(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.HeapObjects
}


func getMemPauseTotalNs(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.PauseTotalNs
}


func getMemEnableGC(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.EnableGC
}


func getMemDebugGC(gcm interface{}) interface{} {
    _gcm := gcm.(GCMemStats)
    return _gcm.mem.DebugGC
}
