package g

import (
	"cfg"
	"log"
	"mp"
	"os"
	"sync"
)

var (
	// TuiKe 第三方平台
	TuiKe       *mp.TuiKeConfig
	wxcfgLock   = new(sync.RWMutex)
	wxTokenLock = new(sync.RWMutex)
	Root        string
)

const (
	// VERSION 版本号
	VERSION = "0.1.0"
)

// InitRootDir 当前路径获取
func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	if err != nil {
		log.Fatalln("getwd fail:", err)
	}
}

//InitWxConfig 初始化WeChat
func InitWxConfig() {
	TuiKe = cfg.Config().TuiKe
	log.Println("g.InitWxConfig ok")
}

// GetTuiKeConfig  通过wxid获取配置信息
func GetTuiKeConfig() *mp.TuiKeConfig {
	wxcfgLock.RLock()
	defer wxcfgLock.RUnlock()
	return TuiKe
}

// SetTuiKeAccessToken 设置accesstoken
func SetTuiKeAccessToken(token string) {
	wxTokenLock.Lock()
	defer wxTokenLock.Unlock()
	c := GetTuiKeConfig()
	c.AccessToken = token
}

// GetTuiKeAccessToken 设置accesstoken
func GetTuiKeAccessToken() string {
	wxTokenLock.Lock()
	defer wxTokenLock.Unlock()
	return GetTuiKeConfig().AccessToken
}

// SetTuiKeTicket 设置accesstoken
func SetTuiKeTicket(ticket string) {
	wxTokenLock.Lock()
	defer wxTokenLock.Unlock()
	c := GetTuiKeConfig()
	c.ComponentVerifyTicket = ticket
}

// GetTuiKeTicket 设置accesstoken
func GetTuiKeTicket() string {
	wxTokenLock.Lock()
	defer wxTokenLock.Unlock()
	return GetTuiKeConfig().ComponentVerifyTicket
}
