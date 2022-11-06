package main

import (
	"context"
	"flag"
	"fmt"
	srv "goservices/goservices"
	"log"
	"net"

	"google.golang.org/grpc"
)

type routeGuideServer struct {
	srv.UnimplementedRouteGuideServer
}

func newrouteGuideServer() *routeGuideServer {
	return &routeGuideServer{}
}

func (s *routeGuideServer) GetHello(ctx context.Context, point *srv.Hello) (*srv.Hello, error) {
	h := srv.Hello{}
	h.Hello = "Go + " + point.Hello
	return &h, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 4040))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	srv.RegisterRouteGuideServer(grpcServer, newrouteGuideServer())
	grpcServer.Serve(lis)
}
