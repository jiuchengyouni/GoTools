package wechat

type AssessTokenJson struct{
	Access_token string `json:"access_token" validate:"required"`
	Expires_in int `json:"expires_in" validate:"required"`
	Refresh_token string `json:"refresh_token" validate:"required"`
	Openid string `json:"openid" validate:"required"`
	Scope string  `json:"scope" validate:"required"`
}