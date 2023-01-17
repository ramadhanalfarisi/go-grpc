package services

import (
	"database/sql"
	"time"
	"github.com/google/uuid"
	"github.com/ramadhanalfarisi/go-grpc/helpers"
	"github.com/ramadhanalfarisi/go-grpc/models"
)

type ProductServiceImpl struct {
	DB *sql.DB
}

func NewProductService(db *sql.DB) ProductService {
	return &ProductServiceImpl{DB: db}
}

func (p *ProductServiceImpl) InsertProduct(product *models.Product) error {
	db := p.DB
	stmt, err := db.Prepare("INSERT INTO products( id, name, price, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NULL);")
	if err != nil {
		helpers.Error(err)
		return err
	}
	defer stmt.Close()

	prod_id := uuid.New()
	created_at := time.Now().Format("2006-01-02 15:04:05")
	if _, err := stmt.Exec(prod_id, product.ProductName, product.ProductPrice, product.ProductCategory, created_at); err != nil {
		helpers.Error(err)
		return err
	}
	return nil
}

func (p *ProductServiceImpl) GetAllProduct(param *models.GetallRequest) *models.ProductList {
	db := p.DB
	var products models.ProductList
	rows, err := db.Query("SELECT id, name, price, category, created_at, updated_at FROM products")
	if err != nil {
		helpers.Error(err)
	}

	for rows.Next() {
		var (
			ProductID       string
			ProductName     string
			ProductPrice    float32
			ProductCategory string
			CreatedAt       string
			UpdatedAt       *string
		)
		if err := rows.Scan(&ProductID,
			&ProductName,
			&ProductPrice,
			&ProductCategory,
			&CreatedAt,
			&UpdatedAt); err != nil {
			helpers.Error(err)
		}
		var productObject models.Product
		productObject.ProductID = ProductID
		productObject.ProductName = ProductName
		productObject.ProductPrice = ProductPrice
		productObject.ProductCategory = ProductCategory
		productObject.CreatedAt = CreatedAt
		productObject.UpdatedAt = UpdatedAt
		products.List = append(products.List, &productObject)
	}
	return &products
}
