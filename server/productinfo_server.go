package main

import (
	"context"
	"log"

	"simple-grpc-example/server/ecommerce"

	"github.com/gofrs/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	productMap map[string]*ecommerce.Product
}

func (s *Server) AddProduct(ctx context.Context, product *ecommerce.Product) (*ecommerce.ProductID, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "프로덕트 아이디 생성 중 에러 발생", err)
	}

	product.Id = id.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*ecommerce.Product)
	}

	s.productMap[product.Id] = product
	log.Println(s)

	return &ecommerce.ProductID{Value: product.Id}, status.New(codes.OK, "").Err()
}

func (s *Server) GetProduct(ctx context.Context, id *ecommerce.ProductID) (*ecommerce.Product, error) {
	value, exist := s.productMap[id.Value]
	if exist {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "프로덕트가 없습니다", id.Value)
}
