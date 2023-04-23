package main

import (
	"fmt"
	"time"
)

const company = "telkom"

var largeCompanyName = make([]string, 100000)

func fill_array(array []string) {
	for index := range array {
		array[index] = company
	}
}

func findCompany(array []string) {
	var start = time.Now()
	for i := 0; i < len(array); i++ {
		if array[i] == "telkom" {
			fmt.Println("Company Found!")
		}
	}
	var ty = time.Since(start)

	fmt.Println(ty)
}
func main() {
	fill_array(largeCompanyName)
	// fmt.Println(largeCompanyName)
	findCompany(largeCompanyName)
}
