package services

import (
	"github.com/ramadhanalfarisi/go-grpc/models"
)

type ProductService interface {
	InsertProduct(*models.Product) error
	GetAllProduct(*models.GetallRequest) *models.ProductList
}

