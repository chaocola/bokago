package Model

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
