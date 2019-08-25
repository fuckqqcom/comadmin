package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func read(c chan bool, i int) {
	fmt.Printf(" go func : %d \n", i)
	<-c
}
