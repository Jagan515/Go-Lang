package main

import "fmt"

// type OrderStatus int

// const (
// 	Recevied OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivery
// )

type OrderStatus string

const (
	Recevied  OrderStatus = "Recevied"
	Confirmed             = "Confirmed"
	Prepared              = "Prepared"
	Delivery              = "Delivery"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("change in order status ", status)
}

func main() {
	changeOrderStatus(Confirmed)
}
