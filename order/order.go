package order

type Order struct {
	tokenID uint64
	amount uint64
}

type LimitOrder struct{
	Order
	price uint64
}
type BuyOrder struct {
	Order
}

type SellOrder struct {
	Order
}
type MarketOrder struct{
	
}
