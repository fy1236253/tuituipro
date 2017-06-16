package section

import (
	"encoding/json"
	"log"
	"time"

	"github.com/toolkits/net/httplib"
)

// GetTuiKeToken 请求第三方accesstoken
func GetTuiKeToken(appID, appSecret string) *AccessTokenInfo {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	r := httplib.Post(url).SetTimeout(3*time.Second, 1*time.Minute)
	tmpjson := "{ \"component_appid\":\"" + appID + "\", \"component_appsecret\":\"" + appSecret + "\",\"component_verify_ticket\":\"" + GetOutRedisTicket() + "\"}"
	r.Body(tmpjson)
	resp, err := r.String() //  {"errcode":40013,"errmsg":"invalid appid hint: [1HtmMa0495vr19]"}
	// log.Println(resp)
	if err != nil {
		log.Println("[ERROR] refresh token", err)
		return nil
	}

	var token AccessTokenInfo

	if err = json.Unmarshal([]byte(resp), &token); err != nil {
		log.Println("[ERROR] json ", err, resp)
		return nil
	}

	if token.ErrCode != 0 {
		//log.Println("[ERROR]", token.ErrCode, token.ErrMsg, appId)
		return nil
	}

	//log.Println(token)

	return &token
}
