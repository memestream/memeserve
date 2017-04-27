package grpcserver

import (
	"fmt"
	"net"

	"github.com/golang/glog"

	"google.golang.org/grpc"
)

// Server is the interface which extends the standard proto server and
// also adds a couple of calls specific to the memeserve
type Server interface {
	// Addr returns the address of the gRPC server
	Addr() string

	// Serve starts listening
	Serve()
}

// GrpcServer implements Server
type GrpcServer struct {
	srv *grpc.Server
	lis net.Listener
}

// New creates a new GrpcServer
func New(port int) (*GrpcServer, error) {
	s := new(GrpcServer)

	var err error
	s.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	// TODO: Register the API proto
	s.srv = grpc.NewServer()

	return s, nil
}

// Addr returns the address that the gRPC server is listening on
func (s *GrpcServer) Addr() string {
	return s.lis.Addr().String()
}

// Serve starts the gRPC server
func (s *GrpcServer) Serve() {
	glog.Infof("Starting gRPC server on %s", s.Addr())
	if err := s.srv.Serve(s.lis); err != nil {
		glog.Fatalf("Unable to start gRPC server: %v", err)
	}
}
