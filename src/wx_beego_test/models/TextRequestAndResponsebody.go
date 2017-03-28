package models

import (
	"encoding/xml"
	"time"
)

//TextRequestBody : text type of request info
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgID        int
}

//TextResponseBody : text type of response info
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      string
	Content      CDATAText
}
