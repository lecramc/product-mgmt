package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "example.com/product-management/product"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

const (
	port       = ":50051"
	address_db = "172.17.0.1"
)

func NewProductManagementServer() *ProductManagementServer {
	return &ProductManagementServer{}
}

type ProductManagementServer struct {
	conn *pgx.Conn
	pb.UnimplementedProductManagementServer
}

func (server *ProductManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func (server *ProductManagementServer) CreateNewProduct(ctx context.Context, in *pb.NewProduct) (*pb.Product, error) {

	createSql := ` CREATE TABLE IF NOT EXISTS Product(
		id SERIAL PRIMARY KEY,
		ref text,
		label text
		);`

	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed %v\n", err)
		os.Exit(1)
	}

	created_product := &pb.Product{Ref: in.GetRef(), Label: in.GetLabel()}
	tx, err := server.conn.Begin(context.Background())

	if err != nil {
		log.Fatalf("conn.Begin failed %v", err)
	}
	_, err = tx.Exec(context.Background(), "INSERT INTO product(ref, label) values ($1, $2)", created_product.Ref, created_product.Label)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return created_product, nil
}

func (server *ProductManagementServer) GetProducts(ctx context.Context, in *pb.GetProductsParams) (*pb.ProductsList, error) {

	var products_list *pb.ProductsList = &pb.ProductsList{}

	rows, err := server.conn.Query(context.Background(), "SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		product := pb.Product{}
		err = rows.Scan(&product.Id, &product.Ref, &product.Label)
		if err != nil {
			return nil, err
		}
		products_list.Products = append(products_list.Products, &product)
	}
	return products_list, nil
}

func (server *ProductManagementServer) GetOneProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	product := &pb.Product{Id: in.GetId(), Ref: in.GetRef(), Label: in.GetLabel()}

	err := server.conn.QueryRow(context.Background(), "SELECT * FROM product WHERE id=$1", product.Id).Scan(&product.Id, &product.Ref, &product.Label)
	log.Printf(product.Label)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (server *ProductManagementServer) DeleteProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {

	product := &pb.Product{Id: in.GetId(), Ref: in.GetRef(), Label: in.GetLabel()}
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed %v", err)
	}
	_, err = tx.Exec(context.Background(), "DELETE FROM product WHERE id=$1", product.Id)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}

	tx.Commit(context.Background())
	return product, nil
}
func main() {
	database_url := "postgres://admin:password@" + address_db + ":5432/database"
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	var product_server *ProductManagementServer = NewProductManagementServer()
	product_server.conn = conn
	if err := product_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
