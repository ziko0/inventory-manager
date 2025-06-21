package cmd

import (
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
)

// UpdateProduct updates a product's quantity and minstock by name
func UpdateProduct(name string, quantity int, minStock int) {
	// Read the JSON file
	file, err := os.ReadFile("data/inventory.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var products []models.Product
	err = json.Unmarshal(file, &products)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Search and update the product
	found := false
	for i, p := range products {
		if p.Name == name {
			products[i].Quantity = quantity
			products[i].MinStock = minStock
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Product not found:", name)
		return
	}

	// Save the update list
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

	fmt.Println("Product update successfully!")
}
