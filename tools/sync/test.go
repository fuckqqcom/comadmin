package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "//http://mp.weixin.qq.com/s?__biz=MzU3ODE2NTMxNQ==&mid=2247485961&idx=1"

	compile := regexp.MustCompile(`biz=(\w*).*?mid=(\w*)\w+&idx=(\d+)`)
	ids := compile.FindAllString(a, -1)
	for k, v := range ids {
		fmt.Println(k, v)
	}
}

func read(c chan bool, i int) {
	fmt.Printf(" go func : %d \n", i)
	<-c
}
