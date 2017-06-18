package http

import (
	"g"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"util"

	"github.com/toolkits/file"
)

// ConfigWebHTTP 对外http
func ConfigWebHTTP() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/Auth", func(w http.ResponseWriter, r *http.Request) {
		log.Println("<----Start of Authority----->")
		cfg := g.Config().TuiKe
		reURL := "http://91coolshe.com/search/authok"
		preCode := util.GetPreAuthCode()
		addr := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=" + cfg.AppID + "&pre_auth_code=" + preCode.Pre_auth_code + "&redirect_uri=" + url.QueryEscape(reURL)
		// log.Println("http.Redirect", addr)

		http.Redirect(w, r, addr, 302)
		return
	})
	http.HandleFunc("/search/authok", func(w http.ResponseWriter, r *http.Request) {
		// queryValues, err := url.ParseQuery(r.URL.RawQuery)
		// // log.Println("ParseQuery", queryValues)
		// if err != nil {
		// 	log.Println("[ERROR] URL.RawQuery", err)
		// 	w.WriteHeader(400)
		// 	return
		// }
		addr := "http://91coolshe.com"
		http.Redirect(w, r, addr, 302)
		return
		// authCode := queryValues.Get("auth_code")
		// authorInfo := util.GetAuthorizationInfo(authCode)
		// if authorInfo == nil {
		// 	log.Println(authorInfo)
		// 	return
		// }
		// authorInfo.AuthorizationInfo.SetAuthorInfo()
	})
}
