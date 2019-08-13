package authorize_net

import "time"

type KeyBlock struct {
	Value KeyValue `json:"value"`
}

type EmailSettings struct {
	Settings []Setting
	Version  int `json:"version,omitempty"`
}

type CreditCardMasked struct {
	CardNumber     string   `json:"cardNumber"`
	ExpirationDate string   `json:"expirationDate"`
	CardType       string   `json:"cardType,omitempty"`
	CardArt        *CardArt `json:"cardArt,omitempty"`
	IssuerNumber   string   `json:"issuerNumber,omitempty"`
	IsPaymentToken bool     `json:"isPaymentToken,omitempty"`
}

type GetSettledBatchListResponse struct {
	ANetApiResponse
	BatchList []BatchDetails `json:"batchList,omitempty"`
}

type CreateCustomerProfileFromTransactionRequest struct {
	Payload CreateCustomerProfileFromTransactionPayload `json:"createCustomerProfileFromTransactionRequest"`
}
type CreateCustomerProfileFromTransactionPayload struct {
	ANetApiRequest
	TransId                string               `json:"transId"`
	Customer               *CustomerProfileBase `json:"customer,omitempty"`
	CustomerProfileId      string               `json:"customerProfileId,omitempty"`
	DefaultPaymentProfile  bool                 `json:"defaultPaymentProfile,omitempty"`
	DefaultShippingAddress bool                 `json:"defaultShippingAddress,omitempty"`
	ProfileType            string               `json:"profileType,omitempty"`
}

type SubscriptionDetail struct {
	Id                        int       `json:"id"`
	Name                      string    `json:"name,omitempty"`
	Status                    string    `json:"status"`
	CreateTimeStampUTC        time.Time `json:"createTimeStampUTC"`
	FirstName                 string    `json:"firstName,omitempty"`
	LastName                  string    `json:"lastName,omitempty"`
	TotalOccurrences          int       `json:"totalOccurrences"`
	PastOccurrences           int       `json:"pastOccurrences"`
	PaymentMethod             string    `json:"paymentMethod"`
	AccountNumber             string    `json:"accountNumber,omitempty"`
	Invoice                   string    `json:"invoice,omitempty"`
	Amount                    float64   `json:"amount"`
	CurrencyCode              string    `json:"currencyCode,omitempty"`
	CustomerProfileId         int       `json:"customerProfileId"`
	CustomerPaymentProfileId  int       `json:"customerPaymentProfileId"`
	CustomerShippingProfileId int       `json:"customerShippingProfileId,omitempty"`
}

type UpdateSplitTenderGroupResponse struct {
	ANetApiResponse
}

type UpdateCustomerShippingAddressRequest struct {
	Payload UpdateCustomerShippingAddressPayload `json:"updateCustomerShippingAddressRequest"`
}
type UpdateCustomerShippingAddressPayload struct {
	ANetApiRequest
	CustomerProfileId      string             `json:"customerProfileId"`
	Address                *CustomerAddressEx `json:"address"`
	DefaultShippingAddress bool               `json:"defaultShippingAddress,omitempty"`
}

type DeleteCustomerPaymentProfileResponse struct {
	ANetApiResponse
}

type GetCustomerShippingAddressResponse struct {
	ANetApiResponse
	DefaultShippingAddress bool               `json:"defaultShippingAddress,omitempty"`
	Address                *CustomerAddressEx `json:"address,omitempty"`
	SubscriptionIds        []string           `json:"subscriptionIds,omitempty"`
}

type CardArt struct {
	CardBrand       string `json:"cardBrand,omitempty"`
	CardImageHeight string `json:"cardImageHeight,omitempty"`
	CardImageUrl    string `json:"cardImageUrl,omitempty"`
	CardImageWidth  string `json:"cardImageWidth,omitempty"`
	CardType        string `json:"cardType,omitempty"`
}

