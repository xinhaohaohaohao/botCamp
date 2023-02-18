package api

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// GetChatId 获取所在所有群聊的 chatID
func GetChatId(client *lark.Client) []string {
	// 创建 Client
	//client := lark.NewClient("cli_a36708e0b038900e", "Tv3X6wFT0CgyIDX5NgZ8VhUF1VTaKO1n")
	// 创建请求对象
	req := larkim.NewListChatReqBuilder().
		UserIdType("open_id").
		PageToken("").
		PageSize(100).
		Build()

	// 发起请求
	resp, err := client.Im.Chat.List(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println("getChatId err\n", err)
		return nil
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println("getChatId resp err\n", resp.Code, resp.Msg, resp.RequestId())
		return nil
	}

	// 业务处理
	//fmt.Println(larkcore.Prettify(resp))
	res := make([]string, 0)
	for _, v := range resp.Data.Items {
		res = append(res, *v.ChatId)
	}
	return res
	//fmt.Printf("%T", resp.Data.Items[0].ChatId)
}
