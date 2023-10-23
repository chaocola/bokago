package Response

type Token struct {
	AccessToken string `json:"accessToken"`
	ShopID      string `json:"shopId"`
	StartTime   int64  `json:"startTime"`
	Error       error  `json:"error"`
}