type SendCustomerTransactionReceiptRequest struct {
	Payload SendCustomerTransactionReceiptPayload `json:"sendCustomerTransactionReceiptRequest"`
}
type SendCustomerTransactionReceiptPayload struct {
	ANetApiRequest
	TransId       string         `json:"transId"`
	CustomerEmail string         `json:"customerEmail"`
	EmailSettings *EmailSettings `json:"emailSettings,omitempty"`
}

type TransactionDetailsTypeEmvDetails struct {
	Tag []TransactionDetailsTypeEmvDetailsTag `json:"tag"` //min=0
}

type UpdateSplitTenderGroupRequest struct {
	Payload UpdateSplitTenderGroupPayload `json:"updateSplitTenderGroupRequest"`
}
type UpdateSplitTenderGroupPayload struct {
	ANetApiRequest
	SplitTenderId     string `json:"splitTenderId"`
	SplitTenderStatus string `json:"splitTenderStatus"`
}

type TransactionDetails struct {
	TransId                   string                            `json:"transId"`
	RefTransId                string                            `json:"refTransId,omitempty"`
	SplitTenderId             string                            `json:"splitTenderId,omitempty"`
	SubmitTimeUTC             time.Time                         `json:"submitTimeUTC"`
	SubmitTimeLocal           time.Time                         `json:"submitTimeLocal"`
	TransactionType           string                            `json:"transactionType"`
	TransactionStatus         string                            `json:"transactionStatus"`
	ResponseCode              int                               `json:"responseCode"`
	ResponseReasonCode        int                               `json:"responseReasonCode"`
	Subscription              *SubscriptionPayment              `json:"subscription,omitempty"`
	ResponseReasonDescription string                            `json:"responseReasonDescription"`
	AuthCode                  string                            `json:"authCode,omitempty"`
	AvsResponse               string                            `json:"avsResponse,omitempty"`
	CardCodeResponse          string                            `json:"cardCodeResponse,omitempty"`
	CavvResponse              string                            `json:"cavvResponse,omitempty"`
	FdsFilterAction           string                            `json:"fdsFilterAction,omitempty"`
	FdsFilters                []FDSFilter                       `json:"fdsFilters,omitempty"`
	Batch                     *BatchDetails                     `json:"batch,omitempty"`
	Order                     *OrderEx                          `json:"order,omitempty"`
	RequestedAmount           float64                           `json:"requestedAmount,omitempty"`
	AuthAmount                float64                           `json:"authAmount"`
	SettleAmount              float64                           `json:"settleAmount"`
	Tax                       *ExtendedAmount                   `json:"tax,omitempty"`
	Shipping                  *ExtendedAmount                   `json:"shipping,omitempty"`
	Duty                      *ExtendedAmount                   `json:"duty,omitempty"`
	LineItems                 []Item                            `json:"lineItems,omitempty"`
	PrepaidBalanceRemaining   float64                           `json:"prepaidBalanceRemaining,omitempty"`
	TaxExempt                 bool                              `json:"taxExempt,omitempty"`
	Payment                   *PaymentMasked                    `json:"payment"`
	Customer                  *CustomerData                     `json:"customer,omitempty"`
	BillTo                    *CustomerAddress                  `json:"billTo,omitempty"`
	ShipTo                    *NameAndAddress                   `json:"shipTo,omitempty"`
	RecurringBilling          bool                              `json:"recurringBilling,omitempty"`
	CustomerIP                string                            `json:"customerIP,omitempty"`
	Product                   string                            `json:"product,omitempty"`
	EntryMode                 string                            `json:"entryMode,omitempty"`
	MarketType                string                            `json:"marketType,omitempty"`
	MobileDeviceId            string                            `json:"mobileDeviceId,omitempty"`
	CustomerSignature         string                            `json:"customerSignature,omitempty"`
	ReturnedItems             []ReturnedItem                    `json:"returnedItems,omitempty"`
	Solution                  *Solution                         `json:"solution,omitempty"`
	EmvDetails                *TransactionDetailsTypeEmvDetails `json:"emvDetails,omitempty"`
	Profile                   *CustomerProfileId                `json:"profile,omitempty"`
	Surcharge                 *ExtendedAmount                   `json:"surcharge,omitempty"`
	EmployeeId                string                            `json:"employeeId,omitempty"`
	Tip                       *ExtendedAmount                   `json:"tip,omitempty"`
	OtherTax                  *OtherTax                         `json:"otherTax,omitempty"`
	ShipFrom                  *NameAndAddress                   `json:"shipFrom,omitempty"`
}

