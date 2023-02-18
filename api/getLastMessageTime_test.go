package api

import (
	"botcamp/botDemo/configs"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"testing"
)

func TestGetLastMessageTime(t *testing.T) {
	var appId, appSecret = configs.AppID, configs.AppSecret
	client := lark.NewClient(appId, appSecret)
	chatId := GetChatId(client)
	for _, v := range chatId {
		lastMessageTime := GetLastMessageTime(client, v)
		fmt.Println(lastMessageTime)
	}
}
