# Inventory Manager CLI (Go)

Simple inventory management system built in Go.

## Features
- ✅ Add products
- ✅ List products
- ✅ Delete products
- ✅ Update product details
- ✅ Search products
- ✅ Low stock alerts
- ✅ Increase/Decrease product quantity
- ✅ Export to CSV (Excel-compatible)
- ✅ Import from CSV (Excel-compatible)

## Commands

### ➕ Add product
```bash
„go run main.go -add -name="Keyboard" -quantity=10 -minstock=5„

###📋 List products
go run main.go -list

❌ Delete product
go run main.go -delete -name="Keyboard"

✏️ Update product
go run main.go -update -name="Keyboard" -quantity=7 -minstock=3

🔍 Search product
go run main.go -search -name="Keyboard"

🔼 Increase product quantity
go run main.go -increase -name="Keyboard" -quantity=3

🔽 Decrease product quantity
go run main.go -decrease -name="Keyboard" -quantity=2

📤 Export to CSV
go run main.go -export

📥 Import from CSV
go run main.go -import

Project Structure
inventory-manager/
├── cmd/
│   ├── add.go
│   ├── delete.go
│   ├── export.go
│   ├── import.go
│   ├── increase.go
│   ├── decrease.go
│   ├── list.go
│   ├── search.go
│   └── update.go
├── data/
│   ├── inventory.json
│   └── inventory.csv
├── models/
│   └── product.go
├── go.mod
└── main.go

Created by [ziko0]
