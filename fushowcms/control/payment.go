package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	mod "fushowcms/models"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Wechat_AppId         = "wx1f9f4614a0377dab"               //微信公众平台AppId
	Wechat_MchId         = "1388166902"                       //微信公众平台商户Id
	Weixin_phoneSign_Key = "shiliulianmengshoujiduan12345678" //微信手机端signKey
	Weixin_AppId         = "wxd4b02d20bb8ac3c4"
	Weixin_MchId         = "1402462202"
)

func wxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if key != "" {
		signStrings = signStrings + "key=" + key
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

type UnifyOrderReq struct {
	Appid            string `xml:"appid"`
	Attach           string `xml:"attach"`
	Body             string `xml:"body"`
	Device_info      string `xml:"device_info"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Notify_url       string `xml:"notify_url"`
	Out_trade_no     string `xml:"out_trade_no"`
	Product_id       string `xml:"product_id"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Total_fee        int    `xml:"total_fee"`
	Trade_type       string `xml:"trade_type"`
	Sign             string `xml:"sign"`
	Openid           string `xml:"openid"`
}

//h5微信下单
func WeiXinUnifiedorder(money int, ip string, uid string) (mod.Unifiedorder, string) {
	var (
		unifiedorder mod.Unifiedorder
		randNum      bytes.Buffer
		yourReq      UnifyOrderReq
	)
	fmt.Println("time", strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Println("now", strconv.Itoa(rand.Int()))
	yourReq.Appid = Weixin_AppId
	yourReq.Mch_id = Weixin_MchId
	yourReq.Nonce_str = strconv.Itoa(rand.Int())
	yourReq.Body = "石榴联盟直播-石榴籽充值"
	randNum.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	randNum.WriteString(strconv.Itoa(rand.Int()))
	yourReq.Out_trade_no = randNum.String() //订单号
	yourReq.Total_fee = money               //单位是分，这里是1毛钱
	yourReq.Spbill_create_ip = ip
	yourReq.Notify_url = "http://www.shiliulianmeng.com/page/wxpaynotice"
	yourReq.Trade_type = "APP"
	yourReq.Attach = uid //用户id
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = yourReq.Appid
	m["mch_id"] = yourReq.Mch_id
	m["nonce_str"] = yourReq.Nonce_str
	m["body"] = yourReq.Body
	m["out_trade_no"] = yourReq.Out_trade_no
	m["total_fee"] = yourReq.Total_fee
	m["spbill_create_ip"] = yourReq.Spbill_create_ip
	m["notify_url"] = yourReq.Notify_url
	m["trade_type"] = yourReq.Trade_type
	m["attach"] = yourReq.Attach
	yourReq.Sign = wxpayCalcSign(m, Weixin_phoneSign_Key)
	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return unifiedorder, ""
	}

	str_req := string(bytes_req)
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
	bytes_req = []byte(str_req)

	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return unifiedorder, ""
	}

	req.Header.Set("Accept", "application/xml")

	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

	c := http.Client{}
	resp, ss := c.Do(req)
	fmt.Println("reso", resp)
	if ss != nil {
		fmt.Println("请求微信支付统一下单接口发送错误, 原因:", ss)
		return unifiedorder, ""
	}

	testReturn := mod.Unifiedorder{}

	bodys, _ := ioutil.ReadAll(resp.Body)

	_ = xml.Unmarshal(bodys, &testReturn)

	return testReturn, yourReq.Out_trade_no
}

