package api

import (
	"context"
	"encoding/json"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type content struct {
	Text string `json:"text"`
}

// SentMessageToGroup 向群发送消息
func SentMessageToGroup(client *lark.Client, receiveIdType string, receiveId string, sendText string) error {
	// content 转换为json 再转换为 string
	text, _ := json.Marshal(&content{sendText})

	// 创建请求对象
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(receiveIdType).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveId).
			MsgType("text").
			Content(string(text)).
			Uuid("").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Im.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return err
	}

	// 业务处理
	//fmt.Println(larkcore.Prettify(resp))
	return nil
}
