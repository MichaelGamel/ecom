package order

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

func (s *Store) CreateOrder(order types.Order) (int, error) {
	stmt, args, err := s.mysql.Insert(utils.TablesConfig.Orders).Columns("userId", "total", "status", "address").Values(order.UserID, order.Total, order.Status, order.Address).ToSql()
	if err != nil {
		return 0, err
	}

	res, err := s.db.Exec(stmt, args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	stmt, args, err := s.mysql.Insert(utils.TablesConfig.OrderItems).Columns("orderId", "productId", "quantity", "price").Values(orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price).ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(stmt, args...)
	return err
}
