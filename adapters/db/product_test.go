package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/javielrezende/go-hexagonal/adapters/db"
	"github.com/javielrezende/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")

	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
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
	insert := `insert into products values("1", "Product 1", 10, "disabled");`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()

	defer Db.Close()

	ProductDb := db.NewProductDb(Db)

	product, err := ProductDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()

	defer Db.Close()

	ProductDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 25

	productResult, err := ProductDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = ProductDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Status, productResult.GetStatus())
}
