package main

import (
	"botcamp/botDemo/api"
	"botcamp/botDemo/configs"
	"botcamp/botDemo/event"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"sync"
)

func init() {
	// 请求地址配置URL
	api.SetURL()
}

func main() {
	var appId, appSecret = configs.AppID, configs.AppSecret
	var timeInterval = configs.TimeInterval
	// 创建 Client
	client := lark.NewClient(appId, appSecret)

	// 获取 机器人所在的 所有 群聊ID
	chatId := api.GetChatId(client)
	var wg sync.WaitGroup
	wg.Add(1)
	// 并发 监听每个群聊 安静时间超过定义时间 发送消息
	for i := 0; i < len(chatId); i++ {
		go event.GroupQuietReminder(client, timeInterval, chatId[i])
	}

	wg.Wait()
}
