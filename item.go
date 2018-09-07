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

// Get Item List from wyw
func GetItemList(basicAuth BasicAuthInfo) RespGetItemList {

	var targetURL string
	if basicAuth.IsSandbox {
		targetURL = SANDBOXURL + URIITEM
	} else {
		targetURL = PRODURL + URIITEM
	}

	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = basicAuth.AppId
	body["gameid"] = basicAuth.GameId
	body["openid"] = basicAuth.OpenId
	body["openkey"] = basicAuth.OpenKey
	body["ts"] = strconv.Itoa(MakeTimestamp())
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))
	body["cmd"] = "1"
	body["mask"] = "1"

	reqBodyStr := addSig(URIITEM, basicAuth.AppKey, &body)

	// 发送POST请求
	resp, err := http.Post(targetURL, "application/x-www-form-urlencoded", strings.NewReader(reqBodyStr))
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	ret := RespGetItemList{}
	json.Unmarshal(respBody, &ret)

	return ret
}

// Get Item List from wyw
func ConsumeItems(basicAuth BasicAuthInfo, itemId int, itemCnt int) RespConsumeItem {

	var targetURL string
	if basicAuth.IsSandbox {
		targetURL = SANDBOXURL + URIITEM
	} else {
		targetURL = PRODURL + URIITEM
	}

	now := strconv.Itoa(MakeTimestamp())
	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = basicAuth.AppId
	body["gameid"] = basicAuth.GameId
	body["openid"] = basicAuth.OpenId
	body["openkey"] = basicAuth.OpenKey
	body["itemids"] = strconv.Itoa(itemId)
	body["itemnums"] = strconv.Itoa(itemCnt)
	body["ts"] = now
	body["itemseqs"] = now
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))
	body["cmd"] = "2"
	body["mask"] = "1"

	reqBodyStr := addSig(URIITEM, basicAuth.AppKey, &body)

	// 发送POST请求
	resp, err := http.Post(targetURL, "application/x-www-form-urlencoded", strings.NewReader(reqBodyStr))
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)


	ret := RespConsumeItem{}
	json.Unmarshal(respBody, &ret)

	return ret
}


// Gift an Item to user
func GiftItems(basicAuth BasicAuthInfo, itemId int, itemCnt int,  actType int) RespGetItemList {

	var targetURL string
	if basicAuth.IsSandbox {
		targetURL = SANDBOXURL + URIITEM
	} else {
		targetURL = PRODURL + URIITEM
	}

	now := strconv.Itoa(MakeTimestamp())
	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = basicAuth.AppId
	body["gameid"] = basicAuth.GameId
	body["openid"] = basicAuth.OpenId
	body["openkey"] = basicAuth.OpenKey
	body["itemids"] = strconv.Itoa(itemId)
	body["itemnums"] = strconv.Itoa(itemCnt)
	body["acttype"] = strconv.Itoa(actType)
	body["ts"] = now
	body["itemseqs"] = now
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))
	body["cmd"] = "3"
	body["mask"] = "1"

	reqBodyStr := addSig(URIITEM, basicAuth.AppKey, &body)

	// 发送POST请求
	resp, err := http.Post(targetURL, "application/x-www-form-urlencoded", strings.NewReader(reqBodyStr))
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)


	ret := RespGetItemList{}
	json.Unmarshal(respBody, &ret)

	return ret
}
