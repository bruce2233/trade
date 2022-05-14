package order

type Order struct {
	amount int
	price  float64
}

type BuyOrder struct{
	Order
}

type SellOrder struct{
	Order
}