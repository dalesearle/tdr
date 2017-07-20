package api

import "time"

type TravelStats struct {
	x          int
	y          int
	name       string
	travelTime time.Duration
}

func NewTravelStats(x, y int, name string, travelTime time.Duration) *TravelStats {
	return &TravelStats{x, y, name, travelTime}
}

func (travelStats *TravelStats) GetX() int {
	return travelStats.x
}

func (travelStats *TravelStats) GetY() int {
	return travelStats.y
}

func (travelStats *TravelStats) GetName() string {
	return travelStats.name
}
