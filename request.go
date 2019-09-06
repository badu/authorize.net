package authorize_net

import (
	"fmt"
	"time"
)

type FingerPrintType struct {
	HashValue    string `json:"hashValue"`
	Timestamp    string `json:"timestamp"`
	Sequence     string `json:"sequence,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
	Amount       string `json:"amount,omitempty"`
}

type ImpersonationAuthenticationType struct {
	PartnerLoginId        string `json:"partnerLoginId"`
	PartnerTransactionKey string `json:"partnerTransactionKey"`
}

type MerchantAuthentication struct {
	Name                        string                           `json:"name"`           // Merchant’s unique API Login ID.
	TransactionKey              string                           `json:"transactionKey"` // Merchant’s unique Transaction Key.
	Password                    string                           `json:"password,omitempty"`
	SessionToken                string                           `json:"sessionToken,omitempty"`
	ImpersonationAuthentication *ImpersonationAuthenticationType `json:"impersonationAuthentication,omitempty"`
	FingerPrint                 *FingerPrintType                 `json:"fingerPrint,omitempty"`
	ClientKey                   string                           `json:"clientKey,omitempty"`
	AccessToken                 string                           `json:"accessToken,omitempty"`
	MobileDeviceId              string                           `json:"mobileDeviceId,omitempty"`
}

type ANetApiRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"` // Contains merchant authentication information.
	ClientId               string                 `json:"clientId,omitempty"`
	RefId                  string                 `json:"refId,omitempty"` // Merchant-assigned reference ID for the request. If included in the request, this value is included in the response. This feature might be especially useful for multi-threaded applications.
}

type Paging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Sorting struct {
	OrderBy         string `json:"orderBy"`
	OrderDescending bool   `json:"orderDescending"`
}

type Date struct {
	time.Time
}

func Now() Date {
	return Date{time.Now()}
}

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.UTC().Format("2006-01-02") + `"`), nil
}

/**
func (t *Date) UnmarshalJSON(data []byte) error {
	return nil
}
**/

const (
	// RFC3339 a subset of the ISO8601 timestamp format. e.g 2014-04-29T18:30:38Z
	ISO8601TimeFormat = "2006-01-02T15:04:05Z"
	// same as above, but no ”Z”
	ISO8601NoZTimeFormat = "2006-01-02T15:04:05"
	// 2018-12-27T11:28:57.467
	ISO8601TimeNano = "2006-01-02T15:04:05.999999999"
)

type DateTime struct {
	time.Time
}

func NowTime() DateTime {
	return DateTime{time.Now()}
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.UTC().Format("2006-01-02") + `"`), nil
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	var err error
	if t.Time, err = time.Parse(ISO8601NoZTimeFormat, string(data[1:len(data)-1])); err != nil {
		return err
	}
	return nil
}

type DateTimeZ struct {
	time.Time
}

func NowTimeZ() DateTimeZ {
	return DateTimeZ{time.Now()}
}

func (t DateTimeZ) MarshalJSON() ([]byte, error) {
	//return []byte(`"` + t.Time.UTC().Format("2006-01-02") + `"`), nil
	return nil, fmt.Errorf("MarshalJSON on DateTimeZ not implemented")
}

func (t *DateTimeZ) UnmarshalJSON(data []byte) error {
	var err error
	if t.Time, err = time.Parse(ISO8601TimeFormat, string(data[1:len(data)-1])); err != nil {
		return err
	}
	return nil
}

type DateTimeNano struct {
	time.Time
}

func NowDateTimeNano() DateTimeNano {
	return DateTimeNano{time.Now()}
}

func (t DateTimeNano) MarshalJSON() ([]byte, error) {
	//return []byte(`"` + t.Time.UTC().Format("2006-01-02") + `"`), nil
	return nil, fmt.Errorf("MarshalJSON on DateTimeNano not implemented")
}

func (t *DateTimeNano) UnmarshalJSON(data []byte) error {
	var err error
	if t.Time, err = time.Parse(ISO8601TimeNano, string(data[1:len(data)-1])); err != nil {
		return err
	}
	return nil
}
