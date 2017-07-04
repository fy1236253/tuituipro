// 二维码接口
package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"mp"
	"net/url"
	"time"

	"github.com/toolkits/net/httplib"
	//"strconv"
)

// 创建二维码
func CreateQRCode(action string, SceneId interface{}, ExpireSeconds int, access_token string) (qrcode *TemporaryQRCode, err error) {

	if SceneId == 0 {
		err = errors.New("SceneId should be greater than 0")
		return
	}

	if ExpireSeconds <= 0 {
		err = errors.New("ExpireSeconds should be greater than 0")
		return
	}

	var request struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneId  uint32 `json:"scene_id,omitempty"`
				SceneStr string `json:"scene_str,omitempty"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ExpireSeconds = ExpireSeconds
	request.ActionName = action

	if action == QR_LIMIT_STR_SCENE {
		switch SceneId.(type) {
		case string:
			request.ActionInfo.Scene.SceneStr = SceneId.(string)
		case int:
			request.ActionInfo.Scene.SceneStr = SceneId.(string)
		}
	} else {
		request.ActionInfo.Scene.SceneId = SceneId.(uint32)
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + url.QueryEscape(access_token)

	buf := bytes.NewBuffer(make([]byte, 0, 16<<10))
	buf.Reset()
	json.NewEncoder(buf).Encode(request)
	tmpjson := buf.String()

	req := httplib.Post(incompleteURL).SetTimeout(3*time.Second, 1*time.Minute)
	req.Body(tmpjson)
	resp, err := req.String()

	log.Println(resp)

	if err != nil {
		log.Println("[ERROR]", err)
		return nil, err
	}

	var result struct {
		mp.Error
		TemporaryQRCode
	}
	err = json.Unmarshal([]byte(resp), &result)
	if result.ErrCode != mp.ErrCodeOK {
		log.Println("[ERROR]", result)
		return
	}

	// if action == QR_LIMIT_STR_SCENE {
	// 	result.TemporaryQRCode.SceneString = strconv.FormatInt(SceneId, 10)
	// } else {
	// 	result.TemporaryQRCode.SceneId = uint32(SceneId)
	// }

	qrcode = &result.TemporaryQRCode

	return
}

// 创建临时二维码
//  SceneId:       场景值ID, 为32位非0整型
//  ExpireSeconds: 二维码有效时间, 以秒为单位.  最大不超过 604800.
func CreateTemporaryQRCode(Scene string, ExpireSeconds int, access_token string) (qrcode *TemporaryQRCode, err error) {

	if ExpireSeconds <= 0 {
		err = errors.New("ExpireSeconds should be greater than 0")
		return
	}

	var request struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneStr string `json:"scene_str"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ExpireSeconds = ExpireSeconds
	request.ActionName = "QR_SCENE"
	request.ActionInfo.Scene.SceneStr = Scene

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + url.QueryEscape(access_token)
	log.Println(request)
	buf := bytes.NewBuffer(make([]byte, 0, 16<<10))
	buf.Reset()
	json.NewEncoder(buf).Encode(request)
	tmpjson := buf.String()

	req := httplib.Post(incompleteURL).SetTimeout(3*time.Second, 1*time.Minute)
	req.Body(tmpjson)
	resp, err := req.String()

	log.Println(resp)

	if err != nil {
		log.Println("[ERROR]", err)
		return nil, err
	}

	var result struct {
		mp.Error
		TemporaryQRCode
	}
	err = json.Unmarshal([]byte(resp), &result)
	if result.ErrCode != mp.ErrCodeOK {
		log.Println("[ERROR]", result)
		return
	}
	qrcode = &result.TemporaryQRCode
	return
}
