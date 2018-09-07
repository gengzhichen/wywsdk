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
func CheckWywAuth(appId string, gameId string, appKey string, openId string, openKey string, isSandbox bool) RespWywLogin {

	var targetURL string

	if isSandbox {
		targetURL = SANDBOXURL + URIAUTH
	} else {
		targetURL = PRODURL + URIAUTH
	}

	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = appId
	body["gameid"] = gameId
	body["openid"] = openId
	body["openkey"] = openKey
	body["ts"] = strconv.Itoa(MakeTimestamp())
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))

	reqBodyStr := addSig(URIAUTH,appKey,&body)

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
