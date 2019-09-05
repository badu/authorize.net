package authorize_net

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	merchantId         = "5KP3u95bQpv"      // sandbox Authorize.net
	merchantSecret     = "346HZ32z3fP4hTG2" // sandbox Authorize.net
	cardExample        = "5424000000000015"
	cardExpirationDate = "2020-12"
	cardCode           = "999"
)

func ExampleAuthenticateTest() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))

	payload := AuthenticateTestRequest{
		ANetApiRequest: ANetApiRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           merchantId,
				TransactionKey: merchantSecret,
			},
		},
	}

	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response ANetApiResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%+v", response)
	// Output: {RefId: Messages:{ResultCode:Ok Messages:[{Code:Successful. Text:Successful. Description:}]} SessionToken:}
}

// Note that credit card information and bank account information are mutually exclusive, so you should not submit both.
func ExampleChargeCreditCard() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))

	payload := CreateTransactionRequest{
		CreateTransactionRequest: TransactionRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           merchantId,
				TransactionKey: merchantSecret,
			},
			Transaction: Transaction{
				TransactionType: AuthCaptureTransaction,
				Amount:          11.3,
				Payment: &Payment{
					CreditCard: &CreditCard{
						CardNumber:     cardExample,
						ExpirationDate: cardExpirationDate,
						CardCode:       cardCode,
					},
				},
			},
		},
	}

	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response CreateTransactionResponse
	if err := client.do(request, &response, false); err != nil {
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%+v", response.TransactionResponse.Messages)
	// Output: [{Code:This transaction has been approved. Text: Description:This transaction has been approved.}]
}

func ExampleRefund() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	voidPayload := CreateTransactionRequest{
		CreateTransactionRequest: TransactionRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           merchantId,
				TransactionKey: merchantSecret,
			},
			Transaction: Transaction{
				TransactionType: RefundTransaction,
				Amount:          123123,
				Payment: &Payment{
					CreditCard: &CreditCard{
						CardNumber:     cardExample,
						ExpirationDate: cardExpirationDate,
						CardCode:       cardCode,
					},
				},
			},
		},
	}

	refundRequest, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &voidPayload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	var voidResponse CreateTransactionResponse
	if err := client.do(refundRequest, &voidResponse, false); err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%+v", voidResponse.TransactionResponse.Messages)
	// Output: [{Code:This transaction has been approved. Text: Description:This transaction has been approved.}]

}

func ExampleVoid() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	amount := 10.31
	payload := CreateTransactionRequest{
		CreateTransactionRequest: TransactionRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           merchantId,
				TransactionKey: merchantSecret,
			},
			Transaction: Transaction{
				TransactionType: AuthCaptureTransaction,
				Amount:          amount,
				Payment: &Payment{
					CreditCard: &CreditCard{
						CardNumber:     cardExample,
						ExpirationDate: cardExpirationDate,
						CardCode:       cardCode,
					},
				},
			},
		},
	}

	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response CreateTransactionResponse
	if err := client.do(request, &response, false); err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	voidPayload := CreateTransactionRequest{
		CreateTransactionRequest: TransactionRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           merchantId,
				TransactionKey: merchantSecret,
			},
			Transaction: Transaction{
				TransactionType: VoidTransaction,
				Amount:          amount,
				Payment: &Payment{
					CreditCard: &CreditCard{
						CardNumber:     cardExample,
						ExpirationDate: cardExpirationDate,
						CardCode:       cardCode,
					},
				},
				RefTransId: response.TransactionResponse.TransId,
			},
		},
	}

	voidRequest, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &voidPayload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	var voidResponse CreateTransactionResponse
	if err := client.do(voidRequest, &voidResponse, false); err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%+v", voidResponse.TransactionResponse.Messages)
	// Output: [{Code:This transaction has been approved. Text: Description:This transaction has been approved.}]

}

func ExampleCreateCustomerProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := CreateCustomerProfileRequest{
		Payload: CreateCustomerProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			Profile: &CustomerProfile{
				CustomerProfileBase: CustomerProfileBase{
					MerchantCustomerId: merchantId,
				},
				PaymentProfiles: []CustomerPaymentProfile{
					{
						CustomerPaymentProfileBase: CustomerPaymentProfileBase{
							CustomerType: Individual,
							BillTo: &CustomerAddress{
								NameAndAddress: NameAndAddress{
									FirstName: "Popescu",
									LastName:  "Grigore",
									Country:   "Romania",
								},
								PhoneNumber: "030030020",
								Email:       "b@gmail.com",
							},
						},
						Payment: &Payment{
							CreditCard: &CreditCard{
								CardNumber:     cardExample,
								ExpirationDate: cardExpirationDate,
								CardCode:       cardCode,
								IsPaymentToken: false,
							},
						},
						DefaultPaymentProfile: true,
					},
				},
				ProfileType: CustomerProfileRegular,
			},
			ValidationMode: ValidationModeTestMode,
		},
	}

	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response CreateCustomerProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%+v", response)
	// Output: {ANetApiResponse:{RefId: Messages:{ResultCode:Error Messages:[{Code:A duplicate record already exists. Text:A duplicate record with ID 1920441543 already exists. Description:}]} SessionToken:} CustomerProfileId: CustomerPaymentProfileIdList:[] CustomerShippingAddressIdList:[] ValidationDirectResponseList:[1,1,1,(TESTMODE) This transaction has been approved.,000000,P,0,none,Test transaction for ValidateCustomerPaymentProfile.,1.00,CC,auth_only,5KP3u95bQpv,Popescu,Grigore,,,,,,Romania,030030020,,,,,,,,,,,0.00,0.00,0.00,FALSE,none,,,,,,,,,,,,,,XXXX0015,MasterCard,,,,,,,,,,,,,,,,,]}

}

func ExampleGetCustomerProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetCustomerProfileRequest{
		Payload: GetCustomerProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "1920441543",
			IncludeIssuerInfo: true,
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetCustomerProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%s\n", response.Profile.CustomerProfileId)
	for _, pp := range response.Profile.PaymentProfiles {
		fmt.Printf("%s", pp.CustomerPaymentProfileId)
	}
	// Output: 1920441543
	// 1833416626
}

func ExampleUpdateCustomerProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := UpdateCustomerProfileRequest{
		Payload: UpdateCustomerProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			Profile: &CustomerProfileInfoEx{
				CustomerProfileEx: CustomerProfileEx{
					CustomerProfileBase: CustomerProfileBase{
						MerchantCustomerId: "custId123",
						Description:        "description",
						Email:              "some@invalid@email",
					},
					CustomerProfileId: "38157432",
				},
				ProfileType: CustomerProfileRegular,
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response UpdateCustomerProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%#v", response)
	// Output: authorize_net.UpdateCustomerProfileResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}}

}
func ExampleDeleteCustomerProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := DeleteCustomerProfileRequest{
		Payload: DeleteCustomerProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "38157410",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response DeleteCustomerProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v\n", response)
	// Output: authorize_net.DeleteCustomerProfileResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}}
}

func ExampleGetCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetCustomerPaymentProfileRequest{
		Payload: GetCustomerPaymentProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId:        "1920441543",
			CustomerPaymentProfileId: "1833416626",
			IncludeIssuerInfo:        true,
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetCustomerPaymentProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%s\n", response.PaymentProfile.CustomerProfileId)
	fmt.Printf("%s\n", response.PaymentProfile.CustomerPaymentProfileId)
	// Output: 1920441543
	// 1833416626
}

func ExampleGetCustomerPaymentProfileList() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetCustomerPaymentProfileListRequest{
		Payload: GetCustomerPaymentProfileListPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			SearchType: "cardsExpiringInMonth", // always the same
			Month:      "2020-12",
			Sorting: Sorting{
				OrderBy:         "id",
				OrderDescending: false,
			},
			Paging: Paging{
				Limit:  10,
				Offset: 1,
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetCustomerPaymentProfileListResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%d\n", response.TotalNumInResultSet)
	for _, pp := range response.PaymentProfiles {
		fmt.Printf("%d\n", pp.CustomerProfileId)
	}

	// Output:118569
	// 37821321
	// 38155405
	// 38155971
	// 38157410
	// 38157423
	// 38157432
	// 38157450
	// 38157457
	// 38157567
	// 38157570
}

func ExampleGetCustomerProfileIds() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetCustomerProfileIdsRequest{
		ANetApiRequest: ANetApiRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           merchantId,
				TransactionKey: merchantSecret,
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetCustomerProfileIdsResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	for idx, id := range response.Ids {
		if idx > 5 {
			break
		}
		fmt.Printf("%s\n", id)
	}
	// Output: 36152127
	// 36152166
	// 36152181
	// 36596285
	// 36715809
	// 36763407
}

func ExampleGetMerchantDetails() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetMerchantDetailsRequest{
		ANetApiRequest: ANetApiRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           merchantId,
				TransactionKey: merchantSecret,
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetMerchantDetailsResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%#v", response)
	// Output: authorize_net.GetMerchantDetailsResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}, IsTestMode:false, Processors:[]authorize_net.Processor{authorize_net.Processor{Name:"First Data Nashville", Id:2, CardTypes:[]string{"A", "D", "M", "V"}}}, MerchantName:"Test Developer", GatewayId:"475314", MarketTypes:[]string{"eCommerce"}, ProductCodes:[]string{"CNP"}, PaymentMethods:[]string{"AmericanExpress", "Discover", "Echeck", "MasterCard", "Paypal", "Visa", "VisaCheckout", "ApplePay", "AndroidPay"}, Currencies:[]string{"USD"}, PublicClientKey:"5FcB6WrfHGS76gHW3v7btBCE3HuuBuke9Pj96Ztfn5R32G5ep42vne7MCWZtAucY", BusinessInformation:(*authorize_net.CustomerAddress)(0xc0003e4000), MerchantTimeZone:"America/Los_Angeles", ContactDetails:[]authorize_net.ContactDetail{authorize_net.ContactDetail{Email:"bmcmanus@visa.com", FirstName:"Sandbox", LastName:"Default"}}}
}

func ExampleUpdateMerchantDetails() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := UpdateMerchantDetailsRequest{
		Payload: UpdateMerchantDetailsPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			IsTestMode: true,
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response UpdateMerchantDetailsResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.UpdateMerchantDetailsResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Error", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x47, Text:"The authentication type is not allowed for this method call.", Description:""}}}, SessionToken:""}}
}

func ExampleGetCustomerShippingAddress() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetCustomerShippingAddressRequest{
		Payload: GetCustomerShippingAddressPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "1920672921",
			//CustomerAddressId: "1877745863",// optional, if you created the record yourself
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetCustomerShippingAddressResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response.Address)
	// Output: &authorize_net.CustomerAddressEx{CustomerAddress:authorize_net.CustomerAddress{NameAndAddress:authorize_net.NameAndAddress{FirstName:"Newfirstname", LastName:"Doe", Company:"", Address:"123 Main St.", City:"Bellevue", State:"WA", Zip:"98004", Country:"USA"}, PhoneNumber:"000-000-0000", FaxNumber:"", Email:""}, CustomerAddressId:"1810861269"}
}

func ExampleCreateCustomerShippingAddress() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := CreateCustomerShippingAddressRequest{
		Payload: CreateCustomerShippingAddressPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "1920672921",
			Address: &CustomerAddress{
				NameAndAddress: NameAndAddress{
					FirstName: "FirstName",
					LastName:  "LastName",
					Company:   "SomeCompany LTD",
					Address:   "Hope Street, ground floor",
					City:      "Oz",
					State:     "Oz",
					Zip:       "300200",
					Country:   "Norway",
				},
				PhoneNumber: "1131131131",
				FaxNumber:   "what?",
				Email:       "36152127@authorize.net@authorize.net",
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response CreateCustomerShippingAddressResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.CreateCustomerShippingAddressResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}, CustomerProfileId:"1920672921", CustomerAddressId:"1877745883"}
}

func ExampleUpdateCustomerShippingAddress() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := UpdateCustomerShippingAddressRequest{
		Payload: UpdateCustomerShippingAddressPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "1920672921",
			Address: &CustomerAddressEx{
				CustomerAddressId: "1877745883",
				CustomerAddress: CustomerAddress{
					NameAndAddress: NameAndAddress{
						FirstName: "FirstName",
						LastName:  "LastName",
						Company:   "SomeCompany LTD",
						Address:   "Hope Street, ground floor",
						City:      "Oz",
						State:     "Oz",
						Zip:       "300200",
						Country:   "China",
					},
					PhoneNumber: "1131131131",
					FaxNumber:   "what?",
					Email:       "36152127@novalidation@authorize.net",
				},
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response UpdateCustomerShippingAddressResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.UpdateCustomerShippingAddressResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}}

}

