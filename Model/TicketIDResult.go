package Model

// TicketIDResult
// @Description: 开单单号
type TicketIDResult struct {
	BillDate     int64  `json:"billDate"`
	BillNo       string `json:"billNo"`
	BillTime     string `json:"billTime"`
	ChangeBillNo int    `json:"changeBillNo"`
}
