package bokago

import "github.com/oddbug/bokago/DefaultConfig"

var (
	BOKA *DefaultConfig.Config
)

// Init
//
//	@Description:
//	@param custID 门店编号
//	@param compID 连锁代码
//	@param userName 用户名
//	@param passWord 密码
//	@param source  来源
//	@param sign   签名
//	@param other 其他参数 Expire int64 token过期时间 默认为5000 Referer string 地址 默认为https://s3.boka.vc/
func Init(custID string, compID string, userName string, passWord string, source string, sign string, other ...interface{}) {

	BOKA = &DefaultConfig.Config{
		CustID:   custID,
		CompID:   compID,
		UserName: userName,
		PassWord: passWord,
		Source:   source,
		Sign:     sign,
	}

	BOKA.Init(other...)

	go BOKA.TokenTask()

}
