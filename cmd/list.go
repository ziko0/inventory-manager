package cmd

import (
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"

	"github.com/fatih/color"
)

func ListProducts() {
	// Open JSON file
	file, err := os.ReadFile("data/inventory.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Load Products from JSON
	var products []models.Product
	err = json.Unmarshal(file, &products)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Display product list
	fmt.Println("Product list")

	for _, p := range products {
		fmt.Printf("Name: %s | Quantity: %d | Min. Stock: %d\n", p.Name, p.Quantity, p.MinStock)
		if p.Quantity <= p.MinStock {
			// Show low stock warning with color
			yellow := color.New(color.FgYellow).SprintFunc()
			fmt.Printf("%s Low stock alert: %s | Quantity: %d | MinStock: %d\n",
				yellow("⚠️"),
				p.Name, p.Quantity, p.MinStock)
		}
	}
}
