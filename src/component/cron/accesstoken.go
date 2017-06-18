package cron

import (
	"cfg"
	"component/g"
	"component/section"
	"log"
	redispool "redis"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 定时检查 token ， 如果发现要过期了，重现请求一个（平台自己的token）
func monitorToken() {
	for {
		c := g.GetTuiKeConfig()
		rc := redispool.ConnPool.Get()
		key := "wx_acc_tkn_" + c.AppID // 判断 redis中这个token 是否存
		t, _ := redis.Int64(rc.Do("TTL", key))
		if t < 600 { //  即将过期
			log.Println("will refresh component accesstoke")
			token := section.GetTuiKeToken(c.AppID, c.AppSecret)
			log.Println(token)
			if token == nil {
				time.Sleep(30 * time.Second)
				continue
			}
			log.Println("wx access token refresh", "***"+token.Component_access_token[:10]+"***", token.ExpiresIn)
			rc.Do("HMSET", key, "token", token.Component_access_token)
			rc.Do("EXPIRE", key, token.ExpiresIn-100)           // 留一个保护间隔
			g.SetTuiKeAccessToken(token.Component_access_token) // 同时保存到 进程内部 提高访问速度
		} else {
			//  每次 同步写入 进程内存中，  这样 多节点，任何一个节点更新后， 都可以实现同步
			token, _ := redis.String(rc.Do("hget", key, "token"))
			g.SetTuiKeAccessToken(token)
		}
		rc.Close()
		time.Sleep(3 * time.Second)
	}
}

// StartToken 开启自动检测token
func StartToken() {
	go monitorToken()
	// time.Sleep(time.Minute * 20)第一次启动时候由于没有接收到微信的推送需要等待至少10分钟
	go CheckAuthorizerAccessToken()

}

//CheckToken 确保所有  access token 都有效
func CheckToken() {
	for {
		if g.GetTuiKeAccessToken() == "" {
			log.Println("[warn] access token not ready, wait", cfg.Config().TuiKe.AppID)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

}

// CheckAuthorizerAccessToken 定时刷新授权方accesstoken（商户的token）
func CheckAuthorizerAccessToken() {
	for {
		rc := redispool.ConnPool.Get()
		user, _ := redis.String(rc.Do("RPOP", "component_user"))
		exist, _ := redis.Bool(rc.Do("EXISTS", user))
		originalID, _ := redis.String(rc.Do("HGET", user, "original_id"))
		if user == "" {
			log.Println("no authorizer")
			break
		}
		t, _ := redis.Int64(rc.Do("TTL", originalID+"_access_token"))
		if t < 600 && exist == true {
			userInfo := section.RefreshAccessToken(user)
			section.CheckIsNil(userInfo)
			userInfo.SetAuthAccessToken()
			log.Println("refresh accesstoke:" + user)
		}
		if exist == true {
			rc.Do("LPUSH", "component_user", user)
		}
		rc.Close()
		time.Sleep(3 * time.Second)
	}
}
