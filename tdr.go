package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"tdr/api"
	"bufio"
)

func main() {
	baseX, err := strconv.Atoi(os.Args[1])
	if err != nil {
		handleInputError()
	}
	baseY, err := strconv.Atoi(os.Args[2])
	if err != nil {
		handleInputError()
	}
	csvPath := os.Args[3]

	readCSV(csvPath)
	fmt.Println(baseX, " ", baseY, " ", csvPath)
}

func handleInputError() {

	fmt.Println("Usage: Sector x Sector y Path to CSV file e.g. 570 -3095 c:\raids.csv")
}

func readCSV(path string) []api.Sector {
	sectors := []api.Sector{}
	data, err := os.Open(path)
	defer data.Close()
	reader := bufio.NewReader(data)
	if err != nil {
		panic(err)
	}
	for true {
		line, err := reader.ReadLine('\n')
		fmt.Println(string(line))
	}
	fmt.Println(string(data))
	return sectors

}
