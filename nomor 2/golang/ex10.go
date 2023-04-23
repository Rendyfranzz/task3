package main

import (
	"fmt"
	"math/rand"
)

var telco = []string{"Telkom", "Indosat", "XL", "Verizon", "AT&T", "Nippon", "Vodafone", "Orange", "KDDI", "Telefonica", "T-Mobile"}

func findCompany(array []string, input int) {
	for i := 0; i < len(array); i++ {
		if array[i] == array[input] {
			fmt.Printf("Company Found: %s\n", array[input])
			break
		}
		fmt.Printf("Searching Company... %d\n", i+1)
	}
}
func main() {
	var company int = rand.Intn(len(telco))
	findCompany(telco, company)
}
