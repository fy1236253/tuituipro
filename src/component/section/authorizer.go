package section

import (
	"cfg"
	"component/g"
	"encoding/json"
	"log"
	"time"

	"github.com/toolkits/net/httplib"
)

// GetPreAuthCode 获取预授权码
func GetPreAuthCode() *PreAuthInfo {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + g.GetTuiKeAccessToken()
	data := "{\"component_appid\":\"" + cfg.Config().TuiKe.AppID + "\" }"
	req := httplib.Post(url).SetTimeout(5*time.Second, 1*time.Minute)
	req.Body(data)
	resp, err := req.String()
	if err != nil {
		log.Println(err)
		return nil
	}
	var preinfo PreAuthInfo
	if err = json.Unmarshal([]byte(resp), &preinfo); err != nil {
		log.Println("[ERROR] json ", err, resp)
		return nil
	}
	return &preinfo
}

// GetAuthorizationInfo 授权成功后换取授权信息
func GetAuthorizationInfo(autorcode string) *AuthorizationInfos {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=" + g.GetTuiKeAccessToken()
	data := "{\"component_appid\":\"" + cfg.Config().TuiKe.AppID + "\",\"authorization_code\": \"" + autorcode + "\"}"
	req := httplib.Post(url).SetTimeout(5*time.Second, 1*time.Minute)
	req.Body(data)
	resp, err := req.String()
	if err != nil {
		log.Println(err)
		return nil
	}
	var authorInfo AuthorizationInfos
	if err = json.Unmarshal([]byte(resp), &authorInfo); err != nil {
		log.Println("[ERROR] json ", err, resp)
		return nil
	}
	return &authorInfo
}

// RefreshAccessToken 刷新令牌
func RefreshAccessToken(appid string) *AuthorizationInfo {
	log.Println("refresh accesstoken")
	url := "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=" + g.GetTuiKeAccessToken()
	refreshToken := GetRefreshToken(appid)
	data := "{\"component_appid\":\"" + cfg.Config().TuiKe.AppID + "\",\"authorizer_appid\": \"" + appid + "\",\"authorizer_refresh_token\": \"" + refreshToken + "\"}"
	req := httplib.Post(url).SetTimeout(5*time.Second, 1*time.Minute)
	req.Body(data)
	resp, err := req.String()
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(resp)
	var authorInfo AuthorizationInfo
	if err = json.Unmarshal([]byte(resp), &authorInfo); err != nil {
		log.Println("[ERROR] json ", err, resp)
		return nil
	}
	authorInfo.AuthorizerAppid = appid
	return &authorInfo
}

// GetAuthorizationBasicInfo 获取授权方基本信息
func GetAuthorizationBasicInfo(appid string) *AuthorizerInfo {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=" + g.GetTuiKeAccessToken()
	data := "{\"component_appid\":\"" + cfg.Config().TuiKe.AppID + "\",\"authorizer_appid\": \"" + appid + "\"}"
	req := httplib.Post(url).SetTimeout(5*time.Second, 1*time.Minute)
	req.Body(data)
	resp, err := req.String()
	if err != nil {
		log.Println(err)
		return nil
	}
	var authorizer AuthorizerInfos
	if err = json.Unmarshal([]byte(resp), &authorizer); err != nil {
		log.Println("[ERROR] json ", err, resp)
		return nil
	}
	return &authorizer.AuthorizerInfo
}

// SetOpenLocation 获取地理位置
func SetOpenLocation(appid string) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_set_authorizer_option?component_access_token=" + g.GetTuiKeAccessToken()
	data := "{\"component_appid\":\"" + cfg.Config().TuiKe.AppID + "\",\"authorizer_appid\":\"" + appid + "\",\"option_name\":\"location_report\",\"option_value\":\"1\"}"
	// data := "{\"component_appid\":\"" + cfg.Config().TuiKe.AppID + "\",\"authorizer_appid\":\"" + appid + "\",\"location_report\":1}"
	log.Println(data)
	req := httplib.Post(url).SetTimeout(5*time.Second, 1*time.Minute)
	req.Body(data)
	resp, err := req.String()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