type TokenMasked struct {
	TokenSource      string `json:"tokenSource,omitempty"`
	TokenNumber      string `json:"tokenNumber"`
	ExpirationDate   string `json:"expirationDate"`
	TokenRequestorId string `json:"tokenRequestorId,omitempty"`
}

type Processor struct {
	Name      string   `json:"name"`
	Id        int      `json:"id"`
	CardTypes []string `json:"cardTypes,omitempty"`
}

type AuUpdate struct {
	AuDetails
	NewCreditCard *CreditCardMasked `json:"newCreditCard"`
	OldCreditCard *CreditCardMasked `json:"oldCreditCard"`
}

type ProfileTransAuthOnly struct {
	ProfileTransOrder
}

type PaymentScheduleTypeInterval struct {
	Length int    `json:"length"`
	Unit   string `json:"unit"`
}

type TransactionSummary struct {
	TransId           string               `json:"transId"`
	SubmitTimeUTC     time.Time            `json:"submitTimeUTC"`
	SubmitTimeLocal   time.Time            `json:"submitTimeLocal"`
	TransactionStatus string               `json:"transactionStatus"`
	InvoiceNumber     string               `json:"invoiceNumber,omitempty"`
	FirstName         string               `json:"firstName,omitempty"`
	LastName          string               `json:"lastName,omitempty"`
	AccountType       string               `json:"accountType"`
	AccountNumber     string               `json:"accountNumber"`
	SettleAmount      float64              `json:"settleAmount"`
	MarketType        string               `json:"marketType,omitempty"`
	Product           string               `json:"product,omitempty"`
	MobileDeviceId    string               `json:"mobileDeviceId,omitempty"`
	Subscription      *SubscriptionPayment `json:"subscription,omitempty"`
	HasReturnedItems  bool                 `json:"hasReturnedItems,omitempty"`
	FraudInformation  *FraudInformation    `json:"fraudInformation,omitempty"`
	Profile           *CustomerProfileId   `json:"profile,omitempty"`
}

type AuDetails struct {
	CustomerProfileID        uint   `json:"customerProfileID"`
	CustomerPaymentProfileID uint   `json:"customerPaymentProfileID"`
	FirstName                string `json:"firstName,omitempty"`
	LastName                 string `json:"lastName,omitempty"`
	UpdateTimeUTC            string `json:"updateTimeUTC"`
	AuReasonCode             string `json:"auReasonCode"`
	ReasonDescription        string `json:"reasonDescription"`
}

type KeyManagementSchemeDUKPT struct {
	Operation     string                                 `json:"operation"`
	Mode          *KeyManagementSchemeDUKPTMode          `json:"mode"`
	DeviceInfo    *KeyManagementSchemeDUKPTDeviceInfo    `json:"deviceInfo"`
	EncryptedData *KeyManagementSchemeDUKPTEncryptedData `json:"encryptedData"`
}

type ProfileTransAmount struct {
	Amount    float64         `json:"amount"`
	Tax       *ExtendedAmount `json:"tax,omitempty"`
	Shipping  *ExtendedAmount `json:"shipping,omitempty"`
	Duty      *ExtendedAmount `json:"duty,omitempty"`
	LineItems []Item          `json:"lineItems,omitempty"` //min=0 max=30
}

type GetBatchStatisticsRequest struct {
	Payload GetBatchStatisticsPayload `json:"getBatchStatisticsRequest"`
}
type GetBatchStatisticsPayload struct {
	ANetApiRequest
	BatchId string `json:"batchId"`
}

