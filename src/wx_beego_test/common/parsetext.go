package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"wx_beego_test/models"
)

//ParseTextRequestBody : Parse info from http.Request
func ParseTextRequestBody(r *http.Request) *models.TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &models.TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}
