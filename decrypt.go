package authorize_net

// -- request --

type DecryptPaymentDataRequest struct {
	Payload DecryptPaymentDataPayload `json:"decryptPaymentDataRequest"`
}
type DecryptPaymentDataPayload struct {
	ANetApiRequest
	OpaqueData *OpaqueData `json:"opaqueData,omitempty"`
	CallId     string      `json:"callId,omitempty"`
}

// -- response --

type PaymentDetails struct {
	Currency         string `json:"currency,omitempty"`
	PromoCode        string `json:"promoCode,omitempty"`
	Misc             string `json:"misc,omitempty"`
	GiftWrap         string `json:"giftWrap,omitempty"`
	Discount         string `json:"discount,omitempty"`
	Tax              string `json:"tax,omitempty"`
	ShippingHandling string `json:"shippingHandling,omitempty"`
	SubTotal         string `json:"subTotal,omitempty"`
	OrderID          string `json:"orderID,omitempty"`
	Amount           string `json:"amount,omitempty"`
}

type DecryptPaymentDataResponse struct {
	ANetApiResponse
	ShippingInfo   *CustomerAddress  `json:"shippingInfo,omitempty"`
	BillingInfo    *CustomerAddress  `json:"billingInfo,omitempty"`
	CardInfo       *CreditCardMasked `json:"cardInfo,omitempty"`
	PaymentDetails *PaymentDetails   `json:"paymentDetails,omitempty"`
}
