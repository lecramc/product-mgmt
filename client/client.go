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

func Menu() int32 {
	var user_choice int32
	var opt []string
	opt = append(opt, "Lister les produits", "Ajouter un produit", "Afficher un produit", "Supprimer un produit")
	fmt.Printf("Que souhaitez vous faire ?\n")
	for i, value := range opt {
		fmt.Printf("%d - %s ?\n", i+1, value)
	}

	fmt.Scan(&user_choice)
	return user_choice
}

func AddProduct() map[string]string {
	new_products := make(map[string]string)
	add_product := true
	var ref, product, one_more string

	for add_product {
		fmt.Printf("Saisir une reference : ")
		fmt.Scan(&ref)
		fmt.Printf("Saisir un nom de produit : ")
		fmt.Scan(&product)

		fmt.Printf("Ajouter un second produit ? (oui/non) : ")
		fmt.Scan(&one_more)
		switch one_more {
		case "non":
			add_product = false

		default:
			add_product = true
		}
		new_products[ref] = product
	}
	fmt.Println("map:", new_products)
	return new_products
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	user_action := Menu()
	switch user_action {
	case 2:
		new_products := make(map[string]string)

		new_products = AddProduct()
		for ref, label := range new_products {
			r, err := c.CreateNewProduct(ctx, &pb.NewProduct{Ref: ref, Label: label})
			if err != nil {
				log.Fatalf("could not create product: %v", err)
			}
			log.Printf(`Product Details:
			ID: %d
			REF: %s
			LABEL: %s`, r.GetId(), r.GetRef(), r.GetLabel())
		}

	case 3:
		var id int32
		fmt.Printf("Saisir l'identifiant de produit à afficher : ")
		fmt.Scan(&id)

		r, err := c.GetOneProduct(ctx, &pb.Product{Id: id})
		if err != nil {
			log.Fatalf("could not print product: %v", err)
		}
		fmt.Printf(`
			ID: %d
			REF: %s
			LABEL: %s`, r.GetId(), r.GetRef(), r.GetLabel())
	case 4:
		var id int32
		fmt.Printf("Saisir l'identifiant de produit à supprimer : ")
		fmt.Scan(&id)

		r, err := c.DeleteProduct(ctx, &pb.Product{Id: id})
		if err != nil {
			log.Fatalf("could not remove product: %v", err)
		}
		fmt.Printf(`Le produit avec l'identifiant %d a bien été supprimé`, r.GetId())
	default:
		params := &pb.GetProductsParams{}
		r, err := c.GetProducts(ctx, params)
		if err != nil {
			log.Fatalf("could not create product: %v", err)
		}
		fmt.Printf("PRODUCT LIST: %v\n", r.GetProducts())
	}

}