func WechatUnifiedorder(money int, ip string, uid string, openid string) (mod.Unifiedorder, string) {
	var (
		unifiedorder mod.Unifiedorder
		randNum      bytes.Buffer
		yourReq      UnifyOrderReq
	)
	fmt.Println("time", strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Println("now", strconv.Itoa(rand.Int()))
	yourReq.Appid = Wechat_AppId
	yourReq.Mch_id = Wechat_MchId
	yourReq.Nonce_str = strconv.Itoa(rand.Int())
	yourReq.Body = "石榴联盟直播-石榴籽充值"
	randNum.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	randNum.WriteString(strconv.Itoa(rand.Int()))
	yourReq.Out_trade_no = randNum.String() //订单号
	yourReq.Total_fee = money               //单位是分，这里是1毛钱
	yourReq.Spbill_create_ip = ip
	yourReq.Notify_url = "http://www.shiliulianmeng.com/page/wxpaynotice"
	yourReq.Trade_type = "JSAPI"
	yourReq.Openid = openid
	yourReq.Attach = uid //用户id
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = yourReq.Appid
	m["mch_id"] = yourReq.Mch_id
	m["nonce_str"] = yourReq.Nonce_str
	m["body"] = yourReq.Body
	m["out_trade_no"] = yourReq.Out_trade_no
	m["total_fee"] = yourReq.Total_fee
	m["spbill_create_ip"] = yourReq.Spbill_create_ip
	m["notify_url"] = yourReq.Notify_url
	m["trade_type"] = yourReq.Trade_type
	m["attach"] = yourReq.Attach
	m["openid"] = yourReq.Openid
	yourReq.Sign = wxpayCalcSign(m, "shiliutvweixinzhifu1234567890123")
	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return unifiedorder, ""
	}
	str_req := string(bytes_req)
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
	bytes_req = []byte(str_req)
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return unifiedorder, ""
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, ss := c.Do(req)
	fmt.Println("reso", resp)
	if ss != nil {
		fmt.Println("请求微信支付统一下单接口发送错误, 原因:", ss)
		return unifiedorder, ""
	}
	testReturn := mod.Unifiedorder{}
	bodys, _ := ioutil.ReadAll(resp.Body)
	_ = xml.Unmarshal(bodys, &testReturn)
	return testReturn, yourReq.Out_trade_no
}

type SelectOrderQueryM struct {
	Appid        string `xml:"appid"`        //公众账号ID
	Mch_id       string `xml:"mch_id"`       //商户号
	Out_trade_no string `xml:"out_trade_no"` //商户订单号
	Nonce_str    string `xml:"nonce_str"`    //随机字符串
	Sign         string `xml:"sign"`         //签名
}

type Refund struct {
	Appid         string `xml:"appid"`         //公众账号ID
	Mch_id        string `xml:"mch_id"`        //商户号
	Out_refund_no string `xml:"out_refund_no"` //商户退款单号
	Out_trade_no  string `xml:"out_trade_no"`  //商户订单号
	Op_user_id    string `xml:"op_user_id"`    //操作员
	Refund_fee    int    `xml:"refund_fee"`    //退款金额
	Total_fee     int    `xml:"total_fee"`     //订单金额
	Nonce_str     string `xml:"nonce_str"`     //随机字符串
	Sign          string `xml:"sign"`          //签名
}

//申请退款  证书无法读取
//out_trade_no 商户订单号 total_fee 订单金额
func QuitMoney(out_trade_no string, total_fee int) (bool, string) {
	var _tlsConfig *tls.Config
	var randNum bytes.Buffer
	var yourReq Refund
	yourReq.Appid = Wechat_AppId
	yourReq.Mch_id = Wechat_MchId
	yourReq.Out_trade_no = out_trade_no
	randNum.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	randNum.WriteString(strconv.Itoa(rand.Int()))
	yourReq.Out_refund_no = randNum.String()
	yourReq.Total_fee = total_fee
	yourReq.Refund_fee = total_fee
	yourReq.Op_user_id = Wechat_MchId
	yourReq.Nonce_str = strconv.Itoa(rand.Int())
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = yourReq.Appid
	m["mch_id"] = yourReq.Mch_id
	m["out_refund_no"] = yourReq.Out_refund_no
	m["out_trade_no"] = yourReq.Out_trade_no
	m["total_fee"] = yourReq.Total_fee
	m["refund_fee"] = yourReq.Refund_fee
	m["op_user_id"] = yourReq.Op_user_id
	m["nonce_str"] = yourReq.Nonce_str
	yourReq.Sign = wxpayCalcSign(m, "shiliutvweixinzhifu1234567890123")
	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return false, "以xml形式编码发送错误"
	}
	str_req := string(bytes_req)
	str_req = strings.Replace(str_req, "Refund", "xml", -1)
	bytes_req = []byte(str_req)
	wechatCertPath := "./static/apiclient_cert.pem"
	wechatKeyPath := "./static/apiclient_key.pem"
	wechatCAPath := "./static/rootca.pem"
	if _tlsConfig != nil {
		fmt.Println("文件解析错误, 原因:", err)
		return false, "文件解析错误"
	}
	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	if err != nil {
		fmt.Println("文件解析错误, 原因:", err)
		return false, "文件解析错误"
	}
	caData, err := ioutil.ReadFile(wechatCAPath)
	if err != nil {
		fmt.Println("文件解析错误, 原因:", err)
		return false, "文件解析错误"
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	tr := &http.Transport{TLSClientConfig: _tlsConfig}
	client := &http.Client{Transport: tr}
	ss, err := client.Post("https://api.mch.weixin.qq.com/secapi/pay/refund", "text/xml", bytes.NewBuffer(bytes_req))
	testReturn := mod.Refund{}
	bodys, _ := ioutil.ReadAll(ss.Body)
	_ = xml.Unmarshal(bodys, &testReturn)
	return true, yourReq.Out_refund_no
}

