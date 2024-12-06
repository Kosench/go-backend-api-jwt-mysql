package order

import (
	"database/sql"
	"go-backend-api-jwt-mysql/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO ecom.orders (userID, total, status, address) VALUES (?,?,?,?)",
		order.UserID, order.Total, order.Status, order.Address)
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
	query := "INSERT INTO ecom.order_items (orderId, productId, quantity, price) " +
		"VALUES (?, ?, ?, ?)"
	_, err := s.db.Exec(query, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
