package menu

import (
	"bytes"
	"component/section"
	"encoding/json"
	"log"
	"mp"
	"net/url"
	"regexp"
	"time"

	"github.com/toolkits/net/httplib"
)

//CreateMenu 创建自定义菜单.
func CreateMenu(obj interface{}, accesstoken string) (err error) {

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + url.QueryEscape(accesstoken)

	req := httplib.Post(incompleteURL).SetTimeout(3*time.Second, 1*time.Minute)
	req.Body(obj)
	resp, err := req.String()

	log.Println(resp)

	if err != nil {
		log.Println("[ERROR]", err)
		return err
	}

	var result mp.Error
	err = json.Unmarshal([]byte(resp), &result)
	if result.ErrCode != mp.ErrCodeOK {
		log.Println("[ERROR]", result)
		return
	}
	return
}

// SearchMenu 查询菜单选项
func SearchMenu(wxid string) {
	url := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=" + section.GetAccessTokenFromRedis(wxid)
	r := httplib.Get(url).SetTimeout(3*time.Second, 1*time.Minute)
	resp, _ := r.String()
	log.Println(resp)
	var menuJson SearchMenuJSON
	var SubBt SubButton
	var Bt Button
	err := json.Unmarshal([]byte(resp), &menuJson)
	if err != nil {
		log.Println(err)
		log.Println("create menu fail")
		return
	}
	buttonLength := len(menuJson.SelfMenuInfo.Buttons)
	log.Println(buttonLength)
	if buttonLength > 2 {
		SubBt.Name = "我要传播"
		SubBt.Type = "click"
		SubBt.Key = "sendNews"
		// var bt []SubButton
		// bt = append(bt, SubBt.SubButtons.List)
		// bt = append(bt, menuJson.SelfMenuInfo.Buttons[buttonLength-1].SubButtons.List)
		// if menuJson.SelfMenuInfo.Buttons[2].SubButtons ==  {
		var list []SubButton
		list = append(list, SubBt)
		menuJson.SelfMenuInfo.Buttons[2].SubButtons = list
		// } else {
		oldMenu := menuJson.SelfMenuInfo.Buttons[2].SubButtons
		oldMenu = append(oldMenu, SubBt)
		// }
		// menuJson.Menu.Buttons[buttonLength-1].SubButtons = append(menuJson.Menu.Buttons[buttonLength-1].SubButtons, Bt)
	} else {
		Bt.Name = "我要传播"
		Bt.Type = "click"
		Bt.Key = "sendNews"
		log.Println("create menu faild")
		menuJson.SelfMenuInfo.Buttons = append(menuJson.SelfMenuInfo.Buttons, Bt)
	}
	log.Println(menuJson)
	// bytes, _ := json.Marshal(menuJson.SelfMenuInfo)
	// log.Println(menuJson.SelfMenuInfo.Buttons[0].SubButtons.List[0].NewsInfo.List[0].ContentURL)
	buf := bytes.NewBuffer(make([]byte, 0, 16<<10))
	buf.Reset()
	json.NewEncoder(buf).Encode(menuJson.SelfMenuInfo)
	tmpjson := buf.String()
	log.Println(tmpjson)
	reg := regexp.MustCompile(`\\u0026`)
	// IsTel := reg.MatchString(tmpjson)
	tmpjson1 := reg.ReplaceAllString(tmpjson, "&")
	reg1 := regexp.MustCompile(`/`)
	tmpjson2 := reg1.ReplaceAllString(tmpjson1, "\\/")
	log.Println(tmpjson2)
	CreateMenu(tmpjson1, section.GetAccessTokenFromRedis(wxid))
}

// DeleteMenu 删除菜单
func DeleteMenu(wxid string) {
	accessToken := section.GetAccessTokenFromRedis(wxid)
	log.Println(accessToken)
	url := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" + url.QueryEscape(accessToken)
	r := httplib.Get(url).SetTimeout(3*time.Second, 1*time.Minute)
	resp, _ := r.String()
	log.Println("[删除菜单]" + resp)
}
