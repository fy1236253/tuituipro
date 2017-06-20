package section

import (
	"time"

	"log"

	"github.com/toolkits/net/httplib"
)

// ReturnAuthorizerInfo 返回web的info
func ReturnAuthorizerInfo(info string) {
	log.Println(info)
	addr := "http://localhost:6062/component/returninfo"
	req := httplib.Post(addr).SetTimeout(5*time.Second, 1*time.Minute)
	req.Body("name")
	resp, err := req.String()
	if err != nil {
		log.Println("return authorizerinfo fail:", err)
	}
	log.Println(resp)
}
