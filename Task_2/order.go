package order

import (
	"errors"
	"strings"
)

const (
	StatusNew        = "new"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusCancelled  = "cancelled"
)

var (
	ErrInvalidId = errors.New("invalid id")
	ErrNilOrder  = errors.New("order can't be nill")
	ErrEmtyName  = errors.New("item name can't be empty")
	ErrNegPrice  = errors.New("price must be positive")
)

var StatusDescription = map[string]string{
	StatusNew:        "new",
	StatusProcessing: "processing",
	StatusCompleted:  "completed",
	StatusCancelled:  "cancelled",
}

type Item struct {
	name  string
	price int
}

type Order struct {
	id     int
	status string
	items  []Item
}

func MakeOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, ErrInvalidId
	}
	return Order{id: id, status: StatusNew}, nil
}

func (order *Order) AddItem(item Item) error {
	if order == nil {
		return ErrNilOrder
	}
	if strings.TrimSpace(item.name) == "" {
		return ErrEmtyName
	}
	if item.price < 0 {
		return ErrNegPrice
	}

	order.items = append(order.items, item)
	return nil
}

func (order *Order) CountTotalPrice() (int, error) {
	if order == nil {
		return 0, ErrNilOrder
	}

	var total_price int = 0
	for _, item := range order.items {
		total_price += item.price
	}

	return total_price, nil
}
