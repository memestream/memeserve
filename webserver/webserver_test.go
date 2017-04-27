package webserver

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

type fakeGrpcServer struct{}

func (fakeGrpcServer) Addr() string {
	return "[::]:23456"
}

func (fakeGrpcServer) Serve() {}

func TestNew(t *testing.T) {
	// Check if an error is returned if a port is already used
	l, _ := net.Listen("tcp", ":")
	defer l.Close()

	http.ListenAndServe(l.Addr().String(), nil)

	port := getPort(l)

	_, err := New(port, &fakeGrpcServer{})
	if err == nil {
		t.Errorf("Expected error about port %d being occupied. Got nil", port)
	}
}

func TestServe(t *testing.T) {
	s, err := New(0, new(fakeGrpcServer))
	if err != nil {
		t.Errorf("Unexpected error %v, should be nil", err)
	}

	go s.Serve()

	_, err = net.Dial("tcp", s.Addr())
	if err != nil {
		t.Errorf("Unexpected error %v, should be nil", err)
	}
}

func getPort(l net.Listener) int {
	sp := strings.Split(l.Addr().String(), ":")
	p, _ := strconv.Atoi(sp[len(sp)-1])
	return p
}
