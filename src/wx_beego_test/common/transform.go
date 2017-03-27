package common

import "wx_beego_test/models"

//Value2CDATA : transform string to CDATAText
func Value2CDATA(v string) models.CDATAText {
	return models.CDATAText{Text: v}
}
