package http

import (
	"cfg"
	"component/g"
	"component/section"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/toolkits/file"
)

// ConfigWebHTTP 对外http
func ConfigWebHTTP() {
	http.HandleFunc("/component/search", func(w http.ResponseWriter, r *http.Request) {
		//log.Println(openid)
		var f string // 模板文件路径
		f = filepath.Join(g.Root, "/public", "index.html")
		if !file.IsExist(f) {
			log.Println("not find", f)
			http.NotFound(w, r)
			return
		}

		data := struct {
		}{}

		t, err := template.ParseFiles(f)
		err = t.Execute(w, data)
		if err != nil {
			log.Println(err)
		}

		return
	})
	http.HandleFunc("/component/Auth", func(w http.ResponseWriter, r *http.Request) {
		log.Println("<----Start of Authority----->")
		r.ParseForm()
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		log.Println("ParseQuery", queryValues)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		cfg := cfg.Config().TuiKe
		reURL := "http://www.91coolshe.com/component/auth/callback?tuitui_code=" + queryValues.Get("tuitui_code")
		// reURL := "http://91coolshe.com"
		preCode := section.GetPreAuthCode()
		addr := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=" + cfg.AppID + "&pre_auth_code=" + preCode.Pre_auth_code + "&redirect_uri=" + url.QueryEscape(reURL)
		// log.Println("http.Redirect", addr)
		http.Redirect(w, r, addr, 302)
		return
	})
	http.HandleFunc("/component/returninfo", func(w http.ResponseWriter, r *http.Request) {
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		log.Println("ParseQuery", queryValues)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		r.ParseForm()
		result, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		log.Println(string(result))
	})
	http.HandleFunc("/component/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		log.Println("ParseQuery", queryValues)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		r.ParseForm()
		authCode := queryValues.Get("auth_code")
		tuiCode := queryValues.Get("tuitui_code")
		// log.Println(authCode)
		authorizer := section.GetAuthorizationInfo(authCode) //获取授权信息
		// log.Println(authorizer)
		// 获取授权号的基本信息
		userInfo := section.GetAuthorizationBasicInfo(authorizer.AuthorizationInfo.AuthorizerAppid)
		section.CheckIsNil(authorizer)
		userInfo.Appid = authorizer.AuthorizationInfo.AuthorizerAppid
		// log.Println(userInfo)
		authorizer.AuthorizationInfo.SetBasicAuthorizerInfo()
		var respdata section.RespData
		respdata.AuthorizerInfo = userInfo
		info, _ := json.Marshal(respdata)
		section.ReturnAuthorizerInfo(string(info), tuiCode)
		// log.Println(string(info))
		addr := "http://www.91coolshe.com/merchant/home"
		http.Redirect(w, r, addr, 302)
		return
	})
	http.HandleFunc("/component/test", func(w http.ResponseWriter, r *http.Request) {
		fullurl := "http://" + r.Host + r.RequestURI
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		log.Println("ParseQuery", queryValues)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		code := queryValues.Get("code") //  摇一摇入口 code 有效
		state := queryValues.Get("state")
		if code == "" && state == "" {
			addr := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxdfac68fcc7a48fca" + "&redirect_uri=" + url.QueryEscape(fullurl) + "&response_type=code&scope=snsapi_userinfo&state=1#wechat_redirect"
			log.Println("http.Redirect", addr)
			http.Redirect(w, r, addr, 302)
			return
		}
		log.Println(code, state)
	})
}
