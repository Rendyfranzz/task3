package main

import "fmt"

func arrange(array []string, memory []string) [][]string {
	results := [][]string{}
	var current string
	if memory == nil {
		memory = []string{}
	}
	for i := 0; i < len(array); i++ {
		current, array = array[i], append(array[:i], array[i+1:]...)
		if len(array) == 0 {
			results = append(results, append(memory, current))
		}
		results = append(results, arrange(append([]string{}, array...), append(memory, current))...)
		array = append(array[:i], append([]string{current}, array[i:]...)...)
	}
	return results
}

func main() {
	candidates := []string{"Baswedan", "Subianto", "Maharani"}
	chairs := arrange(candidates, nil)
	for i := 0; i < len(chairs); i++ {
		fmt.Println(chairs[i])
	}
}
