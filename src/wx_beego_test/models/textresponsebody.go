package models

import (
	"encoding/xml"
	"time"
)

//TextResponseBody : type of response info
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      string
	Content      CDATAText
}
