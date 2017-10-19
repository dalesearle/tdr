package traveltime

import "time"

type TravelStat struct {
	x          int
	y          int
	name       string
	travelTime time.Duration
}

type TravelStats []TravelStat

func NewTravelStats(x, y int, name string, travelTime time.Duration) *TravelStat {
	return &TravelStat{x, y, name, travelTime}
}

func (travelStats *TravelStat) GetX() int {
	return travelStats.x
}

func (travelStats *TravelStat) GetY() int {
	return travelStats.y
}

func (travelStats *TravelStat) GetName() string {
	return travelStats.name
}

func (travelStats *TravelStat) GetTravelTime() time.Duration {
	return travelStats.travelTime
}

func NewTravelStatsSlice() TravelStats {
	return make(TravelStats, 0)
}
func (ts TravelStats) Len() int {
	return len(ts)
}
func (ts TravelStats) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}
func (ts TravelStats) Less(i, j int) bool {
	return ts[i].GetTravelTime() < ts[j].GetTravelTime()
}
