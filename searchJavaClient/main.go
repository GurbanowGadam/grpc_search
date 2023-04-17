package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/GurbanowGadam/grpc_search/searchJava"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:3032", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := searchJava.NewSearchJavaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.SearchPeople(ctx, &searchJava.SearchText{})
	if err != nil {
		fmt.Println("SearchPeople => ", err)
	}
	fmt.Println(res)

}