func ExampleDeleteCustomerShippingAddress() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := DeleteCustomerShippingAddressRequest{
		Payload: DeleteCustomerShippingAddressPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "1920672921",
			CustomerAddressId: "1877745863",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response DeleteCustomerShippingAddressResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.DeleteCustomerShippingAddressResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}}

}

func ExampleCreateCustomerProfileFromTransaction() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := CreateCustomerProfileFromTransactionRequest{
		Payload: CreateCustomerProfileFromTransactionPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			Customer: &CustomerProfileBase{
				MerchantCustomerId: "19328",
				Description:        "1233",
				Email:              "123123",
			},
			TransId:                "60126769903",
			DefaultPaymentProfile:  true,
			DefaultShippingAddress: true,
			ProfileType:            CustomerProfileGuest,
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response CreateCustomerProfileFromTransactionResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.CreateCustomerProfileFromTransactionResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Error", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x70, Text:"Customer profile creation failed. This transaction type does not support profile creation.", Description:""}}}, SessionToken:""}, CustomerProfileID:"", CustomerPaymentProfileIDList:[]string{}, CustomerShippingAddressIDList:[]string{}, ValidationDirectResponseList:[]string{}}

}

func ExampleCreateCustomerProfileTransaction() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := CreateCustomerProfileTransactionRequest{
		Payload: CreateCustomerProfileTransactionPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			Transaction: &ProfileTransaction{
				ProfileTransAuthCapture: &ProfileTransOrder{
					ProfileTransAmount: ProfileTransAmount{
						Amount: 10,
						Tax: &ExtendedAmount{
							Amount:      10,
							Name:        "",
							Description: "",
						},
						Shipping: &ExtendedAmount{
							Amount:      10,
							Name:        "",
							Description: "",
						},
						Duty: &ExtendedAmount{
							Amount:      10,
							Name:        "",
							Description: "",
						},
						LineItems: nil,
					},
					CustomerProfileId:        "37821321",
					CustomerPaymentProfileId: "34382309",
					Order: &OrderEx{
						Order: Order{
							InvoiceNumber:                  "INV001",
							Description:                    "",
							DiscountAmount:                 10,
							TaxIsAfterDiscount:             true,
							TotalTaxTypeCode:               "",
							PurchaserVATRegistrationNumber: "",
							MerchantVATRegistrationNumber:  "",
							VatInvoiceReferenceNumber:      "",
							PurchaserCode:                  "1234",
							SummaryCommodityCode:           "",
							PurchaseOrderDateUTC:           time.Now().UTC().Format("2006-01-02"),
							SupplierOrderReference:         "",
							AuthorizedContactName:          "",
							CardAcceptorRefNumber:          "",
							AmexDataTAA1:                   "",
							AmexDataTAA2:                   "",
							AmexDataTAA3:                   "",
							AmexDataTAA4:                   "",
						},
						PurchaseOrderNumber: "",
					},
					TaxExempt:        true,
					RecurringBilling: false,
					CardCode:         cardCode,
					//SplitTenderId:    "",
					ProcessingOptions: &ProcessingOptions{
						IsFirstRecurringPayment: false,
						IsFirstSubsequentAuth:   false,
						IsSubsequentAuth:        false,
						IsStoredCredentials:     false,
					},
				},
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response CreateCustomerProfileTransactionResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.CreateCustomerProfileTransactionResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}, TransactionResponse:(*authorize_net.TransactionResponse)(nil), DirectResponse:"1,1,1,This transaction has been approved.,N0SU5L,Y,60126781190,INV001,,10.00,CC,auth_capture,,John,Smith,,,,,,,,,,,,,,,,,,10.00,10.00,10.00,TRUE,,,P,2,,,,,,,,,,,XXXX1111,Visa,,,,,,,,,,,,,,,,,"}
}

func ExampleCreateCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := CreateCustomerPaymentProfileRequest{
		Payload: CreateCustomerPaymentProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "1920672070",
			PaymentProfile: &CustomerPaymentProfile{
				CustomerPaymentProfileBase: CustomerPaymentProfileBase{
					CustomerType: 0,
					BillTo: &CustomerAddress{
						NameAndAddress: NameAndAddress{
							FirstName: "Johnny",
							LastName:  "Doe",
							Country:   "USA",
						},
						Email: "gmail@chuck.norris",
					},
				},
				Payment: &Payment{
					CreditCard: &CreditCard{
						CardNumber:     cardExample,
						ExpirationDate: cardExpirationDate,
						CardCode:       cardCode,
					},
				},
				DefaultPaymentProfile: true,
			},
			ValidationMode: ValidationModeTestMode,
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response CreateCustomerPaymentProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.CreateCustomerPaymentProfileResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}, CustomerProfileId:"1920672070", CustomerPaymentProfileId:"1833657059", ValidationDirectResponse:"1,1,1,(TESTMODE) This transaction has been approved.,000000,P,0,none,Test transaction for ValidateCustomerPaymentProfile.,1.00,CC,auth_only,none,Johnny,Doe,,,,,,USA,,,email@example.com,,,,,,,,,0.00,0.00,0.00,FALSE,none,,,,,,,,,,,,,,XXXX0015,MasterCard,,,,,,,,,,,,,,,,,"}

}

func ExampleValidateCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := ValidateCustomerPaymentProfileRequest{
		Payload: ValidateCustomerPaymentProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId:        "1920672070",
			CustomerPaymentProfileId: "1833657059",
			ValidationMode:           ValidationModeTestMode,
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response ValidateCustomerPaymentProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.ValidateCustomerPaymentProfileResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}, DirectResponse:"1,1,1,(TESTMODE) This transaction has been approved.,000000,P,0,none,Test transaction for ValidateCustomerPaymentProfile.,1.00,CC,auth_only,jdoe4801,Johnny,Doe,,,,,,USA,,,4953@mail.com,,,,,,,,,0.00,0.00,0.00,FALSE,none,,,,,,,,,,,,,,XXXX0015,MasterCard,,,,,,,,,,,,,,,,,"}

}

func ExampleUpdateCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := UpdateCustomerPaymentProfileRequest{
		Payload: UpdateCustomerPaymentProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId: "1920672070",
			ValidationMode:    ValidationModeNone,
			PaymentProfile: &CustomerPaymentProfileEx{
				CustomerPaymentProfileId: "1833657059",
				CustomerPaymentProfile: CustomerPaymentProfile{
					CustomerPaymentProfileBase: CustomerPaymentProfileBase{
						CustomerType: Business,
					},
					Payment: &Payment{

						BankAccount: &BankAccount{
							AccountType:   BankAccountTypeSavings,
							RoutingNumber: "133563585",
							AccountNumber: "0123456789",
							NameOnAccount: "SALSA",
							EcheckType:    EcheckCCD,
							BankName:      "SALSA",
							CheckNumber:   "123",
						},
					},
					DriversLicense: &DriversLicense{
						Number:      "12388381813",
						State:       "CA",
						DateOfBirth: "2000-10-10",
					},
					DefaultPaymentProfile: true,
				},
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response UpdateCustomerPaymentProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.UpdateCustomerPaymentProfileResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}, ValidationDirectResponse:""}

}

func ExampleDeleteCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := DeleteCustomerPaymentProfileRequest{
		Payload: DeleteCustomerPaymentProfilePayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			CustomerProfileId:        "1920672070",
			CustomerPaymentProfileId: "1833657059",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response DeleteCustomerPaymentProfileResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Printf("%#v", response)
	// Output: authorize_net.DeleteCustomerPaymentProfileResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Ok", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0x1, Text:"Successful.", Description:""}}}, SessionToken:""}}

}

