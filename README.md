
# Inventory Manager CLI (Go)

Simple inventory management system built in Go.

---

## ✅ Features
- Add products
- List products
- Delete products
- Update product details
- Search products
- Low stock alerts
- Increase/Decrease product quantity
- Export to CSV (Excel-compatible)
- Import from CSV (Excel-compatible)

---

## 🛠️ Usage

### ➕ Add product
Adds a new product to the inventory.

```bash
go run main.go -add -name="Keyboard" -quantity=10 -minstock=5
```

---

### 📋 List products
Displays all products in the inventory.

```bash
go run main.go -list
```

---

### ❌ Delete product
Deletes a product from the inventory by name.

```bash
go run main.go -delete -name="Keyboard"
```

---

### ✏️ Update product
Updates an existing product's quantity and minimum stock.

```bash
go run main.go -update -name="Keyboard" -quantity=7 -minstock=3
```

---

### 🔍 Search product
Searches for a product by name in the inventory.

```bash
go run main.go -search -name="Keyboard"
```

---

### 🔼 Increase product quantity
Increases the quantity of an existing product.

```bash
go run main.go -increase -name="Keyboard" -quantity=3
```

---

### 🔽 Decrease product quantity
Decreases the quantity of an existing product.

```bash
go run main.go -decrease -name="Keyboard" -quantity=2
```

---

### 📤 Export to CSV
Exports the current inventory to a CSV file compatible with Excel.

```bash
go run main.go -export
```

---

### 📥 Import from CSV
Imports products from a CSV file into the inventory.

```bash
go run main.go -import
```

---

## 📂 Project Structure

```
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
```

---

##  Author
Created by **[Zviad Mamardashvili](https://github.com/ziko0)**
