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

func sendPromoDiscount(input1 []string, input2 []string) {
	for i := 0; i < len(input1); i++ {
		fmt.Println("Sending Promo to", input1[i])
	}
	fmt.Println("Its Finished to send Promo to", len(input1), "Customers")
	for i := 0; i < len(input2); i++ {
		fmt.Println("Sending Discount to", input2[i])
	}
	fmt.Println("Its Finished to send Discount to", len(input2), "Customers")
}

func main() {
	dataPromo := generateData(100000000)
	dataDiscount := generateData(1000)
	sendPromoDiscount(dataPromo, dataDiscount)
}
