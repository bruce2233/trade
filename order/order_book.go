package order

import "fmt"

type IOrderBook interface {
	AddSellOrder(order IOrder)
	RemoveSellOrder(order IOrder)
	AddBuyOrder(order IOrder)
	RemoveBuyOrder(order IOrder)
}

type OrderBook struct {
	SellOrders []IOrder
	BuyOrders  []IOrder
}

func match(sellOrder ISellOrder) {
	fmt.Println("match sellOrder")
}