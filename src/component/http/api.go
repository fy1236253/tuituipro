package http

import (
	"component/g"
	"component/model"
	"component/section"
	"fmt"
	"io"
	"log"
	"mp/account"
	"net/http"
	"net/url"
	"os"
)

// ConfigAPIRoutes api相关接口
func ConfigAPIRoutes() {

	http.HandleFunc("/api/v1/upload/image", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.Println(r.Method)
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		log.Println("ParseQuery", queryValues)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		r.ParseMultipartForm(32 << 20)
		// form := r.MultipartForm
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//创建文件
		fW, err := os.Create(g.Root + "/public/img/" + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		log.Println("保存成功" + head.Filename)
	})
	http.HandleFunc("/component/api/v1/send/news", func(w http.ResponseWriter, r *http.Request) {
		log.Println("---> /component/api/v1/send/news")
		r.ParseForm()
		if r.Method == "POST" {
			wxid := r.FormValue("wxid")
			openid := r.FormValue("openid")
			title := r.FormValue("title")
			desc := r.FormValue("desc")
			url := r.FormValue("url")
			pic := r.FormValue("pic")
			if wxid == "" || openid == "" {
				log.Println("[error]:图文消息的参数不正确")
				return
			}
			model.SendMessageNews(wxid, openid, title, desc, url, pic)
		}
	})
	http.HandleFunc("/component/api/v1/qrcode", func(w http.ResponseWriter, r *http.Request) {
		log.Println("---> /api/v1/qrcode")
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		log.Println("ParseQuery", queryValues)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		sence := queryValues.Get("sence")
		wxid := queryValues.Get("wxid")
		log.Println(sence)
		ttl := 604800 //默认30天有效期
		data := map[string]string{}
		data["sence"] = queryValues.Get("sence")
		if sence == "" {
			RenderJson(w, "sence为空")
			return
		}
		qr, e := account.CreateTemporaryQRCode(sence, ttl, section.GetAccessTokenFromRedis(wxid))
		if e == nil {
			data["qrurl"] = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(qr.Ticket)
			log.Println(data)
			RenderDataJson(w, data)
			return
		}
		AutoRender(w, nil, e) // 错误信息提示
		return
	})
	http.HandleFunc("/component/api/v1/send/cash", func(w http.ResponseWriter, r *http.Request) {
		log.Println("---> /api/v1/qrcode")
		queryValues, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Println("[ERROR] URL.RawQuery", err)
			w.WriteHeader(400)
			return
		}
		log.Println("ParseQuery", queryValues)
		openid := queryValues.Get("openid")
		uuid := queryValues.Get("uuid")
		val := queryValues.Get("val")
		model.WeixinPay(uuid, openid, val)
	})

}
