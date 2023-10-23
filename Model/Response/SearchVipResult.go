package Response

type SearchVipResponseData struct {
	BaseResponse
	Page   SearchVipResponsePage `json:"page"`
	Result []SearchVipResult     `json:"result"`
}

type SearchVipResponsePage struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	Total       int `json:"total"`
	TotalPage   int `json:"totalPage"`
}

type SearchVipResult struct {
	ActiveCardDate      string  `json:"activeCardDate"`
	ActiveCardEndDate   string  `json:"activeCardEndDate"`
	ActiveCardStartDate string  `json:"activeCardStartDate"`
	CardKind            int     `json:"cardKind"`
	CardKindName        string  `json:"cardKindName"`
	CardNo              string  `json:"cardNo"`
	CardType            string  `json:"cardType"`
	CardTypeName        string  `json:"cardTypeName"`
	Code                string  `json:"code"`
	CompId              string  `json:"compId"`
	CompIds             string  `json:"compIds"`
	CustId              string  `json:"custId"`
	ExpireDate          string  `json:"expireDate"`
	Gca00C              string  `json:"gca00c"`
	Gca01C              string  `json:"gca01c"`
	Gca02C              string  `json:"gca02c"`
	Gca03I              int     `json:"gca03i"`
	Gca04C              string  `json:"gca04c"`
	Gca05D              string  `json:"gca05d"`
	Gca06D              string  `json:"gca06d"`
	Gca07D              string  `json:"gca07d"`
	Gca08I              int     `json:"gca08i"`
	Gca09C              string  `json:"gca09c"`
	Gca10F              float64 `json:"gca10f"`
	Gca11C              string  `json:"gca11c"`
	Gca12C              string  `json:"gca12c"`
	Gca13D              string  `json:"gca13d"`
	Gca14I              int     `json:"gca14i"`
	Gca15I              int     `json:"gca15i"`
	Gca16D              string  `json:"gca16d"`
	Gca17D              string  `json:"gca17d"`
	Gca18C              string  `json:"gca18c"`
	Gca19F              float64 `json:"gca19f"`
	Gca20C              string  `json:"gca20c"`
	Gca21D              string  `json:"gca21d"`
	Gca22D              string  `json:"gca22d"`
	Gca23I              int     `json:"gca23i"`
	Gca24F              float64 `json:"gca24f"`
	Gca25F              float64 `json:"gca25f"`
	Gca26F              float64 `json:"gca26f"`
	Gca27C              string  `json:"gca27c"`
	Gca28I              int     `json:"gca28i"`
	Gca29F              float64 `json:"gca29f"`
	Gca30D              string  `json:"gca30d"`
	Gca31C              string  `json:"gca31c"`
	Id                  string  `json:"id"`
	MemberMobile        string  `json:"memberMobile"`
	MemberName          string  `json:"memberName"`
	MemberNameAnalyzer  string  `json:"memberNameAnalyzer"`
	MemberNo            string  `json:"memberNo"`
	Period              float64 `json:"period"`
	PeriodUnit          string  `json:"periodUnit"`
	PublishCardCompany  string  `json:"publishCardCompany"`
	RemainAmount        float64 `json:"remainAmount"`
	Remark              string  `json:"remark"`
	SaleDate            string  `json:"saleDate"`
	SaleNo              string  `json:"saleNo"`
	SalesAmount         float64 `json:"salesAmount"`
	SalesMan            string  `json:"salesMan"`
	Status              string  `json:"status"`
	StatusName          string  `json:"statusName"`
}
