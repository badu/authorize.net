package authorize_net

import (
	"time"
)

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

type CardArt struct {
	CardBrand       string `json:"cardBrand,omitempty"`
	CardImageHeight string `json:"cardImageHeight,omitempty"`
	CardImageUrl    string `json:"cardImageUrl,omitempty"`
	CardImageWidth  string `json:"cardImageWidth,omitempty"`
	CardType        string `json:"cardType,omitempty"`
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
	NewCreditCard *CreditCardMasked `json:"newCreditCard,omitempty"`
	OldCreditCard *CreditCardMasked `json:"oldCreditCard,omitempty"`
}

type PaymentScheduleTypeInterval struct {
	Length int    `json:"length"`
	Unit   string `json:"unit"`
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
	Mode          *KeyManagementSchemeDUKPTMode          `json:"mode,omitempty"`
	DeviceInfo    *KeyManagementSchemeDUKPTDeviceInfo    `json:"deviceInfo,omitempty"`
	EncryptedData *KeyManagementSchemeDUKPTEncryptedData `json:"encryptedData,omitempty"`
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
	Data *WebCheckOutData `json:"data,omitempty"`
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

type PaymentSimple struct {
	CreditCard struct {
		CardNumber     string `json:"cardNumber"`
		ExpirationDate string `json:"expirationDate"`
	} `json:"creditCard"`
	BankAccount *BankAccount `json:"bankAccount,omitempty"`
}

type SecurePaymentContainerResponse struct {
	ANetApiResponse
	OpaqueData OpaqueData `json:"opaqueData"`
}

type AuDelete struct {
	AuDetails
	CreditCard *CreditCardMasked `json:"creditCard,omitempty"`
}

type CustomerProfileInfoEx struct {
	CustomerProfileEx
	ProfileType CustomerProfileTypeEnum `json:"profileType,omitempty"`
}

type KeyManagementScheme struct {
	Dukpt *KeyManagementSchemeDUKPT `json:"dukpt,omitempty"`
}

type KeyManagementSchemeDUKPTEncryptedData struct {
	Value string `json:"value"`
}

type HeldTransactionRequest struct {
	Payload HeldTransactionPayload `json:"heldTransactionRequest"`
}
type HeldTransactionPayload struct {
	Action     string `json:"action"`
	RefTransId string `json:"refTransId"`
}

type CustomerPaymentProfileEx struct {
	CustomerPaymentProfile
	CustomerPaymentProfileId string `json:"customerPaymentProfileId,omitempty"`
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
	HeldTransactionRequest *HeldTransactionRequest `json:"heldTransactionRequest,omitempty"`
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

type WebCheckOutDataTypeToken struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	CardCode       string `json:"cardCode,omitempty"`
	Zip            string `json:"zip,omitempty"`
	FullName       string `json:"fullName,omitempty"`
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
	Scheme              *KeyManagementScheme `json:"scheme,omitempty"`
}

type UpdateHeldTransactionResponse struct {
	ANetApiResponse
	TransactionResponse *TransactionResponse `json:"transactionResponse,omitempty"`
}

type EnumCollection struct {
	CustomerProfileSummaryType *CustomerProfileSummary `json:"customerProfileSummaryType,omitempty"`
	PaymentSimpleType          *PaymentSimple          `json:"paymentSimpleType,omitempty"`
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
	TransactionRequest    *Transaction `json:"transactionRequest,omitempty"`
	HostedPaymentSettings []Setting    `json:"hostedPaymentSettings,omitempty"`
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
