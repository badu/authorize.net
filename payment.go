package authorize_net

import (
	"time"
)

// -- request --
type CreditCard struct {
	CardNumber         string `json:"cardNumber"`
	ExpirationDate     string `json:"expirationDate"`
	CardCode           string `json:"cardCode,omitempty"`
	IsPaymentToken     bool   `json:"isPaymentToken,omitempty"`
	Cryptogram         string `json:"cryptogram,omitempty"`
	TokenRequestorName string `json:"tokenRequestorName,omitempty"`
	TokenRequestorId   string `json:"tokenRequestorId,omitempty"`
	TokenRequestorEci  string `json:"tokenRequestorEci,omitempty"`
}

type BankAccount struct {
	AccountType   BankAccountTypeEnum `json:"accountType,omitempty"`
	RoutingNumber string              `json:"routingNumber"`
	AccountNumber string              `json:"accountNumber"`
	NameOnAccount string              `json:"nameOnAccount"`
	EcheckType    EcheckTypeEnum      `json:"echeckType,omitempty"`
	BankName      string              `json:"bankName,omitempty"`
	CheckNumber   string              `json:"checkNumber,omitempty"`
}

type CreditCardTrack struct {
	Track1 string `json:"track1"`
	Track2 string `json:"track2"`
}

type EncryptedTrackData struct {
	FormOfPayment KeyBlock `json:"formOfPayment"`
}

type PayPal struct {
	SuccessUrl         string `json:"successUrl,omitempty"`
	CancelUrl          string `json:"cancelUrl,omitempty"`
	PaypalLc           string `json:"paypalLc,omitempty"`
	PaypalHdrImg       string `json:"paypalHdrImg,omitempty"`
	PaypalPayFlowColor string `json:"paypalPayflowcolor,omitempty"`
	PayerID            string `json:"payerID,omitempty"`
}

type OpaqueData struct {
	DataDescriptor      string    `json:"dataDescriptor"`
	DataValue           string    `json:"dataValue"`
	DataKey             string    `json:"dataKey,omitempty"`
	ExpirationTimeStamp time.Time `json:"expirationTimeStamp,omitempty"`
}

type PaymentEmv struct {
	EmvData       interface{} `json:"emvData"`
	EmvDescriptor interface{} `json:"emvDescriptor"`
	EmvVersion    interface{} `json:"emvVersion"`
}

type Payment struct {
	CreditCard         *CreditCard         `json:"creditCard,omitempty"`
	BankAccount        *BankAccount        `json:"bankAccount,omitempty"`
	TrackData          *CreditCardTrack    `json:"trackData,omitempty"`
	EncryptedTrackData *EncryptedTrackData `json:"encryptedTrackData,omitempty"`
	PayPal             *PayPal             `json:"payPal,omitempty"`
	OpaqueData         *OpaqueData         `json:"opaqueData,omitempty"`
	Emv                *PaymentEmv         `json:"emv,omitempty"`
	DataSource         string              `json:"dataSource,omitempty"`
}

type PaymentProfile struct {
	PaymentProfileId string `json:"paymentProfileId"`
	CardCode         string `json:"cardCode,omitempty"`
}

type CustomerProfilePayment struct {
	CreateProfile     bool            `json:"createProfile,omitempty"`
	CustomerProfileId string          `json:"customerProfileId,omitempty"`
	PaymentProfile    *PaymentProfile `json:"paymentProfile,omitempty"`
	ShippingProfileId string          `json:"shippingProfileId,omitempty"`
}

type Solution struct {
	Id         string `json:"id"`
	Name       string `json:"name,omitempty"`
	VendorName string `json:"vendorName,omitempty"`
}

