package authorize_net

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TransactionTypeEnum uint

const (
	AuthOnlyTransaction            TransactionTypeEnum = 1
	AuthCaptureTransaction         TransactionTypeEnum = 2
	CaptureOnlyTransaction         TransactionTypeEnum = 3
	RefundTransaction              TransactionTypeEnum = 4
	PriorAuthCaptureTransaction    TransactionTypeEnum = 5
	VoidTransaction                TransactionTypeEnum = 6
	GetDetailsTransaction          TransactionTypeEnum = 7
	AuthOnlyContinueTransaction    TransactionTypeEnum = 8
	AuthCaptureContinueTransaction TransactionTypeEnum = 9
)

func (t TransactionTypeEnum) MarshalJSON() ([]byte, error) {
	switch t {
	case AuthOnlyTransaction:
		return json.Marshal("authOnlyTransaction")
	case AuthCaptureTransaction:
		return json.Marshal("authCaptureTransaction")
	case CaptureOnlyTransaction:
		return json.Marshal("captureOnlyTransaction")
	case RefundTransaction:
		return json.Marshal("refundTransaction")
	case PriorAuthCaptureTransaction:
		return json.Marshal("priorAuthCaptureTransaction")
	case VoidTransaction:
		return json.Marshal("voidTransaction")
	case GetDetailsTransaction:
		return json.Marshal("getDetailsTransaction")
	case AuthOnlyContinueTransaction:
		return json.Marshal("authOnlyContinueTransaction")
	case AuthCaptureContinueTransaction:
		return json.Marshal("authCaptureContinueTransaction")
	default:
		return json.Marshal("null")
	}
}

func (t *TransactionTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid TransactionTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "authOnlyTransaction":
		*t = AuthOnlyTransaction
	case "authCaptureTransaction":
		*t = AuthCaptureTransaction
	case "captureOnlyTransaction":
		*t = CaptureOnlyTransaction
	case "refundTransaction":
		*t = RefundTransaction
	case "priorAuthCaptureTransaction":
		*t = PriorAuthCaptureTransaction
	case "voidTransaction":
		*t = VoidTransaction
	case "getDetailsTransaction":
		*t = GetDetailsTransaction
	case "authOnlyContinueTransaction":
		*t = AuthOnlyContinueTransaction
	case "authCaptureContinueTransaction":
		*t = AuthCaptureContinueTransaction
	default:
		return fmt.Errorf("invalid TransactionTypeEnum : %q", string(data))
	}
	return nil
}

type BankAccountTypeEnum uint

const (
	BankAccountTypeChecking         BankAccountTypeEnum = 1
	BankAccountTypeSavings          BankAccountTypeEnum = 2
	BankAccountTypeBusinessChecking BankAccountTypeEnum = 3
)

func (b BankAccountTypeEnum) MarshalJSON() ([]byte, error) {
	switch b {
	case BankAccountTypeChecking:
		return json.Marshal("checking")
	case BankAccountTypeSavings:
		return json.Marshal("savings")
	case BankAccountTypeBusinessChecking:
		return json.Marshal("businessChecking")
	default:
		return json.Marshal("null")
	}
}

func (b *BankAccountTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid BankAccountTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "checking":
		*b = BankAccountTypeChecking
	case "savings":
		*b = BankAccountTypeSavings
	case "businessChecking":
		*b = BankAccountTypeBusinessChecking
	default:
		return fmt.Errorf("invalid BankAccountTypeEnum : %q", string(data))
	}
	return nil
}

type EcheckTypeEnum uint

const (
	EcheckPPD EcheckTypeEnum = 1
	EcheckWEB EcheckTypeEnum = 2
	EcheckCCD EcheckTypeEnum = 3
	EcheckTEL EcheckTypeEnum = 4
	EcheckARC EcheckTypeEnum = 5
	EcheckBOC EcheckTypeEnum = 6
)

func (a EcheckTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case EcheckPPD:
		return json.Marshal("PPD")
	case EcheckWEB:
		return json.Marshal("WEB")
	case EcheckCCD:
		return json.Marshal("CCD")
	case EcheckTEL:
		return json.Marshal("TEL")
	case EcheckARC:
		return json.Marshal("ARC")
	case EcheckBOC:
		return json.Marshal("BOC")
	default:
		return json.Marshal("null")
	}
}

func (a *EcheckTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid EcheckTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "PPD":
		*a = EcheckPPD
	case "WEB":
		*a = EcheckWEB
	case "CCD":
		*a = EcheckCCD
	case "TEL":
		*a = EcheckTEL
	case "ARC":
		*a = EcheckARC
	case "BOC":
		*a = EcheckBOC
	default:
		return fmt.Errorf("invalid EcheckTypeEnum : %q", string(data))
	}
	return nil
}

type PaymentMethodsTypeEnum uint

const (
	PaymentMethodVisa            PaymentMethodsTypeEnum = 1
	PaymentMethodMasterCard      PaymentMethodsTypeEnum = 2
	PaymentMethodAmericanExpress PaymentMethodsTypeEnum = 3
	PaymentMethodDiscover        PaymentMethodsTypeEnum = 4
	PaymentMethodJCB             PaymentMethodsTypeEnum = 5
	PaymentMethodDinersClub      PaymentMethodsTypeEnum = 6
	PaymentMethodEnRoute         PaymentMethodsTypeEnum = 7
	PaymentMethodECheck          PaymentMethodsTypeEnum = 8
	PaymentMethodPaypal          PaymentMethodsTypeEnum = 9
	PaymentMethodVisaCheckout    PaymentMethodsTypeEnum = 10
	PaymentMethodApplePay        PaymentMethodsTypeEnum = 11
	PaymentMethodAndroidPay      PaymentMethodsTypeEnum = 12
)