type WebCheckOutData struct {
	Type      string                    `json:"type"`
	Id        string                    `json:"id"`
	Token     *WebCheckOutDataTypeToken `json:"token,omitempty"`
	BankToken *BankAccount              `json:"bankToken,omitempty"`
}

type Customer struct {
	Type           string          `json:"type,omitempty"`
	Id             string          `json:"id,omitempty"`
	Email          string          `json:"email,omitempty"`
	PhoneNumber    string          `json:"phoneNumber,omitempty"`
	FaxNumber      string          `json:"faxNumber,omitempty"`
	DriversLicense *DriversLicense `json:"driversLicense,omitempty"`
	TaxId          string          `json:"taxId,omitempty"`
}

type GetHostedProfilePageResponse struct {
	ANetApiResponse
	Token string `json:"token"`
}

type SecurePaymentContainerRequest struct {
	Payload SecurePaymentContainerPayload `json:"securePaymentContainerRequest"`
}
type SecurePaymentContainerPayload struct {
	ANetApiRequest
	Data *WebCheckOutData `json:"data"`
}

type IsAliveRequest struct {
	RefId string `json:"refId,omitempty"`
}

type ValidateCustomerPaymentProfileResponse struct {
	ANetApiResponse
	DirectResponse string `json:"directResponse,omitempty"`
}

type GetSettledBatchListRequest struct {
	Payload GetSettledBatchListPayload `json:"getSettledBatchListRequest"`
}
type GetSettledBatchListPayload struct {
	ANetApiRequest
	IncludeStatistics   bool      `json:"includeStatistics,omitempty"`
	FirstSettlementDate time.Time `json:"firstSettlementDate,omitempty"`
	LastSettlementDate  time.Time `json:"lastSettlementDate,omitempty"`
}

type LogoutRequest struct {
	ANetApiRequest `json:"logoutRequest"`
}

type CreateCustomerProfileTransactionResponse struct {
	ANetApiResponse
	TransactionResponse *TransactionResponse `json:"transactionResponse,omitempty"`
	DirectResponse      string               `json:"directResponse,omitempty"`
}

type PaymentSimple struct {
	CreditCard struct {
		CardNumber     string `json:"cardNumber"`
		ExpirationDate string `json:"expirationDate"`
	} `json:"creditCard"`
	BankAccount *BankAccount `json:"bankAccount"`
}

