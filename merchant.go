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

type Processor struct {
	Name      string   `json:"name"`
	Id        int      `json:"id"`
	CardTypes []string `json:"cardTypes,omitempty"`
}

type ContactDetail struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type GetMerchantDetailsResponse struct {
	ANetApiResponse
	IsTestMode          bool             `json:"isTestMode,omitempty"`
	Processors          []Processor      `json:"processors,omitempty"`
	MerchantName        string           `json:"merchantName"`
	GatewayId           string           `json:"gatewayId"`
	MarketTypes         []string         `json:"marketTypes,omitempty"`
	ProductCodes        []string         `json:"productCodes,omitempty"`
	PaymentMethods      []string         `json:"paymentMethods,omitempty"`
	Currencies          []string         `json:"currencies,omitempty"`
	PublicClientKey     string           `json:"publicClientKey,omitempty"`
	BusinessInformation *CustomerAddress `json:"businessInformation,omitempty"`
	MerchantTimeZone    string           `json:"merchantTimeZone,omitempty"`
	ContactDetails      []ContactDetail  `json:"contactDetails,omitempty"`
}
