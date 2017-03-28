package common

import (
	"fmt"
)

//OutputInfo : output info
func OutputInfo(MsgType string) string {
	return fmt.Sprintf("Whchat Service: Receive Msg type : %s", MsgType)
}
