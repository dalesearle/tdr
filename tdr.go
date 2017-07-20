package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/dalesearle/tdr/api"
)

func main() {
	baseSector := createBaseSector()
	sectorData := getSectorData()
	targetSectors := createTargetSectors(sectorData)
	travelStats := createTravelStats(baseSector, targetSectors)
	writeTravelStats(travelStats)
}

func createBaseSector() api.Sector {
	baseX, err := strconv.Atoi(os.Args[1])
	if err != nil {
		handleInputError()
	}
	baseY, err := strconv.Atoi(os.Args[2])
	if err != nil {
		handleInputError()
	}
	return *api.NewSector(baseX, baseY, "main")
}

func handleInputError() {
	fmt.Println("Usage: SectorX SectorY Path_to_CSV_file Path_to_storage_directory e.g. 570 -3095 c:\\raids.csv c:\\")
}

func getSectorData() []string {
	path := os.Args[3]
	sectorData := []string{}
	data, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	if err != nil {
		panic(err)
	}
	for true {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		sectorData = append(sectorData, string(line[:]))
	}
	return sectorData
}

func createTargetSectors(sectorData []string) []api.Sector {
	sectors := []api.Sector{}
	for _, data := range sectorData {
		sd := strings.Split(data, ",")
		x := atoi(sd[0])
		y := atoi(sd[1])

		sectors = append(sectors, *api.NewSector(x, y, sd[2]))
	}
	return sectors
}

func atoi(str string) int {
	igr, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return igr
}

func createTravelStats(baseSector api.Sector, targetSectors []api.Sector) []api.TravelStats {
	travelStats := []api.TravelStats{}
	for _, targetSector := range targetSectors {
		travelTime := time.Duration(baseSector.TravelTimeTo(&targetSector, 480))
		travelStats = append(travelStats, *api.NewTravelStats(targetSector.GetX(), targetSector.GetY(), targetSector.GetName(), travelTime))
	}
	return travelStats
}

func writeTravelStats(travelStats []api.TravelStats) {
	path := os.Args[4] + "/tdr_travel.csv"
	data, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer data.Close()
	writer := bufio.NewWriter(data)
	writer.WriteString("test")
}
