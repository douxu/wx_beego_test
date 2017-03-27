package controllers

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"wx_beego_test/common"
	"wx_beego_test/models"

	"github.com/astaxie/beego"
)

var (
	errForNotDeftypes = errors.New("Not Default Types")
)

//MainController : the main controller
type MainController struct {
	beego.Controller
}

//Get : Get Info from Server
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//Post : Parse info from Client and replay info
func (c *MainController) Post() {
	if !validateURL(c) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	textRequestBody := common.ParseTextRequestBody(c.Ctx.Request)
	if textRequestBody != nil {
		err := checkmsgtype(textRequestBody)
		if err != nil {
			log.Println("Wechat Service: checkmsgtype error: ", err)
			return
		}
		fmt.Printf("Whchat Service: Receive Msg type : %s from user : %s",
			textRequestBody.MsgType, textRequestBody.FromUserName)
		responsecontent := parsecontentfromrequestandreply(textRequestBody.Content)
		responseTextBody, err := common.MakeTextResponseBody(textRequestBody.ToUserName,
			textRequestBody.FromUserName, responsecontent)
		if err != nil {
			log.Println("Wechat Service: makeTextResponseBody error: ", err)
			return
		}
		ResponseWriter := c.Ctx.ResponseWriter.ResponseWriter
		ResponseWriter.Header().Set("Content-Type", "text/xml")
		fmt.Println(string(responseTextBody))
		fmt.Fprintf(ResponseWriter, string(responseTextBody))
	}
}

//ValidateURL : return boolean for check wechat validation
func validateURL(c *MainController) bool {
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	signatureIn := c.GetString("signature")
	signatureGen := common.MakeSignature(timestamp, nonce)

	if signatureGen != signatureIn {
		return false
	}
	echostr := c.GetString("echostr")
	ResponseWriter := c.Ctx.ResponseWriter.ResponseWriter
	fmt.Fprintf(ResponseWriter, echostr)
	return true
}

func checkmsgtype(TextRequest *models.TextRequestBody) error {
	MsgType := TextRequest.MsgType
	switch MsgType {
	case "text":
		//todo
	case "image":
		//todo
	case "voice":
		//todo
	case "video":
		//todo
	case "shortvideo":
		//todo
	case "location":
		//todo
	case "link":
		//todo
	}
	return errForNotDeftypes
}

func parsecontentfromrequestandreply(content string) (reply string) {
	if content == "你好" {
		reply = "你好"
	} else if strings.ToLower(content) == "hello" {
		reply = "Hello"
	} else {
		reply = "暂时不能识别..."
	}
	return reply
}
