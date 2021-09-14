package db

import (
	"database/sql"

	"github.com/javielrezende/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name , price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	// O scan pega os valores e adiciona nos atributos da classe
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
