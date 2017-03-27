package models

import (
	"encoding/xml"
	"time"
)

//TextRequestBody :type of request info
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgID        int
}
