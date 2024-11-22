package services

import (
	"context"
	"encoding/json"
	"go-grpc/cmd/config/helpers"
	"go-grpc/pb/pagination" // Import pagination
	"go-grpc/pb/product"
	"log"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	product.UnimplementedProductServiceServer
	FilePath string // Lokasi JSON
}

// GetProducts retrieves all products with pagination.
func (p *ProductService) GetProducts(ctx context.Context, pageParam *product.Page) (*product.Products, error) {
	var data struct {
		Data       []*product.Product  `json:"data"`
		Pagination *pagination.Pagination `json:"pagination"`
	}
	err := helpers.ReadJSONFile(p.FilePath, &data)
	if err != nil {
		log.Printf("Error reading JSON file: %v", err)
		return nil, err
	}

	return &product.Products{
		Pagination: data.Pagination,
		Data:       data.Data,
	}, nil
}

// GetProduct retrieves a single product by its ID.
func (p *ProductService) GetProduct(ctx context.Context, id *product.Id) (*product.Product, error) {
	var data struct {
		Data []*product.Product `json:"data"`
	}
	err := helpers.ReadJSONFile(p.FilePath, &data)
	if err != nil {
		return nil, err
	}

	for _, product := range data.Data {
		if product.Id == id.Id {
			return product, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Product not found")
}

// CreateProduct adds a new product to the JSON data.
func (p *ProductService) CreateProduct(ctx context.Context, newProduct *product.Product) (*product.Id, error) {
	var data struct {
		Data []*product.Product `json:"data"`
	}

	// Read current data
	err := helpers.ReadJSONFile(p.FilePath, &data)
	if err != nil {
		return nil, err
	}

	// Assign a new ID and add the product to the list
	newProduct.Id = uint64(len(data.Data) + 1) // For simplicity, new ID is len+1
	data.Data = append(data.Data, newProduct)

	// Write the updated data back to the file
	err = WriteJSONFile(p.FilePath, data)
	if err != nil {
		return nil, err
	}

	// Return the new product's ID
	return &product.Id{Id: newProduct.Id}, nil
}

// UpdateProduct updates an existing product by its ID.
func (p *ProductService) UpdateProduct(ctx context.Context, updatedProduct *product.Product) (*product.Status, error) {
	var data struct {
		Data []*product.Product `json:"data"`
	}

	// Read current data
	err := helpers.ReadJSONFile(p.FilePath, &data)
	if err != nil {
		return nil, err
	}

	// Find and update the product by ID
	var found bool
	for _, product := range data.Data {
		if product.Id == updatedProduct.Id {
			product.Name = updatedProduct.Name
			product.Price = updatedProduct.Price
			product.Stock = updatedProduct.Stock
			product.Category = updatedProduct.Category
			found = true
			break
		}
	}

	if !found {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	// Write the updated data back to the file
	err = WriteJSONFile(p.FilePath, data)
	if err != nil {
		return nil, err
	}

	// Return success status
	return &product.Status{Status: 1}, nil
}

// DeleteProduct removes a product from the JSON data by its ID.
func (p *ProductService) DeleteProduct(ctx context.Context, id *product.Id) (*product.Status, error) {
	var data struct {
		Data []*product.Product `json:"data"`
	}

	// Read current data
	err := helpers.ReadJSONFile(p.FilePath, &data)
	if err != nil {
		return nil, err
	}

	// Find and remove the product by ID
	var indexToDelete = -1
	for i, product := range data.Data {
		if product.Id == id.Id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	// Remove the product from the slice
	data.Data = append(data.Data[:indexToDelete], data.Data[indexToDelete+1:]...)

	// Write the updated data back to the file
	err = WriteJSONFile(p.FilePath, data)
	if err != nil {
		return nil, err
	}

	// Return success status
	return &product.Status{Status: 1}, nil
}

// Helper function to write data to the JSON file
func WriteJSONFile(filePath string, v interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}
