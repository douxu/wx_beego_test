package common

import (
	"encoding/xml"
	"time"
	"wx_beego_test/models"
)

//MakeTextResponseBody : response info into xml file
func MakeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &models.TextResponseBody{}
	textResponseBody.FromUserName = Value2CDATA(fromUserName)
	textResponseBody.ToUserName = Value2CDATA(toUserName)
	MsgType := Value2CDATA("text")
	textResponseBody.MsgType = MsgType.Text
	textResponseBody.Content = Value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}
