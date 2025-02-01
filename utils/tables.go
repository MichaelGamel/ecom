package utils

type Tables struct {
	Users      string
	Products   string
	Orders     string
	OrderItems string
}

var TablesConfig = Tables{
	Users:      "users",
	Products:   "products",
	Orders:     "orders",
	OrderItems: "order_items",
}
