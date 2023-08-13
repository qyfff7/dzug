package main

import (
	userclient "dzug/app/gateway/cmd"
	userservice "dzug/app/user/cmd"
)

func main() {

	go userservice.Start()
	userclient.Start()

}
