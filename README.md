# Go-Stats

Generic GoLang internals instrumentation

[![Build Status](https://travis-ci.org/samarudge/go-stats.svg?branch=master)](https://travis-ci.org/samarudge/go-stats)
[![GitHub license](https://img.shields.io/github/license/mashape/apistatus.svg)]()

![](http://f.cl.ly/items/3q3K381r2D0K3O3S2H0Q/Screen%20Shot%202016-04-03%20at%2018.56.08.png?v=9fe2d412)

Usage;

```go
package main

import "github.com/samarudge/go-stats"

func main(){
	gostats.Start("statsd-host:8125", 10, "application-name")
}
```

For a sample Grafana dashboard see graphite.json

Metrics exported;

| Metric                     | Source                           | Description                            | Unit               |
|----------------------------|----------------------------------|----------------------------------------|--------------------|
| cgo.calls                  | runtime.NumCgoCall()             | Number of Cgo Calls                    | calls per second   |
| gc.pauseTimeMs             | runtime.ReadMemStats             | Pause time of last GC run              | MS                 |
| gc.pauseTimeNs             | runtime.ReadMemStats             | Pause time of last GC run              | NS                 |
| gc.period                  | runtime.ReadMemStats             | Time between last two GC runs          | MS                 |
| gc.perSecond               | runtime.ReadMemStats             | Number of GCs per second               | runs per second    |
| goroutines.total           | runtime.NumGoroutine()           | Number of currently running goroutines | total              |
| memory.counters.Frees      | runtime.ReadMemStats.Frees       | Number of frees issued to the system   | frees per second   |
| memory.counters.Mallocs    | runtime.ReadMemStats.Mallocs     | Number of Mallocs issued to the system | mallocs per second |
| memory.heap.Idle           | runtime.ReadMemStats.HeapIdle    | Memory on the heap not in use          | bytes              |
| memory.heap.InUse          | runtime.ReadMemStats.HeapInuse   | Memory on the heap in use              | bytes              |
| memory.objects.HeapObjects | runtime.ReadMemStats.HeapObjects | Total objects on the heap              | # Objects          |
| memory.summary.Alloc       | runtime.ReadMemStats.Alloc       | Total bytes allocated                  | bytes              |
| memory.summary.System      | runtime.ReadMemStats.HeapSys     | Total bytes acquired from system       | bytes              |

More documentation coming soon...