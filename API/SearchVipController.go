package API

import (
	"encoding/json"
	"fmt"
	"github.com/oddbug/bokago/DefaultConfig"
	"github.com/oddbug/bokago/Model/Response"
)

// SearchVip
//
//	@Description: 搜索会员信息
//	- @param searchText 搜索内容
//	- @param page 页码
//	- @param queryType 查询类型
//	- @return Response.SearchVipResponse
func SearchVip(searchText string, page int, queryType int) (data Response.SearchVipResponse) {

	Url := fmt.Sprintf("https://api.bokao2o.com/s3connect/s3memberCard/v2/comp/%v/get", DefaultConfig.BOKA.CompID)

	var Params = map[string]interface{}{
		"searchText": searchText,
		"page":       page,
		"queryType":  queryType,
		"sign":       DefaultConfig.BOKA.Sign,
	}

	res := DefaultConfig.BOKA.GET(Url, Params)
	_ = json.Unmarshal(res, &data)

	return
}
