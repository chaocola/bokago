package bokago

type BaseResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

// -------------------------------------------------------------------------------------------

// AccessTokenResponse   登录获取token
type AccessTokenResponse struct {
	BaseResponse
	Result struct {
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
	} `json:"result"`
}

// -------------------------------------------------------------------------------------------

// TodayTicketOne 消费单据
type TodayTicketOne struct {
	BaseResponse
	Result []TodayTicketOneResult `json:"result"`
}

type TodayTicketOneResult struct {
	UpdateDate    int64         `json:"updateDate"`
	Gender        int           `json:"gender"`
	UpdateDateStr string        `json:"updateDateStr"`
	AliasNumber   string        `json:"aliasNumber"`
	PayWay        int           `json:"payWay"`
	Source        string        `json:"source"`
	Type          int           `json:"type"`
	EmpIdList     []string      `json:"empIdList"`
	BillingType   int           `json:"billingType"`
	CashFlag      int           `json:"cashFlag"`
	ProductInfos  []interface{} `json:"productInfos"`
	CustId        string        `json:"custId"`
	HadRefund     int           `json:"hadRefund"`
	Gga14F        int           `json:"gga14f"`
	ReqUserKey    string        `json:"reqUserKey"`
	HadModify     int           `json:"hadModify"`
	Id            string        `json:"id"`
	ShopId        string        `json:"shopId"`
	NewBilling    int           `json:"newBilling,omitempty"`
	HadFill       int           `json:"hadFill"`
	CreateDate    int64         `json:"createDate"`
	RemainInfoMap struct {
	} `json:"remainInfoMap"`
	Amount        float64 `json:"amount"`
	DurationTime  int     `json:"durationTime"`
	Product       string  `json:"product"`
	CreateDateStr string  `json:"createDateStr"`
	ServiceInfos  []struct {
		TasteFlag   int     `json:"tasteFlag"`
		BillingId   string  `json:"billingId"`
		TotalPrice  float64 `json:"totalPrice"`
		ServiceEmps []struct {
			IsStartCountReserveFlag  string        `json:"isStartCountReserveFlag"`
			ServiceType              string        `json:"serviceType"`
			UnLunFlag                string        `json:"unLunFlag"`
			EmpId                    string        `json:"empId"`
			BillingId                string        `json:"billingId"`
			EmpType                  int           `json:"empType"`
			DepartmentId             string        `json:"departmentId,omitempty"`
			EmpInBillingTime         int           `json:"empInBillingTime"`
			IsAlwaysEndOfQueue       int           `json:"isAlwaysEndOfQueue"`
			JobTitle                 string        `json:"jobTitle"`
			EmpLuntotalNum           int           `json:"empLuntotalNum"`
			Scale                    int           `json:"scale"`
			ReserveNum               int           `json:"reserveNum"`
			AppointFlag              int           `json:"appointFlag"`
			DianNum                  int           `json:"dianNum"`
			EmpIntervalNum           int           `json:"empIntervalNum"`
			ServiceStatus            int           `json:"serviceStatus"`
			EmpName                  string        `json:"empName"`
			CreateDate               int64         `json:"createDate"`
			WorkStatus               int           `json:"workStatus"`
			ReserveEnableInQueueFlag string        `json:"reserveEnableInQueueFlag"`
			EmpInQueuePosition       int           `json:"empInQueuePosition"`
			InitFlag                 int           `json:"initFlag"`
			OvertimeFlag             string        `json:"overtimeFlag"`
			Sex                      int           `json:"sex"`
			Index                    string        `json:"index"`
			EmpModified              int           `json:"empModified"`
			EmpAppointNum            int           `json:"empAppointNum"`
			HasPushMq                bool          `json:"hasPushMq"`
			LunNum                   int           `json:"lunNum"`
			ReserveFlag              string        `json:"reserveFlag"`
			BillingNum               int           `json:"billingNum"`
			CompId                   string        `json:"compId"`
			PushVision               int           `json:"pushVision"`
			Name                     string        `json:"name"`
			EnableServer             int           `json:"enableServer"`
			ProjectTimeDetailBOS     []interface{} `json:"projectTimeDetailBOS"`
			IsExchange               int           `json:"isExchange"`
			ProjectId                string        `json:"projectId"`
			ServiceTypeName          string        `json:"serviceTypeName"`
			Status                   int           `json:"status"`
			CustId                   string        `json:"custId,omitempty"`
			UpdateStatus             int           `json:"updateStatus,omitempty"`
		} `json:"serviceEmps"`
		Discount                float64 `json:"discount"`
		FinalPrice              float64 `json:"finalPrice"`
		HasfreeFlag             int     `json:"hasfreeFlag"`
		SellerId                string  `json:"sellerId"`
		StandPrice              int     `json:"standPrice"`
		CashFlag                int     `json:"cashFlag"`
		Price                   float64 `json:"price"`
		InitDiscount            int     `json:"initDiscount"`
		ServiceCountDown        int     `json:"serviceCountDown"`
		Ggb49I                  int     `json:"ggb49i"`
		CreateDate              int64   `json:"createDate,omitempty"`
		IsOffTreatment          int     `json:"isOffTreatment"`
		RetirementFlag          int     `json:"retirementFlag"`
		Quantity                int     `json:"quantity"`
		FinalPriceD             float64 `json:"finalPriceD"`
		Index                   string  `json:"index"`
		ChangedServiceCountDown int     `json:"changedServiceCountDown"`
		FreeFlag                int     `json:"freeFlag"`
		UpdateStatus            int     `json:"updateStatus"`
		VisitInter              int     `json:"visitInter,omitempty"`
		ProjectName             string  `json:"projectName"`
		PayStatus               int     `json:"payStatus"`
		ProjectId               string  `json:"projectId"`
		Proportion              string  `json:"proportion,omitempty"`
		SortIndex               int     `json:"sortIndex,omitempty"`
		IsAuto                  int     `json:"isAuto,omitempty"`
		PayCodeName             string  `json:"payCodeName,omitempty"`
		PayCode                 string  `json:"payCode,omitempty"`
		Gcf23I                  int     `json:"gcf23i,omitempty"`
		IsAdjustPrice           int     `json:"isAdjustPrice,omitempty"`
		MemberPrice             float64 `json:"memberPrice,omitempty"`
		IsAdjustDiscount        int     `json:"isAdjustDiscount,omitempty"`
	} `json:"serviceInfos"`
	UserId              string        `json:"userId"`
	RoomSvcAndProds     []interface{} `json:"roomSvcAndProds"`
	SettStatus          int           `json:"settStatus"`
	UserRemark          string        `json:"userRemark"`
	BType               int           `json:"bType"`
	BillingDate         int64         `json:"billingDate"`
	CompId              string        `json:"compId"`
	RemainFavourableMap struct {
	} `json:"remainFavourableMap"`
	UserType   int    `json:"userType"`
	ReqSource  string `json:"reqSource"`
	PayStatus  int    `json:"payStatus"`
	ReqPageKey string `json:"reqPageKey"`
	Sellers    []struct {
		BillingId   string `json:"billingId"`
		SellerId    string `json:"sellerId"`
		FirstStatus int    `json:"firstStatus"`
		SellerName  string `json:"sellerName"`
		Avatar      string `json:"avatar"`
	} `json:"sellers"`
	SettleRemark  string `json:"settleRemark"`
	Status        int    `json:"status"`
	CountDown     int    `json:"countDown,omitempty"`
	EndDate       int64  `json:"endDate,omitempty"`
	SettleStatus  int    `json:"settleStatus,omitempty"`
	ConsTime      string `json:"consTime,omitempty"`
	CardNo        string `json:"cardNo"`
	LinkBillingId string `json:"linkBillingId,omitempty"`
	ConsDate      string `json:"consDate,omitempty"`
	Card          struct {
		CardImage    string  `json:"cardImage,omitempty"`
		Balance      float64 `json:"balance"`
		CompId       string  `json:"compId"`
		UserMobile   string  `json:"userMobile"`
		CardTypeName string  `json:"cardTypeName"`
		CardType     string  `json:"cardType"`
		CustId       string  `json:"custId"`
		CardRemark   string  `json:"cardRemark"`
		Accounts     []struct {
			Balance   float64 `json:"balance"`
			AutoLevel int     `json:"autoLevel"`
			Num       float64 `json:"num"`
			ToDate    string  `json:"toDate,omitempty"`
			Name      string  `json:"name"`
			Id        string  `json:"id"`
			PayCode   string  `json:"payCode"`
			Arrears   float64 `json:"arrears,omitempty"`
		} `json:"accounts"`
		UserName string `json:"userName"`
		CardNo   string `json:"cardNo"`
	} `json:"card,omitempty"`
	StoreOrderIds             string        `json:"storeOrderIds,omitempty"`
	StoreOrderId              string        `json:"storeOrderId,omitempty"`
	EmpId                     string        `json:"empId,omitempty"`
	OnlineFlag                int           `json:"onlineFlag,omitempty"`
	BillingQualificationsList []interface{} `json:"billingQualificationsList,omitempty"`
	UseDepositList            []interface{} `json:"useDepositList,omitempty"`
}

// -------------------------------------------------------------------------------------------

type Token struct {
	AccessToken string `json:"access_token"`
	ShopID      string `json:"shop_id"`
	StartTime   int64  `json:"start_time"`
	Error       string `json:"error"`
}
