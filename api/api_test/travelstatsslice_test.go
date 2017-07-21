package api

import (
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/dalesearle/tdr/api"
)

func TestSort(t *testing.T) {
	travelStatsSlice := api.NewTravelStatsSlice()
	travelStatsSlice = append(travelStatsSlice, *api.NewTravelStats(0, 0, "4", time.Hour))
	travelStatsSlice = append(travelStatsSlice, *api.NewTravelStats(0, 0, "3", time.Minute))
	travelStatsSlice = append(travelStatsSlice, *api.NewTravelStats(0, 0, "2", time.Second))
	travelStatsSlice = append(travelStatsSlice, *api.NewTravelStats(0, 0, "1", time.Millisecond))
	sort.Sort(travelStatsSlice)
	for i := 1; i <= len(travelStatsSlice); i++ {
		actual := strconv.Itoa(i)
		if travelStatsSlice[i-1].GetName() != actual {
			t.Errorf("expected %s, got %s", travelStatsSlice[i].GetName(), actual)
		}
	}
}
