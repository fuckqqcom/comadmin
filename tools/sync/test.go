package main

import (
	"fmt"
	"math"
)

func main() {
	var x interface{}
	x = "a"
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是：%T", i)
	case int:
		fmt.Printf("x是 int 类型")
	case float64:
		fmt.Printf("x是 float64 类型")
	case func(int) float64:
		fmt.Printf("x是func(int)类型")
	case bool, string:
		fmt.Printf("x是bool或者string类型")
	default:
		fmt.Printf("未知型")
	}
}

func read(c chan bool, i int) {
	fmt.Printf(" go func : %d \n", i)
	<-c
}
