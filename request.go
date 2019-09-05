package authorize_net

type FingerPrintType struct {
	HashValue    string `json:"hashValue"`
	Timestamp    string `json:"timestamp"`
	Sequence     string `json:"sequence,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
	Amount       string `json:"amount,omitempty"`
}

type ImpersonationAuthenticationType struct {
	PartnerLoginId        string `json:"partnerLoginId"`
	PartnerTransactionKey string `json:"partnerTransactionKey"`
}

type MerchantAuthentication struct {
	Name                        string                           `json:"name"`           // Merchant’s unique API Login ID.
	TransactionKey              string                           `json:"transactionKey"` // Merchant’s unique Transaction Key.
	Password                    string                           `json:"password,omitempty"`
	SessionToken                string                           `json:"sessionToken,omitempty"`
	ImpersonationAuthentication *ImpersonationAuthenticationType `json:"impersonationAuthentication,omitempty"`
	FingerPrint                 *FingerPrintType                 `json:"fingerPrint,omitempty"`
	ClientKey                   string                           `json:"clientKey,omitempty"`
	AccessToken                 string                           `json:"accessToken,omitempty"`
	MobileDeviceId              string                           `json:"mobileDeviceId,omitempty"`
}

type ANetApiRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"` // Contains merchant authentication information.
	ClientId               string                 `json:"clientId,omitempty"`
	RefId                  string                 `json:"refId,omitempty"` // Merchant-assigned reference ID for the request. If included in the request, this value is included in the response. This feature might be especially useful for multi-threaded applications.
}

type Paging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Sorting struct {
	OrderBy         string `json:"orderBy"`
	OrderDescending bool   `json:"orderDescending"`
}
