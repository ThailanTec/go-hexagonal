package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ThailanTec/go-hexagonal/adapter/db"
	repository "github.com/ThailanTec/go-hexagonal/application/repository/product"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory")
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

	smts, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}
	smts.Exec()

}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc, "Product",0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()

}

/*
func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	ProductDb := db.NewProductDb(Db)

	product, err := ProductDb.Get("abc")
	require.Nil(t, err)

	require.Equal(t, "Product", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
} */

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product := repository.NewProduct()

	product.Product.Name = "Product"
	product.Product.Price = 25

	productResult, err := productDb.Save(product)

	require.Nil(t, err)

	require.Equal(t, product.Product.Name, productResult.GetName())
	require.Equal(t, product.Product.Price, productResult.GetPrice())

}
