package main

import (
	"dzug/app/gateway/cmd"
	"dzug/app/user/cmd"
)

func main() {

	go userservice.Start()
	userclient.Start()
	//binding:"required,max=32"
}
