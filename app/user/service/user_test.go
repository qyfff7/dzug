package service

import (
	"context"
	"dzug/protos/user"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client user.ServiceClient

func TestRegister(t *testing.T) {
	should := assert.New(t)
	newuser := new(user.AccountReq)
	newuser.Username = "uuu"
	newuser.Password = "uuu"
	registerResp, err := client.Register(context.Background(), newuser)
	if should.NoError(err) {
		fmt.Println(registerResp)
		fmt.Println(registerResp.UserId)
		fmt.Println(registerResp.Token)
	}
}
