package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/product-management/product"
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
	c := pb.NewProductManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_Product = make(map[string]string)
	new_Product["REF786745"] = "Patates"
	new_Product["REF786754"] = "Courgettes"
	for ref, label := range new_Product {
		r, err := c.CreateNewProduct(ctx, &pb.NewProduct{Ref: ref, Label: label})
		if err != nil {
			log.Fatalf("could not create Product: %v", err)
		}
		log.Printf(`Product Details:
			ID: %d
			REF: %s
			LABEL: %s`, r.GetId(), r.GetRef(), r.GetLabel())
	}
	params := &pb.GetProductsParams{}
	r, err := c.GetProducts(ctx, params)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Print("\nUSER LIST: \n")
	fmt.Printf("r.GetUsers(): %v\n", r.GetProducts())
}