func (a PaymentMethodsTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case PaymentMethodVisa:
		return json.Marshal("Visa")
	case PaymentMethodMasterCard:
		return json.Marshal("MasterCard")
	case PaymentMethodAmericanExpress:
		return json.Marshal("AmericanExpress")
	case PaymentMethodDiscover:
		return json.Marshal("Discover")
	case PaymentMethodJCB:
		return json.Marshal("JCB")
	case PaymentMethodDinersClub:
		return json.Marshal("DinersClub")
	case PaymentMethodEnRoute:
		return json.Marshal("EnRoute")
	case PaymentMethodECheck:
		return json.Marshal("eCheck")
	case PaymentMethodPaypal:
		return json.Marshal("Paypal")
	case PaymentMethodVisaCheckout:
		return json.Marshal("VisaCheckout")
	case PaymentMethodApplePay:
		return json.Marshal("ApplePay")
	case PaymentMethodAndroidPay:
		return json.Marshal("AndroidPay")
	default:
		return json.Marshal("null")
	}
}

func (a *PaymentMethodsTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid PaymentMethodsTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "Visa":
		*a = PaymentMethodVisa
	case "MasterCard":
		*a = PaymentMethodMasterCard
	case "AmericanExpress":
		*a = PaymentMethodAmericanExpress
	case "Discover":
		*a = PaymentMethodDiscover
	case "JCB":
		*a = PaymentMethodJCB
	case "DinersClub":
		*a = PaymentMethodDinersClub
	case "EnRoute":
		*a = PaymentMethodEnRoute
	case "eCheck":
		*a = PaymentMethodECheck
	case "PayPal":
		*a = PaymentMethodPaypal
	case "VisaCheckout":
		*a = PaymentMethodVisaCheckout
	case "ApplePay":
		*a = PaymentMethodApplePay
	case "AndroidPay":
		*a = PaymentMethodAndroidPay
	default:
		return fmt.Errorf("invalid PaymentMethodsTypeEnum : %q", string(data))
	}
	return nil
}

type CardTypeEnum uint

const (
	CardVisa            CardTypeEnum = 1
	CardMasterCard      CardTypeEnum = 2
	CardAmericanExpress CardTypeEnum = 3
	CardDiscover        CardTypeEnum = 4
	CardJCB             CardTypeEnum = 5
	CardDinersClub      CardTypeEnum = 6
)

func (a CardTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case CardVisa:
		return json.Marshal("Visa")
	case CardMasterCard:
		return json.Marshal("MasterCard")
	case CardAmericanExpress:
		return json.Marshal("AmericanExpress")
	case CardDiscover:
		return json.Marshal("Discover")
	case CardJCB:
		return json.Marshal("JCB")
	case CardDinersClub:
		return json.Marshal("DinersClub")
	default:
		return json.Marshal("null")
	}
}

func (a *CardTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid CardTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "Visa":
		*a = CardVisa
	case "MasterCard":
		*a = CardMasterCard
	case "AmericanExpress":
		*a = CardAmericanExpress
	case "Discover":
		*a = CardDiscover
	case "JCB":
		*a = CardJCB
	case "DinersClub":
		*a = CardDinersClub
	default:
		return fmt.Errorf("invalid CardTypeEnum : %q", string(data))
	}
	return nil
}

type AccountTypeEnum uint

const (
	AccountTypeVisa            AccountTypeEnum = 1
	AccountTypeMasterCard      AccountTypeEnum = 2
	AccountTypeAmericanExpress AccountTypeEnum = 3
	AccountTypeDiscover        AccountTypeEnum = 4
	AccountTypeJCB             AccountTypeEnum = 5
	AccountTypeDinersClub      AccountTypeEnum = 6
	AccountTypeECheck          AccountTypeEnum = 7
)

func (a AccountTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case AccountTypeVisa:
		return json.Marshal("Visa")
	case AccountTypeMasterCard:
		return json.Marshal("MasterCard")
	case AccountTypeAmericanExpress:
		return json.Marshal("AmericanExpress")
	case AccountTypeDiscover:
		return json.Marshal("Discover")
	case AccountTypeJCB:
		return json.Marshal("JCB")
	case AccountTypeDinersClub:
		return json.Marshal("DinersClub")
	case AccountTypeECheck:
		return json.Marshal("eCheck")
	default:
		return json.Marshal("null")
	}
}

func (a *AccountTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid AccountTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "Visa":
		*a = AccountTypeVisa
	case "MasterCard":
		*a = AccountTypeMasterCard
	case "AmericanExpress":
		*a = AccountTypeAmericanExpress
	case "Discover":
		*a = AccountTypeDiscover
	case "JCB":
		*a = AccountTypeJCB
	case "DinersClub":
		*a = AccountTypeDinersClub
	case "eCheck":
		*a = AccountTypeECheck
	default:
		return fmt.Errorf("invalid AccountTypeEnum : %q", string(data))
	}
	return nil
}

type CustomerProfileTypeEnum uint

const (
	CustomerProfileRegular CustomerProfileTypeEnum = 1
	CustomerProfileGuest   CustomerProfileTypeEnum = 2
)

func (a CustomerProfileTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case CustomerProfileRegular:
		return json.Marshal("regular")
	case CustomerProfileGuest:
		return json.Marshal("guest")
	default:
		return json.Marshal("null")
	}
}

func (a *CustomerProfileTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid CustomerProfileTypeEnum (empty)")

	}
	switch string(data[1 : len(data)-1]) {
	case "regular":
		*a = CustomerProfileRegular
	case "guest":
		*a = CustomerProfileGuest
	default:
		return fmt.Errorf("invalid CustomerProfileTypeEnum : %q", string(data))
	}
	return nil
}

type MerchantInitTransReasonEnum uint

const (
	ReasonResubmission    MerchantInitTransReasonEnum = 1
	ReasonDelayedCharg    MerchantInitTransReasonEnum = 2
	ReasonReauthorization MerchantInitTransReasonEnum = 3
	ReasonNoShow          MerchantInitTransReasonEnum = 4
)

func (a MerchantInitTransReasonEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case ReasonResubmission:
		return json.Marshal("resubmission")
	case ReasonDelayedCharg:
		return json.Marshal("delayedCharge")
	case ReasonReauthorization:
		return json.Marshal("reauthorization")
	case ReasonNoShow:
		return json.Marshal("noShow")
	default:
		return json.Marshal("null")
	}
}

