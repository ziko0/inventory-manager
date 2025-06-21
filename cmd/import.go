package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"inventory-manager/models"
	"os"
	"strconv"
)

// ImportFromCSV imports products from a CSV file
func ImportFromCSV() {
	csvFile, err := os.Open("data/inventory.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	var products []models.Product

	// Skip header(records[0])
	for i, record := range records {
		if i == 0 {
			continue
		}

		quantity, _ := strconv.Atoi(record[1])
		minStock, _ := strconv.Atoi(record[2])

		product := models.Product{
			Name:     record[0],
			Quantity: quantity,
			MinStock: minStock,
		}

		products = append(products, product)
	}

	// Save to JSON
	data, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	err = os.WriteFile("data/inventory.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("Imported successfully from data/inventory.csv")
}
