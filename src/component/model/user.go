package model

import (
	"mp/user"
	redispool "redis"
)

func SaveUser(self *user.UserInfo) {
	rc := redispool.ConnPool.Get()
	defer rc.Close()

	openid := self.OpenId

	rc.Do("HMSET", openid,
		"sub", "1",
		"sex", self.Sex,
		"nickname", self.Nickname,
		"imgurl", self.HeadImageURL,
		"unionid", self.UnionId)
	rc.Do("EXPIRE", openid, 315360000) // 绑定的数据 永久保存
}
