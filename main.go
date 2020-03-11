package main

import (
	"darlogin/redis_lib"
	"fmt"
	"github.com/google/uuid"
	"time"
)
const (
	fullControl = iota
	read
	write
	update
)

type Admin struct {
	Permission int
}

type UserInfo struct {
	Username string
	Token string
	Admin Admin
}
func main(){
	redisClient := redis_lib.ConnectRedis()
	token:=uuid.New()
	fmt.Println(token)
	input:=&UserInfo{
		Username:"sadad",
		Token: token.String(),
		Admin: Admin{Permission:fullControl},
	}
	fmt.Println(input)
	err:=redisClient.SetKey("key",input,time.Minute*1)
	if err!=nil{
		fmt.Println(err)
	}
	user:=&UserInfo{}
	err=redisClient.GetKey("key",&user)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(user)

}
