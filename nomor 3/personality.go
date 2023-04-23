package main

import (
	"math/rand"
)

func Personality() string {
	var array = []string{"ESTJ", "ENTJ", "ESFJ", "ENFJ", "ISTJ", "ISFJ", "INTJ", "INFJ", "ESTP", "ESFP", "ENTP", "ENFP", "ISTP", "ISFP", "INTP & INFP"}
	input := rand.Intn(len(array))
	return string(array[input])
}
