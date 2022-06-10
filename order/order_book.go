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

func match(orderBook OrderBook, order IOrder) {
	
	fmt.Println("match sellOrder")
}

