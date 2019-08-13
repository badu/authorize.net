package authorize_net

// -- request --
type UpdateMerchantDetailsRequest struct {
	Payload UpdateMerchantDetailsPayload `json:"updateMerchantDetailsRequest"`
}

type UpdateMerchantDetailsPayload struct {
	ANetApiRequest
	IsTestMode bool `json:"isTestMode"`
}

type GetMerchantDetailsRequest struct {
	ANetApiRequest `json:"getMerchantDetailsRequest"`
}

// -- response --

type UpdateMerchantDetailsResponse struct {
	ANetApiResponse
}

type GetMerchantDetailsResponse struct {
	ANetApiResponse
	IsTestMode          bool             `json:"isTestMode,omitempty"`
	Processors          []Processor      `json:"processors"`
	MerchantName        string           `json:"merchantName"`
	GatewayId           string           `json:"gatewayId"`
	MarketTypes         []string         `json:"marketTypes"`
	ProductCodes        []string         `json:"productCodes"`
	PaymentMethods      []string         `json:"paymentMethods"`
	Currencies          []string         `json:"currencies"`
	PublicClientKey     string           `json:"publicClientKey,omitempty"`
	BusinessInformation *CustomerAddress `json:"businessInformation,omitempty"`
	MerchantTimeZone    string           `json:"merchantTimeZone,omitempty"`
	ContactDetails      []ContactDetail  `json:"contactDetails,omitempty"`
}