type FundQueryM struct {
	Appid         string `xml:"appid"`         //公众账号ID
	Mch_id        string `xml:"mch_id"`        //商户号
	Nonce_str     string `xml:"nonce_str"`     //随机字符串
	Out_refund_no string `xml:"out_refund_no"` //商户退款单号
	Out_trade_no  string `xml:"out_trade_no"`  //商户订单号
	Sign          string `xml:"sign"`          //签名
}

//查询退款
//out_trade_no  商户订单号    out_refund_no  商户退款单号
func FundQuery(out_trade_no string, out_refund_no string) {
	var yourReq FundQueryM
	yourReq.Appid = Wechat_AppId
	yourReq.Mch_id = Wechat_MchId
	yourReq.Nonce_str = strconv.Itoa(rand.Int())
	yourReq.Out_trade_no = out_trade_no
	yourReq.Out_refund_no = out_refund_no
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = yourReq.Appid
	m["mch_id"] = yourReq.Mch_id
	m["nonce_str"] = yourReq.Nonce_str
	m["out_trade_no"] = yourReq.Out_trade_no
	m["out_refund_no"] = yourReq.Out_refund_no
	yourReq.Sign = wxpayCalcSign(m, "shiliutvweixinzhifu1234567890123")
	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return
	}
	str_req := string(bytes_req)
	str_req = strings.Replace(str_req, "FundQueryM", "xml", -1)
	bytes_req = []byte(str_req)
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/refundquery", bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, ss := c.Do(req)
	if ss != nil {
		fmt.Println("请求微信查询退款单接口发送错误, 原因:", ss)
		return
	}
	testReturn := mod.FundQuery{}
	bodys, _ := ioutil.ReadAll(resp.Body)
	_ = xml.Unmarshal(bodys, &testReturn)
}

type DownloadBillModel struct {
	Appid     string `xml:"appid"`     //公众账号ID
	Bill_date string `xml:"bill_date"` //对账单日期
	Bill_type string `xml:"bill_type"` //账单类型
	Mch_id    string `xml:"mch_id"`    //商户号
	Nonce_str string `xml:"nonce_str"` //随机字符串
	Sign      string `xml:"sign"`      //签名
}

//对账单下载
func DownloadBill() {
	var (
		downloadBill DownloadBillModel
		randNum      bytes.Buffer
	)
	randNum.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	randNum.WriteString(strconv.Itoa(rand.Int()))
	downloadBill.Appid = Wechat_AppId
	downloadBill.Bill_date = "20161011" //strings.Replace(time.Now().Format("2006-01-02"), "-", "", -1)
	downloadBill.Bill_type = "ALL"
	downloadBill.Mch_id = Wechat_MchId
	downloadBill.Nonce_str = randNum.String()
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = downloadBill.Appid
	m["mch_id"] = downloadBill.Mch_id
	m["nonce_str"] = downloadBill.Nonce_str
	m["bill_date"] = "20161011"
	m["bill_type"] = downloadBill.Bill_type
	downloadBill.Sign = wxpayCalcSign(m, "shiliutvweixinzhifu1234567890123")
	bytes_req, err := xml.Marshal(downloadBill)
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return
	}
	str_req := string(bytes_req)
	str_req = strings.Replace(str_req, "DownloadBillModel", "xml", -1)
	fmt.Println(str_req)
	bytes_req = []byte(str_req)
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/downloadbill", bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, ss := c.Do(req)
	if ss != nil {
		fmt.Println("请求微信查询退款单接口发送错误, 原因:", ss)
		return
	}
	testReturn := mod.CallBackDownBill{}
	bodys, _ := ioutil.ReadAll(resp.Body)
	for index, val := range strings.SplitAfter(string(bodys), "\n") {
		fmt.Println("--------haha--------", strings.SplitAfter(val, ","))
		fmt.Println("-------------", index, val)
	}
	_ = xml.Unmarshal(bodys, &testReturn)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
