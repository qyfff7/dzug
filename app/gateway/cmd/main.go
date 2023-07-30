package main

import (
	"dzug/app/gateway/routes"
	"dzug/app/gateway/rpc"
	"fmt"
)

func main() {
	route := routes.NewRouter()
	rpc.Init()
	defer rpc.Ser.Close()
	defer fmt.Println("我在 Ser.Close 之后运行")
	_ = route.Run(":8001")
}
