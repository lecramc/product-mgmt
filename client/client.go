package main

import (
	"context"
	"log"
	"time"

	pb "example.com/orders-management/orders"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_orders = make(map[int32]string)
	new_orders[786745] = "Patates"
	new_orders[786754] = "Courgettes"
	for ref, name := range new_orders {
		r, err := c.CreateNewOrder(ctx, &pb.NewOrder{Ref: ref, Name: name})
		if err != nil {
			log.Fatalf("could not create order: %v", err)
		}
		log.Printf(`Order Details:
			ID: %d
			REF: %d
			NAME: %s`, r.GetId(), r.GetRef(), r.GetName())

	}
}
