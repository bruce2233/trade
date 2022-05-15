package order

type IOrderBook interface{
	AddSellOrder(order Order)
	RemoveSellOrder(order Order)
	AddBugOrder()
}

type OrderBook struct {
	SellOrders []Order
	BuyOrders  []Order
}
