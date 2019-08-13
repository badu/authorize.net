package authorize_net

// -- request --
type CustomerAddress struct {
	NameAndAddress
	PhoneNumber string `json:"phoneNumber,omitempty"`
	FaxNumber   string `json:"faxNumber,omitempty"`
	Email       string `json:"email,omitempty"`
}

type CustomerPaymentProfileBase struct {
	CustomerType CustomerTypeEnum `json:"customerType,omitempty"`
	BillTo       *CustomerAddress `json:"billTo,omitempty"`
}

type CustomerPaymentProfile struct {
	CustomerPaymentProfileBase
	Payment               *Payment        `json:"payment,omitempty"`
	DriversLicense        *DriversLicense `json:"driversLicense,omitempty"`
	TaxId                 string          `json:"taxId,omitempty"`
	DefaultPaymentProfile bool            `json:"defaultPaymentProfile,omitempty"`
}

type CustomerProfileBase struct {
	MerchantCustomerId string `json:"merchantCustomerId,omitempty"`
	Description        string `json:"description,omitempty"`
	Email              string `json:"email,omitempty"`
}

type CustomerProfile struct {
	CustomerProfileBase
	PaymentProfiles []CustomerPaymentProfile `json:"paymentProfiles,omitempty"` //min=0
	ShipToList      []CustomerAddress        `json:"shipToList,omitempty"`      //min=0
	ProfileType     CustomerProfileTypeEnum  `json:"profileType,omitempty"`
}

type CreateCustomerProfileRequest struct {
	Payload CreateCustomerProfilePayload `json:"createCustomerProfileRequest"`
}

type CreateCustomerProfilePayload struct {
	ANetApiRequest
	Profile        *CustomerProfile   `json:"profile"`
	ValidationMode ValidationModeEnum `json:"validationMode,omitempty"`
}

type GetCustomerProfileRequest struct {
	Payload GetCustomerProfilePayload `json:"getCustomerProfileRequest"`
}

type GetCustomerProfilePayload struct {
	ANetApiRequest
	CustomerProfileId    string `json:"customerProfileId,omitempty"`
	MerchantCustomerId   string `json:"merchantCustomerId,omitempty"`
	Email                string `json:"email,omitempty"`
	UnmaskExpirationDate bool   `json:"unmaskExpirationDate,omitempty"`
	IncludeIssuerInfo    bool   `json:"includeIssuerInfo,omitempty"`
}

type GetCustomerPaymentProfileRequest struct {
	Payload GetCustomerPaymentProfilePayload `json:"getCustomerPaymentProfileRequest"`
}

type GetCustomerPaymentProfilePayload struct {
	ANetApiRequest
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId,omitempty"`
	UnmaskExpirationDate     bool   `json:"unmaskExpirationDate,omitempty"`
	IncludeIssuerInfo        bool   `json:"includeIssuerInfo,omitempty"`
}

type GetCustomerPaymentProfileListRequest struct {
	Payload GetCustomerPaymentProfileListPayload `json:"getCustomerPaymentProfileListRequest"`
}

type CustomerPaymentProfileSorting struct {
	OrderBy         string `json:"orderBy"`
	OrderDescending bool   `json:"orderDescending"`
}

type Paging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetCustomerPaymentProfileListPayload struct {
	ANetApiRequest
	SearchType string                         `json:"searchType"`
	Month      string                         `json:"month"` //String, 7 characters. Use XML gYearMonth (YYYY-MM) formatting.
	Sorting    *CustomerPaymentProfileSorting `json:"sorting,omitempty"`
	Paging     *Paging                        `json:"paging,omitempty"`
}

type GetCustomerProfileIdsRequest struct {
	ANetApiRequest `json:"getCustomerProfileIdsRequest"`
}

type DeleteCustomerProfileRequest struct {
	Payload DeleteCustomerProfilePayload `json:"deleteCustomerProfileRequest"`
}
type DeleteCustomerProfilePayload struct {
	ANetApiRequest
	CustomerProfileId string `json:"customerProfileId"`
}

type UpdateCustomerProfileRequest struct {
	Payload UpdateCustomerProfilePayload `json:"updateCustomerProfileRequest"`
}
type UpdateCustomerProfilePayload struct {
	ANetApiRequest
	Profile *CustomerProfileInfoEx `json:"profile"`
}

