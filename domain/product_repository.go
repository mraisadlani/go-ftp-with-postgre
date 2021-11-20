package domain

import (
	"context"
	"database/sql"
)

type ProductRepository interface {
	GetProduct() (*[]Product, error)
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func BuildProductRepository(db *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

var _ ProductRepository = &ProductRepositoryImpl{}

func (r *ProductRepositoryImpl) GetProduct() (*[]Product, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query, err := r.db.QueryContext(ctx, `SELECT * FROM products ORDER BY product_code ASC`)
	defer query.Close()

	if err != nil {
		return nil, err
	}

	var products []Product
	for query.Next() {
		var product Product

		err = query.Scan(
			&product.ID,
			&product.ProductCode,
			&product.ProductName,
			&product.ProductSlug,
			&product.ProductDescription,
			&product.Qty,
			&product.MinQty,
			&product.MaxQty,
			&product.Weight,
			&product.Volume,
			&product.CreateAt,
			&product.UpdateAt)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return &products, nil
}

func (r *ProductRepositoryImpl) InsertProduct(productDTO ProductDTO) (bool, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query, err := r.db.PrepareContext(ctx, `
		INSERT INTO products (product_code, product_name, product_slug, product_description, qty, min_qty, max_qty, weight, volume) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	defer query.Close()

	if err != nil {
		return false, err
	}

	_, err = query.ExecContext(ctx,
		productDTO.ProductCode,
		productDTO.ProductName,
		productDTO.ProductSlug,
		productDTO.ProductDescription,
		productDTO.Qty,
		productDTO.MinQty,
		productDTO.MaxQty,
		productDTO.Weight,
		productDTO.Volume,
		)

	if err != nil {
		return false, err
	}

	return true, nil
}