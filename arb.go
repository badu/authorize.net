package authorize_net

import "time"

// -- request --
type ARBUpdateSubscriptionRequest struct {
	Payload ARBUpdateSubscriptionPayload `json:"arbUpdateSubscriptionRequest"`
}

type ARBSubscription struct {
	Name            string             `json:"name,omitempty"`
	PaymentSchedule *PaymentSchedule   `json:"paymentSchedule,omitempty"`
	Amount          float64            `json:"amount,omitempty"`
	TrialAmount     float64            `json:"trialAmount,omitempty"`
	Payment         *Payment           `json:"payment,omitempty"`
	Order           *Order             `json:"order,omitempty"`
	Customer        *Customer          `json:"customer,omitempty"`
	BillTo          *NameAndAddress    `json:"billTo,omitempty"`
	ShipTo          *NameAndAddress    `json:"shipTo,omitempty"`
	Profile         *CustomerProfileId `json:"profile,omitempty"`
}

type ARBUpdateSubscriptionPayload struct {
	ANetApiRequest
	SubscriptionId string           `json:"subscriptionId"`
	Subscription   *ARBSubscription `json:"subscription"`
}

type ArbTransaction struct {
	TransId       string    `json:"transId,omitempty"`
	Response      string    `json:"response,omitempty"`
	SubmitTimeUTC time.Time `json:"submitTimeUTC,omitempty"`
	PayNum        int       `json:"payNum,omitempty"`
	AttemptNum    int       `json:"attemptNum,omitempty"`
}

type ARBTransactionList struct {
	ArbTransaction []ArbTransaction `json:"arbTransaction,omitempty"` //min=0
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

type ARBGetSubscriptionListSorting struct {
	OrderBy         string `json:"orderBy"`
	OrderDescending bool   `json:"orderDescending"`
}

type ARBGetSubscriptionListRequest struct {
	Payload ARBGetSubscriptionListPayload `json:"ARBGetSubscriptionListRequest"`
}
type ARBGetSubscriptionListPayload struct {
	ANetApiRequest
	SearchType string                         `json:"searchType"`
	Sorting    *ARBGetSubscriptionListSorting `json:"sorting,omitempty"`
	Paging     *Paging                        `json:"paging,omitempty"`
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
	Subscription *ARBSubscriptionMasked `json:"subscription"`
}

type ARBGetSubscriptionListResponse struct {
	ANetApiResponse
	TotalNumInResultSet int                  `json:"totalNumInResultSet,omitempty"`
	SubscriptionDetails []SubscriptionDetail `json:"subscriptionDetails,omitempty"`
}

type ARBCreateSubscriptionResponse struct {
	ANetApiResponse
	SubscriptionId string             `json:"subscriptionId,omitempty"`
	Profile        *CustomerProfileId `json:"profile,omitempty"`
}
