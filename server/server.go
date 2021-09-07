package main 

import(
	"context"
	"log"
	"math/rand"
	"net"

	pb "./orders-management/orders/orders"
	"google.golang.org/grpc"
)

const (
	port := ":50051"
)

type OrderManagementServer struct {
	pb.UnimplementedOrderManagementServer
}