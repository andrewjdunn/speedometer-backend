package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/andrewjdunn/speedometer-backend/database"
	"github.com/andrewjdunn/speedometer-backend/graph"
	"github.com/andrewjdunn/speedometer-backend/record"
)

var timeNextTestDue time.Time = time.Now()

func recordSpeed() {

	if timeNextTestDue.Before(time.Now()) {
		log.Println("Testing speed")
		record, err := record.Now()
		if err != nil {
			timeNextTestDue = time.Now().Add(time.Duration(time.Minute))
			log.Printf("Could not test speed [%v] Trying again at %v\n", err, timeNextTestDue)
		} else {
			fmt.Printf("Time: %s, Latency: %s, Download: %f Upload: %f, Distance: %f, Ping OK %v\n", record.TimeStamp, record.Latency, record.DownloadSpeed, record.UploadSpeed, record.Distance, record.PingOk)
			id, err := database.StoreRecord(record)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Record added %v\n", id)
			}
			timeNextTestDue = time.Now().Add(time.Minute * time.Duration((15 + rand.Intn(44))))
			fmt.Printf("Next test will be %v\n", timeNextTestDue)
		}
	}
}

func main() {

	fmt.Println("Starting graph")
	go func() { graph.Main() }()

	fmt.Println("Record first speed")
	recordSpeed()
	d, _ := time.ParseDuration("1m")
	ticker := time.NewTicker(d)
	done := make(chan bool)

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			recordSpeed()
		}
	}
}
