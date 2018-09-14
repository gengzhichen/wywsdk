package wywauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
	"time"
)

const SANDBOXURL = "http://openapi-test.hudong.qq.com"
const PRODURL = "https://openapi.hudong.qq.com"
const URIAUTH = "/openapi/apollo_verify_openid_openkey"
const URIITEM = "/openapi/apollo_game_item_proxy"

// Basic Auth Requirement
type BasicAuthInfo struct {
	AppId     string
	GameId    string
	AppKey    string
	OpenId    string
	OpenKey   string
	IsSandbox bool
}

// Wyw Login Response
type RespWywLogin struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

//
type WywItem struct {
	Id       int    `json:"id"`
	Num      int    `json:"num"`
	Name     string `json:"name"`
	Consumed int    `json:"consumed"`
	IconUrl  string `json:"iconurl"`
}

//
type ItemList struct {
	ItemList []WywItem
}

type RespGetItemList struct {
	Ret  int      `json:"ret"`
	Msg  string   `json:"msg"`
	Data ItemList `json:"data"`
}

type ConsumeItem struct {
	Id  int    `json:"id"`
	Seq string `json:"seq"`
}

type ConsumeItemList struct {
	SuccList []ConsumeItem `json:"succList"`
	FailList []ConsumeItem `json:"failList"`
}

type RespConsumeItem struct {
	Ret  int             `json:"ret"`
	Data ConsumeItemList `json:"data"`
}

// 获取当前时间戳
func MakeTimestamp() int {
	return int((time.Now().UnixNano() / int64(time.Millisecond)) / 1000)
}

// Get url according to Env
// sign
// http://wiki.open.qq.com/wiki/
func addSig(uri string, appKey string, parms *map[string]string) string {

	// 第1步：将请求的URI路径进行URL编码（URI不含host，URI示例：/v3/user/get_info）
	escapeURI := strings.Replace(url.QueryEscape(uri), "%26", "_", -1)
	signStr := "POST&" + escapeURI + "&"

	// 第2步：将除“sig”外的所有参数按key进行字典升序排列
	var keys []string
	for k := range *parms {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 第3步：将第2步中排序后的参数(key=value)用&拼接起来，并进行URL编码。
	signParams := []string{}
	for _, k := range keys {
		signParams = append(signParams, k+"="+(*parms)[k])
	}
	reqBodyStr := strings.Join(signParams, "&")
	signStr = signStr + url.QueryEscape(reqBodyStr)

	// 构造密钥
	// 得到密钥的方式：在应用的appkey末尾加上一个字节的“&”，即appkey&, 在输入时直接构造好，
	mac := hmac.New(sha1.New, []byte(appKey))
	mac.Write([]byte(signStr))

	// 生成签名值
	sig := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	// 完成body拼接
	reqBodyStr = reqBodyStr + "&sig=" + sig

	return reqBodyStr

}
