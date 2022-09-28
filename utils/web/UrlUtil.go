package web

// 微信网页授权，获取用户信息
func WechatLogin(access_token string,openid string)string{
	return "https://api.weixin.qq.com/sns/userinfo?access_token="+access_token+"&openid="+openid+"&lang=zh_CN"
}

// 微信网页授权，获取code对应用户的access token
func WechatLoginAccess(appid string,sercet string,code string)string{
	return "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appid + "&secret=" + sercet + "&code=" + code + "&grant_type=authorization_code"
}

//微信小程序的access token
func AccessToken(appid string,secret string)string{
	return "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret
}

//微信网页授权的js sdk ticket
func JsApiTicket(accessToken string)string{
	return "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token="+ accessToken +"&type=jsapi"
}

// 微信模板消息
func TemplateMessage(accessToken string)string{
	return "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + accessToken
}