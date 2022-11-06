package main

import (
	"client/goservices"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// var opts []grpc.DialOption
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer conn.Close()
	client := goservices.NewRouteGuideClient(conn)
	myhello := goservices.Hello{
		Hello: "Test",
	}
	hello, _ := client.GetHello(context.Background(), &myhello)
	println(hello.Hello)

}
