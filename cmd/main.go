package main

import (
	"comadmin/router"
	"fmt"
)

func main() {
	r := router.NewRouter()
	fmt.Println(r.Run(":1234"))
}
