package webserver

import (
	"fmt"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/memestream/memeserve/grpcserver"
	"golang.org/x/net/context"
)

// Server is the interface that the webserver must implement
type Server interface {
	Addr() string
	Serve()
}

// WebServer implements Server
type WebServer struct {
	lis       net.Listener
	webMux    *http.ServeMux
	proxyMux  *runtime.ServeMux
	ctxCancel context.CancelFunc
}

// New creates a new webserver
func New(port int, grpcServer grpcserver.Server) (*WebServer, error) {
	s := new(WebServer)

	var err error
	s.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	ctx, s.ctxCancel = context.WithCancel(ctx)

	s.webMux = http.NewServeMux()
	s.webMux.Handle("/v1", s.proxyMux)

	// TODO: Connect to the gRPC server

	return s, nil
}

// Addr returns the address that the webserver is listening on
func (s *WebServer) Addr() string {
	return s.lis.Addr().String()
}

// Serve Starts the web server
func (s *WebServer) Serve() {
	defer s.ctxCancel()
	glog.Infof("starting webserver on %s", s.Addr())
	http.Serve(s.lis, s.webMux)
}
