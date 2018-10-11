#玩一玩平台 后台验证登录流程模块 Golang 版本
##Usage

###Basic Auth Struct

```go
type BasicAuthInfo struct {
 	AppId     string
 	GameId    string
 	AppKey    string
 	OpenId    string
 	OpenKey   string
 	IsSandbox bool
 }
```

###Auth

```go
wywauth.CheckWywAuth(basicAuth BasicAuthInfo)
```

###Get Item List

```go
wywauth.GetItemList(basicAuth BasicAuthInfo)
```

###Consume Items

```go
wywauth.ConsumeItems(basicAuth BasicAuthInfo, itemId int, itemCnt int)
```

###Gift Items

```go
wywauth.GiftItems(basicAuth BasicAuthInfo, itemId int, itemCnt int,  actType int)
```

Sample Code

```go
// put these in config file
const APPID = "11111111"
const GAMEID = "1111"
const APPKEY = "sample&" // must have & at last
const ISSANDBOX = true

type ReqWithOpenid struct {
	Openid  string
	Openkey string
}

// login
func Login(req ReqWithOpenid) {

	auth := wywauth.BasicAuthInfo{
		AppId:     APPID,
		GameId:    GAMEID,
		AppKey:    APPKEY,
		OpenId:    req.Openid,
		OpenKey:   req.Openkey,
		IsSandbox: ISSANDBOX}

	ret := wywauth.CheckWywAuth(auth)
	...
}

func GetItemList(req ReqWithOpenid) {
	auth := wywauth.BasicAuthInfo{
		AppId:     APPID,
		GameId:    GAMEID,
		AppKey:    APPKEY,
		OpenId:    req.Openid,
		OpenKey:   req.Openkey,
		IsSandbox: ISSANDBOX}
	ret := wywauth.GetItemList(auth)
	...
}

type ReqConsumeItem struct {
	Openid  string
	Openkey string
	Itemid  int
	Itemcnt int
}

func ConsumeItem(req ReqConsumeItem) {

	auth := wywauth.BasicAuthInfo{
		AppId:     APPID,
		GameId:    GAMEID,
		AppKey:    APPKEY,
		OpenId:    req.Openid,
		OpenKey:   req.Openkey,
		IsSandbox: ISSANDBOX}

	ret := wywauth.ConsumeItems(auth, req.Itemid, req.Itemcnt)
	...
}

func GiftItem(req ReqConsumeItem) {

	auth := wywauth.BasicAuthInfo{
		AppId:     APPID,
		GameId:    GAMEID,
		AppKey:    APPKEY,
		OpenId:    req.Openid,
		OpenKey:   req.Openkey,
		IsSandbox: ISSANDBOX}

	ret := wywauth.GiftItems(auth, req.Itemid, req.Itemcnt, 1)
    ...
}

```