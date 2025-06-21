package cmd

import (
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
)

// DeleteProduct removes a product by name and updates the JSON file
func DeleteProduct(name string) {
	// Read JSON file
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

	// Filter out the product to delete
	var updateProducts []models.Product
	found := false
	for _, p := range products {
		if p.Name != name {
			updateProducts = append(updateProducts, p)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Product not found:", name)
		return
	}

	// Save the update list
	updatedData, err := json.MarshalIndent(updateProducts, "", " ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	err = os.WriteFile("data/inventory.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Product deleted successfully")
}
