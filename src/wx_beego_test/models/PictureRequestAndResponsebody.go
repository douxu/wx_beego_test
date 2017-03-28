package models

//PictureRequestBody : picture type of request info
type PictureRequestBody struct {
	CommonToken
	PicURL  string
	MediaID string
	MsgID   int
}

//PictureResponseBody : picture type of response info
type PictureResponseBody struct {
	CommonToken
	MediaID CDATAText
}
