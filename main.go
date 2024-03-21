package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type StationData struct {
	Name string
	Min float64
	Max float64
	Sum float64
	Occurence int
}


func main() {
	start := time.Now()
	
	file, err := os.Open("/Users/bebelino/Downloads/1brc-main/measurements.txt")

	if err != nil {
		log.Fatalf("There was an error opening the file:%v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	map1 := make(map[string]*StationData)

	for scanner.Scan() {
		line := scanner.Text()

		name := strings.Split(line, ";")

		temperature, _ := strconv.ParseFloat(name[1], 64)

		station, ok := map1[name[0]]

		if !ok {
			map1[name[0]] = &StationData{Name: name[0], Min: temperature, Max: temperature, Sum: temperature, Occurence: 1}
		} else {
			if temperature < station.Min {
				station.Min = temperature
			}
			if temperature > station.Max {
				station.Max = temperature
			}
			station.Sum += temperature
			station.Occurence++
		}
	}
	printResults(map1)

	t := time.Now()

	duration := t.Sub(start)

	fmt.Printf("the function took %s to run", duration)
}

func printResults(data map[string]*StationData) {
	result := make(map[string]*StationData, len(data))
	key := make([]string, 0, len(data))

	for _, value := range data {
		key = append(key, value.Name)
		result[value.Name] = value
	}
	sort.Strings(key)

	print("{")
	for _, keys := range key {
		v := result[keys]
		fmt.Printf("%s=%.2f/%.2f/%.2f/%d;",keys, v.Min,v.Sum/float64(v.Occurence),v.Max,v.Occurence)
	}
	print("}\n")
}