func (a *MerchantInitTransReasonEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid MerchantInitTransReasonEnum (empty)")

	}
	switch string(data[1 : len(data)-1]) {
	case "resubmission":
		*a = ReasonResubmission
	case "delayedCharge":
		*a = ReasonDelayedCharg
	case "reauthorization":
		*a = ReasonReauthorization
	case "noShow":
		*a = ReasonNoShow
	default:
		return fmt.Errorf("invalid MerchantInitTransReasonEnum : %q", string(data))
	}
	return nil
}

type AfdsTransactionEnum uint

const (
	AFSDApprove AfdsTransactionEnum = 1
	ASFDDecline AfdsTransactionEnum = 2
)

func (a AfdsTransactionEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case AFSDApprove:
		return json.Marshal("approve")
	case ASFDDecline:
		return json.Marshal("decline")
	default:
		return json.Marshal("null")
	}
}

func (a *AfdsTransactionEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid AfdsTransactionEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "approve":
		*a = AFSDApprove
	case "decline":
		*a = ASFDDecline
	default:
		return fmt.Errorf("invalid AfdsTransactionEnum : %q", string(data))
	}
	return nil
}

type WebCheckOutTypeEnum uint

const (
	WebCheckOutTypeEnumPAN   WebCheckOutTypeEnum = 1
	WebCheckOutTypeEnumTOKEN WebCheckOutTypeEnum = 2
)

func (a WebCheckOutTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case WebCheckOutTypeEnumPAN:
		return json.Marshal("PAN")
	case WebCheckOutTypeEnumTOKEN:
		return json.Marshal("TOKEN")
	default:
		return json.Marshal("null")
	}
}

func (a *WebCheckOutTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid WebCheckOutTypeEnum (empty)")

	}
	switch string(data[1 : len(data)-1]) {
	case "PAN":
		*a = WebCheckOutTypeEnumPAN
	case "TOKEN":
		*a = WebCheckOutTypeEnumTOKEN
	default:
		return fmt.Errorf("invalid WebCheckOutTypeEnum : %q", string(data))
	}
	return nil
}

type ARBSubscriptionStatusEnum uint

const (
	ARBSubscriptionActive     ARBSubscriptionStatusEnum = 1
	ARBSubscriptionExpired    ARBSubscriptionStatusEnum = 2
	ARBSubscriptionSuspended  ARBSubscriptionStatusEnum = 3
	ARBSubscriptionCanceled   ARBSubscriptionStatusEnum = 4
	ARBSubscriptionTerminated ARBSubscriptionStatusEnum = 5
)

func (a ARBSubscriptionStatusEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case ARBSubscriptionActive:
		return json.Marshal("active")
	case ARBSubscriptionExpired:
		return json.Marshal("expired")
	case ARBSubscriptionSuspended:
		return json.Marshal("suspended")
	case ARBSubscriptionCanceled:
		return json.Marshal("canceled")
	case ARBSubscriptionTerminated:
		return json.Marshal("terminated")
	default:
		return json.Marshal("null")
	}
}

func (a *ARBSubscriptionStatusEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid ARBSubscriptionStatusEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "active":
		*a = ARBSubscriptionActive
	case "expired":
		*a = ARBSubscriptionExpired
	case "suspended":
		*a = ARBSubscriptionSuspended
	case "canceled":
		*a = ARBSubscriptionCanceled
	case "terminated":
		*a = ARBSubscriptionTerminated
	default:
		return fmt.Errorf("invalid ARBSubscriptionStatusEnum : %q", string(data))
	}
	return nil
}

type ValidationModeEnum uint

const (
	ValidationModeNone        ValidationModeEnum = 1
	ValidationModeTestMode    ValidationModeEnum = 2
	ValidationModeLiveMode    ValidationModeEnum = 3
	ValidationModeOldLiveMode ValidationModeEnum = 4
)

func (a ValidationModeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case ValidationModeNone:
		return json.Marshal("none")
	case ValidationModeTestMode:
		return json.Marshal("testMode")
	case ValidationModeLiveMode:
		return json.Marshal("liveMode")
	case ValidationModeOldLiveMode:
		return json.Marshal("oldLiveMode")
	default:
		return json.Marshal("null")
	}
}

func (a *ValidationModeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid ValidationModeEnum (empty)")

	}
	switch string(data[1 : len(data)-1]) {
	case "none":
		*a = ValidationModeNone
	case "testMode":
		*a = ValidationModeTestMode
	case "liveMode":
		*a = ValidationModeLiveMode
	case "oldLiveMode":
		*a = ValidationModeOldLiveMode
	default:
		return fmt.Errorf("invalid ValidationModeEnum : %q", string(data))
	}
	return nil
}

type CustomerPaymentProfileSearchTypeEnum uint

const (
	CardsExpiringInMonth CustomerPaymentProfileSearchTypeEnum = 1
)

func (a CustomerPaymentProfileSearchTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case CardsExpiringInMonth:
		return json.Marshal("none")

	default:
		return json.Marshal("null")
	}
}

func (a *CustomerPaymentProfileSearchTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid CustomerPaymentProfileSearchTypeEnum (empty)")

	}
	switch string(data[1 : len(data)-1]) {
	case "none":
		*a = CardsExpiringInMonth
	default:
		return fmt.Errorf("invalid CustomerPaymentProfileSearchTypeEnum : %q", string(data))
	}
	return nil
}

type SplitTenderStatusEnum uint

const (
	SplitTenderCompleted SplitTenderStatusEnum = 1
	SplitTenderHeld      SplitTenderStatusEnum = 2
	SplitTenderVoided    SplitTenderStatusEnum = 3
)

func (a SplitTenderStatusEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case SplitTenderCompleted:
		return json.Marshal("completed")
	case SplitTenderHeld:
		return json.Marshal("held")
	case SplitTenderVoided:
		return json.Marshal("voided")
	default:
		return json.Marshal("null")
	}
}

func (a *SplitTenderStatusEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid SplitTenderStatusEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "completed":
		*a = SplitTenderCompleted
	case "held":
		*a = SplitTenderHeld
	case "voided":
		*a = SplitTenderVoided
	default:
		return fmt.Errorf("invalid SplitTenderStatusEnum : %q", string(data))
	}
	return nil
}

