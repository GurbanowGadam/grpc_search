package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/GurbanowGadam/grpc_search/searchJava"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3032, "The server port")
)

type SearchJavaServer struct {
	searchJava.UnimplementedSearchJavaServer
}

// Implment the notes.NotesServer interface
func (s *SearchJavaServer) SearchJavaPeople(ctx context.Context, n *searchJava.SearchText) (*searchJava.SearchResult, error) {

	return &searchJava.SearchResult{Uuid: "asnjdaksjdhkjaskdka", Username: "GadamGurban", Fullname: "Gadam"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	searchJava.RegisterSearchJavaServer(s, &SearchJavaServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