type ProfileTransPriorAuthCapture struct {
	ProfileTransAmount
	CustomerProfileId         string `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileId  string `json:"customerPaymentProfileId,omitempty"`
	CustomerShippingAddressId string `json:"customerShippingAddressId,omitempty"`
	TransId                   string `json:"transId"`
}

type SecurePaymentContainerResponse struct {
	ANetApiResponse
	OpaqueData OpaqueData `json:"opaqueData"`
}

type GetUnsettledTransactionListRequest struct {
	Payload GetUnsettledTransactionListPayload `json:"getUnsettledTransactionListRequest"`
}
type GetUnsettledTransactionListPayload struct {
	ANetApiRequest
	Status  string                  `json:"status,omitempty"`
	Sorting *TransactionListSorting `json:"sorting,omitempty"`
	Paging  *Paging                 `json:"paging,omitempty"`
}

type AuDelete struct {
	AuDetails
	CreditCard *CreditCardMasked `json:"creditCard"`
}

type CustomerProfileInfoEx struct {
	CustomerProfileEx
	ProfileType string `json:"profileType,omitempty"`
}

type KeyManagementScheme struct {
	Dukpt *KeyManagementSchemeDUKPT `json:"dukpt"`
}

type KeyManagementSchemeDUKPTEncryptedData struct {
	Value string `json:"value"`
}

type ProfileTransCaptureOnly struct {
	ProfileTransOrder
	ApprovalCode string `json:"approvalCode"`
}

type HeldTransactionRequest struct {
	Payload HeldTransactionPayload `json:"heldTransactionRequest"`
}
type HeldTransactionPayload struct {
	Action     string `json:"action"`
	RefTransId string `json:"refTransId"`
}

type GetTransactionDetailsResponse struct {
	ANetApiResponse
	Transaction *TransactionDetails `json:"transaction"`
	ClientId    string              `json:"clientId,omitempty"`
	TransrefId  string              `json:"transrefId,omitempty"`
}

type CustomerPaymentProfileEx struct {
	CustomerPaymentProfile
	CustomerPaymentProfileId string `json:"customerPaymentProfileId,omitempty"`
}

type GetTransactionListResponse struct {
	ANetApiResponse
	Transactions        []TransactionSummary `json:"transactions,omitempty"`
	TotalNumInResultSet int                  `json:"totalNumInResultSet,omitempty"`
}

type PaymentSchedule struct {
	Interval         *PaymentScheduleTypeInterval `json:"interval,omitempty"`
	StartDate        time.Time                    `json:"startDate,omitempty"` //was date
	TotalOccurrences int                          `json:"totalOccurrences,omitempty"`
	TrialOccurrences int                          `json:"trialOccurrences,omitempty"`
}

type ContactDetail struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type GetUnsettledTransactionListResponse struct {
	ANetApiResponse
	Transactions        []CustomerPaymentProfileListItem `json:"transactions,omitempty"`
	TotalNumInResultSet int                              `json:"totalNumInResultSet,omitempty"`
}

type BatchStatistic struct {
	AccountType               string  `json:"accountType"`
	ChargeAmount              float64 `json:"chargeAmount"`
	ChargeCount               int     `json:"chargeCount"`
	RefundAmount              float64 `json:"refundAmount"`
	RefundCount               int     `json:"refundCount"`
	VoidCount                 int     `json:"voidCount"`
	DeclineCount              int     `json:"declineCount"`
	ErrorCount                int     `json:"errorCount"`
	ReturnedItemAmount        float64 `json:"returnedItemAmount,omitempty"`
	ReturnedItemCount         int     `json:"returnedItemCount,omitempty"`
	ChargebackAmount          float64 `json:"chargebackAmount,omitempty"`
	ChargebackCount           int     `json:"chargebackCount,omitempty"`
	CorrectionNoticeCount     int     `json:"correctionNoticeCount,omitempty"`
	ChargeChargeBackAmount    float64 `json:"chargeChargeBackAmount,omitempty"`
	ChargeChargeBackCount     int     `json:"chargeChargeBackCount,omitempty"`
	RefundChargeBackAmount    float64 `json:"refundChargeBackAmount,omitempty"`
	RefundChargeBackCount     int     `json:"refundChargeBackCount,omitempty"`
	ChargeReturnedItemsAmount float64 `json:"chargeReturnedItemsAmount,omitempty"`
	ChargeReturnedItemsCount  int     `json:"chargeReturnedItemsCount,omitempty"`
	RefundReturnedItemsAmount float64 `json:"refundReturnedItemsAmount,omitempty"`
	RefundReturnedItemsCount  int     `json:"refundReturnedItemsCount,omitempty"`
}

type SendCustomerTransactionReceiptResponse struct {
	ANetApiResponse
}

type SubscriptionCustomerProfile struct {
	CustomerProfileEx
	PaymentProfile  *CustomerPaymentProfileMasked `json:"paymentProfile,omitempty"`
	ShippingProfile *CustomerAddressEx            `json:"shippingProfile,omitempty"`
}

type CustomerProfileEx struct {
	CustomerProfileBase
	CustomerProfileId string `json:"customerProfileId,omitempty"`
}

type GetBatchStatisticsResponse struct {
	ANetApiResponse
	Batch *BatchDetails `json:"batch,omitempty"`
}

type SubscriptionPayment struct {
	Id     int `json:"id"`
	PayNum int `json:"payNum"`
}

type UpdateCustomerShippingAddressResponse struct {
	ANetApiResponse
}

type LogoutResponse struct {
	ANetApiResponse
}

type GetTransactionListRequest struct {
	Payload GetTransactionListPayload `json:"getTransactionListRequest"`
}
type GetTransactionListPayload struct {
	ANetApiRequest
	BatchId string                  `json:"batchId,omitempty"`
	Sorting *TransactionListSorting `json:"sorting,omitempty"`
	Paging  *Paging                 `json:"paging,omitempty"`
}

type FDSFilter struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

type CustomerProfileSummary struct {
	CustomerProfileId  string    `json:"customerProfileId,omitempty"`
	Description        string    `json:"description,omitempty"`
	MerchantCustomerId string    `json:"merchantCustomerId"`
	Email              string    `json:"email,omitempty"`
	CreatedDate        time.Time `json:"createdDate"`
}

type UpdateHeldTransactionRequest struct {
	Payload UpdateHeldTransactionPayload `json:"updateHeldTransactionRequest"`
}
type UpdateHeldTransactionPayload struct {
	ANetApiRequest
	HeldTransactionRequest *HeldTransactionRequest `json:"heldTransactionRequest"`
}

type TransactionDetailsTypeEmvDetailsTag struct {
	TagId string `json:"tagId"`
	Data  string `json:"data"`
}

type KeyManagementSchemeDUKPTMode struct {
	Pin  string `json:"pin,omitempty"`
	Data string `json:"data,omitempty"`
}

type CustomerProfileId struct {
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId,omitempty"`
	CustomerAddressId        string `json:"customerAddressId,omitempty"`
}

