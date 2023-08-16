package service

import (
	"context"
	"dzug/app/user/pkg/snowflake"
	"dzug/app/user/redis"
	"dzug/conf"
	"dzug/logger"
	"dzug/protos/user"
	"dzug/repo"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client user.ServiceClient

func iinit() {
	conf.Init()
	logger.Init()
	repo.Init()
	redis.Init()
	snowflake.Init()
}
func TestRegister(t *testing.T) {
	iinit()
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
