package main

import (
	"context"
	pb "grpc-sorter/protobuf/sorter"
	qsort "grpc-sorter/sorter_server/qsort"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) QuickSort(ctx context.Context, in *pb.Numbers) (*pb.Numbers, error) {
	qsort.QuickSort(in.Numbers)
	return in, nil
}
func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterSorterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
