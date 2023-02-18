package event

import (
	"botCamp/api"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"time"
)

var content = "群消息安静提醒"

// GroupQuietReminder 群消息安静提醒
func GroupQuietReminder(client *lark.Client, timeInterval float64, receiveId string) {
	//var receiveId string
	//receiveId = <-receiveIdChan
	//fmt.Println(receiveId)
	for true {
		lastMessageTime := api.GetLastMessageTime(client, receiveId)
		diffTime := getLastToNowDiffTime(lastMessageTime)
		if diffTime > timeInterval {
			err := api.SentMessageToGroup(client, "chat_id", receiveId, content)
			if err != nil {
				panic(err)
			}
		}
	}

}

// 获取当前时间距最后一次消息的时间差  秒钟级别
func getLastToNowDiffTime(lastMessageTimeStr string) float64 {
	messageTime := stringToTime(lastMessageTimeStr)
	t := time.Now().Sub(messageTime)
	return t.Seconds()
}

func stringToTime(timeStr string) time.Time {
	timeTemplate := "2006-01-02 15:04:05"
	timeTime, _ := time.ParseInLocation(timeTemplate, timeStr, time.Local)
	return timeTime
}
