package db

import (
	"database/sql"

	"github.com/ThailanTec/go-hexagonal/application/core/ports"
	repository "github.com/ThailanTec/go-hexagonal/application/repository/product"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{
		db: db,
	}
}

func (p *ProductDb) Get(id string) (ports.ProductInterface, error) {
	var product repository.DProduct

	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Product.ID, &product.Product.Name, &product.Product.Price, &product.Product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
func (p *ProductDb) create(product ports.ProductInterface) (ports.ProductInterface, error) {

	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product ports.ProductInterface) (ports.ProductInterface, error) {
	_, err := p.db.Exec("update products set name =?, price =?, status =? where id =?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) Save(product ports.ProductInterface) (ports.ProductInterface, error) {

	var rows int

	p.db.QueryRow("Select count(*) from products where id=?", product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}
