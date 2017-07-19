package api

import (
	"math"
	"time"
)

type Sector struct {
	x    int
	y    int
	name string
}

const raidTimeOverhead int64 = int64(time.Second * 15)

func NewSector(x, y int, name string) *Sector {
	return &Sector{x, y, name}
}

func (origin *Sector) DistanceTo(destination *Sector) float64 {
	return math.Sqrt(math.Pow(float64(destination.x-origin.x), 2) + math.Pow(float64(destination.y-origin.y), 2))
}

func (origin *Sector) TravelTimeTo(destination *Sector, speed int) int64 {
	distance := origin.DistanceTo(destination)
	actualSpeed := float64(speed) / float64(time.Hour)
	travelTime := RoundToSecond(int64(distance / actualSpeed))
	travelTime = checkMaxTravelTime(travelTime)
	travelTime = checkMinTravelTime(travelTime)
	return travelTime
}

func (origin *Sector) RaidTimeTo(destination *Sector, speed int) int64 {
	travelTime := origin.TravelTimeTo(destination, speed)
	if travelTime > int64(time.Minute) && travelTime < int64(time.Hour*12) {
		travelTime += raidTimeOverhead
	}
	return travelTime
}

func RoundToSecond(base int64) int64 {
	hour := int64(time.Hour)
	minute := int64(time.Minute)
	second := int64(time.Second)
	var rounded int64
	for ; rounded+hour <= base; rounded += hour {
	}
	for ; rounded+minute <= base; rounded += minute {
	}
	for ; rounded+second <= base; rounded += second {
	}
	return rounded
}
func checkMinTravelTime(travelTime int64) int64 {
	if travelTime < int64(time.Minute) {
		travelTime = int64(time.Minute)
	}
	return travelTime
}

func checkMaxTravelTime(travelTime int64) int64 {
	if travelTime > int64(time.Hour*12) {
		travelTime = int64(time.Hour * 12)
	}
	return travelTime
}
