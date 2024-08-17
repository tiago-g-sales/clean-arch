package database

import (
	"database/sql"

	"github.com/tiago-g-sales/clean-arch/internal/entity"
)

type OrderRepository struct{
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository{
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error{
	stmt, err := r.Db.Prepare("INSERT INTRO orders (id, price, tax, final_price) values (?,?,?,?)")
	if err != nil{
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil{
		return err
	}
	return nil
}