func ExampleDecryptPaymentData() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := DecryptPaymentDataRequest{
		Payload: DecryptPaymentDataPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			OpaqueData: &OpaqueData{
				DataDescriptor: "COMMON.VCO.ONLINE.PAYMENT", // always send this
				DataValue:      "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",
				DataKey:        "qGl+69iAWXJ13+18jHgBO2zHCuekawWcApfLhGnbYXD4GsI9EOM2Y5V8zGXvr1lF3hjGMT0NoD2thxzR7CrAvgfgw7lAJzlGIACOZnEkx70ywnJHAR3sGO8hyq9f1Fk",
			},
			CallId: "2186595692635007317",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response DecryptPaymentDataResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	// Note : for some reason it always return an error
	fmt.Printf("%#v", response)
	// Output: authorize_net.DecryptPaymentDataResponse{ANetApiResponse:authorize_net.ANetApiResponse{RefId:"", Messages:authorize_net.Messages{ResultCode:"Error", Messages:[]authorize_net.ErrMessage{authorize_net.ErrMessage{Code:0xd, Text:"An error occurred during processing. Please try again.", Description:""}}}, SessionToken:""}, ShippingInfo:(*authorize_net.CustomerAddress)(nil), BillingInfo:(*authorize_net.CustomerAddress)(nil), CardInfo:(*authorize_net.CreditCardMasked)(nil), PaymentDetails:(*authorize_net.PaymentDetails)(nil)}

}

func ExampleGetTransactionListForCustomer() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetTransactionListForCustomerRequest{
		Payload: GetTransactionListForCustomerPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			Sorting: Sorting{
				OrderBy:         "submitTimeUTC",
				OrderDescending: true,
			},
			Paging: Paging{
				Offset: 1,
				Limit:  30,
			},
			CustomerProfileId: "1920672070",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetTransactionListResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	for _, trans := range response.Transactions {
		fmt.Printf("%#v\n", trans.TransId)
	}
	// Output: "60126763648"
	//"60126763647"
	//"60126763644"
}

func ExampleSendCustomerTransactionReceipt() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := SendCustomerTransactionReceiptRequest{
		Payload: SendCustomerTransactionReceiptPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			TransId:       "60126763648",
			CustomerEmail: "put_valid_email_here",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response SendCustomerTransactionReceiptResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Sprintf("%v", response)
	// Output:
}

func ExampleGetTransactionDetails() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetTransactionDetailsRequest{
		Payload: GetTransactionDetailsPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			TransId: "60126763648",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetTransactionDetailsResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%#v", response.Transaction.Batch.BatchId)
	// Output: "9697004"
}

func ExampleGetTransactionList() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetTransactionListRequest{
		Payload: GetTransactionListPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			Paging: Paging{
				Limit:  50,
				Offset: 1,
			},
			Sorting: Sorting{
				OrderBy:         "submitTimeUTC",
				OrderDescending: false,
			},
			BatchId: "9697004",
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetTransactionListResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%#v", response.TotalNumInResultSet)
	// Output: 50

}

func ExampleGetUnsettledTransactionList() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	payload := GetUnsettledTransactionListRequest{
		Payload: GetUnsettledTransactionListPayload{
			ANetApiRequest: ANetApiRequest{
				MerchantAuthentication: MerchantAuthentication{
					Name:           merchantId,
					TransactionKey: merchantSecret,
				},
			},
			Paging: Paging{
				Limit:  50,
				Offset: 1,
			},
			Sorting: Sorting{
				OrderBy:         "submitTimeUTC",
				OrderDescending: false,
			},
		},
	}
	request, err := client.prepareRequest(context.Background(), SandboxEndpoint, http.MethodPost, &payload)
	if err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}

	var response GetUnsettledTransactionListResponse
	if err := client.do(request, &response, false); err != nil {
		client.log.Printf("Error : %v\n", err)
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%#v", response.TotalNumInResultSet)
	// Output: 50
}

func ExampleANetApi() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleARBCancelSubscription() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleARBCreateSubscription() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleARBGetSubscription() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleARBGetSubscriptionList() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleARBGetSubscriptionStatus() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleARBUpdateSubscription() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetAUJobDetails() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetAUJobSummary() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetBatchStatistics() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetHostedPaymentPage() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetHostedProfilePage() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetSettledBatchList() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleIsAlive() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleLogout() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleMobileDeviceLogin() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleMobileDeviceRegistration() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleSecurePaymentContainer() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateHeldTransaction() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateSplitTenderGroup() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}