type SettingNameEnum uint

const (
	SettingEmailCustomer                       SettingNameEnum = 1
	SettingMerchantEmail                       SettingNameEnum = 2
	SettingAllowPartialAuth                    SettingNameEnum = 3
	SettingHeaderEmailReceipt                  SettingNameEnum = 4
	SettingFooterEmailReceipt                  SettingNameEnum = 5
	SettingRecurringBilling                    SettingNameEnum = 6
	SettingDuplicateWindow                     SettingNameEnum = 7
	SettingTestRequest                         SettingNameEnum = 8
	SettingHostedProfileReturnUrl              SettingNameEnum = 9
	SettingHostedProfileReturnUrlText          SettingNameEnum = 10
	SettingHostedProfilePageBorderVisible      SettingNameEnum = 11
	SettingHostedProfileIFrameCommunicatorUrl  SettingNameEnum = 12
	SettingHostedProfileHeadingBgColor         SettingNameEnum = 13
	SettingHostedProfileValidationMode         SettingNameEnum = 14
	SettingHostedProfileBillingAddressRequired SettingNameEnum = 15
	SettingHostedProfileCardCodeRequired       SettingNameEnum = 16
	SettingHostedProfileBillingAddressOptions  SettingNameEnum = 17
	SettingHostedProfileManageOptions          SettingNameEnum = 18
	SettingHostedPaymentIFrameCommunicatorUrl  SettingNameEnum = 19
	SettingHostedPaymentButtonOptions          SettingNameEnum = 20
	SettingHostedPaymentReturnOptions          SettingNameEnum = 21
	SettingHostedPaymentOrderOptions           SettingNameEnum = 22
	SettingHostedPaymentPaymentOptions         SettingNameEnum = 23
	SettingHostedPaymentBillingAddressOptions  SettingNameEnum = 24
	SettingHostedPaymentShippingAddressOptions SettingNameEnum = 25
	SettingHostedPaymentSecurityOptions        SettingNameEnum = 26
	SettingHostedPaymentCustomerOptions        SettingNameEnum = 27
	SettingHostedPaymentStyleOptions           SettingNameEnum = 28
	SettingTypeEmailReceipt                    SettingNameEnum = 29
	SettingHostedProfilePaymentOptions         SettingNameEnum = 30
	SettingHostedProfileSaveButtonText         SettingNameEnum = 31
)

func (a SettingNameEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case SettingEmailCustomer:
		return json.Marshal("emailCustomer")
	case SettingMerchantEmail:
		return json.Marshal("merchantEmail")
	case SettingAllowPartialAuth:
		return json.Marshal("allowPartialAuth")
	case SettingHeaderEmailReceipt:
		return json.Marshal("headerEmailReceipt")
	case SettingFooterEmailReceipt:
		return json.Marshal("footerEmailReceipt")
	case SettingRecurringBilling:
		return json.Marshal("recurringBilling")
	case SettingDuplicateWindow:
		return json.Marshal("duplicateWindow")
	case SettingTestRequest:
		return json.Marshal("testRequest")
	case SettingHostedProfileReturnUrl:
		return json.Marshal("hostedProfileReturnUrl")
	case SettingHostedProfileReturnUrlText:
		return json.Marshal("hostedProfileReturnUrlText")
	case SettingHostedProfilePageBorderVisible:
		return json.Marshal("hostedProfilePageBorderVisible")
	case SettingHostedProfileIFrameCommunicatorUrl:
		return json.Marshal("hostedProfileIFrameCommunicatorUrl")
	case SettingHostedProfileHeadingBgColor:
		return json.Marshal("hostedProfileHeadingBgColor")
	case SettingHostedProfileValidationMode:
		return json.Marshal("hostedProfileValidationMode")
	case SettingHostedProfileBillingAddressRequired:
		return json.Marshal("hostedProfileBillingAddressRequired")
	case SettingHostedProfileCardCodeRequired:
		return json.Marshal("hostedProfileCardCodeRequired")
	case SettingHostedProfileBillingAddressOptions:
		return json.Marshal("hostedProfileBillingAddressOptions")
	case SettingHostedProfileManageOptions:
		return json.Marshal("hostedProfileManageOptions")
	case SettingHostedPaymentIFrameCommunicatorUrl:
		return json.Marshal("hostedPaymentIFrameCommunicatorUrl")
	case SettingHostedPaymentButtonOptions:
		return json.Marshal("hostedPaymentButtonOptions")
	case SettingHostedPaymentReturnOptions:
		return json.Marshal("hostedPaymentReturnOptions")
	case SettingHostedPaymentOrderOptions:
		return json.Marshal("hostedPaymentOrderOptions")
	case SettingHostedPaymentPaymentOptions:
		return json.Marshal("hostedPaymentPaymentOptions")
	case SettingHostedPaymentBillingAddressOptions:
		return json.Marshal("hostedPaymentBillingAddressOptions")
	case SettingHostedPaymentShippingAddressOptions:
		return json.Marshal("hostedPaymentShippingAddressOptions")
	case SettingHostedPaymentSecurityOptions:
		return json.Marshal("hostedPaymentSecurityOptions")
	case SettingHostedPaymentCustomerOptions:
		return json.Marshal("hostedPaymentCustomerOptions")
	case SettingHostedPaymentStyleOptions:
		return json.Marshal("hostedPaymentStyleOptions")
	case SettingTypeEmailReceipt:
		return json.Marshal("typeEmailReceipt")
	case SettingHostedProfilePaymentOptions:
		return json.Marshal("hostedProfilePaymentOptions")
	case SettingHostedProfileSaveButtonText:
		return json.Marshal("hostedProfileSaveButtonText")
	default:
		return json.Marshal("null")
	}
}

