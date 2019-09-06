package authorize_net

// -- request --
type ARBUpdateSubscriptionRequest struct {
	Payload ARBUpdateSubscriptionPayload `json:"ARBUpdateSubscriptionRequest"`
}

type PaymentScheduleTypeInterval struct {
	Length int                     `json:"length"`
	Unit   ARBSubscriptionUnitEnum `json:"unit"`
}

type PaymentSchedule struct {
	Interval         *PaymentScheduleTypeInterval `json:"interval,omitempty"`
	StartDate        DateTime                     `json:"startDate,omitempty"`
	TotalOccurrences int                          `json:"totalOccurrences,omitempty"`
	TrialOccurrences int                          `json:"trialOccurrences,omitempty"`
}

type ARBSubscription struct {
	Name            string             `json:"name,omitempty"`
	PaymentSchedule *PaymentSchedule   `json:"paymentSchedule,omitempty"`
	Amount          float64            `json:"amount,omitempty"`
	TrialAmount     float64            `json:"trialAmount"`
	Payment         *Payment           `json:"payment,omitempty"`
	Order           *Order             `json:"order,omitempty"`
	Customer        *Customer          `json:"customer,omitempty"`
	BillTo          *NameAndAddress    `json:"billTo,omitempty"`
	ShipTo          *NameAndAddress    `json:"shipTo,omitempty"`
	Profile         *CustomerProfileId `json:"profile,omitempty"`
}

type ARBUpdateSubscriptionPayload struct {
	ANetApiRequest
	SubscriptionId string          `json:"subscriptionId"`
	Subscription   ARBSubscription `json:"subscription"`
}

type ArbTransaction struct {
	TransId       string    `json:"transId,omitempty"`
	Response      string    `json:"response,omitempty"`
	SubmitTimeUTC DateTimeZ `json:"submitTimeUTC,omitempty"`
	PayNum        int       `json:"payNum,omitempty"`
	AttemptNum    int       `json:"attemptNum,omitempty"`
}

type ARBTransactionList struct {
	ArbTransaction []ArbTransaction `json:"arbTransaction,omitempty"` //min=0
}

type CustomerProfileEx struct {
	CustomerProfileBase
	CustomerProfileId string `json:"customerProfileId,omitempty"`
}

type SubscriptionCustomerProfile struct {
	CustomerProfileEx
	PaymentProfile  *CustomerPaymentProfileMasked `json:"paymentProfile,omitempty"`
	ShippingProfile *CustomerAddressEx            `json:"shippingProfile,omitempty"`
}

type ARBSubscriptionMasked struct {
	Name            string                       `json:"name,omitempty"`
	PaymentSchedule *PaymentSchedule             `json:"paymentSchedule,omitempty"`
	Amount          float64                      `json:"amount,omitempty"`
	TrialAmount     float64                      `json:"trialAmount,omitempty"`
	Status          string                       `json:"status,omitempty"`
	Profile         *SubscriptionCustomerProfile `json:"profile,omitempty"`
	Order           *Order                       `json:"order,omitempty"`
	ArbTransactions *ARBTransactionList          `json:"arbTransactions,omitempty"`
}

type ARBGetSubscriptionStatusRequest struct {
	Payload ARBGetSubscriptionStatusPayload `json:"ARBGetSubscriptionStatusRequest"`
}
type ARBGetSubscriptionStatusPayload struct {
	ANetApiRequest
	SubscriptionId string `json:"subscriptionId"`
}

type ARBGetSubscriptionRequest struct {
	Payload ARBGetSubscriptionPayload `json:"ARBGetSubscriptionRequest"`
}
type ARBGetSubscriptionPayload struct {
	ANetApiRequest
	SubscriptionId      string `json:"subscriptionId"`
	IncludeTransactions bool   `json:"includeTransactions,omitempty"`
}

type ARBCreateSubscriptionRequest struct {
	Payload ARBCreateSubscriptionPayload `json:"ARBCreateSubscriptionRequest"`
}
type ARBCreateSubscriptionPayload struct {
	ANetApiRequest
	Subscription ARBSubscription `json:"subscription"`
}

type ARBGetSubscriptionListRequest struct {
	Payload ARBGetSubscriptionListPayload `json:"ARBGetSubscriptionListRequest"`
}
type ARBGetSubscriptionListPayload struct {
	ANetApiRequest
	SearchType string  `json:"searchType"`
	Sorting    Sorting `json:"sorting,omitempty"`
	Paging     Paging  `json:"paging,omitempty"`
}

type ARBCancelSubscriptionRequest struct {
	Payload ARBCancelSubscriptionPayload `json:"ARBCancelSubscriptionRequest"`
}
type ARBCancelSubscriptionPayload struct {
	ANetApiRequest
	SubscriptionId string `json:"subscriptionId"`
}

// -- response --
type ARBCancelSubscriptionResponse struct {
	ANetApiResponse
}

type ARBUpdateSubscriptionResponse struct {
	ANetApiResponse
	Profile CustomerProfileId `json:"profile,omitempty"`
}

type ARBGetSubscriptionStatusResponse struct {
	ANetApiResponse
	Status string `json:"status,omitempty"`
}

type ARBGetSubscriptionResponse struct {
	ANetApiResponse
	Subscription *ARBSubscriptionMasked `json:"subscription,omitempty"`
}

type SubscriptionDetail struct {
	Id                        int          `json:"id"`
	Name                      string       `json:"name,omitempty"`
	Status                    string       `json:"status"`
	CreateTimeStampUTC        DateTimeNano `json:"createTimeStampUTC"`
	FirstName                 string       `json:"firstName,omitempty"`
	LastName                  string       `json:"lastName,omitempty"`
	TotalOccurrences          int          `json:"totalOccurrences"`
	PastOccurrences           int          `json:"pastOccurrences"`
	PaymentMethod             string       `json:"paymentMethod"`
	AccountNumber             string       `json:"accountNumber,omitempty"`
	Invoice                   string       `json:"invoice,omitempty"`
	Amount                    float64      `json:"amount"`
	CurrencyCode              string       `json:"currencyCode,omitempty"`
	CustomerProfileId         int          `json:"customerProfileId"`
	CustomerPaymentProfileId  int          `json:"customerPaymentProfileId"`
	CustomerShippingProfileId int          `json:"customerShippingProfileId,omitempty"`
}

type ARBGetSubscriptionListResponse struct {
	ANetApiResponse
	TotalNumInResultSet int                  `json:"totalNumInResultSet,omitempty"`
	SubscriptionDetails []SubscriptionDetail `json:"subscriptionDetails,omitempty"`
}

type CustomerProfileId struct {
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId,omitempty"`
	CustomerAddressId        string `json:"customerAddressId,omitempty"`
}

type ARBCreateSubscriptionResponse struct {
	ANetApiResponse
	SubscriptionId string             `json:"subscriptionId,omitempty"`
	Profile        *CustomerProfileId `json:"profile,omitempty"`
}
