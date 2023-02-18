package api

import (
	"botCamp/configs"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"testing"
)

func TestGetChatId(t *testing.T) {
	var appId, appSecret = configs.AppID, configs.AppSecret
	client := lark.NewClient(appId, appSecret)
	chatId := GetChatId(client)
	for _, v := range chatId {
		fmt.Println(v)
	}
}
