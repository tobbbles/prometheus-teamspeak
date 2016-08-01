# A Prometheus monitor for TeamSpeak3

A simple [Prometheus](https://prometheus.io) tool for checking the status of a TeamSpeak3 server via pinging the ServerQuery port.

This is a GB-style project. You can [get GB here](https://getgb.io)

## Installing:
```
git clone https://github.com/mnzt/prometheus-teamspeak
cd prometheus-teamspeak
gb build
./bin/prometheus-teamspeak
```

## Usage

There are three environmental variables utilised:
- Address - The address of your TeamSpeak3 server (Be sure to specify the port with a colon).
`export ADDR="localhost:10011"`
- Interval (optional) - How often the TeamSpeak3 server is pinged.
`export INTERVAL=5`
- Port (optional) - The port on which this service runs on.
`export PORT=8010`

### Manually via binary
Simply specify the envvars above and run the service via `./bin/prometheus-teamspeak`.
Visiting [http://localhost:8010/metrics](http://127.0.0.1:8010/metrics) in your browser should present you with the metrics (if you specified a different port, navigate to that URL instead.)

### Via Docker

1. Pull the docker image via `docker pull mnzt/prometheus-teamspeak`
  - Alternatively build the image from source with `docker build -t prometheus-teamspeak .`
2. Run the container, ensuring to specify the environment variables. _NOTE:_ If you are running a TeamSpeak3 server and prometheus-teamspeak on the same host, be sure to set `--net=host`.
```
docker run -d \
            -e ADDR=ts.awesome.server:10011 \
            -e INTERVAL=5 \
            -p 1337:8010 \
            prometheus-teamspeak
```
