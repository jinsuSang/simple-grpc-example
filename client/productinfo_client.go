package main

import (
	"context"
	"log"
	"simple-grpc-example/client/ecommerce"

	"time"

	"google.golang.org/grpc"
)

const address = "localhost:3000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("연결 안됨: %v", err)
	}
	defer conn.Close()

	c := ecommerce.NewProductInfoClient(conn)
	name := "Apple Macbook M1"

	description := "I want buy M1X"
	price := float32(2500000.0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	productId, err := c.AddProduct(ctx, &ecommerce.Product{Name: name, Description: description, Price: price})

	if err != nil {
		log.Fatalf("프로덕트 추가 에러: %v", err)
	}
	log.Printf("프로덕트 아이디 %s 추가 성공", productId.Value)

	product, err := c.GetProduct(ctx, &ecommerce.ProductID{Value: productId.Value})

	if err != nil {
		log.Fatalf("프로덕트 없음: %v", err)
	}
	log.Printf("프로덕트: %s", product.String())
}
