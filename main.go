package main

import (
	"flag"
	"log"

	"github.com/golang/glog"
	"github.com/memestream/memeserve/grpcserver"
	"github.com/memestream/memeserve/webserver"
)

var (
	grpcPortFlag = flag.Int("grpcport", 4200, "Port to start gRPC server on")
	webPortFlag  = flag.Int("port", 8080, "Port for the standard web server")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Info("Starting memeserve")

	gSrv, err := grpcserver.New(*grpcPortFlag)
	if err != nil {
		log.Fatalf("Unable to start the gRPC server: %v", err)
	}

	wSrv, err := webserver.New(*webPortFlag, gSrv)
	if err != nil {
		log.Fatalf("Unable to start the web server: %v", err)
	}

	go gSrv.Serve()
	go wSrv.Serve()

	for {
	}
}
