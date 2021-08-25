package repo

import (
	"database/sql"

	"github.com/danisbagus/edagang-package/errs"
	"github.com/danisbagus/edagang-package/logger"
	"github.com/danisbagus/edagang-product/internal/core/domain"
	"github.com/danisbagus/edagang-product/internal/core/port"
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

func (r ProductRepo) Create(data *domain.ProductModel) (*domain.ProductModel, *errs.AppError) {
	insertSql := "insert into products (product_name, product_category, quantity) values (?,?,?)"

	result, err := r.db.Exec(insertSql, data.ProductName, data.ProductCategory, data.Quantity)
	if err != nil {
		logger.Error("Error while creating new employee " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while get last insert id for new employee" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.ProductID = id

	return data, nil
}
