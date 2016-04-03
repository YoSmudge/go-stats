package gostats

import (
	"sync"
	"time"
)

type counter struct {
	LastUpdate time.Time
	LastValue  int64
}

var counterTracker = make(map[string]counter)
var trackerMutex = &sync.Mutex{}

func perSecondCounter(name string, value int64) float64 {
	trackerMutex.Lock()
	defer trackerMutex.Unlock()
	updateTime := time.Now()
	var tracker counter

	tracker, found := counterTracker[name]
	if !found {
		tracker.LastUpdate = updateTime
		tracker.LastValue = value
		counterTracker[name] = tracker
	}

	secondsSince := updateTime.Sub(tracker.LastUpdate).Seconds()
	valueDiff := value - tracker.LastValue

	tracker.LastUpdate = updateTime
	tracker.LastValue = value
	counterTracker[name] = tracker

	if secondsSince == 0 {
		return float64(0)
	}

	changePerSecond := float64(valueDiff) / secondsSince
	return changePerSecond
}
