package Response

// BaseResponse 基础返回结构体
type BaseResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

// Response 基础返回结构体
type Response[T any] struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Result  T      `json:"result"`
}

// AccessTokenResponse 获取token
type AccessTokenResponse = Response[AccessTokenResult]

// TicketIDResponse 获取开单单号
type TicketIDResponse = Response[TicketIDResult]

// SearchVipResponse 搜索会员信息
type SearchVipResponse = SearchVipResponseData