type Order struct {
	InvoiceNumber                  string    `json:"invoiceNumber,omitempty"`
	Description                    string    `json:"description,omitempty"`
	DiscountAmount                 float64   `json:"discountAmount,omitempty"`
	TaxIsAfterDiscount             bool      `json:"taxIsAfterDiscount,omitempty"`
	TotalTaxTypeCode               string    `json:"totalTaxTypeCode,omitempty"`
	PurchaserVATRegistrationNumber string    `json:"purchaserVATRegistrationNumber,omitempty"`
	MerchantVATRegistrationNumber  string    `json:"merchantVATRegistrationNumber,omitempty"`
	VatInvoiceReferenceNumber      string    `json:"vatInvoiceReferenceNumber,omitempty"`
	PurchaserCode                  string    `json:"purchaserCode,omitempty"`
	SummaryCommodityCode           string    `json:"summaryCommodityCode,omitempty"`
	PurchaseOrderDateUTC           time.Time `json:"purchaseOrderDateUTC,omitempty"` //was Date
	SupplierOrderReference         string    `json:"supplierOrderReference,omitempty"`
	AuthorizedContactName          string    `json:"authorizedContactName,omitempty"`
	CardAcceptorRefNumber          string    `json:"cardAcceptorRefNumber,omitempty"`
	AmexDataTAA1                   string    `json:"amexDataTAA1,omitempty"`
	AmexDataTAA2                   string    `json:"amexDataTAA2,omitempty"`
	AmexDataTAA3                   string    `json:"amexDataTAA3,omitempty"`
	AmexDataTAA4                   string    `json:"amexDataTAA4,omitempty"`
}

type Item struct {
	ItemId                  string  `json:"itemId"`
	Name                    string  `json:"name"`
	Description             string  `json:"description,omitempty"`
	Quantity                float64 `json:"quantity"`
	UnitPrice               float64 `json:"unitPrice"`
	Taxable                 bool    `json:"taxable,omitempty"`
	UnitOfMeasure           string  `json:"unitOfMeasure,omitempty"`
	TypeOfSupply            string  `json:"typeOfSupply,omitempty"`
	TaxRate                 float64 `json:"taxRate,omitempty"`
	TaxAmount               float64 `json:"taxAmount,omitempty"`
	NationalTax             float64 `json:"nationalTax,omitempty"`
	LocalTax                float64 `json:"localTax,omitempty"`
	VatRate                 float64 `json:"vatRate,omitempty"`
	AlternateTaxId          string  `json:"alternateTaxId,omitempty"`
	AlternateTaxType        string  `json:"alternateTaxType,omitempty"`
	AlternateTaxTypeApplied string  `json:"alternateTaxTypeApplied,omitempty"`
	AlternateTaxRate        float64 `json:"alternateTaxRate,omitempty"`
	AlternateTaxAmount      float64 `json:"alternateTaxAmount,omitempty"`
	TotalAmount             float64 `json:"totalAmount,omitempty"`
	CommodityCode           string  `json:"commodityCode,omitempty"`
	ProductCode             string  `json:"productCode,omitempty"`
	ProductSKU              string  `json:"productSKU,omitempty"`
	DiscountRate            float64 `json:"discountRate,omitempty"`
	DiscountAmount          float64 `json:"discountAmount,omitempty"`
	TaxIncludedInTotal      bool    `json:"taxIncludedInTotal,omitempty"`
	TaxIsAfterDiscount      bool    `json:"taxIsAfterDiscount,omitempty"`
}

type ExtendedAmount struct {
	Amount      float64 `json:"amount"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
}

type DriversLicense struct {
	Number      string `json:"number"`
	State       string `json:"state"`
	DateOfBirth string `json:"dateOfBirth"`
}

type CustomerData struct {
	Type           string          `json:"type,omitempty"`
	Id             string          `json:"id,omitempty"`
	Email          string          `json:"email,omitempty"`
	DriversLicense *DriversLicense `json:"driversLicense,omitempty"`
	TaxId          string          `json:"taxId,omitempty"`
}

type NameAndAddress struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Company   string `json:"company,omitempty"`
	Address   string `json:"address,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
}

type CCAuthentication struct {
	AuthenticationIndicator       string `json:"authenticationIndicator"`
	CardholderAuthenticationValue string `json:"cardholderAuthenticationValue"`
}

type TransRetailInfo struct {
	MarketType        string `json:"marketType,omitempty"`
	DeviceType        string `json:"deviceType,omitempty"`
	CustomerSignature string `json:"customerSignature,omitempty"`
	TerminalNumber    string `json:"terminalNumber,omitempty"`
}

type Setting struct {
	SettingName  string `json:"settingName,omitempty"`
	SettingValue string `json:"settingValue,omitempty"`
}

type UserFields struct {
	UserField []UserField `json:"userField,omitempty"` //min=0 max=20
}

