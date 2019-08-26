package main

import (
	"fmt"
	"time"
)

func main() {
	local := time.Now().Local()
	fmt.Println(local)
}

func read(c chan bool, i int) {
	fmt.Printf(" go func : %d \n", i)
	<-c
}
