package cmd

import (
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
	"strings"
)

// IncreaseProduct increases product quantity
func IncreaseProduct(name string, qty int) {
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
			products[i].Quantity += qty
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

	err = os.WriteFile("data/inventory.json", updateData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Product quantity increased successfully!")
}
