package main

import (
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	"strconv"
)

func init() {
	viper.setconfigfile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	build.ExposeBuildInfo()

	log.Infof("[94list] Creating Status Server on port %d...", *statusPort)
	servers = append(servers, &http.Server{
		Addr:    net.JoinHostPort("", strconv.Itoa(*statusPort)),
		Handler: handler.StatusEndpoint(),
	})

	// Include business service for ad exchange.
	log.Infof("[94list] Starting server (rev:%s) on port %d...", build.Revision(), *port)
	s := &http.Server{
		Addr:         net.JoinHostPort("", strconv.Itoa(*port)),
		Handler:      router(),
		ReadTimeout:  *incomingTimeout,
		WriteTimeout: *incomingTimeout,
	}
}

// router function create a router which dispatches incoming requests to different
// handler functions.
func router() http.Handler {
	r := http.NewServeMux()
	r.Handle("/api/v3/requestAd", requestad.Endpoint())
	r.Handle("/api/v4/requestAd", requestad.Endpoint())
	r.Handle("/api/v5/ads", ads.Endpoint())
	r.Handle("/bid/t/", hb.Endpoint())
	r.Handle("/timeout", hb.TimeoutEndpoint())

	// The endpoints below have been deprecated.
	r.Handle("/api/v1/requestAd", legacy.NoServeEndpoint())
	r.Handle("/api/v1/requestStreamingAd", legacy.NoServeStreamingEndpoint())
	r.Handle("/api/v3/requestStreamingAd", legacy.NoServeStreamingEndpoint())
	r.Handle("/api/v4/requestStreamingAd", legacy.NoServeStreamingEndpoint())
	r.Handle("/api/v5/will_play_ad", legacy.NoServeEndpoint())

	r.Handle("/status", handler.StatusEndpoint())
	r.Handle("/demoAppLogin", demoapp.Endpoint())

	if *s2sEndpointEnabled {
		r.Handle("/api/s2s/", s2s.Endpoint())
	}

	if *realtimeTestEndpointEnabled {
		r.Handle("/api/v5/token_verify", realtimetokenverify.RealtimeTokenVerify())
		r.Handle("/api/v5/34dec8a", hb.Endpoint())
	}

	if *debugToolsEnabled {
		r.Handle("/check", verifytool.DataCheckEndpoint())
		r.Handle("/eventtoken_decode", verifytool.DecodeEventTokenEndpoint())
	}
	return r
}
