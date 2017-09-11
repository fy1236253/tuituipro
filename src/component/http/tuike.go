package http

import (
	"component/g"
	"component/model"
	"component/section"
	"encoding/base64"
	"encoding/xml"
	"log"
	"mp"
	"mp/message"
	"mp/message/request"
	"net/http"
	"net/url"
	"util"
)

// ConfigWechatRoutes 微信页面路由
func ConfigWechatRoutes() {
	http.HandleFunc("/tuike", func(w http.ResponseWriter, req *http.Request) {

		// 捕获异常
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Runtime error caught: %v", r)
				w.WriteHeader(400)
				w.Write([]byte(""))
				return
			}
		}()
		var wxcfg *mp.TuiKeConfig
		var queryValues url.Values
		queryValues, _ = url.ParseQuery(req.URL.RawQuery)
		model.WechatQueryParamsValid(queryValues)
		wxcfg = g.GetTuiKeConfig()
		switch req.Method {
		case "GET":
			{
			}

		case "POST":
			{

				if queryValues.Get("encrypt_type") == "aes" {
					var aesBody message.AesRequestBody
					var aeskey [32]byte // 秘钥
					var mixedMsg message.MixedMessage
					// 非加密码模式 不接入
					model.WechatStrValid(queryValues.Get("encrypt_type"), "aes", "[ERROR] encryptType not support") // xml 解析验证
					model.WechatMessageXMLValid(req, &aesBody)
					model.WechatSignEncryptValid(wxcfg, queryValues, aesBody.EncryptedMsg)
					k, _ := util.AESKeyDecode(wxcfg.Aeskey)
					copy(aeskey[:], k)
					// 解密
					encryptedMsgBytes, _ := base64.StdEncoding.DecodeString(aesBody.EncryptedMsg)
					_, rawMsgXML, appid, _ := util.AESDecryptMsg(encryptedMsgBytes, aeskey)
					model.WechatStrValid(string(appid), wxcfg.AppID, "[Warn] AppId mismatch")
					// 解密ok
					// log.Println(string(rawMsgXML))
					if err := xml.Unmarshal(rawMsgXML, &mixedMsg); err != nil {
						log.Println("[Warn] rawMsgXML Unmarshal", err)
						w.WriteHeader(400)
						return
					}
					// log.Println(mixedMsg.InfoType)
					switch mixedMsg.InfoType {

					// 十分钟的ticket推送
					case "component_verify_ticket":
						{
							log.Println("[refresh ComponentVerifyTicket]:ok")
							ComponentVerifyTicket := mixedMsg.ComponentVerifyTicket
							g.SetTuiKeTicket(ComponentVerifyTicket)
							section.SetInRedisTicket(ComponentVerifyTicket)
						}
					//授权
					case "authorized":
						{
							authorInfo := section.GetAuthorizationInfo(mixedMsg.AuthorizationCode)
							section.CheckIsNil(authorInfo)
							authorInfo.AuthorizationInfo.SetBasicAuthorizerInfo()
							section.SetOpenLocation(mixedMsg.AuthorizerAppid)
							// menu.SearchMenu(mixedMsg.ToUserName)
							log.Println("<-------Authority finished------->")
						}
					// 取消授权
					case "unauthorized":
						{
							log.Println("[cancel Authority]:" + mixedMsg.AuthorizerAppid)
							// menu.DeleteMenu(mixedMsg.AuthorizerAppid)
							section.DelAuthorInfo(mixedMsg.AuthorizerAppid)

						}
					case "updateauthorized":
						{
							log.Println("updateauthorized")
						}

					}
					RenderText(w, "success")
					return
				}
			}
		}

	})
	http.HandleFunc("/wxpush/", func(w http.ResponseWriter, req *http.Request) {
		// 捕获异常
		// defer func() {
		// 	if r := recover(); r != nil {
		// 		log.Printf("Runtime error caught: %v", r)
		// 		w.WriteHeader(400)
		// 		w.Write([]byte(""))
		// 		return
		// 	}
		// }()
		var wxcfg *mp.TuiKeConfig
		var queryValues url.Values
		queryValues, _ = url.ParseQuery(req.URL.RawQuery)
		model.WechatQueryParamsValid(queryValues)
		wxcfg = g.GetTuiKeConfig()
		switch req.Method {
		case "GET":
			{
			}
		case "POST":
			{
				if queryValues.Get("encrypt_type") == "aes" {
					var aesBody message.AesRequestBody
					var aeskey [32]byte // 秘钥
					var mixedMsg message.MixedMessage
					// 非加密码模式 不接入
					model.WechatStrValid(queryValues.Get("encrypt_type"), "aes", "[ERROR] encryptType not support") // xml 解析验证
					model.WechatMessageXMLValid(req, &aesBody)
					model.WechatSignEncryptValid(wxcfg, queryValues, aesBody.EncryptedMsg)
					k, _ := util.AESKeyDecode(wxcfg.Aeskey)
					copy(aeskey[:], k)
					// 解密
					encryptedMsgBytes, _ := base64.StdEncoding.DecodeString(aesBody.EncryptedMsg)
					_, rawMsgXML, appid, _ := util.AESDecryptMsg(encryptedMsgBytes, aeskey)
					model.WechatStrValid(string(appid), wxcfg.AppID, "[Warn] AppId mismatch")
					// 解密ok
					if err := xml.Unmarshal(rawMsgXML, &mixedMsg); err != nil {
						log.Println("[Warn] rawMsgXML Unmarshal", err)
						w.WriteHeader(400)
						return
					}
					log.Println(string(rawMsgXML))
					switch mixedMsg.MsgType {

					// text
					case request.MsgTypeText:
						{
							model.ProcessWechatText(&mixedMsg) // 文本消息的处理逻辑
						}
					// event
					case request.MsgTypeEvent:
						{
							model.ProcessWechatEvent(&mixedMsg)
						}

					}
					RenderText(w, "success")
					return
				}
			}
		}

	})
}
