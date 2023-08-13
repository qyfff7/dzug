package main

import (
	"dzug/app/gateway/cmd"
	"dzug/app/user/cmd"
	"dzug/app/video/cmd"
	"time"
)

func main() {

	go userservice.Start()
	time.Sleep(time.Second * 3)
	go videoservice.Start()
	client.Start()
	//binding:"required,max=32"
}
