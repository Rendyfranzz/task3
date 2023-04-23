package main

import (
	"fmt"
	"time"
)

func fill_array(array []string) {
	address := "DKI Jakarta"
	for index := range array {
		array[index] = address
	}
}
func findAddress(addresses []string) {
	tx := time.Now()
	fmt.Println("The default Addres is " + (addresses[0]))
	ty := time.Since(tx)
	fmt.Println(ty)

}
func main() {
	addresses := make([]string, 10)
	fill_array(addresses)
	findAddress(addresses)
}
