package main

import (
	"log"
	"math"
	"strings"
	"time"

	"gopkg.in/zeromq/goczmq.v4"
)

func send_data(dealer *goczmq.Sock, data []byte, itr int) {
	start := time.Now()
	for i := 0; i < itr; i++ {
		dealer.SendFrame(data, goczmq.FlagNone)
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("Time took: %f", elapsed)
}

func main() {
	const base = 10_000
	const num_tests = 4
	const data_size = 1_0

	data := []byte(strings.Repeat("1", data_size))
	dealer, err := goczmq.NewPub("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()

	var iterations float64
	for i := 0; i < num_tests; i++ {
		iterations = base * math.Pow(10, float64(i))
		send_data(dealer, data, int(iterations))
	}
}
