package wywauth

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Check login
// Ref to http://wiki.open.qq.com/wiki/
func CheckWywAuth(basicAuth BasicAuthInfo) RespWywLogin {

	var targetURL string

	if basicAuth.IsSandbox {
		targetURL = SANDBOXURL + URIAUTH
	} else {
		targetURL = PRODURL + URIAUTH
	}

	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = basicAuth.AppId
	body["gameid"] = basicAuth.GameId
	body["openid"] = basicAuth.OpenId
	body["openkey"] = basicAuth.OpenKey
	body["ts"] = strconv.Itoa(MakeTimestamp())
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))

	reqBodyStr := addSig(URIAUTH, basicAuth.AppKey, &body)

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
