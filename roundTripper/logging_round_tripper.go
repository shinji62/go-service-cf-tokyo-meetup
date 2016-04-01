package roundTripper

import (
	"crypto/tls"
	"errors"
	"github.com/CrowdSurge/banner"
	"github.com/Pivotal-Japan/route-service-cf/headers"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
)

type LoggingRoundTripper struct {
	transport http.RoundTripper
	debug     bool
}

func NewLoggingRoundTripper(debug bool) *LoggingRoundTripper {
	return &LoggingRoundTripper{
		transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		debug: debug,
	}
}

// forward to the URL
// Send response back to the Router

func (lrt *LoggingRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	var err error
	var res *http.Response
	start := time.Now()
	if request.Host == "No Host" {
		return nil, errors.New("Invalid Header")
	}
	res, err = lrt.transport.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	bannerS := banner.PrintS(string(contents[:]))
	r := strings.NewReader(bannerS)
	body := ioutil.NopCloser(r)
	res.Body = body

	if lrt.debug {
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Fatalln(err.Error())
		}

		log.Printf("%q", dump)
		log.Printf("Time Elapsed RoundTrip %v", time.Since(start))

	}
	log.Printf("res %q", res.Body)
	res.ContentLength = int64(len(bannerS))
	res.Header.Set("Content-Length", strconv.Itoa(len(bannerS)))
	//Adding CF headers
	//res.Header.(headers.RouteServiceMetadata, request.Header.Get(headers.RouteServiceMetadata))
	res.Header.Add(headers.RouteServiceMetadata, request.Header.Get(headers.RouteServiceMetadata))
	res.Header.Add(headers.RouteServiceSignature, request.Header.Get(headers.RouteServiceSignature))
	return res, err
}
