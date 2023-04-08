package main

import "fmt"

var results [][]string

func arrange(array []string, memory []string) [][]string {
	var current string
	if memory == nil {
		memory = []string{}
	}
	for i := 0; i < len(array); i++ {
		current, array = array[i], append(array[:i], array[i+1:]...)
		if len(array) == 0 {
			results = append(results, append(memory, current))
		}
		arrange(append([]string(nil), array...), append(memory, current))
		array = append(array[:i], append([]string{current}, array[i:]...)...)
	}
	return results
}

func main() {
	candidates := []string{"Baswedan", "Subianto", "Maharani", "Ganjar"}
	chairs := arrange(candidates, nil)
	fmt.Println(chairs)
}
