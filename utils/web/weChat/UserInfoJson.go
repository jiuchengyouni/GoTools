package wechat

type UserInfoJson struct {
	Openid     string   `json:"openid" validate:"required"`
	Nickname   string   `json:"nickname" validate:"required"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}