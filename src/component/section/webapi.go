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
	// addr := "http://localhost:6062/component/returninfo"
	req := httplib.Post(addr).SetTimeout(5*time.Second, 1*time.Minute)
	req.Param("data", info)
	req.Param("tuitui_code", "testcode")
	resp, err := req.String()
	if err != nil {
		log.Println("return authorizerinfo fail:", err)
	}
	log.Println(resp)
}