type GetCustomerShippingAddressRequest struct {
	Payload GetCustomerShippingAddressPayload `json:"getCustomerShippingAddressRequest"`
}
type GetCustomerShippingAddressPayload struct {
	ANetApiRequest
	CustomerProfileId string `json:"customerProfileId"`
	CustomerAddressId string `json:"customerAddressId,omitempty"`
}

type GetTransactionListForCustomerRequest struct {
	Payload GetTransactionListForCustomerPayload `json:"getTransactionListForCustomerRequest"`
}
type GetTransactionListForCustomerPayload struct {
	ANetApiRequest
	CustomerProfileId        string                  `json:"customerProfileId"`
	CustomerPaymentProfileId string                  `json:"customerPaymentProfileId,omitempty"`
	Sorting                  *TransactionListSorting `json:"sorting,omitempty"`
	Paging                   Paging                  `json:"paging,omitempty"`
}

type DeleteCustomerPaymentProfileRequest struct {
	Payload DeleteCustomerPaymentProfilePayload `json:"deleteCustomerPaymentProfileRequest"`
}
type DeleteCustomerPaymentProfilePayload struct {
	ANetApiRequest
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}

type CreateCustomerShippingAddressRequest struct {
	Payload CreateCustomerShippingAddressPayload `json:"createCustomerShippingAddressRequest"`
}
type CreateCustomerShippingAddressPayload struct {
	ANetApiRequest
	CustomerProfileId      string           `json:"customerProfileId"`
	Address                *CustomerAddress `json:"address"`
	DefaultShippingAddress bool             `json:"defaultShippingAddress,omitempty"`
}

type GetCustomerPaymentProfileNonceRequest struct {
	Payload GetCustomerPaymentProfileNoncePayload `json:"getCustomerPaymentProfileNonceRequest"`
}
type GetCustomerPaymentProfileNoncePayload struct {
	ANetApiRequest
	ConnectedAccessToken     string `json:"connectedAccessToken"`
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}

type UpdateCustomerPaymentProfileRequest struct {
	Payload UpdateCustomerPaymentProfilePayload `json:"updateCustomerPaymentProfileRequest"`
}
type UpdateCustomerPaymentProfilePayload struct {
	ANetApiRequest
	CustomerProfileId string                    `json:"customerProfileId"`
	PaymentProfile    *CustomerPaymentProfileEx `json:"paymentProfile"`
	ValidationMode    string                    `json:"validationMode,omitempty"`
}

type ValidateCustomerPaymentProfileRequest struct {
	Payload ValidateCustomerPaymentProfilePayload `json:"validateCustomerPaymentProfileRequest"`
}
type ValidateCustomerPaymentProfilePayload struct {
	ANetApiRequest
	CustomerProfileId         string `json:"customerProfileId"`
	CustomerPaymentProfileId  string `json:"customerPaymentProfileId"`
	CustomerShippingAddressId string `json:"customerShippingAddressId,omitempty"`
	CardCode                  string `json:"cardCode,omitempty"`
	ValidationMode            string `json:"validationMode"`
}

type CreateCustomerProfileTransactionRequest struct {
	Payload CreateCustomerProfileTransactionPayload `json:"createCustomerProfileTransactionRequest"`
}
type CreateCustomerProfileTransactionPayload struct {
	ANetApiRequest
	Transaction  *ProfileTransaction `json:"transaction"`
	ExtraOptions string              `json:"extraOptions,omitempty"`
}

type DeleteCustomerShippingAddressRequest struct {
	Payload DeleteCustomerShippingAddressPayload `json:"deleteCustomerShippingAddressRequest"`
}
type DeleteCustomerShippingAddressPayload struct {
	ANetApiRequest
	CustomerProfileId string `json:"customerProfileId"`
	CustomerAddressId string `json:"customerAddressId"`
}

type CreateCustomerPaymentProfileRequest struct {
	Payload CreateCustomerPaymentProfilePayload `json:"createCustomerPaymentProfileRequest"`
}
type CreateCustomerPaymentProfilePayload struct {
	ANetApiRequest
	CustomerProfileId string                  `json:"customerProfileId"`
	PaymentProfile    *CustomerPaymentProfile `json:"paymentProfile"`
	ValidationMode    string                  `json:"validationMode,omitempty"`
}