func (a *SettingNameEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid SettingNameEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "emailCustomer":
		*a = SettingEmailCustomer
	case "merchantEmail":
		*a = SettingMerchantEmail
	case "allowPartialAuth":
		*a = SettingAllowPartialAuth
	case "headerEmailReceipt":
		*a = SettingHeaderEmailReceipt
	case "footerEmailReceipt":
		*a = SettingFooterEmailReceipt
	case "recurringBilling":
		*a = SettingRecurringBilling
	case "duplicateWindow":
		*a = SettingDuplicateWindow
	case "testRequest":
		*a = SettingTestRequest
	case "hostedProfileReturnUrl":
		*a = SettingHostedProfileReturnUrl
	case "hostedProfileReturnUrlText":
		*a = SettingHostedProfileReturnUrlText
	case "hostedProfilePageBorderVisible":
		*a = SettingHostedProfilePageBorderVisible
	case "hostedProfileIFrameCommunicatorUrl":
		*a = SettingHostedProfileIFrameCommunicatorUrl
	case "hostedProfileHeadingBgColor":
		*a = SettingHostedProfileHeadingBgColor
	case "hostedProfileValidationMode":
		*a = SettingHostedProfileValidationMode
	case "hostedProfileBillingAddressRequired":
		*a = SettingHostedProfileBillingAddressRequired
	case "hostedProfileCardCodeRequired":
		*a = SettingHostedProfileCardCodeRequired
	case "hostedProfileBillingAddressOptions":
		*a = SettingHostedProfileBillingAddressOptions
	case "hostedProfileManageOptions":
		*a = SettingHostedProfileManageOptions
	case "hostedPaymentIFrameCommunicatorUrl":
		*a = SettingHostedPaymentIFrameCommunicatorUrl
	case "hostedPaymentButtonOptions":
		*a = SettingHostedPaymentButtonOptions
	case "hostedPaymentReturnOptions":
		*a = SettingHostedPaymentReturnOptions
	case "hostedPaymentOrderOptions":
		*a = SettingHostedPaymentOrderOptions
	case "hostedPaymentPaymentOptions":
		*a = SettingHostedPaymentPaymentOptions
	case "hostedPaymentBillingAddressOptions":
		*a = SettingHostedPaymentBillingAddressOptions
	case "hostedPaymentShippingAddressOptions":
		*a = SettingHostedPaymentShippingAddressOptions
	case "hostedPaymentSecurityOptions":
		*a = SettingHostedPaymentSecurityOptions
	case "hostedPaymentCustomerOptions":
		*a = SettingHostedPaymentCustomerOptions
	case "hostedPaymentStyleOptions":
		*a = SettingHostedPaymentStyleOptions
	case "typeEmailReceipt":
		*a = SettingTypeEmailReceipt
	case "hostedProfilePaymentOptions":
		*a = SettingHostedProfilePaymentOptions
	case "hostedProfileSaveButtonText":
		*a = SettingHostedProfileSaveButtonText
	default:
		return fmt.Errorf("invalid SettingNameEnum : %q", string(data))
	}
	return nil
}

type EncryptionAlgorithmType uint

const (
	EncryptionAlgoTDES EncryptionAlgorithmType = 1
	EncryptionAlgoAES  EncryptionAlgorithmType = 2
	EncryptionAlgoRSA  EncryptionAlgorithmType = 3
)

func (a EncryptionAlgorithmType) MarshalJSON() ([]byte, error) {
	switch a {
	case EncryptionAlgoTDES:
		return json.Marshal("TDES")
	case EncryptionAlgoAES:
		return json.Marshal("AES")
	case EncryptionAlgoRSA:
		return json.Marshal("RSA")
	default:
		return json.Marshal("null")
	}
}

func (a *EncryptionAlgorithmType) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid EncryptionAlgorithmType (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "TDES":
		*a = EncryptionAlgoTDES
	case "AES":
		*a = EncryptionAlgoAES
	case "RSA":
		*a = EncryptionAlgoRSA
	default:
		return fmt.Errorf("invalid EncryptionAlgorithmType : %q", string(data))
	}
	return nil
}

type FDSFilterActionEnum uint

const (
	FDSFilterActionReject      FDSFilterActionEnum = 1
	FDSFilterActionDecline     FDSFilterActionEnum = 2
	FDSFilterActionHold        FDSFilterActionEnum = 3
	FDSFilterActionAuthAndHold FDSFilterActionEnum = 4
	FDSFilterActionReport      FDSFilterActionEnum = 5
)

func (a FDSFilterActionEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case FDSFilterActionReject:
		return json.Marshal("reject")
	case FDSFilterActionDecline:
		return json.Marshal("decline")
	case FDSFilterActionHold:
		return json.Marshal("hold")
	case FDSFilterActionAuthAndHold:
		return json.Marshal("authAndHold")
	case FDSFilterActionReport:
		return json.Marshal("report")
	default:
		return json.Marshal("null")
	}
}

func (a *FDSFilterActionEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid FDSFilterActionEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "reject":
		*a = FDSFilterActionReject
	case "decline":
		*a = FDSFilterActionDecline
	case "hold":
		*a = FDSFilterActionHold
	case "authAndHold":
		*a = FDSFilterActionAuthAndHold
	case "report":
		*a = FDSFilterActionReport
	default:
		return fmt.Errorf("invalid FDSFilterActionEnum : %q", string(data))
	}
	return nil
}

type PermissionsEnum uint

const (
	PermissionAPIMerchantBasicReporting PermissionsEnum = 1
	PermissionSubmitCharge              PermissionsEnum = 2
	PermissionSubmitRefund              PermissionsEnum = 3
	PermissionSubmitUpdate              PermissionsEnum = 4
	PermissionMobileAdmin               PermissionsEnum = 5
)

func (a PermissionsEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case PermissionAPIMerchantBasicReporting:
		return json.Marshal("API_Merchant_BasicReporting")
	case PermissionSubmitCharge:
		return json.Marshal("Submit_Charge")
	case PermissionSubmitRefund:
		return json.Marshal("Submit_Refund")
	case PermissionSubmitUpdate:
		return json.Marshal("Submit_Update")
	case PermissionMobileAdmin:
		return json.Marshal("Mobile_Admin")
	default:
		return json.Marshal("null")
	}
}

