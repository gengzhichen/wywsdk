#玩一玩平台 后台验证登录流程模块 Golang 版本
##Usage
###认证
`wywauth.CheckWywAuth(APPID, GAMEID, APPKEY, Openid, Openkey, true)`

###获取道具一览
`wywauth.GetItemList(appId string, gameId string, appKey string, openId string, openKey string, isSandbox bool)`

###消耗道具
`wywauth.ConsumeItems(appId string, gameId string, appKey string, openId string, openKey string, itemId int, itemCnt int, isSandbox bool)`

###赠送道具
`wywauth.GiftItems(appId string, gameId string, appKey string, openId string, openKey string, itemId int, itemCnt int,  actType int,isSandbox bool)`