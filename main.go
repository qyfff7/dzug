package main

import (
	"dzug/app/gateway/cmd"
	"dzug/app/user/cmd"
	"dzug/app/video/cmd"
	"time"
)

func main() {

	go userservice.Start()
	time.Sleep(time.Second)
	go videoservice.Start()
	client.Start()
	//binding:"required,max=32"
	//fmt.Println(time.Now().Unix()) //8.14日13.25fen 1691990786
}
