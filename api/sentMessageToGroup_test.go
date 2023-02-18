package api

import (
	"botcamp/botDemo/configs"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"testing"
)

func TestSentMessageToGroup(t *testing.T) {
	//mobile := []string{"15253423554"}
	var appId, appSecret = configs.AppID, configs.AppSecret
	client := lark.NewClient(appId, appSecret)
	//var userId = GetOpenId(client, mobile)
	var text = "你好啊，群主"
	chatId := GetChatId(client)

	for _, v := range chatId {
		err := SentMessageToGroup(client, "chat_id", v, text)
		if err != nil {
			fmt.Println(err)
		}
	}
}