type OrderEx struct {
	Order
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`
}

type ProfileTransRefund struct {
	ProfileTransAmount
	CustomerProfileId         string   `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileId  string   `json:"customerPaymentProfileId,omitempty"`
	CustomerShippingAddressId string   `json:"customerShippingAddressId,omitempty"`
	CreditCardNumberMasked    string   `json:"creditCardNumberMasked,omitempty"`
	BankRoutingNumberMasked   string   `json:"bankRoutingNumberMasked,omitempty"`
	BankAccountNumberMasked   string   `json:"bankAccountNumberMasked,omitempty"`
	Order                     *OrderEx `json:"order,omitempty"`
	TransId                   string   `json:"transId,omitempty"`
}

type WebCheckOutDataTypeToken struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	CardCode       string `json:"cardCode,omitempty"`
	Zip            string `json:"zip,omitempty"`
	FullName       string `json:"fullName,omitempty"`
}

type TransactionListSorting struct {
	OrderBy         string `json:"orderBy"`
	OrderDescending bool   `json:"orderDescending"`
}

type GetHostedProfilePageRequest struct {
	Payload GetHostedProfilePagePayload `json:"getHostedProfilePageRequest"`
}
type GetHostedProfilePagePayload struct {
	ANetApiRequest
	CustomerProfileId     string    `json:"customerProfileId"`
	HostedProfileSettings []Setting `json:"hostedProfileSettings,omitempty"`
}

type IsAliveResponse struct {
	ANetApiResponse
}

type KeyValue struct {
	Encoding            string               `json:"encoding"`
	EncryptionAlgorithm string               `json:"encryptionAlgorithm"`
	Scheme              *KeyManagementScheme `json:"scheme"`
}

type ProfileTransOrder struct {
	ProfileTransAmount
	CustomerProfileId         string                     `json:"customerProfileId"`
	CustomerPaymentProfileId  string                     `json:"customerPaymentProfileId"`
	CustomerShippingAddressId string                     `json:"customerShippingAddressId,omitempty"`
	Order                     *OrderEx                   `json:"order,omitempty"`
	TaxExempt                 bool                       `json:"taxExempt,omitempty"`
	RecurringBilling          bool                       `json:"recurringBilling,omitempty"`
	CardCode                  string                     `json:"cardCode,omitempty"`
	SplitTenderId             string                     `json:"splitTenderId,omitempty"`
	ProcessingOptions         *ProcessingOptions         `json:"processingOptions,omitempty"`
	SubsequentAuthInformation *SubsequentAuthInformation `json:"subsequentAuthInformation,omitempty"`
}

