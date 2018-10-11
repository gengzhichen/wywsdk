#玩一玩平台 后台验证登录流程模块 Golang 版本
##Usage

###基础人证信息
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
###认证
```go
wywauth.CheckWywAuth(basicAuth BasicAuthInfo)
```
###获取道具一览
```go
wywauth.GetItemList(basicAuth BasicAuthInfo)
```
###消耗道具
```go
wywauth.ConsumeItems(basicAuth BasicAuthInfo, itemId int, itemCnt int)
```
###赠送道具
```go
wywauth.GiftItems(basicAuth BasicAuthInfo, itemId int, itemCnt int,  actType int)
```
