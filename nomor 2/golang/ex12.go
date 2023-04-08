package main

import (
	"fmt"
	"strconv"
)

func generateData(n int) []string {
	baseNumber := "882"
	var customers []string
	var mobileNumber string
	for i := 0; i < n; i++ {
		mobileNumber = baseNumber
		mobileNumber = fmt.Sprintf("%s%09s", mobileNumber, "0")
		mobileNumber += strconv.Itoa(i)
		customers = append(customers, mobileNumber)
	}
	return customers
}

func sendPromoDiscount(input []string) {
	for i := 0; i < len(input); i++ {
		fmt.Println("Sending Promo to", input[i])
	}
	fmt.Println("Its Finished to send Promo to", len(input), "Customers")
	for i := 0; i < len(input); i++ {
		fmt.Println("Sending Discount to", input[i])
	}
	fmt.Println("Its Finished to send Discount to", len(input), "Customers")
}

func main() {
	data := generateData(1000)
	sendPromoDiscount(data)
}
