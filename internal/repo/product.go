package repo

import (
	"database/sql"

	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/internal/core/port"
	"github.com/danisbagus/semimarket-product/pkg/errs"
	"github.com/danisbagus/semimarket-product/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) port.IProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r ProductRepo) FindAll() ([]domain.ProductModel, *errs.AppError) {
	products := make([]domain.ProductModel, 0)

	findAllSql := "select product_id, product_name, product_category, quantity from products"
	err := r.db.Select(&products, findAllSql)

	if err != nil {
		logger.Error("Error while quering find all product " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return products, nil
}

func (r ProductRepo) FindOneByID(productID int64) (*domain.ProductModel, *errs.AppError) {
	var data domain.ProductModel

	findOneByIDSql := "select product_id, product_name, product_category, quantity from products where product_id = ?"

	err := r.db.Get(&data, findOneByIDSql, productID)

	if err != nil {
		logger.Error("Error while get find one by id product " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Product not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &data, nil
}
