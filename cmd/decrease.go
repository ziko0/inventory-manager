package cmd

import (
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
	"strings"
)

// DecreaseProduct decreases product quantity
func DecreaseProduct(name string, qty int) {
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

	found := false
	for i, p := range products {
		if strings.EqualFold(p.Name, name) {
			if products[i].Quantity-qty < 0 {
				fmt.Println("Error: Quantity cannot be negative.")
				return
			}
			products[i].Quantity -= qty
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Product not found:", name)
		return
	}

	updateData, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	err = os.WriteFile("data/invnetory.json", updateData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Product quantity decreased successfully")
}