func (a *PermissionsEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid PermissionsEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "API_Merchant_BasicReporting":
		*a = PermissionAPIMerchantBasicReporting
	case "Submit_Charge":
		*a = PermissionSubmitCharge
	case "Submit_Refund":
		*a = PermissionSubmitRefund
	case "Submit_Update":
		*a = PermissionSubmitUpdate
	case "Mobile_Admin":
		*a = PermissionMobileAdmin
	default:
		return fmt.Errorf("invalid PermissionsEnum : %q", string(data))
	}
	return nil
}

type OperationType uint

const (
	OperationTypeDECRYPT OperationType = 1
)

func (a OperationType) MarshalJSON() ([]byte, error) {
	switch a {
	case OperationTypeDECRYPT:
		return json.Marshal("DECRYPT")
	default:
		return json.Marshal("null")
	}
}

func (a *OperationType) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid OperationType (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "DECRYPT":
		*a = OperationTypeDECRYPT
	default:
		return fmt.Errorf("invalid OperationType : %q", string(data))
	}
	return nil
}

type TransactionStatusEnum uint

const (
	TransactionAuthorizedPendingCapture   TransactionStatusEnum = 1
	TransactionCapturedPendingSettlement  TransactionStatusEnum = 2
	TransactionCommunicationError         TransactionStatusEnum = 3
	TransactionRefundSettledSuccessfully  TransactionStatusEnum = 4
	TransactionRefundPendingSettlement    TransactionStatusEnum = 5
	TransactionApprovedReview             TransactionStatusEnum = 6
	TransactionDeclined                   TransactionStatusEnum = 7
	TransactionCouldNotVoid               TransactionStatusEnum = 8
	TransactionExpired                    TransactionStatusEnum = 9
	TransactionGeneralError               TransactionStatusEnum = 10
	TransactionPendingFinalSettlement     TransactionStatusEnum = 11
	TransactionPendingSettlement          TransactionStatusEnum = 12
	TransactionFailedReview               TransactionStatusEnum = 13
	TransactionSettledSuccessfully        TransactionStatusEnum = 14
	TransactionSettlementError            TransactionStatusEnum = 15
	TransactionUnderReview                TransactionStatusEnum = 16
	TransactionUpdatingSettlement         TransactionStatusEnum = 17
	TransactionVoided                     TransactionStatusEnum = 18
	TransactionFDSPendingReview           TransactionStatusEnum = 19
	TransactionFDSAuthorizedPendingReview TransactionStatusEnum = 20
	TransactionReturnedItem               TransactionStatusEnum = 21
	TransactionChargeback                 TransactionStatusEnum = 22
	TransactionChargebackReversal         TransactionStatusEnum = 23
	TransactionAuthorizedPendingRelease   TransactionStatusEnum = 24
)

func (a TransactionStatusEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case TransactionAuthorizedPendingCapture:
		return json.Marshal("authorizedPendingCapture")
	case TransactionCapturedPendingSettlement:
		return json.Marshal("capturedPendingSettlement")
	case TransactionCommunicationError:
		return json.Marshal("communicationError")
	case TransactionRefundSettledSuccessfully:
		return json.Marshal("refundSettledSuccessfully")
	case TransactionRefundPendingSettlement:
		return json.Marshal("refundPendingSettlement")
	case TransactionApprovedReview:
		return json.Marshal("approvedReview")
	case TransactionDeclined:
		return json.Marshal("declined")
	case TransactionCouldNotVoid:
		return json.Marshal("couldNotVoid")
	case TransactionExpired:
		return json.Marshal("expired")
	case TransactionGeneralError:
		return json.Marshal("generalError")
	case TransactionPendingFinalSettlement:
		return json.Marshal("pendingFinalSettlement")
	case TransactionPendingSettlement:
		return json.Marshal("pendingSettlement")
	case TransactionFailedReview:
		return json.Marshal("failedReview")
	case TransactionSettledSuccessfully:
		return json.Marshal("settledSuccessfully")
	case TransactionSettlementError:
		return json.Marshal("settlementError")
	case TransactionUnderReview:
		return json.Marshal("underReview")
	case TransactionUpdatingSettlement:
		return json.Marshal("updatingSettlement")
	case TransactionVoided:
		return json.Marshal("voided")
	case TransactionFDSPendingReview:
		return json.Marshal("FDSPendingReview")
	case TransactionFDSAuthorizedPendingReview:
		return json.Marshal("FDSAuthorizedPendingReview")
	case TransactionReturnedItem:
		return json.Marshal("returnedItem")
	case TransactionChargeback:
		return json.Marshal("chargeback")
	case TransactionChargebackReversal:
		return json.Marshal("chargebackReversal")
	case TransactionAuthorizedPendingRelease:
		return json.Marshal("authorizedPendingRelease")
	default:
		return json.Marshal("null")
	}
}

func (a *TransactionStatusEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid TransactionStatusEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "authorizedPendingCapture":
		*a = TransactionAuthorizedPendingCapture
	case "capturedPendingSettlement":
		*a = TransactionCapturedPendingSettlement
	case "communicationError":
		*a = TransactionCommunicationError
	case "refundSettledSuccessfully":
		*a = TransactionRefundSettledSuccessfully
	case "refundPendingSettlement":
		*a = TransactionRefundPendingSettlement
	case "approvedReview":
		*a = TransactionApprovedReview
	case "declined":
		*a = TransactionDeclined
	case "couldNotVoid":
		*a = TransactionCouldNotVoid
	case "expired":
		*a = TransactionExpired
	case "generalError":
		*a = TransactionGeneralError
	case "pendingFinalSettlement":
		*a = TransactionPendingFinalSettlement
	case "pendingSettlement":
		*a = TransactionPendingSettlement
	case "failedReview":
		*a = TransactionFailedReview
	case "settledSuccessfully":
		*a = TransactionSettledSuccessfully
	case "settlementError":
		*a = TransactionSettlementError
	case "underReview":
		*a = TransactionUnderReview
	case "updatingSettlement":
		*a = TransactionUpdatingSettlement
	case "voided":
		*a = TransactionVoided
	case "FDSPendingReview":
		*a = TransactionFDSPendingReview
	case "FDSAuthorizedPendingReview":
		*a = TransactionFDSAuthorizedPendingReview
	case "returnedItem":
		*a = TransactionReturnedItem
	case "chargeback":
		*a = TransactionChargeback
	case "chargebackReversal":
		*a = TransactionChargebackReversal
	case "authorizedPendingRelease":
		*a = TransactionAuthorizedPendingRelease
	default:
		return fmt.Errorf("invalid TransactionStatusEnum : %q", string(data))
	}
	return nil
}

