package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
)

// ExportToCSV exports products to a csv file
func ExportToCSV() {
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

	csvFile, err := os.Create("data/inventory.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Quantity", "MinStock"})

	// Write CSV rows
	for _, p := range products {
		record := []string{p.Name, fmt.Sprintf("%d", p.Quantity), fmt.Sprintf("%d", p.MinStock)}
		writer.Write(record)
	}
	fmt.Println("Exported successfully to data/inventory.csv")
}
