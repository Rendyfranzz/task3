package main

import "fmt"

func generateData() []string {
	const number = "0821234567"
	var customers []string
	var mobileNumber string
	for i := 0; i < 100; i++ {
		if i < 10 {
			mobileNumber = number
		} else {
			mobileNumber = number + fmt.Sprintf("%d", i)
		}
		customers = append(customers, mobileNumber)
	}
	return customers
}

func sendPromoDiscount(array []string) {
	for i := 0; i < len(array); i++ {
		fmt.Println("Sending Promo to", array[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Sending Discount to Chosen Customer", i+1)
	}
}

func main() {
	customers := generateData()
	sendPromoDiscount(customers)
}
