package lpr

import (
	"encoding/json"
	"math/big"
	"os"
	"sort"
	"strings"
	"sync/atomic"
	"time"
)

var (
	loaded   atomic.Bool
	events1y []lprEvent
	events5y []lprEvent
)

type lprEvent struct {
	date time.Time
	rate *big.Float
}

// Get5Y 返回5年期 LPR。LPR 随时间波动，t 对应要查询的时间点
func Get5Y(t time.Time) *big.Float {
	ensureLoaded()
	return binarySearch(t, events5y)
}

func Get1Y(t time.Time) *big.Float {
	ensureLoaded()
	return binarySearch(t, events1y)
}

func ensureLoaded() {
	loadEventsData()
}

func loadEventsData() {
	if loaded.Load() {
		return
	}

	data, err := os.ReadFile("events.json")
	if err != nil {
		panic(err)
	}

	type LRPEvents struct {
		LPR struct {
			Source string            `json:"source"`
			Y5     map[string]string `json:"5Y"`
			Y1     map[string]string `json:"1Y"`
		} `json:"lpr"`
	}

	var events LRPEvents
	err = json.Unmarshal(data, &events)
	if err != nil {
		panic(err)
	}

	parseEvent := func(kvs map[string]string, es *[]lprEvent) {
		for k, v := range kvs {
			t, err := time.ParseInLocation("2006-01-02", k, time.Local)
			if err != nil {
				panic(err)
			}
			rate, _, err := big.ParseFloat(strings.ReplaceAll(v, "%", ""), 10, 256, big.ToNearestEven)
			if err != nil || rate == nil {
				panic(err)
			}
			rate = new(big.Float).Mul(rate, big.NewFloat(0.01))
			*es = append(*es, lprEvent{
				date: t,
				rate: rate,
			})
		}
		sort.Slice(*es, func(i, j int) bool {
			return (*es)[i].date.Before((*es)[j].date)
		})

	}

	parseEvent(events.LPR.Y1, &events1y)
	parseEvent(events.LPR.Y5, &events5y)

	loaded.Store(true)
}

func binarySearch(t time.Time, events []lprEvent) *big.Float {
	lo, hi := 0, len(events)-1

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if events[mid].date.After(t) {
			hi = mid - 1
		} else {
			lo = mid
		}
	}

	return events[lo].rate
}