type SubMerchant struct {
	Identifier                 string `json:"identifier"`
	DoingBusinessAs            string `json:"doingBusinessAs,omitempty"`
	PaymentServiceProviderName string `json:"paymentServiceProviderName,omitempty"`
	PaymentServiceFacilitator  string `json:"paymentServiceFacilitator,omitempty"`
	StreetAddress              string `json:"streetAddress,omitempty"`
	Phone                      string `json:"phone,omitempty"`
	Email                      string `json:"email,omitempty"`
	PostalCode                 string `json:"postalCode,omitempty"`
	City                       string `json:"city,omitempty"`
	RegionCode                 string `json:"regionCode,omitempty"`
	CountryCode                string `json:"countryCode,omitempty"`
}

type ProcessingOptions struct {
	IsFirstRecurringPayment bool `json:"isFirstRecurringPayment,omitempty"`
	IsFirstSubsequentAuth   bool `json:"isFirstSubsequentAuth,omitempty"`
	IsSubsequentAuth        bool `json:"isSubsequentAuth,omitempty"`
	IsStoredCredentials     bool `json:"isStoredCredentials,omitempty"`
}

type SubsequentAuthInformation struct {
	OriginalNetworkTransId string `json:"originalNetworkTransId,omitempty"`
	Reason                 string `json:"reason,omitempty"`
}

type OtherTax struct {
	AlternateTaxId     string  `json:"alternateTaxId,omitempty"`
	NationalTaxAmount  float64 `json:"nationalTaxAmount,omitempty"`
	LocalTaxAmount     float64 `json:"localTaxAmount,omitempty"`
	AlternateTaxAmount float64 `json:"alternateTaxAmount,omitempty"`
	VatTaxRate         float64 `json:"vatTaxRate,omitempty"`
	VatTaxAmount       float64 `json:"vatTaxAmount,omitempty"`
}

type Transaction struct {
	TransactionType           TransactionTypeEnum        `json:"transactionType"`
	Amount                    float64                    `json:"amount,omitempty"`
	CurrencyCode              string                     `json:"currencyCode,omitempty"`
	Payment                   *Payment                   `json:"payment,omitempty"`
	Profile                   *CustomerProfilePayment    `json:"profile,omitempty"`
	Solution                  *Solution                  `json:"solution,omitempty"`
	CallId                    string                     `json:"callId,omitempty"`
	TerminalNumber            string                     `json:"terminalNumber,omitempty"`
	AuthCode                  string                     `json:"authCode,omitempty"`
	RefTransId                string                     `json:"refTransId,omitempty"`
	SplitTenderId             string                     `json:"splitTenderId,omitempty"`
	Order                     *Order                     `json:"order,omitempty"`
	LineItems                 []Item                     `json:"lineItems,omitempty"` //min=0
	Tax                       *ExtendedAmount            `json:"tax,omitempty"`
	Duty                      *ExtendedAmount            `json:"duty,omitempty"`
	Shipping                  *ExtendedAmount            `json:"shipping,omitempty"`
	Surcharge                 *ExtendedAmount            `json:"surcharge,omitempty"`
	Tip                       *ExtendedAmount            `json:"tip,omitempty"`
	TaxExempt                 bool                       `json:"taxExempt,omitempty"`
	PoNumber                  string                     `json:"poNumber,omitempty"`
	Customer                  *CustomerData              `json:"customer,omitempty"`
	BillTo                    *CustomerAddress           `json:"billTo,omitempty"`
	ShipTo                    *NameAndAddress            `json:"shipTo,omitempty"`
	CustomerIP                string                     `json:"customerIP,omitempty"`
	CardholderAuthentication  *CCAuthentication          `json:"cardholderAuthentication,omitempty"`
	Retail                    *TransRetailInfo           `json:"retail,omitempty"`
	EmployeeId                string                     `json:"employeeId,omitempty"`
	TransactionSettings       []Setting                  `json:"transactionSettings,omitempty"` //min=0
	UserFields                *UserFields                `json:"userFields,omitempty"`
	MerchantDescriptor        string                     `json:"merchantDescriptor,omitempty"`
	SubMerchant               *SubMerchant               `json:"subMerchant,omitempty"`
	ProcessingOptions         *ProcessingOptions         `json:"processingOptions,omitempty"`
	SubsequentAuthInformation *SubsequentAuthInformation `json:"subsequentAuthInformation,omitempty"`
	OtherTax                  *OtherTax                  `json:"otherTax,omitempty"`
	ShipFrom                  *NameAndAddress            `json:"shipFrom,omitempty"`
}

type TransactionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Transaction            Transaction            `json:"transactionRequest"`
}

type CreateTransactionRequest struct {
	CreateTransactionRequest TransactionRequest `json:"createTransactionRequest"`
}

// -- response --
type TransactionResponsePrePaidCard struct {
	RequestedAmount string `json:"requestedAmount,omitempty"`
	ApprovedAmount  string `json:"approvedAmount,omitempty"`
	BalanceOnCard   string `json:"balanceOnCard,omitempty"`
}

type SplitTenderPayment struct {
	TransId            string          `json:"transId,omitempty"`
	ResponseCode       string          `json:"responseCode,omitempty"`
	ResponseToCustomer string          `json:"responseToCustomer,omitempty"`
	AuthCode           string          `json:"authCode,omitempty"`
	AccountNumber      string          `json:"accountNumber,omitempty"`
	AccountType        AccountTypeEnum `json:"accountType,omitempty"`
	RequestedAmount    string          `json:"requestedAmount,omitempty"`
	ApprovedAmount     string          `json:"approvedAmount,omitempty"`
	BalanceOnCard      string          `json:"balanceOnCard,omitempty"`
}

type UserField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type SecureAcceptance struct {
	SecureAcceptanceUrl string `json:"secureAcceptanceUrl,omitempty"`
	PayerID             string `json:"payerID,omitempty"`
	PayerEmail          string `json:"payerEmail,omitempty"`
}

type EmvTag struct {
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Formatted string `json:"formatted,omitempty"`
}

type EmvResponse struct {
	TlvData string   `json:"tlvData,omitempty"`
	Tags    []EmvTag `json:"tags,omitempty"`
}

type TransactionResponse struct {
	ResponseCode        string                          `json:"responseCode,omitempty"`
	RawResponseCode     string                          `json:"rawResponseCode,omitempty"`
	AuthCode            string                          `json:"authCode,omitempty"`
	AvsResultCode       string                          `json:"avsResultCode,omitempty"`
	CvvResultCode       string                          `json:"cvvResultCode,omitempty"`
	CavvResultCode      string                          `json:"cavvResultCode,omitempty"`
	TransId             string                          `json:"transId,omitempty"`
	RefTransID          string                          `json:"refTransID,omitempty"`
	TransHash           string                          `json:"transHash,omitempty"`
	TestRequest         string                          `json:"testRequest,omitempty"`
	AccountNumber       string                          `json:"accountNumber,omitempty"`
	EntryMode           string                          `json:"entryMode,omitempty"`
	AccountType         AccountTypeEnum                 `json:"accountType,omitempty"`
	SplitTenderId       string                          `json:"splitTenderId,omitempty"`
	PrePaidCard         *TransactionResponsePrePaidCard `json:"prePaidCard,omitempty"`
	Messages            []ErrMessage                    `json:"messages,omitempty"`
	Errors              []Error                         `json:"errors,omitempty"` //min=0
	SplitTenderPayments []SplitTenderPayment            `json:"splitTenderPayments,omitempty"`
	UserFields          []UserField                     `json:"userFields,omitempty"`
	ShipTo              *NameAndAddress                 `json:"shipTo,omitempty"`
	SecureAcceptance    *SecureAcceptance               `json:"secureAcceptance,omitempty"`
	EmvResponse         *EmvResponse                    `json:"emvResponse,omitempty"`
	TransHashSha2       string                          `json:"transHashSha2,omitempty"`
	Profile             *CustomerProfileId              `json:"profile,omitempty"`
	NetworkTransId      string                          `json:"networkTransId,omitempty"`
}

type CreateProfileResponse struct {
	Messages                      Messages `json:"messages"`
	CustomerProfileId             string   `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileIdList  []string `json:"customerPaymentProfileIdList,omitempty"`
	CustomerShippingAddressIdList []string `json:"customerShippingAddressIdList,omitempty"`
}

type CreateTransactionResponse struct {
	ANetApiResponse
	TransactionResponse TransactionResponse   `json:"transactionResponse"`
	ProfileResponse     CreateProfileResponse `json:"profileResponse,omitempty"`
}
