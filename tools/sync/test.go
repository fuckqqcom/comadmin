package main

import (
	"fmt"
	"math"
)

func main() {

	userCount := math.MaxInt64
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		ch <- true
		go read(ch, i)
	}
}

func read(c chan bool, i int) {
	fmt.Printf(" go func : %d \n", i)
	<-c
}
