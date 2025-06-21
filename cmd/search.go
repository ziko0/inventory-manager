package cmd

import (
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
	"strings"
)

// SearchProduct finds and displays product by name
func SearchProduct(name string) {
	// read JSON file
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

	// search the product
	found := false
	for _, p := range products {
		if strings.EqualFold(p.Name, name) {
			fmt.Printf("Product found:\nName: %s | Quantity: %d | Min. Stock: %d\n", p.Name, p.Quantity, p.MinStock)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Product not found:", name)
	}
}
