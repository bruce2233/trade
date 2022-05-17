package order

import (
	"encoding/json"
	"math/rand"
)

type IOrder interface {
	ToJSON() ([]byte, error)
}

type Order struct {
	TokenID uint64
	Amount  uint64
}

func (order Order) ToJSON() ([]byte, error) {
	a, err := json.Marshal(order)
	return a, err
}

type ILimitOrder interface {
}

type LimitOrder struct {
	Order
	price uint64
}

type IBuyOrder struct {
}

type BuyOrder struct {
	Order
}

type ISellOrder interface {
	IOrder
}

type SellOrder struct {
	Order
}

type MarketOrder struct {
}

func GenerateOrder() Order {
	order := Order{
		TokenID: rand.Uint64() % 100,
		Amount:  rand.Uint64() % 100000,
	}
	return order
}
