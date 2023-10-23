package Response

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
