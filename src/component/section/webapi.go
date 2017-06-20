package section

import (
	"time"

	"log"

	"github.com/toolkits/net/httplib"
)

// ReturnAuthorizerInfo 返回web的info
func ReturnAuthorizerInfo(info string) {
	log.Println(info)
	addr := "http://127.0.0.1:4200/authorizer_info"
	req := httplib.Post(addr).SetTimeout(5*time.Second, 1*time.Minute)
	req.Body(info)
	resp, err := req.String()
	if err != nil {
		log.Println("return authorizerinfo fail:", err)
	}
	log.Println(resp)
}
