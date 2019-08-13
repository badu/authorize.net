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

// -- response --
type ARBCancelSubscriptionResponse struct {
	ANetApiResponse
}