type ProfileTransVoid struct {
	CustomerProfileId         string `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileId  string `json:"customerPaymentProfileId,omitempty"`
	CustomerShippingAddressId string `json:"customerShippingAddressId,omitempty"`
	TransId                   string `json:"transId"`
}

type ProfileTransaction struct {
	ProfileTransAuthCapture      *ProfileTransOrder            `json:"profileTransAuthCapture"`
	ProfileTransAuthOnly         *ProfileTransAuthOnly         `json:"profileTransAuthOnly"`
	ProfileTransPriorAuthCapture *ProfileTransPriorAuthCapture `json:"profileTransPriorAuthCapture"`
	ProfileTransCaptureOnly      *ProfileTransCaptureOnly      `json:"profileTransCaptureOnly"`
	ProfileTransRefund           *ProfileTransRefund           `json:"profileTransRefund"`
	ProfileTransVoid             *ProfileTransVoid             `json:"profileTransVoid"`
}

type UpdateHeldTransactionResponse struct {
	ANetApiResponse
	TransactionResponse *TransactionResponse `json:"transactionResponse,omitempty"`
}

type EnumCollection struct {
	CustomerProfileSummaryType *CustomerProfileSummary `json:"customerProfileSummaryType"`
	PaymentSimpleType          *PaymentSimple          `json:"paymentSimpleType"`
	AccountTypeEnum            string                  `json:"accountTypeEnum"`
	CardTypeEnum               string                  `json:"cardTypeEnum"`
	FdsFilterActionEnum        string                  `json:"fdsFilterActionEnum"`
	PermissionsEnum            string                  `json:"permissionsEnum"`
	SettingNameEnum            string                  `json:"settingNameEnum"`
	SettlementStateEnum        string                  `json:"settlementStateEnum"`
	TransactionStatusEnum      string                  `json:"transactionStatusEnum"`
	TransactionTypeEnum        string                  `json:"transactionTypeEnum"`
}

type GetHostedPaymentPageRequest struct {
	Payload GetHostedPaymentPagePayload `json:"getHostedPaymentPageRequest"`
}
type GetHostedPaymentPagePayload struct {
	ANetApiRequest
	TransactionRequest    *Transaction `json:"transactionRequest"`
	HostedPaymentSettings []Setting    `json:"hostedPaymentSettings,omitempty"`
}

type GetTransactionDetailsRequest struct {
	Payload GetTransactionDetailsPayload `json:"getTransactionDetailsRequest"`
}
type GetTransactionDetailsPayload struct {
	ANetApiRequest
	TransId string `json:"transId"`
}

type KeyManagementSchemeDUKPTDeviceInfo struct {
	Description string `json:"description"`
}

type GetHostedPaymentPageResponse struct {
	ANetApiResponse
	Token string `json:"token"`
}

type UpdateCustomerProfileResponse struct {
	ANetApiResponse
}

type UpdateCustomerPaymentProfileResponse struct {
	ANetApiResponse
	ValidationDirectResponse string `json:"validationDirectResponse,omitempty"`
}

type BatchDetails struct {
	BatchId             string           `json:"batchId"`
	SettlementTimeUTC   time.Time        `json:"settlementTimeUTC,omitempty"`
	SettlementTimeLocal time.Time        `json:"settlementTimeLocal,omitempty"`
	SettlementState     string           `json:"settlementState"`
	PaymentMethod       string           `json:"paymentMethod,omitempty"`
	MarketType          string           `json:"marketType,omitempty"`
	Product             string           `json:"product,omitempty"`
	Statistics          []BatchStatistic `json:"statistics,omitempty"`
}

type FraudInformation struct {
	FraudFilterList []string `json:"fraudFilterList"`
	FraudAction     string   `json:"fraudAction"`
}

type ReturnedItem struct {
	ErrMessage
	Id        string    `json:"id"`
	DateUTC   time.Time `json:"dateUTC"`
	DateLocal time.Time `json:"dateLocal"`
}
