package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	address  = os.Getenv("ADDR")
	port     = os.Getenv("PORT")
	interval = os.Getenv("INTERVAL")

	ts3Status = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "teamspeak",
		Name:      "is_up",
		Help:      "Whether TeamSpeak3 server is up. (1 = yes, 0 = no)",
	})
)

func main() {
	// Parse our flags
	prometheus.MustRegister(ts3Status)

	if address == "" {
		fmt.Println(ErrNoAddress)
		os.Exit(-1)
	}

	if port == "" {
		port = "8000"
	}

	if interval == "" {
		interval = "5"
	}

	fmt.Printf("Listening to TeamSpeak3 server at: %s\n", address)

	// starts the ping loop
	go ping()

	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":8010", nil)

}

func ping() {
	for {
		conn, err := net.DialTimeout("tcp", address, 200*time.Millisecond)
		if err != nil {
			// Most likely the connection timed out.
			ts3Status.Set(0)
			fmt.Println(err)
			continue
		}

		// Allocate a byteslice and read the response into it
		data := make([]byte, 32)
		b, _ := conn.Read(data)

		if b >= 1 && len(data) != 0 {
			// The connected returned, therefore we can assume TS3 is up
			ts3Status.Set(1)
		}

		// Sleep for the specified period

		intrvl, err := strconv.ParseInt(interval, 0, 8)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		<-time.After(time.Duration(intrvl) * time.Second)
	}
}

var (
	// ErrNoAddress is returned when no teamspeak3 server address is specified
	ErrNoAddress = errors.New("no teamspeak3 server address given")
)
