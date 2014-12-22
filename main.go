package main

import (
	"fmt"
	"bytes"
	"time"
)

type Train struct {
	Name string
	Destination string
	Status string
}

const traveling =  "TRAVELING"
const stop = "TRAVELING"

func (train *Train) goingToSummaryInfo() (string) {
	var buffer bytes.Buffer
	buffer.WriteString(train.Name)
	buffer.WriteString(" going to: ")
	buffer.WriteString(train.Destination)
	return buffer.String()
}

func stopSummaryInfo() () {
	var buffer bytes.Buffer
	buffer.WriteString(train.Name)
	buffer.WriteString(" going to: ")
	buffer.WriteString(train.Station.Name)
	return buffer.String()
}

func (train *Train) summaryInfo() (string) {
	if train.Status == traveling {
		return train.travelingSummaryInfo()
	} else if train.Status == stop {
		return train.stopSummaryInfo()
	} else {
		return train.goingToSummaryInfo()
	}
}

type Station struct {
	Name string
}
func startSimulation() {
	trainOne := Train{Name: "ICE 899", Destination: "Oberhausen", Status: traveling}
	trainTwo := Train{Name: "ICE 746", Destination: "Dortmund", Status: traveling}

	trains := make(map[string]Train)
	trains[trainOne.Name] = trainOne
	trains[trainTwo.Name] = trainTwo

	for {
		time.Sleep(3000 * time.Millisecond)
		for _, v := range trains {
			fmt.Println(v.summaryInfo())
		}
	}
}

func main() {
	startSimulation()
}
