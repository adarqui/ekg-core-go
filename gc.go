package ekg_core

import (
	"runtime"
	"runtime/debug"
)

type GCMemStats struct {
	gcs *debug.GCStats
	mem *runtime.MemStats
}

func (store *Store) RegisterGCMetrics() {
	m := make(map[string]Metric)
	registerMetric(m, "gcs.last_gc", getGcsLastGC, TIMESTAMP)
	registerMetric(m, "gcs.next_gc", getMemNextGC, GAUGE)
	registerMetric(m, "gcs.num_gc", getGcsNumGC, COUNTER)
	registerMetric(m, "gcs.alloc", getMemAlloc, GAUGE)
	registerMetric(m, "gcs.total_alloc", getMemTotalAlloc, COUNTER)
	registerMetric(m, "gcs.sys", getMemSys, COUNTER)
	registerMetric(m, "gcs.lookups", getMemLookups, COUNTER)
	registerMetric(m, "gcs.mallocs", getMemMallocs, COUNTER)
	registerMetric(m, "gcs.frees", getMemFrees, COUNTER)
	registerMetric(m, "gcs.heap_alloc", getMemHeapAlloc, GAUGE)
	registerMetric(m, "gcs.heap_sys", getMemHeapSys, GAUGE)
	registerMetric(m, "gcs.heap_idle", getMemHeapIdle, GAUGE)
	registerMetric(m, "gcs.heap_inuse", getMemHeapInuse, GAUGE)
	registerMetric(m, "gcs.heap_released", getMemHeapReleased, GAUGE)
	registerMetric(m, "gcs.heap_objects", getMemHeapObjects, COUNTER)
	registerMetric(m, "gcs.pause_total_ns", getMemPauseTotalNs, COUNTER)
	registerMetric(m, "gcs.enable_gc", getMemEnableGC, BOOL)
	registerMetric(m, "gcs.debug_gc", getMemDebugGC, BOOL)
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
