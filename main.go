package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/shinji62/go-service-cf-tokyo-meetup/roundTripper"
	"github.com/shinji62/route-service-cf/proxy"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var (
	port  = kingpin.Flag("port", "Port to listen").Envar("PORT").Short('p').Required().Int()
	debug = kingpin.Flag("debug", "Port to listen").Envar("DEBUG").Short('d').Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	httpClient := &http.Client{}
	roundTripper := roundTripper.NewLoggingRoundTripper(*debug)
	proxy := proxy.NewReverseProxy(roundTripper, httpClient, *debug)

	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%v", *port), proxy))
}
