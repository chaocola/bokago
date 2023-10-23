package bokago

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

type AccessTokenResult struct {
	EmpID        string `json:"empId"`
	SuperManager int    `json:"superManager"`
	CompName     string `json:"compName"`
	Demo         int    `json:"demo"`
	Token        string `json:"token"`
	ExpiryDate   int64  `json:"expiryDate"`
	Password     string `json:"password"`
	CompID       string `json:"compId"`
	EmpName      string `json:"empName"`
	CustID       string `json:"custId"`
	StaffEmpID   string `json:"staffEmpId"`
	CustType     string `json:"custType"`
	IsMore       bool   `json:"isMore"`
	ShopID       string `json:"shopId"`
}

type Token struct {
	AccessToken string `json:"accessToken"`
	ShopID      string `json:"shopId"`
	StartTime   int64  `json:"startTime"`
	Error       error  `json:"error"`
}
