package http

import (
	"component/g"
	"component/model"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
		r.ParseForm()
		if r.Method == "POST" {
			result, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()
			log.Println(string(result))
			wxid := r.FormValue("wxid")
			openid := r.FormValue("openid")
			title := r.FormValue("openid")
			desc := r.FormValue("desc")
			url := r.FormValue("url")
			pic := r.FormValue("pic")
			model.SendMessageNews(wxid, openid, title, desc, url, pic)
		}
	})

}
