package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	address   = flag.String("address", "localhost:10011", `-address="localhost:10011"`)
	ts3Status = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "teamspeak_up",
		Help: "Current status of the TeamSpeak3.",
	})
)

func init() {
	// registers the gauge
	prometheus.MustRegister(ts3Status)
}

func main() {
	flag.Parse()
	//
	fmt.Println("Starting listener on: ", *address)
	// starts the ping loop
	go ping()
	// attaches the prometheus handler to the /metrics path
	http.Handle("/metrics", prometheus.Handler())
	// serves the application
	http.ListenAndServe(":8080", nil)

}

func ping() {
	for {
		// set a granularity of 3 seconds
		time.Sleep(3 * time.Second)
		// 2 second timeout allowing for some latency.
		conn, err := net.DialTimeout("tcp", *address, 200*time.Millisecond)
		if err != nil {
			// the connection timedout, therefor the server is down
			fmt.Println(err)
			ts3Status.Set(0)
			continue
		}
		// allocate byte scape for the respnse to be read into
		data := make([]byte, 32)

		b, _ := conn.Read(data)
		//checks the bytesread and the length of the read data for a response
		if b != 0 && len(data) != 0 {
			fmt.Println("setting status to 1")
			// the connection returned a string, therefore the server is online
			ts3Status.Set(1)
		}

	}
}