// -- response --
type CreateCustomerProfileResponse struct {
	ANetApiResponse
	CustomerProfileId             string   `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileIdList  []string `json:"customerPaymentProfileIdList"`
	CustomerShippingAddressIdList []string `json:"customerShippingAddressIdList"`
	ValidationDirectResponseList  []string `json:"validationDirectResponseList"`
}

type BankAccountMasked struct {
	AccountType   BankAccountTypeEnum `json:"accountType,omitempty"`
	RoutingNumber string              `json:"routingNumber"`
	AccountNumber string              `json:"accountNumber"`
	NameOnAccount string              `json:"nameOnAccount"`
	EcheckType    string              `json:"echeckType,omitempty"`
	BankName      string              `json:"bankName,omitempty"`
}

type PaymentMasked struct {
	CreditCard       *CreditCardMasked  `json:"creditCard"`
	BankAccount      *BankAccountMasked `json:"bankAccount"`
	TokenInformation *TokenMasked       `json:"tokenInformation"`
}

type DriversLicenseMasked struct {
	Number      string `json:"number"`
	State       string `json:"state"`
	DateOfBirth string `json:"dateOfBirth"`
}

type CustomerPaymentProfileMasked struct {
	CustomerPaymentProfileBase
	CustomerProfileId        string                `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileId string                `json:"customerPaymentProfileId"`
	DefaultPaymentProfile    bool                  `json:"defaultPaymentProfile,omitempty"`
	Payment                  *PaymentMasked        `json:"payment,omitempty"`
	DriversLicense           *DriversLicenseMasked `json:"driversLicense,omitempty"`
	TaxId                    string                `json:"taxId,omitempty"`
	SubscriptionIds          []string              `json:"subscriptionIds,omitempty"`
}

type CustomerAddressEx struct {
	CustomerAddress
	CustomerAddressId string `json:"customerAddressId,omitempty"`
}

type CustomerProfileMasked struct {
	CustomerProfileEx
	PaymentProfiles []CustomerPaymentProfileMasked `json:"paymentProfiles,omitempty"` //min=0
	ShipToList      []CustomerAddressEx            `json:"shipToList,omitempty"`      //min=0
	ProfileType     string                         `json:"profileType,omitempty"`
}

type GetCustomerProfileResponse struct {
	ANetApiResponse
	Profile         *CustomerProfileMasked `json:"profile,omitempty"`
	SubscriptionIds []string               `json:"subscriptionIds,omitempty"`
}

type GetCustomerPaymentProfileResponse struct {
	ANetApiResponse
	PaymentProfile *CustomerPaymentProfileMasked `json:"paymentProfile,omitempty"`
}

type CustomerPaymentProfileListItem struct {
	DefaultPaymentProfile    bool             `json:"defaultPaymentProfile,omitempty"`
	CustomerPaymentProfileId int              `json:"customerPaymentProfileId"`
	CustomerProfileId        int              `json:"customerProfileId"`
	BillTo                   *CustomerAddress `json:"billTo"`
	Payment                  *PaymentMasked   `json:"payment"`
}

type GetCustomerPaymentProfileListResponse struct {
	ANetApiResponse
	TotalNumInResultSet int                              `json:"totalNumInResultSet"`
	PaymentProfiles     []CustomerPaymentProfileListItem `json:"paymentProfiles,omitempty"`
}

type GetCustomerProfileIdsResponse struct {
	ANetApiResponse
	Ids []string `json:"ids"`
}

type CreateCustomerPaymentProfileResponse struct {
	ANetApiResponse
	CustomerProfileId        string `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId,omitempty"`
	ValidationDirectResponse string `json:"validationDirectResponse,omitempty"`
}

type DeleteCustomerProfileResponse struct {
	ANetApiResponse
}

type DeleteCustomerShippingAddressResponse struct {
	ANetApiResponse
}

type GetCustomerPaymentProfileNonceResponse struct {
	ANetApiResponse
	OpaqueData *OpaqueData `json:"opaqueData,omitempty"`
}

type CreateCustomerShippingAddressResponse struct {
	ANetApiResponse
	CustomerProfileId string `json:"customerProfileId,omitempty"`
	CustomerAddressId string `json:"customerAddressId,omitempty"`
}
