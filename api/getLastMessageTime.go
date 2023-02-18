package api

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"strconv"
	"time"
)

// GetLastMessageTime 获取群内的最后消息时间
func GetLastMessageTime(client *lark.Client, chatId string) string {
	pageToken := ""
	messageTime := ""
	for true {
		// 创建请求对象
		req := larkim.NewListMessageReqBuilder().
			ContainerIdType("chat").
			ContainerId(chatId).
			StartTime("").
			EndTime("").
			PageSize(20).
			PageToken(pageToken).
			Build()
		// 发起请求
		resp, err := client.Im.Message.List(context.Background(), req)

		// 处理错误
		if err != nil {
			fmt.Println("GetLastMessageTime err \n", err)
			return ""
		}

		// 服务端错误处理
		if !resp.Success() {
			fmt.Println("GetLastMessageTime resp err \n", resp.Code, resp.Msg, resp.RequestId())
			return "0"
		}

		// 业务处理
		items := resp.Data.Items // 相应报文中的items
		if len(items) != 0 {
			unixTime, _ := strconv.Atoi(*items[len(items)-1].UpdateTime)
			messageTime = unixToTime(int64(unixTime / 1000))
			//fmt.Println(unixTime)
			//fmt.Println(messageTime)
			pageToken = *resp.Data.PageToken
			hasMore := resp.Data.HasMore // 判断是否还有下一页
			if *hasMore == false {
				break
			}
		} else {
			fmt.Println("没有消息")
			messageTime = "0"
			break
		}

	}

	return messageTime
}

// 时间戳 转换为 时间
// 单位：秒
func unixToTime(unix int64) string {
	timeTemplate := "2006-01-02 15:04:05"
	t := time.Unix(unix, 0)
	return t.Format(timeTemplate)
}
