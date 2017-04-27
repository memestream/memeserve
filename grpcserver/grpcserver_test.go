package grpcserver

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"google.golang.org/grpc"
)

func TestNew(t *testing.T) {
	// Check if an error is returned if a port is already used
	l, _ := net.Listen("tcp", ":0")
	defer l.Close()

	http.ListenAndServe(l.Addr().String(), nil)

	port := getPort(l)

	_, err := New(port)
	if err == nil {
		t.Errorf("Expected error about port %d being occupied. Got nil", port)
	}
}

func TestAddr(t *testing.T) {
	lis, _ := net.Listen("tcp", ":0")
	s := &GrpcServer{
		lis: lis,
	}

	if s.Addr() != lis.Addr().String() {
		t.Errorf("Server address should be %s, have %s", lis.Addr(), s.Addr())
	}
}

func TestServe(t *testing.T) {
	s, err := New(0)
	if err != nil {
		t.Errorf("Unexpected error %v, should be nil", err)
	}

	go s.Serve()

	_, err = grpc.Dial(s.Addr(), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Unexpected error %v, should be nil", err)
	}
}

func getPort(l net.Listener) int {
	sp := strings.Split(l.Addr().String(), ":")
	p, _ := strconv.Atoi(sp[len(sp)-1])
	return p
}
