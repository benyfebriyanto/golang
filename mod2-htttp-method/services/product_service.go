package services

import "mod2-http-method/models"

// In-memory database
var products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 15000, Qty: 10, Category: "PC"},
	{ID: 2, Name: "Smartphone", Price: 7000, Qty: 20, Category: "HP"},
	{ID: 3, Name: "Personnal Computer", Price: 1000000, Qty: 5, Category: "PC"},
	{ID: 4, Name: "Tab", Price: 2000, Qty: 20, Category: "HP"},
}

type ProductService struct{}

// GetAllProduct returns all products
func (s *ProductService) GetAllProducts() []models.Product {
	return products
}

// GetProductByID returns an product by its ID
func (s *ProductService) GetProductByID(id int) (models.Product, bool) {
	for _, product := range products {
		if product.ID == id {
			return product, true
		}
	}
	return models.Product{}, false
}

// CreateProduct adds a new product
func (s *ProductService) CreateProduct(newProduct models.Product) models.Product {
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	return newProduct
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(id int, updatedProduct models.Product) (models.Product, bool) {
	for i, product := range products {
		if product.ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct
			return updatedProduct, true
		}
	}
	return models.Product{}, false
}

// DeleteProduct deletes an product by ID
func (s *ProductService) DeleteProduct(id int) bool {
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return true
		}
	}
	return false
}
