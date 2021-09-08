package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/orders-management/orders"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type OrderManagementServer struct {
	pb.UnimplementedOrderManagementServer
}

func (s *OrderManagementServer) CreateNewOrder(ctx context.Context, in *pb.NewOrder) (*pb.Order, error) {
	log.Printf("received: %v", in.GetName())
	var order_id int32 = int32(rand.Intn(100))
	return &pb.Order{Id: order_id, Ref: in.GetRef(), Name: in.GetName()}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &OrderManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
