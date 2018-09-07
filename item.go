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
func GetItemList(appId string, gameId string, appKey string, openId string, openKey string, isSandbox bool) RespGetItemList {

	var targetURL string
	if isSandbox {
		targetURL = SANDBOXURL + URIITEM
	} else {
		targetURL = PRODURL + URIITEM
	}

	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = appId
	body["gameid"] = gameId
	body["openid"] = openId
	body["openkey"] = openKey
	body["ts"] = strconv.Itoa(MakeTimestamp())
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))
	body["cmd"] = "1"
	body["mask"] = "1"

	reqBodyStr := addSig(URIITEM, appKey, &body)

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
func ConsumeItems(appId string, gameId string, appKey string, openId string, openKey string, itemId int, itemCnt int, isSandbox bool) RespConsumeItem {

	var targetURL string
	if isSandbox {
		targetURL = SANDBOXURL + URIITEM
	} else {
		targetURL = PRODURL + URIITEM
	}

	now := strconv.Itoa(MakeTimestamp())
	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = appId
	body["gameid"] = gameId
	body["openid"] = openId
	body["openkey"] = openKey
	body["itemids"] = strconv.Itoa(itemId)
	body["itemnums"] = strconv.Itoa(itemCnt)
	body["ts"] = now
	body["itemseqs"] = now
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))
	body["cmd"] = "2"
	body["mask"] = "1"

	reqBodyStr := addSig(URIITEM, appKey, &body)

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
func GiftItems(appId string, gameId string, appKey string, openId string, openKey string, itemId int, itemCnt int,  actType int,isSandbox bool) RespGetItemList {

	var targetURL string
	if isSandbox {
		targetURL = SANDBOXURL + URIITEM
	} else {
		targetURL = PRODURL + URIITEM
	}

	now := strconv.Itoa(MakeTimestamp())
	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = appId
	body["gameid"] = gameId
	body["openid"] = openId
	body["openkey"] = openKey
	body["itemids"] = strconv.Itoa(itemId)
	body["itemnums"] = strconv.Itoa(itemCnt)
	body["acttype"] = strconv.Itoa(actType)
	body["ts"] = now
	body["itemseqs"] = now
	body["rnd"] = strconv.Itoa(rand.Intn(9999999))
	body["cmd"] = "3"
	body["mask"] = "1"

	reqBodyStr := addSig(URIITEM, appKey, &body)

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
