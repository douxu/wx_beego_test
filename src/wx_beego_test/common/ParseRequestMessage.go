package common

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"wx_beego_test/models"
)

//ParseRequestMessage : parse request message from weixin
func ParseRequestMessage(r *http.Request) (msg *models.MixMessage, err error) {
	var body []byte
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = xml.Unmarshal(body, msg)
	return
}
