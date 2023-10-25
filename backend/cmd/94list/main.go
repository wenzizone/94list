package main

import (
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/wenzizone/94list/backend/internal/build"
	"github.com/wenzizone/94list/backend/internal/router/handler/api"
	"log/slog"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

var (
	port             = flag.Int("port", 7000, "Port for which the service will run on.")
	statusPort       = flag.Int("statusport", 6688, "Port of which the status service will run on.")
	debugPort        = flag.Int("debugport", -1, "Port of which the service will export debug information.")
	blockProfileRate = flag.Int("blockprofilerate", 0, "Rate at which the profiler profiles for blocking contentions; see 'go doc runtime.SetBlockProfileRate'.")
)

func init() {
	flag.Parse()
}

// incomingTimeout flag contains the duration that the server will wait for activities over an incoming
// connection from the Vungle SDK before timing out and move on.
// it configures a commandline flag that configures the router to time out the
// connections by a different interval. For example,
//
//	-connectiontimeout=30s
//
// will timeout each connection in 30 seconds.
var incomingTimeout = flag.Duration(
	"incoming_timeout",
	30*time.Second,
	"Timeout duration for which the server allows before timing out and move on.",
)

// noop defines a function to do nothing.
func noop() {
	// Do nothing as designed.
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	//defer shutdown()
	build.ExposeBuildInfo()
	servers := make([]*http.Server, 0, 4)
	/*
		log.Info("[94list] Creating Status Server on port %d...", *statusPort)
		servers = append(servers, &http.Server{
			Addr:    net.JoinHostPort("", strconv.Itoa(*statusPort)),
			Handler: handler.StatusEndpoint(),
		})
	*/
	// Include business service for ad exchange.
	log.Info("[94list] Starting server (rev:%s) on port %d...", build.Revision(), *port)
	s := &http.Server{
		Addr:         net.JoinHostPort("", strconv.Itoa(*port)),
		Handler:      router(),
		ReadTimeout:  *incomingTimeout,
		WriteTimeout: *incomingTimeout,
	}

	servers = append(servers, s)

	// Include debug service if debugPort is defined.
	if *debugPort > 0 {
		log.Info("[94list] Creating Debug server on port %d...", *debugPort)
		runtime.SetBlockProfileRate(*blockProfileRate)
		debugServer := &http.Server{
			Addr:    net.JoinHostPort("", strconv.Itoa(*debugPort)),
			Handler: http.DefaultServeMux,
		}
		servers = append(servers, debugServer)
	}
	if err := gracehttp.Serve(servers...); err != nil {
		log.Error("[94list] Failed to serve: %v.", err)
	}
}

// router function create a router which dispatches incoming requests to different
// handler functions.
func router() http.Handler {
	r := http.NewServeMux()
	r.Handle("/api/getList", api.getList)
	//r.Handle("/api/getSign", api.getSign)
	return r
}
