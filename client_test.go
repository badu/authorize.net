package authorize_net

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	merchantId         = "5KP3u95bQpv"      // sandbox Authorize.net
	merchantSecret     = "346HZ32z3fP4hTG2" // sandbox Authorize.net
	aCreditCard        = "5424000000000015"
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
	// Output: {RefId: Messages:{ResultCode:Ok Messages:[{Code:Successful. Text:Successful.}]} SessionToken:}
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
						CardNumber:     aCreditCard,
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
						CardNumber:     aCreditCard,
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
						CardNumber:     aCreditCard,
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
						CardNumber:     aCreditCard,
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
								CardNumber:     aCreditCard,
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
			Sorting: &CustomerPaymentProfileSorting{
				OrderBy:         "id",
				OrderDescending: false,
			},
			Paging: &Paging{
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

func ExampleCreateCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleCreateCustomerProfileFromTransaction() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleCreateCustomerProfileTransaction() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleCreateCustomerShippingAddress() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleDecryptPaymentData() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleDeleteCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleDeleteCustomerProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleDeleteCustomerShippingAddress() {
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

func ExampleGetCustomerShippingAddress() {
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

func ExampleGetMerchantDetails() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetSettledBatchList() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetTransactionDetails() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetTransactionList() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetTransactionListForCustomer() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleGetUnsettledTransactionList() {
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

func ExampleSendCustomerTransactionReceipt() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateCustomerProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateCustomerShippingAddress() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateHeldTransaction() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateMerchantDetails() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleUpdateSplitTenderGroup() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}

func ExampleValidateCustomerPaymentProfile() {
	client := NewAPIClient(nil, log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	fmt.Sprintf("%v", client)

}
