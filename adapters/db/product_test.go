package db

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"log"
	"testing"

	"github.com/guilhermewoelke/arquitetura-hexagonal/application"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string, 
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("abc", "Product Test", 0, "disabled");`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := NewProductDB(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, "abc", product.GetID())
	require.Equal(t, float64(0), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDB := NewProductDB(Db)

	product := application.NewProduct()
	product.Name = "Product test"
	product.Price = 25

	pr, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, pr.GetName())
	require.Equal(t, product.Price, pr.GetPrice())
	require.Equal(t, product.Status, pr.GetStatus())

	product.Status = "enabled"
	pr, err = productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, pr.GetName())
	require.Equal(t, product.Price, pr.GetPrice())
	require.Equal(t, product.Status, pr.GetStatus())

}
