package models

import (
	"encoding/xml"
	"time"
)

//PictureRequestBody : picture type of request info
type PictureRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	PicURL       string
	MediaID      int
	MsgID        int
}

//PictureResponseBody : picture type of response info
type PictureResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      string
	PicURL       CDATAText
	MediaID      CDATAText
}
