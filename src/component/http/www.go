package http

import (
	"cfg"
	"component/g"
	"component/section"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"encoding/json"

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
	http.HandleFunc("/component/auth", func(w http.ResponseWriter, r *http.Request) {
		log.Println("<----Start of Authority----->")
		cfg := cfg.Config().TuiKe
		reURL := "http://91coolshe.com/component/authccallback?num=13618075393"
		// reURL := "http://91coolshe.com"
		preCode := section.GetPreAuthCode()
		addr := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=" + cfg.AppID + "&pre_auth_code=" + preCode.Pre_auth_code + "&redirect_uri=" + url.QueryEscape(reURL)
		// log.Println("http.Redirect", addr)
		http.Redirect(w, r, addr, 302)
		return
	})
	http.HandleFunc("/component/authcallback", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		log.Println("ParseQuery", queryValues)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		num := queryValues.Get("num")
		authCode := queryValues.Get("auth_code")
		authorizer := section.GetAuthorizationInfo(authCode)
		userInfo := section.GetAuthorizationBasicInfo(authorizer.AuthorizationInfo.AuthorizerAppid)
		info, _ := json.Marshal(userInfo)
		log.Println(string(info))
		log.Println(num)
		addr := "http://91coolshe.com/main"
		http.Redirect(w, r, addr, 302)
		return
	})
}
