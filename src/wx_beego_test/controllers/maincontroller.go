package controllers

import (
	"errors"
	"fmt"
	"log"
	"wx_beego_test/common"
	"wx_beego_test/parse"
	"wx_beego_test/response"

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
	if !validateURL(c) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	log.Println("Wechat Service: validateUrl Ok!")
}

//Post : Parse info from Client and replay info
func (c *MainController) Post() {
	MsgType := c.GetString("MsgType")
	err := checkMsgTypeAndCallFunction(MsgType, c)
	if err != nil {
		log.Println("Wechat Service: checkmsgtype error: ", err)
		return
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

func checkMsgTypeAndCallFunction(MsgType string, c *MainController) error {
	switch MsgType {
	case "text":
		common.OutputInfo(MsgType)
		textRequestBody := parse.ParseTextRequestBody(c.Ctx.Request)
		response.ResTextInfo(c.Ctx.ResponseWriter, textRequestBody)
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
