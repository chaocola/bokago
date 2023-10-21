package bokago

import (
	"encoding/json"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ShopID      string `json:"shop_id"`
	StartTime   int64  `json:"start_time"`
	Error       string `json:"error"`
}

// GetAccessToken
//
//	@Description: 登录获取token
//	@param custID 门店编号
//	@param compID 连锁代码
//	@param userName 用户名
//	@param passWord 密码
//	@param source 来源
//	@return TokenContent token信息
func GetAccessToken(custID string, compID string, userName string, passWord string, source string) Token {
	BASEURL := "https://api.bokao2o.com/auth/merchant/v2/user/login"
	headers := map[string]interface{}{
		"referer": "https://s3.boka.vc/",
	}

	res := Requests.POST(BASEURL, nil, headers, map[string]interface{}{
		"custId":   custID,   // 门店编号
		"compId":   compID,   // 连锁代码
		"userName": userName, // 用户名
		"passWord": passWord, // 密码
		"source":   source,
	})
	var data AccessTokenResponse
	_ = json.Unmarshal(res, &data)

	if data.Code == 200 || data.Success {
		return Token{
			data.Result.Token,
			data.Result.ShopID,
			time.Now().Unix(),
			"",
		}
	}
	return Token{
		"",
		"",
		0,
		data.Msg,
	}
}
