package authorize_net

type GetSettledBatchListResponse struct {
	ANetApiResponse
	BatchList []BatchDetails `json:"batchList,omitempty"`
}

type UpdateSplitTenderGroupResponse struct {
	ANetApiResponse
}

type UpdateSplitTenderGroupRequest struct {
	Payload UpdateSplitTenderGroupPayload `json:"updateSplitTenderGroupRequest"`
}
type UpdateSplitTenderGroupPayload struct {
	ANetApiRequest
	SplitTenderId     string `json:"splitTenderId"`
	SplitTenderStatus string `json:"splitTenderStatus"`
}

type AuUpdate struct {
	AuDetails
	NewCreditCard *CreditCardMasked `json:"newCreditCard,omitempty"`
	OldCreditCard *CreditCardMasked `json:"oldCreditCard,omitempty"`
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
	IncludeStatistics   bool `json:"includeStatistics,omitempty"`
	FirstSettlementDate Date `json:"firstSettlementDate,omitempty"`
	LastSettlementDate  Date `json:"lastSettlementDate,omitempty"`
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

type GetBatchStatisticsResponse struct {
	ANetApiResponse
	Batch *BatchDetails `json:"batch,omitempty"`
}

type UpdateCustomerShippingAddressResponse struct {
	ANetApiResponse
}

type CustomerProfileSummary struct {
	CustomerProfileId  string `json:"customerProfileId,omitempty"`
	Description        string `json:"description,omitempty"`
	MerchantCustomerId string `json:"merchantCustomerId"`
	Email              string `json:"email,omitempty"`
	CreatedDate        Date   `json:"createdDate"`
}

type UpdateHeldTransactionRequest struct {
	Payload UpdateHeldTransactionPayload `json:"updateHeldTransactionRequest"`
}
type UpdateHeldTransactionPayload struct {
	ANetApiRequest
	HeldTransactionRequest *HeldTransactionRequest `json:"heldTransactionRequest,omitempty"`
}

type KeyManagementSchemeDUKPTMode struct {
	Pin  string `json:"pin,omitempty"`
	Data string `json:"data,omitempty"`
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
