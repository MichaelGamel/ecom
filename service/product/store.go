package product

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/MichaelGamel/ecom/types"
	"github.com/MichaelGamel/ecom/utils"
)

type Store struct {
	db    *sql.DB
	mysql sq.StatementBuilderType
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db,
		mysql: sq.StatementBuilder.PlaceholderFormat(sq.Question),
	}
}

func (s *Store) GetProductByID(productID int) (*types.Product, error) {
	q, args, err := s.mysql.Select("*").From(utils.TablesConfig.Products).Where(sq.Eq{"id": productID}).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) GetProductsByID(productIDs []int) ([]types.Product, error) {
	q, args, err := s.mysql.Select("*").From(utils.TablesConfig.Products).Where(sq.Eq{"id": productIDs}).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil

}

func (s *Store) GetProducts() ([]*types.Product, error) {
	q, _, err := s.mysql.Select("*").From(utils.TablesConfig.Products).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) (int64, error) {
	stmt, args, err := s.mysql.Insert(utils.TablesConfig.Products).Columns("name", "price", "image", "description", "quantity").Values(product.Name, product.Price, product.Image, product.Description, product.Quantity).ToSql()
	if err != nil {
		return 0, err
	}

	result, err := s.db.Exec(stmt, args...)
	if err != nil {
		return 0, err
	}

	id, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return id, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	stmt, args, err := s.mysql.Update(utils.TablesConfig.Products).Set("name", product.Name).Set("price", product.Price).Set("image", product.Image).Set("description", product.Description).Set("quantity", product.Quantity).Where(sq.Eq{"id": product.ID}).ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(stmt, args...)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
