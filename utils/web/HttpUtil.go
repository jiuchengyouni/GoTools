package web

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
	"go.uber.org/zap"
)

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var client = HttpRequest.Transport(transport)

// Get 发送 get 请求
func Get(ctx *gin.Context, url string, param map[string]interface{}) (*HttpRequest.Response, error) {
	if res, err := client.Get(url, param); err != nil {
		return res, err
	} else {
		zap.L().Info(
			"HttpRequest",
			zap.String("url", res.Url()),
			zap.String("method", "Get"),
			zap.Int("status", res.StatusCode()),
		)
		return res, nil
	}
}
// func DoGet(urlPath string, parms map[string]interface{}) string {
// 	var sb strings.Builder
// 	if parms !=nil&&len(parms)!=0{
// 		sb.WriteString("?")
// 		valueUrl:=""
// 		for key,value:=range parms{
// 			if value!=nil {
// 				valueUrl= fmt.Sprintf("%v", value)
// 				value=url.QueryEscape(valueUrl)
// 			}
// 			sb.WriteString(key)
// 			sb.WriteString("=")
// 			sb.WriteString(valueUrl)
// 		}
// 	}
// 	// 删除最后一个&
// 	str:=sb.String()
// 	if len(str)>0{
// 		str=str[:len(str)-1]
// 	}
// 	http.Get(urlPath+str)
// 	return ""
// }
