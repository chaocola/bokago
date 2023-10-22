package bokago

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

var (
	BOKA            *Config
	Expire          int64 = 5000
	Referer               = `https://s3.boka.vc/`
	TokenTaskSignal       = make(chan bool, 1)
)

//
//  Config
//  @Description: 博卡主要请求配置
//

type Config struct {
	CustID   string `json:"custId"`
	CompID   string `json:"compId"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Source   string `json:"source"`
	Sign     string `json:"sign"`

	// 微信登录用户绑定的 WechatOpenid
	DeviceID string `json:"deviceId"`
	// 程序内部使用
	Token   Token  `json:"token"`
	EmpName string `json:"empName"`
}

func (config *Config) GetHeaders() map[string]interface{} {
	token := config.Token

	if time.Now().Unix()-Expire > token.StartTime {
		// TODO: - 重新获取token
		TokenTaskSignal <- true
		BOKA.Token = GetAccessToken(BOKA.CustID, BOKA.CompID, BOKA.UserName, BOKA.PassWord, BOKA.Source)
	}
	empName := url.QueryEscape(config.EmpName)
	return map[string]interface{}{
		"Cookie": fmt.Sprintf(`subCustType=; token=%s; custCode=%v; custId=%v; compId=%v; shopId=%s; empId=%v; empName=%v;`,
			token.AccessToken,
			config.CustID,
			config.CustID,
			config.CompID,
			token.ShopID,
			config.UserName,
			empName,
		),
		"referer":      Referer,
		"origin":       Referer,
		"access_token": token.AccessToken,
		"accesstoken":  token.ShopID,
		"device_id":    config.DeviceID,
		"deviceid":     config.DeviceID,
		"tenant":       config.CustID,
	}
}

// GET
//
//	@Description: GET请求
//	@receiver config
//	@param url
//	@param params
//	@return []byte
func (config *Config) GET(url string, params map[string]interface{}) []byte {
	data := Requests.GET(url, params, config.GetHeaders())
	if !checkDataCode(data) {
		data = Requests.GET(url, params, config.GetHeaders())
	}

	return data

}

// POST
//
//	@Description: POST请求
//	@receiver config
//	@param url
//	@param params
//	@param body
//	@return []byte
func (config *Config) POST(url string, params map[string]interface{}, body interface{}) []byte {

	data := Requests.POST(url, params, config.GetHeaders(), body)
	if !checkDataCode(data) {
		data = Requests.POST(url, params, config.GetHeaders(), body)
	}
	return data
}

// checkDataCode
//
//	@Description: 检查返回的数据是否是403
//	@receiver config
//	@param data
//	@return bool
func checkDataCode(data []byte) bool {
	var res BaseResponse
	_ = json.Unmarshal(data, &res)
	if res.Code == 403 {
		TokenTaskSignal <- true
		SetToken()
		return false
	}
	return true
}

// SetToken
//
//	@Description: 设置token
func SetToken() {
	BOKA.Token = GetAccessToken(BOKA.CustID, BOKA.CompID, BOKA.UserName, BOKA.PassWord, BOKA.Source)
}

// TokenTask
//
//	@Description: 定时任务
func TokenTask() {
	ticker := time.NewTicker(time.Duration(Expire) * time.Second)
	defer func() {
		if err := recover(); err != nil {
			ticker.Stop()
			time.Sleep(time.Second)
			TokenTask()

		}
	}()
	for {
		select {
		case <-ticker.C:
			SetToken()
		case <-TokenTaskSignal:
			ticker.Reset(time.Duration(Expire) * time.Second)
		}
	}
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

// Init
//
//	@Description: 初始化 主要方法 必须要在一开始调用
//	@param custID 门店编号
//	@param compID 连锁代码
//	@param userName 用户名
//	@param passWord 密码
//	@param source 来源
//	@param sign 签名
func Init(custID string, compID string, userName string, passWord string, source string, sign string) {

	BOKA = &Config{
		CustID:   custID,
		CompID:   compID,
		UserName: userName,
		PassWord: passWord,
		Source:   source,
		Sign:     sign,
		Token:    GetAccessToken(custID, compID, userName, passWord, source),
	}

	go TokenTask()

}
