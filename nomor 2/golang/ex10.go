package main

import (
	"fmt"
	"math/rand"
)

var telco = []string{"Telkom", "Indosat"}

func findCompany(array []string, input int) {
	for i := 0; i < len(array); i++ {
		if array[i] == array[input] {
			fmt.Println("found", array[input])
			break
		}
		fmt.Println("serach")
	}
}
func main() {
	var company int = rand.Intn(len(telco))
	findCompany(telco, company)
}
