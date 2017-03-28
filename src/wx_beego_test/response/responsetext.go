package response

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"wx_beego_test/common"
	"wx_beego_test/models"
)

//ResTextInfo : response text from struct TextRequestBody
func ResTextInfo(w http.ResponseWriter, textrequest *models.TextRequestBody) {
	if textrequest != nil {
		responsecontent := parseTextContentFromRequestAndReply(textrequest.Content)
		responseTextBody, err := common.MakeTextResponseBody(textrequest.ToUserName,
			textrequest.FromUserName, responsecontent)
		if err != nil {
			log.Println("Wechat Service: makeTextResponseBody error: ", err)
			return
		}
		w.Header().Set("Content-Type", "text/xml")
		fmt.Println(string(responseTextBody))
		fmt.Fprintf(w, string(responseTextBody))
	}
}

func parseTextContentFromRequestAndReply(content string) (reply string) {
	if content == "你好" {
		reply = "你好"
	} else if strings.ToLower(content) == "hello" {
		reply = "Hello"
	} else {
		reply = "暂时不能识别..."
	}
	return reply
}
