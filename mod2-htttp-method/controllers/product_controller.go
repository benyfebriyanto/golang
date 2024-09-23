package controllers

import (
	"encoding/json"
	"mod2-http-method/models"
	"mod2-http-method/services"
	"net/http"
	"strconv"
)

// ItemController handles HTTP requests related to items
type ProductController struct {
	ProductService services.ProductService
}

// GetAllProductsHandler handles the GET request to retrieve all products
func (c *ProductController) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	items := c.ProductService.GetAllProducts()
	sendJSONResponse(w, "200", "Products retrieved successfully", items)
}

// GetProductHandler handles the GET request to retrieve an product by ID
func (c *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	product, found := c.ProductService.GetProductByID(id)
	if !found {
		sendJSONResponse(w, "404", "Product not found", nil)
		return
	}
	sendJSONResponse(w, "200", "Product retrieved successfully", product)
}

// CreateProductHandler handles the POST request to create a new product
func (c *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid request body", nil)
		return
	}
	createdProduct := c.ProductService.CreateProduct(newProduct)
	sendJSONResponse(w, "201", "Product created successfully", createdProduct)
}

// UpdateItemHandler handles the PUT request to update an item by ID
func (c *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid request body", nil)
		return
	}
	product, updated := c.ProductService.UpdateProduct(id, updatedProduct)
	if !updated {
		sendJSONResponse(w, "404", "product not found", nil)
		return
	}
	sendJSONResponse(w, "200", "product updated successfully", product)
}

// DeleteItemHandler handles the DELETE request to delete an item by ID
func (c *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	deleted := c.ProductService.DeleteProduct(id)
	if !deleted {
		sendJSONResponse(w, "404", "product not found", nil)
		return
	}
	sendJSONResponse(w, "204", "product deleted successfully", nil)
}
