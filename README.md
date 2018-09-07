#玩一玩平台 后台验证登录流程模块 Golang 版本
##Usage

###基础人证信息
`type BasicAuthInfo struct {
 	AppId     string
 	GameId    string
 	AppKey    string
 	OpenId    string
 	OpenKey   string
 	IsSandbox bool
 }`

###认证
`wywauth.CheckWywAuth(basicAuth BasicAuthInfo)`

###获取道具一览
`wywauth.GetItemList(basicAuth BasicAuthInfo)`

###消耗道具
`wywauth.ConsumeItems(basicAuth BasicAuthInfo, itemId int, itemCnt int)`

###赠送道具
`wywauth.GiftItems(basicAuth BasicAuthInfo, itemId int, itemCnt int,  actType int)`