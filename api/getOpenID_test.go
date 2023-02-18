package api

import (
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"testing"
)

func TestGetOpenId(t *testing.T) {
	mobile := []string{"15253423554"}
	var appId, appSecret = "cli_a36708e0b038900e", "Tv3X6wFT0CgyIDX5NgZ8VhUF1VTaKO1n"
	client := lark.NewClient(appId, appSecret)
	fmt.Println(GetOpenId(client, mobile))
}
