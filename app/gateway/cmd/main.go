package main

import "dzug/app/gateway/routes"

func main() {
	route := routes.NewRouter()
	_ = route.Run(":8001")
}
