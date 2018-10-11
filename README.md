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