type ARBGetSubscriptionListOrderFieldEnum uint

const (
	ARBGetSubscriptionListOrderFieldId                 ARBGetSubscriptionListOrderFieldEnum = 1
	ARBGetSubscriptionListOrderFieldName               ARBGetSubscriptionListOrderFieldEnum = 2
	ARBGetSubscriptionListOrderFieldStatus             ARBGetSubscriptionListOrderFieldEnum = 3
	ARBGetSubscriptionListOrderFieldCreateTimeStampUTC ARBGetSubscriptionListOrderFieldEnum = 4
	ARBGetSubscriptionListOrderFieldLastName           ARBGetSubscriptionListOrderFieldEnum = 5
	ARBGetSubscriptionListOrderFieldFirstName          ARBGetSubscriptionListOrderFieldEnum = 6
	ARBGetSubscriptionListOrderFieldAccountNumber      ARBGetSubscriptionListOrderFieldEnum = 7
	ARBGetSubscriptionListOrderFieldAmount             ARBGetSubscriptionListOrderFieldEnum = 8
	ARBGetSubscriptionListOrderFieldPastOccurrences    ARBGetSubscriptionListOrderFieldEnum = 9
)

func (a ARBGetSubscriptionListOrderFieldEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case ARBGetSubscriptionListOrderFieldId:
		return json.Marshal("id")
	case ARBGetSubscriptionListOrderFieldName:
		return json.Marshal("name")
	case ARBGetSubscriptionListOrderFieldStatus:
		return json.Marshal("status")
	case ARBGetSubscriptionListOrderFieldCreateTimeStampUTC:
		return json.Marshal("createTimeStampUTC")
	case ARBGetSubscriptionListOrderFieldLastName:
		return json.Marshal("lastName")
	case ARBGetSubscriptionListOrderFieldFirstName:
		return json.Marshal("firstName")
	case ARBGetSubscriptionListOrderFieldAccountNumber:
		return json.Marshal("accountNumber")
	case ARBGetSubscriptionListOrderFieldAmount:
		return json.Marshal("amount")
	case ARBGetSubscriptionListOrderFieldPastOccurrences:
		return json.Marshal("pastOccurrences")
	default:
		return json.Marshal("null")
	}
}

func (a *ARBGetSubscriptionListOrderFieldEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid ARBGetSubscriptionListOrderFieldEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "id":
		*a = ARBGetSubscriptionListOrderFieldId
	case "name":
		*a = ARBGetSubscriptionListOrderFieldName
	case "status":
		*a = ARBGetSubscriptionListOrderFieldStatus
	case "createTimeStampUTC":
		*a = ARBGetSubscriptionListOrderFieldCreateTimeStampUTC
	case "lastName":
		*a = ARBGetSubscriptionListOrderFieldLastName
	case "firstName":
		*a = ARBGetSubscriptionListOrderFieldFirstName
	case "accountNumber":
		*a = ARBGetSubscriptionListOrderFieldAccountNumber
	case "amount":
		*a = ARBGetSubscriptionListOrderFieldAmount
	case "pastOccurrences":
		*a = ARBGetSubscriptionListOrderFieldPastOccurrences

	default:
		return fmt.Errorf("invalid ARBGetSubscriptionListOrderFieldEnum : %q", string(data))
	}
	return nil
}

type CustomerPaymentProfileOrderFieldEnum uint

const (
	CustomerPaymentProfileOrderFieldID CustomerPaymentProfileOrderFieldEnum = 1
)

func (a CustomerPaymentProfileOrderFieldEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case CustomerPaymentProfileOrderFieldID:
		return json.Marshal("id")
	default:
		return json.Marshal("null")
	}
}

func (a *CustomerPaymentProfileOrderFieldEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid CustomerPaymentProfileOrderFieldEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "id":
		*a = CustomerPaymentProfileOrderFieldID
	default:
		return fmt.Errorf("invalid CustomerPaymentProfileOrderFieldEnum : %q", string(data))
	}
	return nil
}

type CustomerTypeEnum uint

const (
	Individual CustomerTypeEnum = 1
	Business   CustomerTypeEnum = 2
)

func (a CustomerTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case Individual:
		return json.Marshal("individual")
	case Business:
		return json.Marshal("business")
	default:
		return json.Marshal("null")
	}
}

func (a *CustomerTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid CustomerTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "individual":
		*a = Individual
	case "business":
		*a = Business

	default:
		return fmt.Errorf("invalid CustomerTypeEnum : %q", string(data))
	}
	return nil
}

type ARBSubscriptionUnitEnum uint

const (
	Days   ARBSubscriptionUnitEnum = 1
	Months ARBSubscriptionUnitEnum = 2
)

func (a ARBSubscriptionUnitEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case Days:
		return json.Marshal("days")
	case Months:
		return json.Marshal("months")
	default:
		return json.Marshal("null")
	}
}

func (a *ARBSubscriptionUnitEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid ARBSubscriptionUnitEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "days":
		*a = Days
	case "months":
		*a = Months

	default:
		return fmt.Errorf("invalid ARBSubscriptionUnitEnum : %q", string(data))
	}
	return nil
}

type TransactionGroupStatusEnum uint

const (
	TransactionGroupAny             TransactionGroupStatusEnum = 1
	TransactionGroupPendingApproval TransactionGroupStatusEnum = 2
)

func (a TransactionGroupStatusEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case TransactionGroupAny:
		return json.Marshal("any")
	case TransactionGroupPendingApproval:
		return json.Marshal("pendingApproval")
	default:
		return json.Marshal("null")
	}
}

