package handlers

import (
	"context"
	"database/sql"

	"github.com/ramadhanalfarisi/go-grpc/models"
	"github.com/ramadhanalfarisi/go-grpc/services/product"
)

type ProductsServer struct {
	models.UnimplementedProductsServer
	DB             *sql.DB
	productService services.ProductService
}

func RegisterProductHandler(DB *sql.DB) ProductsServer {
	prodServ := services.NewProductService(DB)
	return ProductsServer{DB: DB, productService: prodServ}
}

func (p ProductsServer) Insert(ctx context.Context, product *models.Product) (*models.Response, error) {
	var response *models.Response
	err := p.productService.InsertProduct(product)
	if err != nil {
		response = &models.Response{Code: 500, Status: "failed", Message: []string{"Insert product failed"}}
	} else {
		response = &models.Response{Code: 200, Status: "success", Message: []string{"Insert product successfully"}}
	}
	return response, nil
}

func (p ProductsServer) GetAll(ctx context.Context, param *models.GetallRequest) (*models.ProductList, error) {
	var productList *models.ProductList
	productList = p.productService.GetAllProduct(param)
	return productList, nil
}
