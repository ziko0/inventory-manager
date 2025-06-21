package cmd

import (
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
)

// AddProduct adds a new product to the inventory and displays the update list
func AddProduct(name string, quantity int, minStock int) {
	// Read existing products
	file, err := os.ReadFile("data/inventory.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var products []models.Product
	json.Unmarshal(file, &products)

	// Create new product
	newProduct := models.Product{
		Name:     name,
		Quantity: quantity,
		MinStock: minStock,
	}

	// Add the product to the invenotry
	products = append(products, newProduct)

	// Save update list to file
	updateData, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	err = os.WriteFile("data/inventory.json", updateData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Product successfully added!")

}
