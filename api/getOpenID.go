package api

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

// GetOpenId 获取 userId 以发送消息
func GetOpenId(client *lark.Client, mobile []string) string {
	//mobile := []string{"15253423554"}
	//var appID, appSecret = "cli_a36708e0b038900e", "Tv3X6wFT0CgyIDX5NgZ8VhUF1VTaKO1n"
	// 创建 Client
	//client := lark.NewClient(appId, appSecret)
	// 创建请求对象
	req := larkcontact.NewBatchGetIdUserReqBuilder().
		UserIdType("open_id").
		Body(larkcontact.NewBatchGetIdUserReqBodyBuilder().
			Emails([]string{}).
			Mobiles(mobile).
			Build()).
		Build()
	// 发起请求
	resp, err := client.Contact.User.BatchGetId(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return ""
	}

	// 业务处理
	//fmt.Println(larkcore.Prettify(resp))
	return *resp.Data.UserList[0].UserId
}
