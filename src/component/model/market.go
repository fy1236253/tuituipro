// 营销活动的 逻辑处理

package model

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"sort"
	"strings"
	//"io/ioutil"
)

//SendCashBill 送企业现金红包 接口
func SendCashBill(openid, mobile string, val int) {
	//rand.Seed(time.Now().UnixNano())
	//uid := time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(999))
	//tmpjson := "{\"cmd\":\"hongbao\", \"uuid\":\""+ uid +"\", \"openid\":\"" + openid + "\", \"value\":\"1\", \"type\":\"0\" }"
	//mq.Publish("ka007.exchange", "direct", "rechange.key", tmpjson, true)
	go WeixinPay(mobile, openid, "1")
}

//WeixinRedPack 微信红包相关
type WeixinRedPack struct {
	XMLName     struct{} `xml:"xml"`
	Sign        string   `xml:"sign"`
	MchBillno   string   `xml:"mch_billno"`
	MchId       string   `xml:"mch_id"`
	Wxappid     string   `xml:"wxappid"`
	SendName    string   `xml:"send_name"`
	Openid      string   `xml:"re_openid"`
	TotalAmount string   `xml:"total_amount"`
	TotalNum    string   `xml:"total_num"`
	Wishing     string   `xml:"wishing"`
	ClientIp    string   `xml:"client_ip"`
	ActName     string   `xml:"act_name"`
	Remark      string   `xml:"remark"`
	NonceStr    string   `xml:"nonce_str"`
}

//WeixinPay 红包金额 val 元
func WeixinPay(uuid, openid, val string) {
	log.Println("=========wexinpay")
	// rand.Seed(time.Now().UnixNano())
	// nonce := strconv.Itoa(rand.Intn(999999999))
	// o := &WeixinRedPack{
	// 	Sign:        "",
	// 	MchBillno:   uuid,
	// 	MchId:       "1484374812",
	// 	Wxappid:     "wxb7f7a24ef49a4263",
	// 	SendName:    "推推平台",
	// 	Openid:      openid,
	// 	TotalAmount: val + "00",
	// 	TotalNum:    "1",
	// 	Wishing:     "云喇叭祝大家恭喜发财",
	// 	ClientIp:    cfg.Config().WeiXinPay.Ip, //"121.43.102.49",
	// 	ActName:     "新年红包",
	// 	Remark:      "新年红包",
	// 	NonceStr:    nonce,
	// }
	var o WeixinRedPack
	o.MchBillno = uuid
	// o.MchId = "1484374812"
	// o.Wxappid = "wxb7f7a24ef49a4263"
	// o.SendName = "推推平台"
	// o.Openid = openid
	// o.TotalAmount = val + "00"
	// o.TotalNum = "1"
	// o.Wishing = "感谢支持推推平台"
	// o.ClientIp = cfg.Config().WeiXinPay.Ip
	// o.ActName = "推推积分兑换"
	// o.Remark = "积分兑换"
	// o.NonceStr = nonce
	log.Println(o)
	// o.Sign = sign(o, cfg.Config().WeiXinPay.Key)
	// buf := bytes.NewBuffer(make([]byte, 0, 16<<10))
	// buf.Reset()
	// xml.NewEncoder(buf).Encode(o)
	// log.Println(o)
	// body := buf.String()
	// log.Println(body)

	// cert, err := tls.LoadX509KeyPair("/data/pay.weixin/apiclient_cert.pem", "/data/pay.weixin/apiclient_key.pem")
	// if err != nil {
	// 	log.Fatalf("server: loadkeys: %s", err)
	// }
	// r := httplib.Post("https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack").SetTimeout(3*time.Second, 1*time.Minute)
	// r.SetTLSClientConfig(&tls.Config{Certificates: []tls.Certificate{cert}})
	// r.Header("Content-Type", "application/xml;charset=UTF-8")
	// r.Body(body)

	// resp, err := r.String()
	// if err != nil {
	// 	log.Println("[ERROR] weixinpay", err)
	// 	return
	// }
	// log.Println("weixin pay result", resp, openid)
}

func sign(o *WeixinRedPack, key string) string {
	strs := sort.StringSlice{"mch_billno=" + o.MchBillno,
		"mch_id=" + o.MchId,
		"wxappid=" + o.Wxappid,
		"send_name=" + o.SendName,
		"re_openid=" + o.Openid,
		"total_amount=" + o.TotalAmount,
		"total_num=" + o.TotalNum,
		"wishing=" + o.Wishing,
		"client_ip=" + o.ClientIp,
		"act_name=" + o.ActName,
		"remark=" + o.Remark,
		"nonce_str=" + o.NonceStr}
	strs.Sort()

	strA := strings.Join(strs[:], "&")

	//log.Println(strA)

	strB := strA + "&key=" + key

	//log.Println(strB)

	md5Sum := md5.Sum([]byte(strB))
	sig := strings.ToUpper(hex.EncodeToString(md5Sum[:]))

	//log.Println(sig)

	return sig
}
