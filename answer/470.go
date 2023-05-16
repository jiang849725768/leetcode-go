package answer

import (
	"fmt"
	"math/rand"
)

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	x, y := rand7(), rand7()
	if y <= 4 {
		return x
	}
	if x <= 4 {
		return y + 3
	}
	return rand10()
}

func (sol *Solution) Title470() {
	tmap := make(map[int]int, 10)
	for i := 0; i < 1000000; i++ {
		tmap[rand10()]++
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d: %d\n", i, tmap[i])
	}
}
