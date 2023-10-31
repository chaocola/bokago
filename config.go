package bokago

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"time"
)

//
//  Config
//  @Description: 博卡主要请求配置
//

type Config struct {
	CustID            string `json:"custId"`
	CompID            string `json:"compId"`
	UserName          string `json:"userName"`
	PassWord          string `json:"passWord"`
	Source            string `json:"source"`
	Sign              string `json:"sign"`
	AdminWechatOpenID string `json:"adminWechatOpenId"`
	Token             Token  `json:"token"`
	EmpName           string `json:"empName"`
	CompName          string `json:"compName"`
}

func (config *Config) GetHeaders() map[string]interface{} {
	token := config.Token

	if time.Now().Unix()-privateExpire > token.StartTime {
		// TODO: - 重新获取token
		privateTokenTaskSignal <- true
		config.GetAccessToken()
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
		"referer":      privateReferer,
		"origin":       privateReferer,
		"access_token": token.AccessToken,
		"accesstoken":  token.ShopID,
		"device_id":    config.AdminWechatOpenID,
		"deviceid":     config.AdminWechatOpenID,
		"tenant":       config.CustID,
	}
}

// GET
//
//	@Description: GET请求
//	@receiver DefaultConfig
//	@param url
//	@param params
//	@return []byte
func (config *Config) GET(url string, params map[string]interface{}, other ...time.Duration) []byte {
	return Client.GET(url, params, config.GetHeaders(), other...)
}

// POST
//
//	@Description: POST请求
//	@receiver DefaultConfig
//	@param url
//	@param params
//	@param body
//	@return []byte
func (config *Config) POST(url string, params map[string]interface{}, body interface{}, other ...time.Duration) []byte {
	return Client.POST(url, params, config.GetHeaders(), body, other...)
}

// TokenTask
//
//	@Description: 定时任务
func (config *Config) TokenTask() {
	ticker := time.NewTicker(time.Duration(privateExpire) * time.Second)
	defer func() {
		if err := recover(); err != nil {
			ticker.Stop()
			time.Sleep(time.Second)
			go config.TokenTask()

		}
	}()
	for {
		select {
		case <-ticker.C:
			config.GetAccessToken()
		case <-privateTokenTaskSignal:
			ticker.Reset(time.Duration(privateExpire) * time.Second)
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
func (config *Config) GetAccessToken() bool {

	var data AccessTokenResponse

	BASEURL := "https://api.bokao2o.com/auth/merchant/v2/user/login"

	res := Client.POST(
		BASEURL,
		nil,
		map[string]interface{}{
			"referer": "https://s3.boka.vc/",
		},
		map[string]interface{}{
			"custId":   config.CustID,   // 门店编号
			"compId":   config.CompID,   // 连锁代码
			"userName": config.UserName, // 用户名
			"passWord": config.PassWord, // 密码
			"source":   config.Source,
		})

	err := json.Unmarshal(res, &data)

	if err != nil {
		config.Token = Token{
			Error: err,
		}
		return false
	}

	if data.Code == 200 || data.Success {

		config.CompName = data.Result.CompName
		config.Token = Token{
			AccessToken: data.Result.Token,
			ShopID:      data.Result.ShopID,
			StartTime:   time.Now().Unix(),
			Error:       nil,
		}
		return true
	} else {
		config.Token = Token{
			Error: errors.New(data.Msg),
		}
		return false
	}

}

func (config *Config) SetOther(other ...interface{}) {
	SetConfigData := func(data interface{}) {
		switch reflect.TypeOf(data).String() {
		case "string":
			privateReferer = data.(string)
		case "int":
			privateExpire = int64(data.(int))
		case "int64":
			privateExpire = data.(int64)
		}
	}

	if len(other) > 0 && len(other) == 1 {
		SetConfigData(other[0])
		return

	} else if len(other) > 1 {
		SetConfigData(reflect.TypeOf(other[0]).String())
		SetConfigData(reflect.TypeOf(other[1]).String())

		return
	}

}

func (config *Config) Init(other ...interface{}) *Config {
	config.GetAccessToken()
	config.SetOther(other...)
	return config
}
