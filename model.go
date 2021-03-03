package main

type fieldDescription struct {
	ContentType string `yaml:"ContentType"`
	MaxLen      int    `yaml:"MaxLen"`
	MinLen      int    `yaml:"MinLen"`
	LenType     string `yaml:"LenType"`
	Label       string `yaml:"Label"`
}

type PPOBInquiryRequest struct {
	TransactionID string `json:"transaction_id"`
	PartnerID     string `json:"partner_id"`
	ProductCode   string `json:"product_code"`
	CustomerNo    string `json:"customer_no"`
	Periode       string `json:"periode"`
	MerchantCode  string `json:"merchant_code"`
	RequestTime   string `json:"request_time"`
}

type PPOBInquiryResponse struct {
	Rc           string `json:"rc"`
	Msg          string `json:"msg"`
	Produk       string `json:"produk"`
	Nopel        string `json:"nopel"`
	Nama         string `json:"nama"`
	Tagihan      int    `json:"tagihan"`
	Admin        int    `json:"admin"`
	TotalTagihan int    `json:"total_tagihan"`
	Reffid       string `json:"reffid"`
	Data         string `json:"data"`
	Restime      string `json:"restime"`
}

type PPOBPaymentRequest struct {
	TransactionID string `json:"transaction_id"`
	PartnerID     string `json:"partner_id"`
	ProductCode   string `json:"product_code"`
	CustomerNo    string `json:"customer_no"`
	MerchantCode  string `json:"merchant_code"`
	ReffID        string `json:"reff_id"`
	Amount        int    `json:"amount"`
	RequestTime   string `json:"request_time"`
}

type PPOBPaymentResponse struct {
	Rc           string   `json:"rc"`
	Msg          string   `json:"msg"`
	Produk       string   `json:"produk"`
	Nopel        string   `json:"nopel"`
	Nama         string   `json:"nama"`
	Tagihan      int      `json:"tagihan"`
	Admin        int      `json:"admin"`
	TotalTagihan int      `json:"total_tagihan"`
	Reffid       string   `json:"reffid"`
	TglLunas     string   `json:"tgl_lunas"`
	Struk        []string `json:"struk"`
	ReffNo       string   `json:"Reff_no"`
	Restime      string   `json:"restime"`
}

type PPOBStatusRequest struct {
	TransactionID string `json:"transaction_id"`
	PartnerID     string `json:"partner_id"`
	ProductCode   string `json:"product_code"`
	CustomerNo    string `json:"customer_no"`
	MerchantCode  string `json:"merchant_code"`
	ReffID        string `json:"reff_id"`
	Amount        int    `json:"amount"`
	RequestTime   string `json:"request_time"`
}

type PPOBStatusResponse struct {
	Rc           string   `json:"rc"`
	Msg          string   `json:"msg"`
	Produk       string   `json:"produk"`
	Nopel        string   `json:"nopel"`
	Nama         string   `json:"nama"`
	Tagihan      int      `json:"tagihan"`
	Admin        int      `json:"admin"`
	TotalTagihan int      `json:"total_tagihan"`
	Reffid       string   `json:"reffid"`
	TglLunas     string   `json:"tgl_lunas"`
	Struk        []string `json:"struk"`
	ReffNo       string   `json:"Reff_no"`
	Status       string   `json:"status"`
}

type TopupBuyRequest struct {
	TransactionID string `json:"transaction_id"`
	PartnerID     string `json:"partner_id"`
	ProductCode   string `json:"product_code"`
	CustomerNo    string `json:"customer_no"`
	MerchantCode  string `json:"merchant_code"`
	RequestTime   string `json:"request_time"`
}

type TopupBuyResponse struct {
	Rc      string `json:"rc"`
	Msg     string `json:"msg"`
	Restime string `json:"restime"`
	SN      string `json:"sn"`
	Price   string `json:"price"`
}

type TopupCheckRequest struct {
	TransactionID string `json:"transaction_id"`
	PartnerID     string `json:"partner_id"`
	ProductCode   string `json:"product_code"`
	CustomerNo    string `json:"customer_no"`
	MerchantCode  string `json:"merchant_code"`
	RequestTime   string `json:"request_time"`
}

type TopupCheckResponse struct {
	Rc      string `json:"rc"`
	Msg     string `json:"msg"`
	Restime string `json:"restime"`
	SN      string `json:"sn"`
	Price   string `json:"price"`
}

type UnsuccessfulChipsakti struct {
	Rc      string `json:"rc"`
	Msg     string `json:"msg"`
	Restime string `json:"restime"`
}
