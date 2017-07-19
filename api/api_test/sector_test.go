package api_test

import (
	"math"
	"tdr/api"
	"testing"
	"time"
)

var mainSector = api.NewSector(570, -3094, "Main")
var distanceData = map[string]struct {
	x  int
	y  int
	fn func() float64
}{
	"TestOne": {571, -3095, func() float64 { return math.Sqrt(2) }},
	"TestTwo": {572, -3096, func() float64 { return math.Sqrt(8) }},
}

var roundingData = map[string]struct {
	value    int64
	expected int64
}{
	"NoChange": {int64(time.Hour), int64(time.Hour)},
	"Over":     {int64(time.Hour + time.Millisecond), int64(time.Hour)},
	"Under":    {int64(time.Hour - time.Millisecond), int64(time.Hour - time.Second)},
}

var travelData = map[string]struct {
	x        int
	y        int
	speed    int
	expected int64
}{
	"Min_Drone": {571, -3095, 480, int64(time.Minute)},
	"Drone":     {70, -2594, 480, int64(time.Hour + (time.Minute * 28) + (time.Second * 23))},
	"Max_Drone": {570, 3100, 480, int64(time.Hour * 12)},
}

var raidData = map[string]struct {
	x        int
	y        int
	speed    int
	expected int64
}{
	"Min_Exterminator": {571, -3095, 600, int64(time.Minute)},
	"Exterminator":     {70, -2594, 600, int64(time.Hour + (time.Minute * 10) + (time.Second * 57))},
	"Max_bc":           {570, 3490, 360, int64(time.Hour * 12)},
}

func TestDistance(t *testing.T) {
	for key, data := range distanceData {
		testSector := api.NewSector(data.x, data.y, "")
		expected := data.fn()
		result := testSector.DistanceTo(mainSector)
		if result != expected {
			t.Errorf("[%s]: expected %f, got %f", key, expected, result)
		}
	}
}

func TestSecondRounding(t *testing.T) {
	for key, data := range roundingData {
		result := api.RoundToSecond(data.value)
		if result != data.expected {
			t.Errorf("[%s]: expected %d, got %d", key, data.expected, result)
		}
	}
}

func TestTravelTime(t *testing.T) {
	for key, data := range travelData {
		destination := api.NewSector(data.x, data.y, "Test")
		result := mainSector.TravelTimeTo(destination, data.speed)
		if result != data.expected {
			t.Errorf("[%s]: expected %d, got %d", key, data.expected, result)
		}
	}
}

func TestRaidTime(t *testing.T) {
	for key, data := range raidData {
		destination := api.NewSector(data.x, data.y, "Test")
		result := mainSector.RaidTimeTo(destination, data.speed)
		if result != data.expected {
			t.Errorf("[%s]: expected %d(%v), got %d (%v)", key, data.expected, time.Duration(data.expected), result, time.Duration(result))
		}
	}
}
