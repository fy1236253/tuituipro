package http

import (
	"cfg"
	"component/g"
	"encoding/xml"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// Start 路由相关的启动
func Start() {
	// 静态资源请求
	ConfigWebHTTP()
	ConfigAPIRoutes()
	ConfigWechatRoutes()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(filepath.Join(g.Root, "/public"))).ServeHTTP(w, r)
	})
	// start http server
	addr := cfg.Config().HTTP.Listen

	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}

	log.Println("http.Start ok, listening on", addr)
	log.Fatalln(s.ListenAndServe())
}

//RenderText200 只返回200和描述
func RenderText200(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(200)
	w.Write([]byte(s))
}

//RenderXML 只返回200和描述
func RenderXML(w http.ResponseWriter, v interface{}) {
	bs, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
	w.Write(bs)
}

//RenderText 只返回描述
func RenderText(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.Write([]byte(s))
}
func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Dto{Msg: "success", Ts: time.Now().Format("20060102150405"), Data: data})
}

func RenderMsgJson(w http.ResponseWriter, msg string) {
	RenderJson(w, map[string]string{"msg": msg})
}

func AutoRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		RenderMsgJson(w, err.Error())
		return
	}
	RenderDataJson(w, data)
}

func StdRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		w.WriteHeader(400)
		RenderMsgJson(w, err.Error())
		return
	}
	RenderJson(w, data)
}
