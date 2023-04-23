package main

import "fmt"

func MatchPerson(s string) string {

	switch s {
	case "ISTP":
		return "ESTP"
	case "ESFJ":
		return "ESTP"
	case "ISFJ":
		return "ESTP"
	case "ESTP":
		return "ISTP"
	default:
		fmt.Println("Invalid day")
	}
	return ""
}
