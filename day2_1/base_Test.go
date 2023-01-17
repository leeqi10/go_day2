package day2_1

import (
	"math/rand"
	_ "math/rand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}
func Select() int {
	return ServerIndex[rand.Intn(10)]
}
