package main

import (
	"flag"
	"fmt"
	"inventory-manager/cmd"
)

func main() {
	// Define the CLI flags
	list := flag.Bool("list", false, "See product list")
	add := flag.Bool("add", false, "Add a new product")
	deleteFlag := flag.Bool("delete", false, "Delete a product")
	update := flag.Bool("update", false, "Update a product")
	search := flag.Bool("search", false, "Search a product")
	export := flag.Bool("export", false, "Export inventory file to CSV file")
	importCSV := flag.Bool("import", false, "Import inventory from CSV file")

	increase := flag.Bool("increase", false, "Increase product quantity")
	decrease := flag.Bool("decrease", false, "Decrease product quantity")

	// Additional flags for adding product
	name := flag.String("name", "", "Product name")
	quantity := flag.Int("quantity", 0, "Product quantity")
	minstock := flag.Int("minstock", 0, "Product minimum stock")

	// parse the CLI flags
	flag.Parse()

	// Check which command to run
	switch {
	case *list:
		cmd.ListProducts()
	case *add:
		if *name == "" || *quantity <= 0 || *minstock < 0 {
			fmt.Println("Pleas provide valid product details: -name=\"Product Name\" -quantity=Number -minstock=Number")
			return
		}
		cmd.AddProduct(*name, *quantity, *minstock)
	case *deleteFlag:
		if *name == "" {
			fmt.Println("Please provide the product name to delete: -name=\"Product Name\"")
			return
		}
		cmd.DeleteProduct(*name)
	case *update:
		if *name == "" || *quantity <= 0 || *minstock < 0 {
			fmt.Println("Please provide valid product details to update: -name=\"Product Name\" -quantity=Number -minstock=Number")
		}
		cmd.UpdateProduct(*name, *quantity, *minstock)
	case *search:
		if *name == "" {
			fmt.Println("Please provide the product name to search: -name=\"Product Name\"")
			return
		}
		cmd.SearchProduct(*name)
	case *increase:
		if *name == "" || *quantity <= 0 {
			fmt.Println("Please provide product name and quantity to increase.")
			return
		}
		cmd.IncreaseProduct(*name, *quantity)
	case *decrease:
		if *name == "" || *quantity <= 0 {
			fmt.Println("Please provide product name and quantity to decrease.")
			return
		}
		cmd.DecreaseProduct(*name, *quantity)
	case *export:
		cmd.ExportToCSV()
	case *importCSV:
		cmd.ImportFromCSV()
	default:
		fmt.Println("Pleas provide a valid flag. Example: -list or -add")
	}
}
