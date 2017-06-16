package section

import (
	"log"
)

// CheckIsNil 检查是否为空
func CheckIsNil(data interface{}) {
	if data == nil {
		log.Println("[error]:found nil pointer")
		return
	}
}
