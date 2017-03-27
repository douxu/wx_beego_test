package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	token = "wechat4go"
)

//MakeSignature : return hashcode
func MakeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	hashcode := sha1.New()
	io.WriteString(hashcode, strings.Join(sl, ""))
	return fmt.Sprintf("%x", hashcode.Sum(nil))
}
