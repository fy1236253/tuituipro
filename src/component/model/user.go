package model

import (
	"encoding/json"
	"io"
	"mp/user"
	"os"
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
		"imgurl", self.HeadImageURL)
	rc.Do("EXPIRE", openid, 315360000) // 绑定的数据 永久保存
}

// SaveUserLocal 保存用户信息到本地文本
func SaveUserLocal(wxid string, self *user.UserInfo) {
	f, _ := os.OpenFile("user/"+wxid+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	user, _ := json.Marshal(self)
	io.WriteString(f, string(user))
}
