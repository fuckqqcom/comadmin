package main

import (
	"comadmin/router"
	"flag"
	"fmt"
)

func main() {

	path := flag.String("-c", "config/config.json", "config.conf")
	r := router.NewRouter(*path)
	fmt.Println(r.Run(":1234"))
}
