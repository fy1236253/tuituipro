package section

import (
	"time"

	"log"

	"encoding/json"

	"github.com/toolkits/net/httplib"
)

// ReturnAuthorizerInfo 返回web的info
func ReturnAuthorizerInfo(info, tuiCode string) {
	addr := "http://127.0.0.1:4200/authorizer_info"
	// addr := "http://localhost:6062/component/returninfo"
	req := httplib.Post(addr).SetTimeout(5*time.Second, 1*time.Minute)
	req.Param("data", info)
	req.Param("tuitui_code", tuiCode)
	resp, err := req.String()
	if err != nil {
		log.Println("return authorizerinfo fail:", err)
	}
	log.Println(resp)

}

// IsBindInfo 判断是否绑定
func IsBindInfo(wxid, openid string) (respson *BindInfoRes) {
	addr := "http://127.0.0.1:4200/binds/user?wxid=" + wxid + "&openid=" + openid
	req := httplib.Get(addr).SetTimeout(5*time.Second, 1*time.Minute)
	resp, err := req.String()
	if err != nil {
		log.Println("[error]:", err)
	}
	json.Unmarshal([]byte(resp), &respson)
	log.Println(string(resp))
	return respson
}

// SubscribeFeedback 关注回调
func SubscribeFeedback(openid, sence string) {
	addr := "http://127.0.0.1:4200/subscribe/user"
	req := httplib.Post(addr).SetTimeout(5*time.Second, 1*time.Minute)
	req.Param("openid", openid)
	req.Param("sence", sence)
	resp, err := req.String()
	if err != nil {
		log.Println("[error]:", err)
	}
	log.Println(string(resp))
}
