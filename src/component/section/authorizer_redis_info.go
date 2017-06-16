package section

import (
	"cfg"
	"component/g"
	"log"
	redispool "redis"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// SetAuthAccessToken 设置授权方的token信息
func (author *AuthorizationInfo) SetAuthAccessToken() {
	log.Println("set accesstoken")
	rc := redispool.ConnPool.Get()
	defer rc.Close()
	originalID, _ := redis.String(rc.Do("HGET", author.AuthorizerAppid, "original_id"))
	rc.Do("HMSET", author.AuthorizerAppid, "refresh_token", author.AuthorizerRefreshToken)
	rc.Do("HMSET", originalID+"_access_token", "authorizer_access_token", author.AuthorizerAccessToken)
	rc.Do("EXPIRE", originalID+"_access_token", 7200)
}

// SetBasicAuthorizerInfo 设置授权方基础信息
func (author *AuthorizationInfo) SetBasicAuthorizerInfo() {
	info := GetAuthorizationBasicInfo(author.AuthorizerAppid)
	CheckIsNil(info)
	rc := redispool.ConnPool.Get()
	defer rc.Close()
	servicetype := strconv.FormatInt(info.ServiceTypeInfo.ID, 10)
	rc.Do("HMSET", author.AuthorizerAppid,
		"original_id", info.UserName,
		"nick_name", info.NickName,
		"head_img", info.HeadImg,
		"service_type", servicetype,
		"qrcode_url", info.QrcodeURL,
		"refresh_token", author.AuthorizerRefreshToken)
	rc.Do("EXPIRE", author.AuthorizerAppid, 315360000) //十年的保存时间
	rc.Do("HMSET", info.UserName+"_access_token", "authorizer_access_token", author.AuthorizerAccessToken)
	rc.Do("EXPIRE", info.UserName+"_access_token", 7200)
	rc.Do("LPUSH", "component_user", author.AuthorizerAppid)
}

// GetRefreshToken 获取刷新token
func GetRefreshToken(appid string) string {
	rc := redispool.ConnPool.Get()
	defer rc.Close()
	token, _ := redis.String(rc.Do("HGET", appid, "refresh_token"))
	return token
}

// GetAccessTokenFromRedis 从redis中获取对应用户的accesstoken
func GetAccessTokenFromRedis(wxid string) string {
	rc := redispool.ConnPool.Get()
	defer rc.Close()
	token, _ := redis.String(rc.Do("HGET", wxid+"_access_token", "authorizer_access_token"))
	return token
}

// DelAuthorInfo 取消授权后删除信息
func DelAuthorInfo(appid string) {
	rc := redispool.ConnPool.Get()
	defer rc.Close()
	wxid, _ := redis.String(rc.Do("HGET", appid, "original_id"))
	accesstokenkey := wxid + "_access_token"
	rc.Do("del", appid)
	rc.Do("del", accesstokenkey)
}

// SetInRedisTicket 将ticket放入redis中
func SetInRedisTicket(ticket string) {
	rc := redispool.ConnPool.Get()
	defer rc.Close()
	key := "wx_acc_tkn_" + cfg.Config().TuiKe.AppID
	rc.Do("HMSET", key, "ticket", ticket)
	g.SetTuiKeTicket(ticket)
}

// GetOutRedisTicket 将ticket放入redis中
func GetOutRedisTicket() string {
	rc := redispool.ConnPool.Get()
	defer rc.Close()
	if g.GetTuiKeTicket() != "" {
		return g.GetTuiKeTicket()
	}
	key := "wx_acc_tkn_" + cfg.Config().TuiKe.AppID
	ticket, _ := redis.String(rc.Do("hget", key, "ticket"))
	g.SetTuiKeTicket(ticket)
	return ticket
}
