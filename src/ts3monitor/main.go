package main

import (
	"flag"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	ts3Status = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "teamspeak_up",
		Help: "Current status of the TeamSpeak3.",
	})

	address = flag.String("addr", "localhost:10011", `-addr="localhost:10011"`)
)

func init() {
	// registers the gauge
	prometheus.MustRegister(ts3Status)
}

func main() {
	// starts the ping loop
	ping()
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
			ts3Status.Set(0)
			continue
		}
		// allocate byte scape for the respnse to be read into
		data := make([]byte, 32)

		b, _ := conn.Read(data)
		//checks the bytesread and the length of the read data for a response
		if b != 0 && len(data) != 0 {
			// the connection returned a string, therefore the server is online
			ts3Status.Set(1)
		}

	}
}