func (a *TransactionGroupStatusEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid TransactionGroupStatusEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "any":
		*a = TransactionGroupAny
	case "pendingApproval":
		*a = TransactionGroupPendingApproval
	default:
		return fmt.Errorf("invalid TransactionGroupStatusEnum : %q", string(data))
	}
	return nil
}

type ARBGetSubscriptionListSearchTypeEnum uint

const (
	CardExpiringThisMonth         ARBGetSubscriptionListSearchTypeEnum = 1
	SubscriptionActive            ARBGetSubscriptionListSearchTypeEnum = 2
	SubscriptionExpiringThisMonth ARBGetSubscriptionListSearchTypeEnum = 3
	SubscriptionInactive          ARBGetSubscriptionListSearchTypeEnum = 4
)

func (a ARBGetSubscriptionListSearchTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case CardExpiringThisMonth:
		return json.Marshal("cardExpiringThisMonth")
	case SubscriptionActive:
		return json.Marshal("subscriptionActive")
	case SubscriptionExpiringThisMonth:
		return json.Marshal("subscriptionExpiringThisMonth")
	case SubscriptionInactive:
		return json.Marshal("subscriptionInactive")
	default:
		return json.Marshal("null")
	}
}

func (a *ARBGetSubscriptionListSearchTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid ARBGetSubscriptionListSearchTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "cardExpiringThisMonth":
		*a = CardExpiringThisMonth
	case "subscriptionActive":
		*a = SubscriptionActive
	case "subscriptionExpiringThisMonth":
		*a = SubscriptionExpiringThisMonth
	case "subscriptionInactive":
		*a = SubscriptionInactive
	default:
		return fmt.Errorf("invalid ARBGetSubscriptionListSearchTypeEnum : %q", string(data))
	}
	return nil
}

type TransactionListOrderFieldEnum uint

const (
	TransactionListFieldID       TransactionListOrderFieldEnum = 1
	TransactionListSubmitTimeUTC TransactionListOrderFieldEnum = 2
)

func (a TransactionListOrderFieldEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case TransactionListFieldID:
		return json.Marshal("id")
	case TransactionListSubmitTimeUTC:
		return json.Marshal("submitTimeUTC")
	default:
		return json.Marshal("null")
	}
}

func (a *TransactionListOrderFieldEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid TransactionListOrderFieldEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "id":
		*a = TransactionListFieldID
	case "submitTimeUTC":
		*a = TransactionListSubmitTimeUTC
	default:
		return fmt.Errorf("invalid TransactionListOrderFieldEnum : %q", string(data))
	}
	return nil
}

type DeviceActivationEnum uint

const (
	Activate DeviceActivationEnum = 1
	Disable  DeviceActivationEnum = 2
)

func (a DeviceActivationEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case Activate:
		return json.Marshal("Activate")
	case Disable:
		return json.Marshal("Disable")
	default:
		return json.Marshal("null")
	}
}

func (a *DeviceActivationEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid DeviceActivationEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "Activate":
		*a = Activate
	case "Disable":
		*a = Disable

	default:
		return fmt.Errorf("invalid DeviceActivationEnum : %q", string(data))
	}
	return nil
}

type PaymentMethodEnum uint

const (
	CC        PaymentMethodEnum = 1
	Check     PaymentMethodEnum = 2
	WebPayPal PaymentMethodEnum = 3
)

func (a PaymentMethodEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case CC:
		return json.Marshal("creditCard")
	case Check:
		return json.Marshal("eCheck")
	case WebPayPal:
		return json.Marshal("payPal")
	default:
		return json.Marshal("null")
	}
}

func (a *PaymentMethodEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid PaymentMethodEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "creditCard":
		*a = CC
	case "eCheck":
		*a = Check
	case "payPal":
		*a = WebPayPal

	default:
		return fmt.Errorf("invalid PaymentMethodEnum : %q", string(data))
	}
	return nil
}

type SettlementStateEnum uint

const (
	StateSettledSuccessfully SettlementStateEnum = 1
	StateSettlementError     SettlementStateEnum = 2
	StatePendingSettlement   SettlementStateEnum = 3
)

func (a SettlementStateEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case StateSettledSuccessfully:
		return json.Marshal("settledSuccessfully")
	case StateSettlementError:
		return json.Marshal("settlementError")
	case StatePendingSettlement:
		return json.Marshal("pendingSettlement")
	default:
		return json.Marshal("null")
	}
}

func (a *SettlementStateEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid SettlementStateEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "settledSuccessfully":
		*a = StateSettledSuccessfully
	case "settlementError":
		*a = StateSettlementError
	case "pendingSettlement":
		*a = StatePendingSettlement
	default:
		return fmt.Errorf("invalid SettlementStateEnum : %q", string(data))
	}
	return nil
}

type EncodingType uint

const (
	EncBase64 EncodingType = 1
	EncHex    EncodingType = 2
)

func (a EncodingType) MarshalJSON() ([]byte, error) {
	switch a {
	case EncBase64:
		return json.Marshal("Base64")
	case EncHex:
		return json.Marshal("Hex")
	default:
		return json.Marshal("null")
	}
}

func (a *EncodingType) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid EncodingType (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "Base64":
		*a = EncBase64
	case "Hex":
		*a = EncHex

	default:
		return fmt.Errorf("invalid EncodingType : %q", string(data))
	}
	return nil
}

type AUJobTypeEnum uint

const (
	AUJobAll     AUJobTypeEnum = 1
	AUJobUpdates AUJobTypeEnum = 2
	AUJobDeletes AUJobTypeEnum = 3
)

func (a AUJobTypeEnum) MarshalJSON() ([]byte, error) {
	switch a {
	case AUJobAll:
		return json.Marshal("all")
	case AUJobUpdates:
		return json.Marshal("updates")
	case AUJobDeletes:
		return json.Marshal("deletes")
	default:
		return json.Marshal("null")
	}
}

func (a *AUJobTypeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid AUJobTypeEnum (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "all":
		*a = AUJobAll
	case "updates":
		*a = AUJobUpdates
	case "deletes":
		*a = AUJobDeletes

	default:
		return fmt.Errorf("invalid AUJobTypeEnum : %q", string(data))
	}
	return nil
}
