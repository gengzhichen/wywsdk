package main

import (

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"net/url"
	"sort"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"io/ioutil"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})



	v1 := app.Party("/", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
	{

		v1.Post("/login", Login)


	}

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to ！</h1>")
	})

	//app.Post("/encode", serverhandler.Test)

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	//app.Get("/ping", func(ctx iris.Context) {
	//	ctx.WriteString("pong")
	//})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	//app.Get("/hello", func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	//})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}


// 这部分放在配置文件中
const APPID = "101497571"
const GAMEID = "3398"
const APPKEY = "WZDMXZCWL123CZWL&"
const SANDBOXURL = "http://openapi-test.hudong.qq.com"
const PRODURL = "https://openapi.hudong.qq.com"
const URI = "/openapi/apollo_verify_openid_openkey"
const ISSANDBOX = true

/**
签名
http://wiki.open.qq.com/wiki/
*/
func checkWywAuth(openId string, openKey string) string {

	var targetURL string
	if ISSANDBOX {
		targetURL = SANDBOXURL + URI
	} else {
		targetURL = PRODURL + URI
	}

	rand.Seed(time.Now().UnixNano())
	body := map[string]string{}
	body["appid"] = APPID
	body["gameid"] = GAMEID
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
	mac := hmac.New(sha1.New, []byte(APPKEY))
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

	return string(respBody)
}

/**
获取当前时间戳
*/
func MakeTimestamp() int {
	return int((time.Now().UnixNano() / int64(time.Millisecond)) / 1000)
}

type ReqWithOpenid struct {
	Openid  string
	Openkey string
}

func Login(ctx iris.Context) {

	var req ReqWithOpenid
	err := ctx.ReadJSON(&req)

	if err != nil {
		ctx.Write([]byte("error"))
		return
	}

	// ctx.JSON(genCommon())
	ctx.Write([]byte(checkWywAuth(req.Openid, req.Openkey)))
}