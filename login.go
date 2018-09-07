package wywauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const SANDBOXURL = "http://openapi-test.hudong.qq.com"
const PRODURL = "https://openapi.hudong.qq.com"
const URI = "/openapi/apollo_verify_openid_openkey"

type RespWywLogin struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

/**
获取当前时间戳
*/
func MakeTimestamp() int {
	return int((time.Now().UnixNano() / int64(time.Millisecond)) / 1000)
}

/**
签名
http://wiki.open.qq.com/wiki/
*/
func CheckWywAuth(appId string, gameId string, appKey string, openId string, openKey string, isSandbox bool) RespWywLogin {

	var targetURL string
	if isSandbox {
		targetURL = SANDBOXURL + URI
	} else {
		targetURL = PRODURL + URI
	}

	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = appId
	body["gameid"] = gameId
	body["openid"] = openId
	body["openkey"] = openKey
	body["ts"] = strconv.Itoa(MakeTimestamp())
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))

	// 第1步：将请求的URI路径进行URL编码（URI不含host，URI示例：/v3/user/get_info）
	escapeURI := strings.Replace(url.QueryEscape(URI), "%26", "_", -1)
	signStr := "POST&" + escapeURI + "&"

	// 第2步：将除“sig”外的所有参数按key进行字典升序排列
	var keys []string
	for k := range body {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 第3步：将第2步中排序后的参数(key=value)用&拼接起来，并进行URL编码。
	signParams := []string{}
	for _, k := range keys {
		signParams = append(signParams, k+"="+body[k])
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

	// 发送POST请求
	resp, err := http.Post(targetURL, "application/x-www-form-urlencoded", strings.NewReader(reqBodyStr))
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	ret := RespWywLogin{}
	json.Unmarshal(respBody, &ret)

	return ret
}
