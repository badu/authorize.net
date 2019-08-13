package authorize_net

import "errors"

type ErrCodeEnum uint

const (
	// The request was processed successfully.
	I00001     = 1 // original "I00001"
	EnUsI00001 = "Successful."
	// The subscription has already been canceled.
	I00002     = 2 // original "I00002"
	EnUsI00002 = "The subscription has already been canceled."
	// The record has already been deleted.
	I00003     = 3 // original "I00003"
	EnUsI00003 = "The record has already been deleted."
	// No records have been found that match your query.
	I00004     = 4 // original "I00004"
	EnUsI00004 = "No records found."
	// The mobile device was successfully inserted into the database.
	I00005     = 5 // original "I00005"
	EnUsI00005 = "The mobile device has been submitted for approval by the account administrator."
	// The mobile device was successfully registered and approved by the account administrator.
	I00006     = 6 // original "I00006"
	EnUsI00006 = "The mobile device is approved and ready for use."
	// The Payment Gateway Account service (id=8) has already been accepted.
	I00007     = 7 // original "I00007"
	EnUsI00007 = "The Payment Gateway Account service (id&#x3D;8) has already been accepted."
	// The Payment Gateway Account service (id=8) has already been declined.
	I00008     = 8 // original "I00008"
	EnUsI00008 = "The Payment Gateway Account service (id&#x3D;8) has already been declined."
	// The APIUser already exists.
	I00009     = 9 // original "I00009"
	EnUsI00009 = "The APIUser already exists."
	// The merchant is activated successfully.
	I00010     = 10 // original "I00010"
	EnUsI00010 = "The merchant is activated successfully."
	// The merchant is not activated.
	I00011     = 11 // original "I00011"
	EnUsI00011 = "The merchant is not activated."
	// An unexpected system error occurred while processing this request.
	E00001     = 13 // original "E00001"
	EnUsE00001 = "An error occurred during processing. Please try again."
	// The only supported content-types are text/xml and application/xml.
	E00002     = 14 // original "E00002"
	EnUsE00002 = "The content-type specified is not supported."
	// This is the result of an XML parser error.
	E00003     = 15 // original "E00003"
	EnUsE00003 = "An error occurred while parsing the XML request."
	// The name of the root node of the XML request is the API method being called. It is not valid.
	E00004     = 16 // original "E00004"
	EnUsE00004 = "The name of the requested API method is invalid."
	// User authentication requires a valid value for transaction key or API key.
	E00005     = 17 // original "E00005"
	EnUsE00005 = "The transaction key or API key is invalid or not present."
	// User authentication requires a valid value for API user name.
	E00006     = 18 // original "E00006"
	EnUsE00006 = "The API user name is invalid or not present."
	// The API user name is invalid and/or the transaction key or API key is invalid.
	E00007     = 19 // original "E00007"
	EnUsE00007 = "User authentication failed due to invalid authentication values."
	// The payment gateway, reseller, or user account is not currently active.
	E00008     = 20 // original "E00008"
	EnUsE00008 = "User authentication failed. The account or API user is inactive."
	// The requested API method cannot be executed while the payment gateway account is in Test Mode.
	// To disable Test Mode, log into the Merchant Interface at https://account.authorize.net/ and click Account > Test Mode > Turn Test OFF.
	E00009     = 21 // original "E00009"
	EnUsE00009 = "The payment gateway account is in Test Mode. The request cannot be processed."
	// The user does not have permission to call the API.
	E00010     = 22 // original "E00010"
	EnUsE00010 = "User authentication failed. You do not have the appropriate permissions."
	// The user does not have permission to call the API method.
	E00011     = 23 // original "E00011"
	EnUsE00011 = "Access denied. You do not have the appropriate permissions."
	// A duplicate of the subscription was already submitted.
	E00012     = 24 // original "E00012"
	EnUsE00012 = "A duplicate subscription already exists."
	// One of the field values is not valid.
	E00013     = 25 // original "E00013"
	EnUsE00013 = "The field is invalid."
	// One of the required fields was not present.
	E00014     = 26 // original "E00014"
	EnUsE00014 = "A required field is not present."
	// One of the fields has an invalid length.
	E00015     = 27 // original "E00015"
	EnUsE00015 = "The field length is invalid."
	// The field type is not valid.
	E00016     = 28 // original "E00016"
	EnUsE00016 = "The field type is invalid."
	// The subscription start date cannot occur before the subscription submission date.
	E00017     = 29 // original "E00017"
	EnUsE00017 = "The start date cannot occur in the past."
	// The credit card is not valid as of the start date of the subscription.
	E00018     = 30 // original "E00018"
	EnUsE00018 = "The credit card expires before the subscription start date."
	// The customer tax ID or driver's license information (driver's license number, driver's license state, driver's license DOB) is required for the subscription.
	E00019     = 31 // original "E00019"
	EnUsE00019 = "The customer tax id or drivers license information is required."
	// The payment gateway account is not set up to process eCheck.Net subscriptions.
	E00020     = 32 // original "E00020"
	EnUsE00020 = "The payment gateway account is not enabled for eCheck.Net subscriptions."
	// The payment gateway account is not set up to process credit card subscriptions.
	E00021     = 33 // original "E00021"
	EnUsE00021 = "The payment gateway account is not enabled for credit card subscriptions."
	// The interval length must be 7 to 365 days or 1 to 12 months.
	E00022     = 34 // original "E00022"
	EnUsE00022 = "The interval length cannot exceed 365 days or 12 months."
	// The number of total occurrences cannot extend the duration of the subscription beyond three years from the start date.
	E00023     = 35 // original "E00023"
	EnUsE00023 = "The subscription duration cannot exceed three years."
	// The number of trial occurrences cannot be zero if a valid trial amount is submitted.
	E00024     = 36 // original "E00024"
	EnUsE00024 = "Trial Occurrences is required when Trial Amount is specified."
	// The payment gateway account is not enabled for Automated Recurring Billing.
	E00025     = 37 // original "E00025"
	EnUsE00025 = "Automated Recurring Billing is not enabled."
	// If either a trial amount or number of trial occurrences is specified then values for both must be submitted.
	E00026     = 38 // original "E00026"
	EnUsE00026 = "Both Trial Amount and Trial Occurrences are required."
	// An approval was not returned for the transaction.
	// For more information, check the errorCode field in the response.
	E00027     = 39 // original "E00027"
	EnUsE00027 = "The transaction was unsuccessful."
	// The number of trial occurrences specified must be less than the number of total occurrences specified.
	E00028     = 40 // original "E00028"
	EnUsE00028 = "Trial Occurrences must be less than Total Occurrences."
	// Payment information is required when creating a subscription or payment profile.
	E00029     = 41 // original "E00029"
	EnUsE00029 = "Payment information is required."
	// A payment schedule is required when creating a subscription.
	E00030     = 42 // original "E00030"
	EnUsE00030 = "The payment schedule is required."
	// The subscription amount is required when creating a subscription.
	E00031     = 43 // original "E00031"
	EnUsE00031 = "The amount is required."
	// The subscription start date is required to create a subscription.
	E00032     = 44 // original "E00032"
	EnUsE00032 = "The start date is required."
	// Once a subscription is created the start date cannot be changed.
	E00033     = 45 // original "E00033"
	EnUsE00033 = "The start date cannot be changed."
	// Once a subscription is created the interval cannot be changed.
	E00034     = 46 // original "E00034"
	EnUsE00034 = "The interval information cannot be changed."
	// The subscription ID for this request is not valid for this merchant.
	E00035     = 47 // original "E00035"
	EnUsE00035 = "The subscription cannot be found."
	// Changing the subscription payment type between credit card and eCheck.Net is not currently supported.
	E00036     = 48 // original "E00036"
	EnUsE00036 = "The payment type cannot be changed."
	// Subscriptions that are expired, canceled or terminated cannot be updated.
	E00037     = 49 // original "E00037"
	EnUsE00037 = "The subscription cannot be updated."
	// Subscriptions that are expired or terminated cannot be canceled.
	E00038     = 50 // original "E00038"
	EnUsE00038 = "The subscription cannot be canceled."
	// A duplicate of the customer profile, customer payment profile, or customer address was already submitted.
	E00039     = 51 // original "E00039"
	EnUsE00039 = "A duplicate record already exists."
	// The customer profile ID, payment profile ID, shipping address ID, or transaction ID for this request is not valid for this merchant.
	E00040     = 52 // original "E00040"
	EnUsE00040 = "The record cannot be found."
	// All of the fields were empty or missing.
	E00041     = 53 // original "E00041"
	EnUsE00041 = "One or more fields must contain a value."
	// The maximum number of payment profiles for the customer profile has been reached.
	E00042     = 54 // original "E00042"
	EnUsE00042 = "You cannot add more than {0} payment profiles."
	// The maximum number of shipping addresses for the customer profile has been reached.
	E00043     = 55 // original "E00043"
	EnUsE00043 = "You cannot add more than {0} shipping addresses."
	// The payment gateway account is not enabled for Customer Information Manager (CIM).
	E00044     = 56 // original "E00044"
	EnUsE00044 = "Customer Information Manager is not enabled."
	// The root node does not reference a valid XML namespace.
	E00045     = 57 // original "E00045"
	EnUsE00045 = "The root node does not reference a valid XML namespace."
	// Generic InsertNewMerchant failure.
	E00046     = 58 // original "E00046"
	EnUsE00046 = "Generic InsertNewMerchant failure."
	// The reseller account is not enabled for Merchant Boarding API.
	E00047     = 59 // original "E00047"
	EnUsE00047 = "Merchant Boarding API is not enabled."
	// The merchant account must be set up to accept credit card payments, eCheck payments, or both.
	E00048     = 60 // original "E00048"
	EnUsE00048 = "At least one payment method must be set in payment types or an echeck service must be provided."
	// The database operation timed out before it could be completed.
	E00049     = 61 // original "E00049"
	EnUsE00049 = "The operation timed out before it could be completed."
	// Cannot set a buyrate to less than the sellrate
	E00050     = 62 // original "E00050"
	EnUsE00050 = "Sell Rates cannot be less than Buy Rates"
	// If customer profile ID, payment profile ID, and shipping address ID are included, they must match the original transaction.
	E00051     = 63 // original "E00051"
	EnUsE00051 = "The original transaction was not issued for this payment profile."
	// The maximum number of elements for an array has been reached.
	E00052     = 64 // original "E00052"
	EnUsE00052 = "The maximum number of elements for an array {0} is {1}."
	// The server is currently too busy, please try again later.
	E00053     = 65 // original "E00053"
	EnUsE00053 = "Server too busy"
	// The mobile device identifier is not associated with the merchant account.
	E00054     = 66 // original "E00054"
	EnUsE00054 = "The mobile device is not registered with this merchant account."
	// The mobile device exists but is in a pending status.
	E00055     = 67 // original "E00055"
	EnUsE00055 = "The mobile device has already been registered but is pending approval by the account administrator."
	// The mobile device exists but has a status of disabled.
	E00056     = 68 // original "E00056"
	EnUsE00056 = "The mobile device has been disabled for use with this account."
	// The user does not have sufficient permissions to use a mobile device with this merchant account.
	E00057     = 69 // original "E00057"
	EnUsE00057 = "The user does not have permissions to submit requests from a mobile device."
	// The merchant has too many devices in a pending status.
	E00058     = 70 // original "E00058"
	EnUsE00058 = "The merchant has met or exceeded the number of pending mobile devices permitted for this account."
	// The authentication type is not allowed for the requested method call.
	E00059     = 71 // original "E00059"
	EnUsE00059 = "The authentication type is not allowed for this method call."
	// The transaction type is invalid.
	E00060     = 72 // original "E00060"
	EnUsE00060 = "The transaction type is invalid."
	// Could not decrypt DUKPT blobs and returned error.
	E00061     = 73 // original "E00061"
	EnUsE00061 = "{0}({1})."
	// Fatal error when calling web service.
	E00062     = 74 // original "E00062"
	EnUsE00062 = "Fatal error when calling web service."
	// Calling web service return error.
	E00063     = 75 // original "E00063"
	EnUsE00063 = "Calling web service return error."
	// Client authorization denied.
	E00064     = 76 // original "E00064"
	EnUsE00064 = "Client authorization denied."
	// Prerequisite failed.
	E00065     = 77 // original "E00065"
	EnUsE00065 = "Prerequisite failed."
	// Invalid value.
	E00066     = 78 // original "E00066"
	EnUsE00066 = "Invalid value."
	// This is the result of an XML parser error.  Too many nodes specified.
	E00067     = 79 // original "E00067"
	EnUsE00067 = "An error occurred while parsing the XML request.  Too many {0} specified."
	// This is the result of an XML parser error.  The node is invalid.
	E00068     = 80 // original "E00068"
	EnUsE00068 = "An error occurred while parsing the XML request.  {0} is invalid."
	// The Payment Gateway Account service (id=8) has already been accepted.  Decline is not allowed.
	E00069     = 81 // original "E00069"
	EnUsE00069 = "The Payment Gateway Account service (id&#x3D;8) has already been accepted.  Decline is not allowed."
	// The Payment Gateway Account service (id=8) has already been declined.  Agree is not allowed.
	E00070     = 82 // original "E00070"
	EnUsE00070 = "The Payment Gateway Account service (id&#x3D;8) has already been declined.  Agree is not allowed."
	// All of the fields were empty or missing.
	E00071     = 83 // original "E00071"
	EnUsE00071 = "{0} must contain data."
	// Required node missing.
	E00072     = 84 // original "E00072"
	EnUsE00072 = "Node {0} is required."
	// One of the field values is not valid.
	E00073     = 85 // original "E00073"
	EnUsE00073 = "{0} is invalid."
	// This merchant is not associated with this reseller.
	E00074     = 86 // original "E00074"
	EnUsE00074 = "This merchant is not associated with this reseller."
	// This is the result of an XML parser error.  Missing field(s).
	E00075     = 87 // original "E00075"
	EnUsE00075 = "An error occurred while parsing the XML request.  Missing field(s) {0}."
	// Invalid value.
	E00076     = 88 // original "E00076"
	EnUsE00076 = "{0} contains invalid value."
	// Value too long.
	E00077     = 89 // original "E00077"
	EnUsE00077 = "The value of {0} is too long.  The length of value should be {1}"
	// Pending Status (not completed).
	E00078     = 90 // original "E00078"
	EnUsE00078 = "Pending Status (not completed)."
	// Impersonation partner login ID is invalid or not present.
	E00079     = 91 // original "E00079"
	EnUsE00079 = "The impersonation login ID is invalid or not present."
	// Impersonation API Key is invalid or not present.
	E00080     = 92 // original "E00080"
	EnUsE00080 = "The impersonation API Key is invalid or not present."
	// The partner account is not authorized to impersonate the login account.
	E00081     = 93 // original "E00081"
	EnUsE00081 = "Partner account is not authorized to impersonate the login account."
	// Country is not valid.
	E00082     = 94 // original "E00082"
	EnUsE00082 = "Country for {0} is not valid."
	// Bank payment method is not accepted for the selected business country.
	E00083     = 95 // original "E00083"
	EnUsE00083 = "Bank payment method is not accepted for the selected business country."
	// Credit card payment method is not accepted for the selected business country.
	E00084     = 96 // original "E00084"
	EnUsE00084 = "Credit card payment method is not accepted for the selected business country."
	// State is not valid.
	E00085     = 97 // original "E00085"
	EnUsE00085 = "State for {0} is not valid."
	// Merchant has declined authorization to resource.
	E00086     = 98 // original "E00086"
	EnUsE00086 = "Merchant has declined authorization to resource."
	// There are no subscriptions available for the merchant account for the type of subscriptions requested.
	E00087     = 99 // original "E00087"
	EnUsE00087 = "No subscriptions found for the given request."
	// CreateProfile and profileIds are mutually exclusive, only one of them can be provided at a time.
	E00088     = 100 // original "E00088"
	EnUsE00088 = "ProfileIds cannot be sent when requesting CreateProfile."
	// When requesting CreateProfile payment data cannot be null.
	E00089     = 101 // original "E00089"
	EnUsE00089 = "Payment data is required when requesting CreateProfile."
	// PaymentProfile and PaymentData are mutually exclusive, only one of them can be provided at a time.
	E00090     = 102 // original "E00090"
	EnUsE00090 = "PaymentProfile cannot be sent with payment data."
	// PaymentProfileId and payment data are mutually exclusive, only one of them can be provided at a time.
	E00091     = 103 // original "E00091"
	EnUsE00091 = "PaymentProfileId cannot be sent with payment data."
	// ShippingProfileId and ShipToAddress are mutually exclusive, only one of them can be provided at a time.
	E00092     = 104 // original "E00092"
	EnUsE00092 = "ShippingProfileId cannot be sent with ShipTo data."
	// PaymentProfile and Billing information are mutually exclusive, only one of them can be provided at a time.
	E00093     = 105 // original "E00093"
	EnUsE00093 = "PaymentProfile cannot be sent with billing data."
	// Paging Offset exceeds allowed value. Check and lower the value.
	E00094     = 106 // original "E00094"
	EnUsE00094 = "Paging Offset exceeds the maximum allowed value."
	// When using Customer Profile with Credit Card Info to specify Shipping Profile, Shipping Profile Id must be included.
	E00095     = 107 // original "E00095"
	EnUsE00095 = "ShippingProfileId is not provided within Customer Profile."
	// Finger Print value is not valid.
	E00096     = 108 // original "E00096"
	EnUsE00096 = "Finger Print value is not valid."
	// Finger Print can't be generated.
	E00097     = 109 // original "E00097"
	EnUsE00097 = "Finger Print can&#x27;t be generated."
	// Search for shipping profile using customer profile id and shipping profile id did not find any records.
	E00098     = 110 // original "E00098"
	EnUsE00098 = "Customer Profile ID or Shipping Profile ID not found."
	// Customer profile creation failed. This transaction ID is invalid.
	E00099     = 111 // original "E00099"
	EnUsE00099 = "Customer profile creation failed. This transaction ID is invalid."
	// Customer profile creation failed. This transaction type does not support profile creation.
	E00100     = 112 // original "E00100"
	EnUsE00100 = "Customer profile creation failed. This transaction type does not support profile creation."
	// Error creating a customer payment profile from transaction.
	E00101     = 113 // original "E00101"
	EnUsE00101 = "Customer profile creation failed."
	// Error creating a customer profile from transaction.
	E00102     = 114 // original "E00102"
	EnUsE00102 = "Customer Info is missing."
	// Customer profile creation failed. This payment method does not support profile creation.
	E00103     = 115 // original "E00103"
	EnUsE00103 = "Customer profile creation failed. This payment method does not support profile creation."
	// The server is in maintenance, so the requested method is unavailable. Please try again later.
	E00104     = 116 // original "E00104"
	EnUsE00104 = "Server in maintenance. Please try again later."
	// The specified payment profile is associated with an active or suspended subscription and cannot be deleted.
	E00105     = 117 // original "E00105"
	EnUsE00105 = "The specified payment profile is associated with an active or suspended subscription and cannot be deleted."
	// The specified customer profile is associated with an active or suspended subscription and cannot be deleted.
	E00106     = 118 // original "E00106"
	EnUsE00106 = "The specified customer profile is associated with an active or suspended subscription and cannot be deleted."
	// The specified shipping profile is associated with an active or suspended subscription and cannot be deleted.
	E00107     = 119 // original "E00107"
	EnUsE00107 = "The specified shipping profile is associated with an active or suspended subscription and cannot be deleted."
	// CustomerProfileId and Customer data are mutually exclusive, only one of them can be provided for any single request.
	E00108     = 120 // original "E00108"
	EnUsE00108 = "CustomerProfileId cannot be sent with customer data."
	//  Shipping Address ID and Shipping data are mutually exclusive, only one of them can be provided for any single request.
	E00109     = 121 // original "E00109"
	EnUsE00109 = "CustomerAddressId cannot be sent with shipTo data."
	// When using Customer Profile, CustomerPaymentProfileId must be included.
	E00110     = 122 // original "E00110"
	EnUsE00110 = "CustomerPaymentProfileId is not provided within Customer Profile."
	// If Customer Profile ID is included, it must match the Customer Profile ID used for the original subscription.
	E00111     = 123 // original "E00111"
	EnUsE00111 = "The original subscription was not created with this Customer Profile."
	// Reports cannot be generated for future dates, thus the specified date is invalid.
	E00112     = 124 // original "E00112"
	EnUsE00112 = "The specified month should not be in the future."
	// The specified OTS Token Data is invalid.
	E00113     = 125 // original "E00113"
	EnUsE00113 = "Invalid OTS Token Data."
	// The specified OTS Token is invalid.
	E00114     = 126 // original "E00114"
	EnUsE00114 = "Invalid OTS Token."
	// The specified OTS Token has expired.
	E00115     = 127 // original "E00115"
	EnUsE00115 = "Expired OTS Token."
	// The authenticated merchant does not have access to the specified OTS Token.
	E00116     = 128 // original "E00116"
	EnUsE00116 = "OTS Token access violation"
	// The OTS Service cannot complete the request due to a validation or configuration error.
	E00117     = 129 // original "E00117"
	EnUsE00117 = "OTS Service Error &#x27;{0}&#x27;"
	// The transaction was submitted from a blocked IP address.
	E00118     = 130 // original "E00118"
	EnUsE00118 = "The transaction has been declined."
	// Hosted Payment Page will capture the payment (bank/card) information so this information should not be included with this request.
	E00119     = 131 // original "E00119"
	EnUsE00119 = "Payment information should not be sent to Hosted Payment Page request."
	// Payment and Shipping Profile IDs cannot be specified when creating new profiles.
	E00120     = 132 // original "E00120"
	EnUsE00120 = "Payment and Shipping Profile IDs cannot be specified when creating new profiles."
	// The customer profile does not have a default payment/shipping profile.
	E00121     = 133 // original "E00121"
	EnUsE00121 = "No default payment/shipping profile found."
	// Signature key missing.
	E00122     = 134 // original "E00122"
	EnUsE00122 = "Please use Merchant Interface settings (API Credentials and Keys) to generate a signature key."
	// The access token provided has expired.
	// See the <a href&#x3D;`https://developer.authorize.net/api/reference/features/oauth.html`>OAuth documentation</a> for details.
	E00123     = 135 // original "E00123"
	EnUsE00123 = "The provided access token has expired"
	// The access token used to validate the request is insufficient to do so.
	// See the <a href&#x3D;`https://developer.authorize.net/api/reference/features/oauth.html`>OAuth documentation</a> for more details.
	E00124     = 136 // original "E00124"
	EnUsE00124 = "The provided access token is invalid"
	// Hash doesnâ€™t match.
	E00125     = 137 // original "E00125"
	EnUsE00125 = "Hash doesnâ€™t match"
	// Failed shared key validation.
	E00126     = 138 // original "E00126"
	EnUsE00126 = "Failed shared key validation"
	// Invoice number did not find any records.
	E00127     = 139 // original "E00127"
	EnUsE00127 = "Invoice does not exist"
	// Requested action is not allowed due to current status of the object.
	E00128     = 140 // original "E00128"
	EnUsE00128 = "Requested action is not allowed"
	// Failed sending email.
	E00129     = 141 // original "E00129"
	EnUsE00129 = "Failed sending email"
	// Valid Customer Profile ID or Email is required
	E00130     = 142 // original "E00130"
	EnUsE00130 = "Valid Customer Profile ID or Email is required"
	// Invoice created but failed send email and update status
	E00131     = 143 // original "E00131"
	EnUsE00131 = "Invoice created but not processed completely"
	// The payment gateway account is not enabled for Invoicing or CIM service.
	E00132     = 144 // original "E00132"
	EnUsE00132 = "Invoicing or CIM service is not enabled."
	// Server error
	E00133     = 145 // original "E00133"
	EnUsE00133 = "Server error."
	// Due date is past date or not specified.
	E00134     = 146 // original "E00134"
	EnUsE00134 = "Due date is invalid"
	// Merchant has not yet provided processor information to set test mode flag to false.
	E00135     = 147 // original "E00135"
	EnUsE00135 = "Merchant has not provided processor information."
	// Processor account has not been setup to set test mode flag to false.
	E00136     = 148 // original "E00136"
	EnUsE00136 = "Processor account is still in process, please try again later."
	// Only either CreditCard or Bank is allowed.
	E00137     = 149 // original "E00137"
	EnUsE00137 = "Multiple payment types are not allowed."
	// Payment and Shipping Profile IDs cannot be specified when requesting a hosted payment page.
	E00138     = 150 // original "E00138"
	EnUsE00138 = "Payment and Shipping Profile IDs cannot be specified when requesting a hosted payment page."
	// The Access token does not have permission to call the API method.
	E00139     = 151 // original "E00139"
	EnUsE00139 = "Access denied. Access Token does not have correct permissions for this API."
	// Reference Id not found.
	E00140     = 152 // original "E00140"
	EnUsE00140 = "Reference Id not found"
	// Payment Profile creation with this OpaqueData descriptor requires transactionMode to be set to liveMode.
	E00141     = 153 // original "E00141"
	EnUsE00141 = "Payment Profile creation with this OpaqueData descriptor requires transactionMode to be set to liveMode."
	// RecurringBilling setting is a required field for recurring tokenized payment transactions.
	E00142     = 154 // original "E00142"
	EnUsE00142 = "RecurringBilling setting is a required field for recurring tokenized payment transactions."
	// Failed to parse the MerchantId to integer
	E00143     = 155 // original "E00143"
	EnUsE00143 = "Failed to parse MerchantId to integer"
	// We are currently holding the last transaction for review. Before you reactivate the subscription, review the transaction.
	// Alternately, the merchant can log into the <a href&#x3D;`https://account.authorize.net/`>Merchant Interface</a> and click Tools > Fraud Detection Suite to view the Suspicious Transaction Reports and approve or decline the held transaction.
	E00144     = 156 // original "E00144"
	EnUsE00144 = "We are currently holding the last transaction for review. Before you reactivate the subscription, review the transaction."
	// This invoice has been canceled by the sender. Please contact the sender directly if you have questions.
	E00145     = 157 // original "E00145"
	EnUsE00145 = "This invoice has been canceled by the sender. Please contact the sender directly if you have questions. "
	Err0       = 158 // original "0"
	EnUsErr0   = "Unknown Error."
	Err1       = 159 // original "1"
	EnUsErr1   = "This transaction has been approved."
	Err2       = 160 // original "2"
	EnUsErr2   = "This transaction has been declined."
	Err3       = 161 // original "3"
	EnUsErr3   = "This transaction has been declined."
	// The code returned from the processor indicating that the card used needs to be picked up.
	Err4     = 162 // original "4"
	EnUsErr4 = "This transaction has been declined."
	// The value submitted in the amount field did not pass validation for a number.
	Err5     = 163 // original "5"
	EnUsErr5 = "A valid amount is required."
	Err6     = 164 // original "6"
	EnUsErr6 = "The credit card number is invalid."
	// The format of the date submitted was incorrect.
	Err7     = 165 // original "7"
	EnUsErr7 = "Credit card expiration date is invalid."
	Err8     = 166 // original "8"
	EnUsErr8 = "The credit card has expired."
	// The value submitted in the routingNumber field did not pass validation or was not for a valid financial institution.
	Err9     = 167 // original "9"
	EnUsErr9 = "The ABA code is invalid"
	// The value submitted in the accountNumber field did not pass validation.
	Err10     = 168 // original "10"
	EnUsErr10 = "The account number is invalid"
	// A transaction with identical amount and credit card information was submitted within the previous two minutes.
	Err11     = 169 // original "11"
	EnUsErr11 = "A duplicate transaction has been submitted."
	// The transaction request required the field authCode but either it was not submitted, or it was submitted without a value.
	Err12     = 170 // original "12"
	EnUsErr12 = "An authorization code is required but not present."
	Err13     = 171 // original "13"
	EnUsErr13 = "The merchant login ID or password is invalid or the account is inactive."
	// <para>Applicable only to SIM API. The Relay Response or Referrer URL does not match the merchant&#x27;s configured value(s) or is absent.</para>
	//
	//<para><b>NOTE:</b> Parameterized URLs are not permitted.</para>
	Err14     = 172 // original "14"
	EnUsErr14 = "The referrer, relay response or receipt link URL is invalid."
	// The transaction ID value is non-numeric or was not present for a transaction that requires it (i.e., VOID, PRIOR_AUTH_CAPTURE, and CREDIT).
	Err15     = 173 // original "15"
	EnUsErr15 = "The transaction ID is invalid or not present."
	// The transaction ID sent in was properly formatted but the gateway had no record of the transaction.
	Err16     = 174 // original "16"
	EnUsErr16 = "The transaction cannot be found."
	// If you encounter this error on an Authorize.Net Sandbox account, please contact <a href&x3D;`https://developer.authorize.net/support/contact_us/&quot>Developer Support</a> to enable this card type on your account.
	//
	Err17     = 175 // original "17"
	EnUsErr17 = "The merchant does not accept this type of credit card."
	// The merchant does not accept electronic checks.
	Err18     = 176 // original "18"
	EnUsErr18 = "ACH transactions are not accepted by this merchant."
	Err19     = 177 // original "19"
	EnUsErr19 = "An error occurred during processing.  Please try again."
	Err20     = 178 // original "20"
	EnUsErr20 = "An error occurred during processing.  Please try again."
	Err21     = 179 // original "21"
	EnUsErr21 = "An error occurred during processing.  Please try again."
	Err22     = 180 // original "22"
	EnUsErr22 = "An error occurred during processing.  Please try again."
	Err23     = 181 // original "23"
	EnUsErr23 = "An error occurred during processing.  Please try again."
	Err24     = 182 // original "24"
	EnUsErr24 = "The Elavon bank number or terminal ID is incorrect. Call Merchant Service Provider."
	Err25     = 183 // original "25"
	EnUsErr25 = "An error occurred during processing.  Please try again."
	Err26     = 184 // original "26"
	EnUsErr26 = "An error occurred during processing.  Please try again."
	Err27     = 185 // original "27"
	EnUsErr27 = "The transaction has been declined because of an AVS mismatch. The address provided does not match billing address of cardholder."
	// If you encounter this error on an Authorize.Net Sandbox account, please contact <a href&x3D;`https://developer.authorize.net/support/contact_us/&quot>Developer Support</a> for assistance.
	//
	Err28     = 186 // original "28"
	EnUsErr28 = "The merchant does not accept this type of credit card."
	// Invalid Paymentech client number, merchant number or terminal number.
	Err29     = 187 // original "29"
	EnUsErr29 = "The Paymentech identification numbers are incorrect. Call Merchant Service Provider."
	Err30     = 188 // original "30"
	EnUsErr30 = "The configuration with processor is invalid. Call Merchant Service Provider."
	// The merchant was incorrectly set up at the processor.
	Err31     = 189 // original "31"
	EnUsErr31 = "The FDC Merchant ID or Terminal ID is incorrect. Call Merchant Service Provider."
	Err32     = 190 // original "32"
	EnUsErr32 = "The merchant password is invalid or not present."
	// This error indicates that a field the merchant specified as required was not filled in.
	Err33     = 191 // original "33"
	EnUsErr33 = "%s cannot be left blank."
	// The merchant was incorrectly set up at the processor.
	Err34     = 192 // original "34"
	EnUsErr34 = "The VITAL identification numbers are incorrect. Call Merchant Service Provider."
	// The merchant was incorrectly set up at the processor.
	Err35     = 193 // original "35"
	EnUsErr35 = "An error occurred during processing. Call Merchant Service Provider."
	// The customer was approved at the time of authorization, but failed at settlement.
	Err36     = 194 // original "36"
	EnUsErr36 = "The authorization was approved but settlement failed."
	Err37     = 195 // original "37"
	EnUsErr37 = "The credit card number is invalid."
	// The merchant was incorrectly set up at the processor.
	Err38     = 196 // original "38"
	EnUsErr38 = "The Global Payment System identification numbers are incorrect. Call Merchant Service Provider."
	Err39     = 197 // original "39"
	EnUsErr39 = "The supplied currency code is either invalid, not supported, not allowed for this merchant or doesnt have an exchange rate."
	Err40     = 198 // original "40"
	EnUsErr40 = "This transaction must be encrypted."
	// Only merchants set up for the FraudScreen.Net service would receive this decline. This code will be returned if a given transaction&#x27;s fraud score is higher than the threshold set by the merchant.
	Err41     = 199 // original "41"
	EnUsErr41 = "This transaction has been declined."
	// This is applicable only to merchants processing through the Wells Fargo SecureSource product who have requirements for transaction submission that are different from merchants not processing through Wells Fargo.
	Err42     = 200 // original "42"
	EnUsErr42 = "There is missing or invalid information in a required field."
	// The merchant was incorrectly set up at the processor.
	Err43     = 201 // original "43"
	EnUsErr43 = "The merchant was incorrectly set up at the processor. Call Merchant Service Provider."
	// The card code submitted with the transaction did not match the card code on file at the card issuing bank and the transaction was declined.
	Err44     = 202 // original "44"
	EnUsErr44 = "This transaction has been declined."
	// This error would be returned if the transaction received a code from the processor that matched the rejection criteria set by the merchant for both the AVS and Card Code filters.
	Err45     = 203 // original "45"
	EnUsErr45 = "This transaction has been declined."
	Err46     = 204 // original "46"
	EnUsErr46 = "Your session has expired or does not exist. You must log in again to continue working."
	// This occurs if the merchant tries to capture funds greater than the amount of the original authorization-only transaction.
	Err47     = 205 // original "47"
	EnUsErr47 = "The amount requested for settlement cannot be greater than the original amount authorized."
	// The merchant attempted to settle for less than the originally authorized amount.
	Err48     = 206 // original "48"
	EnUsErr48 = "This processor does not accept partial reversals."
	Err49     = 207 // original "49"
	EnUsErr49 = "The transaction amount submitted was greater than the maximum amount allowed."
	// Credits or refunds may only be performed against settled transactions. The transaction against which the credit/refund was submitted has not been settled, so a credit cannot be issued.
	Err50     = 208 // original "50"
	EnUsErr50 = "This transaction is awaiting settlement and cannot be refunded."
	Err51     = 209 // original "51"
	EnUsErr51 = "The sum of all credits against this transaction is greater than the original transaction amount."
	Err52     = 210 // original "52"
	EnUsErr52 = "The transaction was authorized but the client could not be notified; it will not be settled."
	// If payment type is bankAccount, transactionType cannot be set to captureOnlyTransaction.
	Err53     = 211 // original "53"
	EnUsErr53 = "The transaction type is invalid for ACH transactions."
	Err54     = 212 // original "54"
	EnUsErr54 = "The referenced transaction does not meet the criteria for issuing a credit."
	// The transaction is rejected if the sum of this credit and prior credits exceeds the original debit amount.
	Err55     = 213 // original "55"
	EnUsErr55 = "The sum of credits against the referenced transaction would exceed original debit amount."
	// The merchant processes eCheck.Net transactions only and does not accept credit cards.
	Err56     = 214 // original "56"
	EnUsErr56 = "Credit card transactions are not accepted by this merchant."
	Err57     = 215 // original "57"
	EnUsErr57 = "An error occurred during processing.  Please try again."
	Err58     = 216 // original "58"
	EnUsErr58 = "An error occurred during processing.  Please try again."
	Err59     = 217 // original "59"
	EnUsErr59 = "An error occurred during processing.  Please try again."
	Err60     = 218 // original "60"
	EnUsErr60 = "An error occurred during processing.  Please try again."
	Err61     = 219 // original "61"
	EnUsErr61 = "An error occurred during processing.  Please try again."
	Err62     = 220 // original "62"
	EnUsErr62 = "An error occurred during processing.  Please try again."
	Err63     = 221 // original "63"
	EnUsErr63 = "An error occurred during processing.  Please try again."
	// This error is applicable to Wells Fargo SecureSource merchants only. Credits or refunds cannot be issued against transactions that were not authorized.
	Err64     = 222 // original "64"
	EnUsErr64 = "The referenced transaction was not approved."
	// The transaction was declined because the merchant configured their account through the Merchant Interface to reject transactions with certain values for a Card Code mismatch.
	Err65     = 223 // original "65"
	EnUsErr65 = "This transaction has been declined."
	// If you are using the SIM connection method, make sure your code is providing values for the SIM required fields listed below:
	//  <ul>      The sequence number of the transaction (x_fp_sequence)       The time when the sequence number was generated (x_fp_timestamp)    The Fingerprint Hash (x_fp_hash)
	Err66     = 224 // original "66"
	EnUsErr66 = "This transaction cannot be accepted for processing."
	// This error code is applicable to merchants using the Wells Fargo SecureSource product only. This product does not allow transactions of type CAPTURE_ONLY.
	Err67     = 225 // original "67"
	EnUsErr67 = "The given transaction type is not supported for this merchant."
	// The value submitted in x_version was invalid.
	Err68     = 226 // original "68"
	EnUsErr68 = "The version parameter is invalid"
	// The value submitted in transactionType was invalid.
	Err69     = 227 // original "69"
	EnUsErr69 = "The transaction type is invalid"
	// The value submitted in x_method was invalid.
	Err70     = 228 // original "70"
	EnUsErr70 = "The transaction method is invalid."
	// The value submitted in accountType was invalid.
	Err71     = 229 // original "71"
	EnUsErr71 = "The bank account type is invalid."
	// The value submitted in authCode was more than six characters in length.
	Err72     = 230 // original "72"
	EnUsErr72 = "The authorization code is invalid."
	// The format of the value submitted in x_drivers_license_num was invalid.
	Err73     = 231 // original "73"
	EnUsErr73 = "The drivers license date of birth is invalid."
	// The value submitted in duty failed format validation.
	Err74     = 232 // original "74"
	EnUsErr74 = "The duty amount is invalid."
	// The value submitted in freight failed format validation.
	Err75     = 233 // original "75"
	EnUsErr75 = "The freight amount is invalid."
	// The value submitted in tax failed format validation.
	Err76     = 234 // original "76"
	EnUsErr76 = "The tax amount is invalid."
	// The value submitted in x_customer_tax_id failed validation.
	Err77     = 235 // original "77"
	EnUsErr77 = "The SSN or tax ID is invalid."
	// The value submitted in cardCode failed format validation.
	Err78     = 236 // original "78"
	EnUsErr78 = "The card code is invalid."
	// The value submitted in x_drivers_license_num failed format validation.
	Err79     = 237 // original "79"
	EnUsErr79 = "The drivers license number is invalid."
	// The value submitted in x_drivers_license_state failed format validation.
	Err80     = 238 // original "80"
	EnUsErr80 = "The drivers license state is invalid."
	// The merchant requested an integration method not compatible with the AIM API.
	Err81     = 239 // original "81"
	EnUsErr81 = "The requested form type is invalid."
	// The system no longer supports version 2.5; requests cannot be posted to scripts.
	Err82     = 240 // original "82"
	EnUsErr82 = "Scripts are only supported in version 2.5."
	// The system no longer supports version 2.5; requests cannot be posted to scripts.
	Err83     = 241 // original "83"
	EnUsErr83 = "The requested script is either invalid or no longer supported."
	// Invalid value for deviceType.
	Err84     = 242 // original "84"
	EnUsErr84 = "The device type is invalid or missing."
	// Invalid value for marketType.
	Err85     = 243 // original "85"
	EnUsErr85 = "The market type is invalid"
	// Invalid value for x_response_format.
	Err86     = 244 // original "86"
	EnUsErr86 = "The Response Format is invalid"
	Err87     = 245 // original "87"
	EnUsErr87 = "Transactions of this market type cannot be processed on this system."
	// Invalid value for track1.
	Err88     = 246 // original "88"
	EnUsErr88 = "Track1 data is not in a valid format."
	// Invalid value for track2.
	Err89     = 247 // original "89"
	EnUsErr89 = "Track2 data is not in a valid format."
	Err90     = 248 // original "90"
	EnUsErr90 = "ACH transactions cannot be accepted by this system."
	Err91     = 249 // original "91"
	EnUsErr91 = "Version 2.5 is no longer supported."
	//
	Err92     = 250 // original "92"
	EnUsErr92 = "The gateway no longer supports the requested method of integration."
	// This code is applicable to Wells Fargo SecureSource merchants only. Country is a required field and must contain the value of a supported country.
	Err93     = 251 // original "93"
	EnUsErr93 = "A valid country is required."
	// This code is applicable to Wells Fargo SecureSource merchants only.
	Err94     = 252 // original "94"
	EnUsErr94 = "The shipping state or country is invalid."
	// This code is applicable to Wells Fargo SecureSource merchants only.
	Err95     = 253 // original "95"
	EnUsErr95 = "A valid state is required."
	// This code is applicable to Wells Fargo SecureSource merchants only. Country is a required field and must contain the value of a supported country.
	Err96     = 254 // original "96"
	EnUsErr96 = "This country is not authorized for buyers."
	// Applicable only to the SIM API. Fingerprints are only valid for a short period of time. This code indicates that the transaction fingerprint has expired.
	Err97     = 255 // original "97"
	EnUsErr97 = "This transaction cannot be accepted."
	// Applicable only to the SIM API. The transaction fingerprint has already been used.
	Err98     = 256 // original "98"
	EnUsErr98 = "This transaction cannot be accepted."
	// Applicable only to the SIM API. The server-generated fingerprint does not match the merchant-specified fingerprint in the x_fp_hash field.
	Err99     = 257 // original "99"
	EnUsErr99 = "This transaction cannot be accepted."
	// Applicable only to eCheck.Net. The value specified in the echeckType field is invalid.
	Err100     = 258 // original "100"
	EnUsErr100 = "The eCheck type parameter is invalid."
	// Applicable only to eCheck.Net. The specified name on the account and/or the account type do not match the NOC record for this account.
	Err101     = 259 // original "101"
	EnUsErr101 = "The given name on the account and/or the account type does not match the actual account."
	// A transaction key was submitted with this WebLink request.
	Err102     = 260 // original "102"
	EnUsErr102 = "This request cannot be accepted."
	Err103     = 261 // original "103"
	EnUsErr103 = "This transaction cannot be accepted."
	// Applicable only to eCheck.Net. The value submitted for country failed validation.
	Err104     = 262 // original "104"
	EnUsErr104 = "The transaction is currently under review."
	// Applicable only to eCheck.Net. The values submitted for city and country failed validation.
	Err105     = 263 // original "105"
	EnUsErr105 = "The transaction is currently under review."
	// Applicable only to eCheck.Net. The value submitted for company failed validation.
	Err106     = 264 // original "106"
	EnUsErr106 = "The transaction is currently under review."
	// The value submitted for bank account name failed validation.
	Err107     = 265 // original "107"
	EnUsErr107 = "The transaction is currently under review."
	// Applicable only to eCheck.Net. The values submitted for firstName and lastName failed validation.
	Err108     = 266 // original "108"
	EnUsErr108 = "The transaction is currently under review."
	// Applicable only to eCheck.Net. The values submitted for firstName and lastName failed validation.
	Err109     = 267 // original "109"
	EnUsErr109 = "The transaction is currently under review."
	// The value submitted for accountName does not contain valid characters.
	Err110     = 268 // original "110"
	EnUsErr110 = "The transaction is currently under review."
	// This code is applicable to Wells Fargo SecureSource merchants only.
	Err111     = 269 // original "111"
	EnUsErr111 = "A valid billing country is required."
	// This code is applicable to Wells Fargo SecureSource merchants only.
	Err112     = 270 // original "112"
	EnUsErr112 = "A valid billing state/province is required."
	Err113     = 271 // original "113"
	EnUsErr113 = "The commercial card type is invalid."
	Err114     = 272 // original "114"
	EnUsErr114 = "The merchant account is in test mode. This automated payment will not be processed."
	Err115     = 273 // original "115"
	EnUsErr115 = "The merchant account is not active. This automated payment will not be processed."
	// This code is applicable only to merchants that include the authenticationIndicator in the transaction request.  The ECI value for a Visa transaction; or the UCAF indicator for a Mastercard transaction submitted in the authenticationIndicator field is invalid.
	Err116     = 274 // original "116"
	EnUsErr116 = "The authentication indicator is invalid."
	// This code is applicable only to merchants that include the cardholderAuthenticationValue in the transaction request. The CAVV for a Visa transaction or the AVV/UCAF for a Mastercard transaction is invalid or contains an invalid character.
	Err117     = 275 // original "117"
	EnUsErr117 = "The cardholder authentication value is invalid."
	// This code is applicable only to merchants that include the authenticationIndicator and cardholderAuthenticationValue in the transaction request. The combination of authenticationIndicator and cardholderAuthenticationValue is invalid.
	Err118     = 276 // original "118"
	EnUsErr118 = "The combination of card type, authentication indicator and cardholder authentication value is invalid."
	// This code is applicable only to merchants that include the authenticationIndicator and recurringBilling in the transaction request. Transactions submitted with a value in authenticationIndicator while recurringBilling is set to true will be rejected.
	Err119     = 277 // original "119"
	EnUsErr119 = "Transactions having cardholder authentication values cannot be marked as recurring."
	// The system-generated void for the original timed-out transaction failed. The original transaction timed out while waiting for a response from the authorizer.
	Err120     = 278 // original "120"
	EnUsErr120 = "An error occurred during processing. Please try again."
	// The system-generated void for the original errored transaction failed. The original transaction experienced a database error.
	Err121     = 279 // original "121"
	EnUsErr121 = "An error occurred during processing. Please try again."
	// The system-generated void for the original errored transaction failed. The original transaction experienced a processing error.
	Err122     = 280 // original "122"
	EnUsErr122 = "An error occurred during processing. Please try again."
	// The transaction request must include the API login ID associated with the payment gateway account.
	Err123     = 281 // original "123"
	EnUsErr123 = "This account has not been given the permission(s) required for this request."
	Err124     = 282 // original "124"
	EnUsErr124 = "This processor does not accept recurring transactions."
	Err125     = 283 // original "125"
	EnUsErr125 = "The surcharge amount is invalid."
	Err126     = 284 // original "126"
	EnUsErr126 = "The Tip amount is invalid."
	// The system-generated void for the original AVS-rejected transaction failed.
	Err127     = 285 // original "127"
	EnUsErr127 = "The transaction resulted in an AVS mismatch. The address provided does not match billing address of cardholder."
	// The customer&#x27;s financial institution does not currently allow transactions for this account.
	Err128     = 286 // original "128"
	EnUsErr128 = "This transaction cannot be processed."
	// The payment gateway account status is Blacklisted.
	Err130     = 287 // original "130"
	EnUsErr130 = "This merchant account has been closed."
	// The payment gateway account status is Suspended-STA.
	Err131     = 288 // original "131"
	EnUsErr131 = "This transaction cannot be accepted at this time."
	// The payment gateway account status is Suspended - Blacklist.
	Err132     = 289 // original "132"
	EnUsErr132 = "This transaction cannot be accepted at this time."
	// The system-generated void for the original FraudScreen-rejected transaction failed.
	Err141     = 290 // original "141"
	EnUsErr141 = "This transaction has been declined."
	// The system-generated void for the original card code-rejected and AVS-rejected transaction failed.
	Err145     = 291 // original "145"
	EnUsErr145 = "This transaction has been declined."
	// The system-generated void for the original transaction failed. The response for the original transaction could not be communicated to the client.
	Err152     = 292 // original "152"
	EnUsErr152 = "The transaction was authorized but the client could not be notified; it will not be settled."
	// <ul>Set marketType to `0` to flag the transaction as e-commerce.Set transactionType to authCaptureTransaction or authOnlyTransaction.Specify both opaque data parameters.Do not include card number, expiration date, or track data.Do not include 3DS data. Ensure that you submit data that can be successfully decrypted.Only submit decrypted data that belongs to the merchant submitting the request.Encode the submitted data in Base64. </ul>
	Err153     = 293 // original "153"
	EnUsErr153 = "There was an error processing the payment data."
	Err154     = 294 // original "154"
	EnUsErr154 = "Processing Apple Payments is not enabled for this merchant account."
	Err155     = 295 // original "155"
	EnUsErr155 = "This processor does not support this method of submitting payment data."
	Err156     = 296 // original "156"
	EnUsErr156 = "The cryptogram is either invalid or cannot be used in combination with other parameters."
	// System void failed. CVV2 Code mismatch based on the CVV response and the merchant settings.
	Err165     = 297 // original "165"
	EnUsErr165 = "This transaction has been declined."
	// Concord EFS - Provisioning at the processor has not been completed.
	Err170     = 298 // original "170"
	EnUsErr170 = "An error occurred during processing. Please contact the merchant."
	// Concord EFS - This request is invalid.
	Err171     = 299 // original "171"
	EnUsErr171 = "An error occurred during processing. Please contact the merchant."
	// Concord EFS - The store ID is invalid.
	Err172     = 300 // original "172"
	EnUsErr172 = "An error occurred during processing. Please contact the merchant."
	// Concord EFS - The store key is invalid.
	Err173     = 301 // original "173"
	EnUsErr173 = "An error occurred during processing. Please contact the merchant."
	// Concord EFS - This transaction type is not accepted by the processor.
	Err174     = 302 // original "174"
	EnUsErr174 = "The transaction type is invalid. Please contact the merchant."
	// Concord EFS - This transaction is not allowed. The Concord EFS processing platform does not support voiding credit transactions. Please debit the credit card instead of voiding the credit.
	Err175     = 303 // original "175"
	EnUsErr175 = "This processor does not allow voiding of credits."
	// The processor response format is invalid.
	Err180     = 304 // original "180"
	EnUsErr180 = "An error occurred during processing.  Please try again."
	// The system-generated void for the original invalid transaction failed. (The original transaction included an invalid processor response format.)
	Err181     = 305 // original "181"
	EnUsErr181 = "An error occurred during processing.  Please try again."
	Err182     = 306 // original "182"
	EnUsErr182 = "One or more of the sub-merchant values are invalid."
	Err183     = 307 // original "183"
	EnUsErr183 = "One or more of the required sub-merchant values are missing."
	Err184     = 308 // original "184"
	EnUsErr184 = "Invalid Token Requestor Name."
	// Merchant is not configured for VPOS.
	Err185     = 309 // original "185"
	EnUsErr185 = "This transaction cannot be processed."
	Err186     = 310 // original "186"
	EnUsErr186 = "Invalid Token Requestor ID Length."
	Err187     = 311 // original "187"
	EnUsErr187 = "Invalid Token Requestor ECI Length."
	Err191     = 312 // original "191"
	EnUsErr191 = "This transaction has been declined."
	Err192     = 313 // original "192"
	EnUsErr192 = "An error occurred during processing. Please try again."
	Err193     = 314 // original "193"
	EnUsErr193 = "The transaction is currently under review."
	Err195     = 315 // original "195"
	EnUsErr195 = "One or more of the HTML type configuration fields do not appear to be safe."
	// This error code applies only to merchants on FDC Omaha. The credit card number is invalid.
	Err200     = 316 // original "200"
	EnUsErr200 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The expiration date is invalid.
	Err201     = 317 // original "201"
	EnUsErr201 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The transaction type is invalid.
	Err202     = 318 // original "202"
	EnUsErr202 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The value submitted in the amount field is invalid.
	Err203     = 319 // original "203"
	EnUsErr203 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The department code is invalid.
	Err204     = 320 // original "204"
	EnUsErr204 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The value submitted in the merchant number field is invalid.
	Err205     = 321 // original "205"
	EnUsErr205 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The merchant is not on file.
	Err206     = 322 // original "206"
	EnUsErr206 = "This transaction has been declined."
	// This error code applies only to merchants on FDC Omaha. The merchant account is closed.
	Err207     = 323 // original "207"
	EnUsErr207 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The merchant is not on file.
	Err208     = 324 // original "208"
	EnUsErr208 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. Communication with the processor could not be established.
	Err209     = 325 // original "209"
	EnUsErr209 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The merchant type is incorrect.
	Err210     = 326 // original "210"
	EnUsErr210 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The cardholder is not on file.
	Err211     = 327 // original "211"
	EnUsErr211 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The bank configuration is not on file.
	Err212     = 328 // original "212"
	EnUsErr212 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The merchant assessment code is incorrect.
	Err213     = 329 // original "213"
	EnUsErr213 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. This function is currently unavailable.
	Err214     = 330 // original "214"
	EnUsErr214 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The encrypted PIN field format is invalid.
	Err215     = 331 // original "215"
	EnUsErr215 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The ATM term ID is invalid.
	Err216     = 332 // original "216"
	EnUsErr216 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. This transaction experienced a general message format problem.
	Err217     = 333 // original "217"
	EnUsErr217 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The PIN block format or PIN availability value is invalid.
	Err218     = 334 // original "218"
	EnUsErr218 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The ETC void is unmatched.
	Err219     = 335 // original "219"
	EnUsErr219 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The primary CPU is not available.
	Err220     = 336 // original "220"
	EnUsErr220 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. The SE number is invalid.
	Err221     = 337 // original "221"
	EnUsErr221 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. Duplicate auth request (from INAS).
	Err222     = 338 // original "222"
	EnUsErr222 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. This transaction experienced an unspecified error.
	Err223     = 339 // original "223"
	EnUsErr223 = "This transaction has been declined"
	// This error code applies only to merchants on FDC Omaha. Please re-enter the transaction.
	Err224     = 340 // original "224"
	EnUsErr224 = "This transaction has been declined"
	// The transaction has an invalid dynamic currency conversion (DCC) action.
	Err225     = 341 // original "225"
	EnUsErr225 = "This transaction cannot be processed."
	// Incomplete set of Dynamic Currency Conversion (DCC) parameters.
	Err226     = 342 // original "226"
	EnUsErr226 = "This transaction cannot be processed."
	// Merchant is not configured for Dynamic Currency Conversion (DCC).
	Err227     = 343 // original "227"
	EnUsErr227 = "This transaction cannot be processed."
	// Dynamic Currency Conversion (DCC) is not allowed for this transaction type.
	Err228     = 344 // original "228"
	EnUsErr228 = "This transaction cannot be processed."
	Err229     = 345 // original "229"
	EnUsErr229 = "Conversion rate for this card is available."
	Err230     = 346 // original "230"
	EnUsErr230 = "This transaction cannot be processed."
	// Incoming data is different than the referenced Dynamic Currency Conversion (DCC) transaction.
	Err231     = 347 // original "231"
	EnUsErr231 = "This transaction cannot be processed."
	// Merchant is not configured for Dynamic Currency Conversion (DCC).
	Err232     = 348 // original "232"
	EnUsErr232 = "This transaction cannot be processed."
	// Cannot perform Dynamic Currency Conversion (DCC) action on this transaction.
	Err233     = 349 // original "233"
	EnUsErr233 = "This transaction cannot be processed."
	// This transaction has been voided.
	Err234     = 350 // original "234"
	EnUsErr234 = "This transaction cannot be processed."
	// This transaction has been captured previously.
	Err235     = 351 // original "235"
	EnUsErr235 = "This transaction cannot be processed."
	// Dynamic Currency Conversion (DCC) data for the referenced transaction is not available.
	Err236     = 352 // original "236"
	EnUsErr236 = "This transaction cannot be processed."
	// The referenced transaction has expired.
	Err237     = 353 // original "237"
	EnUsErr237 = "This transaction cannot be processed."
	// The transaction version does not support Dynamic Currency Conversion (DCC).
	Err238     = 354 // original "238"
	EnUsErr238 = "This transaction cannot be processed."
	// The response format for this transaction does not support Dynamic Currency Conversion (DCC).
	Err239     = 355 // original "239"
	EnUsErr239 = "This transaction cannot be processed."
	// Currency for Dynamic Currency Conversion (DCC) transactions must be US dollars.
	Err240     = 356 // original "240"
	EnUsErr240 = "This transaction cannot be processed."
	// Invalid response from card processor.
	Err241     = 357 // original "241"
	EnUsErr241 = "This transaction cannot be processed."
	// Recurring billing flag not allowed on Dynamic Currency Conversion (DCC).
	Err242     = 358 // original "242"
	EnUsErr242 = "This transaction cannot be processed."
	// The combination of values submitted for recurringBilling and echeckType is not allowed.
	Err243     = 359 // original "243"
	EnUsErr243 = "Recurring billing is not allowed for this eCheck.Net type."
	// The combination of values submitted for accountType and echeckType is not allowed.
	Err244     = 360 // original "244"
	EnUsErr244 = "This eCheck.Net type is not allowed for this Bank Account Type."
	// The value submitted for echeckType is not allowed when using the payment gateway hosted payment form.
	Err245     = 361 // original "245"
	EnUsErr245 = "This eCheck.Net type is not allowed when using the payment gateway hosted payment form."
	// The merchant&#x27;s payment gateway account is not enabled to submit the eCheck.Net type.
	Err246     = 362 // original "246"
	EnUsErr246 = "This eCheck.Net type is not allowed."
	// The combination of values submitted for transactionType and echeckType is not allowed.
	Err247     = 363 // original "247"
	EnUsErr247 = "This eCheck.Net type is not allowed."
	// Invalid check number. Check numbers can only consist of letters and numbers, up to 15 characters.
	Err248     = 364 // original "248"
	EnUsErr248 = "The check number is invalid."
	// This transaction was submitted from a blocked IP address.
	Err250     = 365 // original "250"
	EnUsErr250 = "This transaction has been declined."
	// The transaction was declined as a result of triggering a Fraud Detection Suite filter.
	Err251     = 366 // original "251"
	EnUsErr251 = "This transaction has been declined."
	// The transaction was accepted, but is being held for merchant review.  The merchant may customize the customer response in the Merchant Interface.
	Err252     = 367 // original "252"
	EnUsErr252 = "Your order has been received. Thank you for your business!"
	// The transaction was accepted and was authorized, but is being held for merchant review.  The merchant may customize the customer response in the Merchant Interface.
	Err253     = 368 // original "253"
	EnUsErr253 = "Your order has been received. Thank you for your business!"
	// The transaction was declined after manual review.
	Err254     = 369 // original "254"
	EnUsErr254 = "This transaction has been declined."
	// Only authorizations and credits can be reversed.
	Err260     = 370 // original "260"
	EnUsErr260 = "Reversal not supported for this transaction type."
	// The transaction experienced an error during sensitive data encryption and was not processed.  Please try again.
	Err261     = 371 // original "261"
	EnUsErr261 = "An error occurred during processing.  Please try again."
	Err262     = 372 // original "262"
	EnUsErr262 = "The PayformMask is invalid."
	Err265     = 373 // original "265"
	EnUsErr265 = "This transaction cannot be processed."
	// A value submitted in lineItem for the item referenced is invalid.
	Err270     = 374 // original "270"
	EnUsErr270 = "Line item %1 is invalid."
	// The number of line items submitted in lineItem exceeds the allowed maximum of 30.
	Err271     = 375 // original "271"
	EnUsErr271 = "The number of line items submitted is not allowed. A maximum of %1 line items can be submitted."
	Err280     = 376 // original "280"
	EnUsErr280 = "The auction platform name is invalid."
	Err281     = 377 // original "281"
	EnUsErr281 = "The auction platform ID is invalid."
	Err282     = 378 // original "282"
	EnUsErr282 = "The auction listing type is invalid."
	Err283     = 379 // original "283"
	EnUsErr283 = "The auction listing ID is invalid."
	Err284     = 380 // original "284"
	EnUsErr284 = "The auction seller ID is invalid."
	Err285     = 381 // original "285"
	EnUsErr285 = "The auction buyer ID is invalid."
	Err286     = 382 // original "286"
	EnUsErr286 = "One or more required auction values were not submitted."
	Err287     = 383 // original "287"
	EnUsErr287 = "The combination of auction platform ID and auction platform name is invalid."
	Err288     = 384 // original "288"
	EnUsErr288 = "This transaction cannot be accepted."
	Err289     = 385 // original "289"
	EnUsErr289 = "This processor does not accept zero dollar authorization for this card type."
	Err290     = 386 // original "290"
	EnUsErr290 = "There is one or more missing or invalid required fields."
	Err295     = 387 // original "295"
	EnUsErr295 = "The amount of this request was only partially approved on the given prepaid card. An additional payment is required to fulfill the balance of this transaction."
	Err296     = 388 // original "296"
	EnUsErr296 = "The specified SplitTenderID is invalid."
	Err297     = 389 // original "297"
	EnUsErr297 = "Transaction ID and Split Tender ID cannot both be used in the same request."
	Err298     = 390 // original "298"
	EnUsErr298 = "This order has already been released or voided therefore new transaction associations cannot be accepted."
	// Invalid deviceId value.
	Err300     = 391 // original "300"
	EnUsErr300 = "The device ID is invalid."
	// Invalid batchId value.
	Err301     = 392 // original "301"
	EnUsErr301 = "The device batch ID is invalid."
	// Invalid x_reversal value.
	Err302     = 393 // original "302"
	EnUsErr302 = "The reversal flag is invalid."
	// The current device batch must be closed manually from the POS device.
	Err303     = 394 // original "303"
	EnUsErr303 = "The device batch is full. Please close the batch."
	// The original transaction has been settled and cannot be reversed.
	Err304     = 395 // original "304"
	EnUsErr304 = "The original transaction is in a closed batch."
	// This merchant is configured for auto-close and cannot manually close batches.
	Err305     = 396 // original "305"
	EnUsErr305 = "The merchant is configured for auto-close."
	Err306     = 397 // original "306"
	EnUsErr306 = "The batch is already closed."
	Err307     = 398 // original "307"
	EnUsErr307 = "The reversal was processed successfully."
	// The transaction submitted for reversal was not found.
	Err308     = 399 // original "308"
	EnUsErr308 = "Original transaction for reversal not found."
	Err309     = 400 // original "309"
	EnUsErr309 = "The device has been disabled."
	Err310     = 401 // original "310"
	EnUsErr310 = "This transaction has already been voided."
	Err311     = 402 // original "311"
	EnUsErr311 = "This transaction has already been captured."
	Err312     = 403 // original "312"
	EnUsErr312 = "The specified security code was invalid."
	Err313     = 404 // original "313"
	EnUsErr313 = "A new security code was requested."
	Err314     = 405 // original "314"
	EnUsErr314 = "This transaction cannot be processed."
	// This is a processor-issued decline.
	Err315     = 406 // original "315"
	EnUsErr315 = "The credit card number is invalid."
	// This is a processor-issued decline.
	Err316     = 407 // original "316"
	EnUsErr316 = "Credit card expiration date is invalid."
	// This is a processor-issued decline.
	Err317     = 408 // original "317"
	EnUsErr317 = "The credit card has expired."
	// This is a processor-issued decline.
	Err318     = 409 // original "318"
	EnUsErr318 = "A duplicate transaction has been submitted."
	// This is a processor-issued decline.
	Err319     = 410 // original "319"
	EnUsErr319 = "The transaction cannot be found."
	Err320     = 411 // original "320"
	EnUsErr320 = "The device identifier is either not registered or not enabled."
	// <para>Merchants processing transactions via one of the following processors (AIBMS UK, Barclays, Cardnet, HBOS, HSBC, Streamline, FdiAus and Westpac) are required to submit the following billing information fields:</para><ul> First Name (firstName) Last Name (lastName) Address (address) City (city) State (state) Postal Code Country (country) Email Address (email) </ul>    <para><b>NOTE:</b> State and Postal Code are optional if the billing address is not in the U.S. or Canada. If the address is in the U.S. or Canada, the two-digit State/Province code must be provided, along with the Zip/Postal Code.</para>
	Err325      = 412 // original "325"
	EnUsErr325  = "The request data did not pass the required fields check for this application."
	Err326      = 413 // original "326"
	EnUsErr326  = "The request field(s) are either invalid or missing."
	Err327      = 414 // original "327"
	EnUsErr327  = "The void request failed. Either the original transaction type does not support void, or the transaction is in the process of being settled."
	Err328      = 415 // original "328"
	EnUsErr328  = "A validation error occurred at the processor."
	Err330      = 416 // original "330"
	EnUsErr330  = "V.me transactions are not accepted by this merchant."
	Err331      = 417 // original "331"
	EnUsErr331  = "The x_call_id value is missing."
	Err332      = 418 // original "332"
	EnUsErr332  = "The x_call_id value is not found or invalid."
	Err333      = 419 // original "333"
	EnUsErr333  = "A validation error was returned from V.me."
	Err334      = 420 // original "334"
	EnUsErr334  = "The V.me transaction is in an invalid state."
	Err339      = 421 // original "339"
	EnUsErr339  = "Use x_method to specify method or send only x_call_id or card/account information."
	Err340      = 422 // original "340"
	EnUsErr340  = "V.me by Visa does not support voids on captured or credit transactions. Please allow the transaction to settle, then process a refund for the captured transaction."
	Err341      = 423 // original "341"
	EnUsErr341  = "The x_discount value is invalid."
	Err342      = 424 // original "342"
	EnUsErr342  = "The x_giftwrap value is invalid."
	Err343      = 425 // original "343"
	EnUsErr343  = "The x_subtotal value is invalid."
	Err344      = 426 // original "344"
	EnUsErr344  = "The x_misc value is invalid."
	Err350      = 427 // original "350"
	EnUsErr350  = "Country must be a valid two or three-character value if specified."
	Err351      = 428 // original "351"
	EnUsErr351  = "Employee ID must be 1 to %x characters in length."
	Err355      = 429 // original "355"
	EnUsErr355  = "An error occurred while parsing the EMV data."
	Err356      = 430 // original "356"
	EnUsErr356  = "EMV-based transactions are not currently supported for this processor and card type."
	Err357      = 431 // original "357"
	EnUsErr357  = "Opaque Descriptor is required."
	Err358      = 432 // original "358"
	EnUsErr358  = "EMV data is not supported with this transaction type."
	Err359      = 433 // original "359"
	EnUsErr359  = "EMV data is not supported with this market type."
	Err360      = 434 // original "360"
	EnUsErr360  = "An error occurred while decrypting the EMV data."
	Err361      = 435 // original "361"
	EnUsErr361  = "The EMV version is invalid."
	Err362      = 436 // original "362"
	EnUsErr362  = "The EMV version is required."
	Err363      = 437 // original "363"
	EnUsErr363  = "The POS Entry Mode value is invalid."
	Err370      = 438 // original "370"
	EnUsErr370  = "Signature data is too large."
	Err371      = 439 // original "371"
	EnUsErr371  = "Signature must be PNG formatted data."
	Err375      = 440 // original "375"
	EnUsErr375  = "Terminal/lane number must be numeric."
	Err380      = 441 // original "380"
	EnUsErr380  = "KSN is duplicated."
	Err901      = 442 // original "901"
	EnUsErr901  = "This transaction cannot be accepted at this time due to system maintenance.  Please try again later."
	Err2000     = 443 // original "2000"
	EnUsErr2000 = "Need payer consent."
	Err2001     = 444 // original "2001"
	EnUsErr2001 = "PayPal transactions are not accepted by this merchant."
	Err2002     = 445 // original "2002"
	EnUsErr2002 = "PayPal transactions require x_version of at least 3.1."
	Err2003     = 446 // original "2003"
	EnUsErr2003 = "Request completed successfully"
	Err2004     = 447 // original "2004"
	EnUsErr2004 = "Success  URL is required."
	Err2005     = 448 // original "2005"
	EnUsErr2005 = "Cancel URL is required."
	Err2006     = 449 // original "2006"
	EnUsErr2006 = "Payer ID is required."
	Err2007     = 450 // original "2007"
	EnUsErr2007 = "This processor does not accept zero dollar authorizations."
	Err2008     = 451 // original "2008"
	EnUsErr2008 = "The referenced transaction does not meet the criteria for issuing a Continued Authorization."
	Err2009     = 452 // original "2009"
	EnUsErr2009 = "The referenced transaction does not meet the criteria for issuing a Continued Authorization  w/ Auto Capture."
	Err2100     = 453 // original "2100"
	EnUsErr2100 = "PayPal transactions require valid URL for success_url"
	Err2101     = 454 // original "2101"
	EnUsErr2101 = "PayPal transactions require valid URL for cancel_url"
	Err2102     = 455 // original "2102"
	EnUsErr2102 = "Payment not authorized.  Payment has not been authorized by the user."
	Err2103     = 456 // original "2103"
	EnUsErr2103 = "This transaction has already been authorized."
	Err2104     = 457 // original "2104"
	EnUsErr2104 = "The totals of the cart item amounts do not match order amounts. Be sure the total of the payment detail item parameters add up to the order total."
	Err2105     = 458 // original "2105"
	EnUsErr2105 = "PayPal has rejected the transaction.Invalid Payer ID."
	Err2106     = 459 // original "2106"
	EnUsErr2106 = "PayPal has already captured this transaction."
	Err2107     = 460 // original "2107"
	EnUsErr2107 = "PayPal has rejected the transaction. This Payer ID belongs to a different customer."
	Err2108     = 461 // original "2108"
	EnUsErr2108 = "PayPal has rejected the transaction. x_paypal_hdrimg exceeds maximum allowable length."
	Err2109     = 462 // original "2109"
	EnUsErr2109 = "PayPal has rejected the transaction. x_paypal_payflowcolor must be a 6 character hexadecimal value."
	// This error applies to WePay merchants only if the Prior Authorization Capture request amount is different than the Authorization Only request amount.
	Err2200     = 463 // original "2200"
	EnUsErr2200 = "The amount requested for settlement cannot be different than the original amount authorized. Please void transaction if required"
	// This is a work in progress. This message will not be released to production. It's just a dev placeholder.
	ErrUnknown  = 99999 // original "I99999"
	EnUsUnknown = "This feature has not yet been completed. One day it will be but it looks like today is not that day."
)

func (e *ErrCodeEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("invalid error code (empty)")
	}
	switch string(data[1 : len(data)-1]) {
	case "I00001":
		*e = I00001
	case "I00002":
		*e = I00002
	case "I00003":
		*e = I00003
	case "I00004":
		*e = I00004
	case "I00005":
		*e = I00005
	case "I00006":
		*e = I00006
	case "I00007":
		*e = I00007
	case "I00008":
		*e = I00008
	case "I00009":
		*e = I00009
	case "I00010":
		*e = I00010
	case "I00011":
		*e = I00011
	case "E00001":
		*e = E00001
	case "E00002":
		*e = E00002
	case "E00003":
		*e = E00003
	case "E00004":
		*e = E00004
	case "E00005":
		*e = E00005
	case "E00006":
		*e = E00006
	case "E00007":
		*e = E00007
	case "E00008":
		*e = E00008
	case "E00009":
		*e = E00009
	case "E00010":
		*e = E00010
	case "E00011":
		*e = E00011
	case "E00012":
		*e = E00012
	case "E00013":
		*e = E00013
	case "E00014":
		*e = E00014
	case "E00015":
		*e = E00015
	case "E00016":
		*e = E00016
	case "E00017":
		*e = E00017
	case "E00018":
		*e = E00018
	case "E00019":
		*e = E00019
	case "E00020":
		*e = E00020
	case "E00021":
		*e = E00021
	case "E00022":
		*e = E00022
	case "E00023":
		*e = E00023
	case "E00024":
		*e = E00024
	case "E00025":
		*e = E00025
	case "E00026":
		*e = E00026
	case "E00027":
		*e = E00027
	case "E00028":
		*e = E00028
	case "E00029":
		*e = E00029
	case "E00030":
		*e = E00030
	case "E00031":
		*e = E00031
	case "E00032":
		*e = E00032
	case "E00033":
		*e = E00033
	case "E00034":
		*e = E00034
	case "E00035":
		*e = E00035
	case "E00036":
		*e = E00036
	case "E00037":
		*e = E00037
	case "E00038":
		*e = E00038
	case "E00039":
		*e = E00039
	case "E00040":
		*e = E00040
	case "E00041":
		*e = E00041
	case "E00042":
		*e = E00042
	case "E00043":
		*e = E00043
	case "E00044":
		*e = E00044
	case "E00045":
		*e = E00045
	case "E00046":
		*e = E00046
	case "E00047":
		*e = E00047
	case "E00048":
		*e = E00048
	case "E00049":
		*e = E00049
	case "E00050":
		*e = E00050
	case "E00051":
		*e = E00051
	case "E00052":
		*e = E00052
	case "E00053":
		*e = E00053
	case "E00054":
		*e = E00054
	case "E00055":
		*e = E00055
	case "E00056":
		*e = E00056
	case "E00057":
		*e = E00057
	case "E00058":
		*e = E00058
	case "E00059":
		*e = E00059
	case "E00060":
		*e = E00060
	case "E00061":
		*e = E00061
	case "E00062":
		*e = E00062
	case "E00063":
		*e = E00063
	case "E00064":
		*e = E00064
	case "E00065":
		*e = E00065
	case "E00066":
		*e = E00066
	case "E00067":
		*e = E00067
	case "E00068":
		*e = E00068
	case "E00069":
		*e = E00069
	case "E00070":
		*e = E00070
	case "E00071":
		*e = E00071
	case "E00072":
		*e = E00072
	case "E00073":
		*e = E00073
	case "E00074":
		*e = E00074
	case "E00075":
		*e = E00075
	case "E00076":
		*e = E00076
	case "E00077":
		*e = E00077
	case "E00078":
		*e = E00078
	case "E00079":
		*e = E00079
	case "E00080":
		*e = E00080
	case "E00081":
		*e = E00081
	case "E00082":
		*e = E00082
	case "E00083":
		*e = E00083
	case "E00084":
		*e = E00084
	case "E00085":
		*e = E00085
	case "E00086":
		*e = E00086
	case "E00087":
		*e = E00087
	case "E00088":
		*e = E00088
	case "E00089":
		*e = E00089
	case "E00090":
		*e = E00090
	case "E00091":
		*e = E00091
	case "E00092":
		*e = E00092
	case "E00093":
		*e = E00093
	case "E00094":
		*e = E00094
	case "E00095":
		*e = E00095
	case "E00096":
		*e = E00096
	case "E00097":
		*e = E00097
	case "E00098":
		*e = E00098
	case "E00099":
		*e = E00099
	case "E00100":
		*e = E00100
	case "E00101":
		*e = E00101
	case "E00102":
		*e = E00102
	case "E00103":
		*e = E00103
	case "E00104":
		*e = E00104
	case "E00105":
		*e = E00105
	case "E00106":
		*e = E00106
	case "E00107":
		*e = E00107
	case "E00108":
		*e = E00108
	case "E00109":
		*e = E00109
	case "E00110":
		*e = E00110
	case "E00111":
		*e = E00111
	case "E00112":
		*e = E00112
	case "E00113":
		*e = E00113
	case "E00114":
		*e = E00114
	case "E00115":
		*e = E00115
	case "E00116":
		*e = E00116
	case "E00117":
		*e = E00117
	case "E00118":
		*e = E00118
	case "E00119":
		*e = E00119
	case "E00120":
		*e = E00120
	case "E00121":
		*e = E00121
	case "E00122":
		*e = E00122
	case "E00123":
		*e = E00123
	case "E00124":
		*e = E00124
	case "E00125":
		*e = E00125
	case "E00126":
		*e = E00126
	case "E00127":
		*e = E00127
	case "E00128":
		*e = E00128
	case "E00129":
		*e = E00129
	case "E00130":
		*e = E00130
	case "E00131":
		*e = E00131
	case "E00132":
		*e = E00132
	case "E00133":
		*e = E00133
	case "E00134":
		*e = E00134
	case "E00135":
		*e = E00135
	case "E00136":
		*e = E00136
	case "E00137":
		*e = E00137
	case "E00138":
		*e = E00138
	case "E00139":
		*e = E00139
	case "E00140":
		*e = E00140
	case "E00141":
		*e = E00141
	case "E00142":
		*e = E00142
	case "E00143":
		*e = E00143
	case "E00144":
		*e = E00144
	case "E00145":
		*e = E00145
	case "0":
		*e = Err0
	case "1":
		*e = Err1
	case "2":
		*e = Err2
	case "3":
		*e = Err3
	case "4":
		*e = Err4
	case "5":
		*e = Err5
	case "6":
		*e = Err6
	case "7":
		*e = Err7
	case "8":
		*e = Err8
	case "9":
		*e = Err9
	case "10":
		*e = Err10
	case "11":
		*e = Err11
	case "12":
		*e = Err12
	case "13":
		*e = Err13
	case "14":
		*e = Err14
	case "15":
		*e = Err15
	case "16":
		*e = Err16
	case "17":
		*e = Err17
	case "18":
		*e = Err18
	case "19":
		*e = Err19
	case "20":
		*e = Err20
	case "21":
		*e = Err21
	case "22":
		*e = Err22
	case "23":
		*e = Err23
	case "24":
		*e = Err24
	case "25":
		*e = Err25
	case "26":
		*e = Err26
	case "27":
		*e = Err27
	case "28":
		*e = Err28
	case "29":
		*e = Err29
	case "30":
		*e = Err30
	case "31":
		*e = Err31
	case "32":
		*e = Err32
	case "33":
		*e = Err33
	case "34":
		*e = Err34
	case "35":
		*e = Err35
	case "36":
		*e = Err36
	case "37":
		*e = Err37
	case "38":
		*e = Err38
	case "39":
		*e = Err39
	case "40":
		*e = Err40
	case "41":
		*e = Err41
	case "42":
		*e = Err42
	case "43":
		*e = Err43
	case "44":
		*e = Err44
	case "45":
		*e = Err45
	case "46":
		*e = Err46
	case "47":
		*e = Err47
	case "48":
		*e = Err48
	case "49":
		*e = Err49
	case "50":
		*e = Err50
	case "51":
		*e = Err51
	case "52":
		*e = Err52
	case "53":
		*e = Err53
	case "54":
		*e = Err54
	case "55":
		*e = Err55
	case "56":
		*e = Err56
	case "57":
		*e = Err57
	case "58":
		*e = Err58
	case "59":
		*e = Err59
	case "60":
		*e = Err60
	case "61":
		*e = Err61
	case "62":
		*e = Err62
	case "63":
		*e = Err63
	case "64":
		*e = Err64
	case "65":
		*e = Err65
	case "66":
		*e = Err66
	case "67":
		*e = Err67
	case "68":
		*e = Err68
	case "69":
		*e = Err69
	case "70":
		*e = Err70
	case "71":
		*e = Err71
	case "72":
		*e = Err72
	case "73":
		*e = Err73
	case "74":
		*e = Err74
	case "75":
		*e = Err75
	case "76":
		*e = Err76
	case "77":
		*e = Err77
	case "78":
		*e = Err78
	case "79":
		*e = Err79
	case "80":
		*e = Err80
	case "81":
		*e = Err81
	case "82":
		*e = Err82
	case "83":
		*e = Err83
	case "84":
		*e = Err84
	case "85":
		*e = Err85
	case "86":
		*e = Err86
	case "87":
		*e = Err87
	case "88":
		*e = Err88
	case "89":
		*e = Err89
	case "90":
		*e = Err90
	case "91":
		*e = Err91
	case "92":
		*e = Err92
	case "93":
		*e = Err93
	case "94":
		*e = Err94
	case "95":
		*e = Err95
	case "96":
		*e = Err96
	case "97":
		*e = Err97
	case "98":
		*e = Err98
	case "99":
		*e = Err99
	case "100":
		*e = Err100
	case "101":
		*e = Err101
	case "102":
		*e = Err102
	case "103":
		*e = Err103
	case "104":
		*e = Err104
	case "105":
		*e = Err105
	case "106":
		*e = Err106
	case "107":
		*e = Err107
	case "108":
		*e = Err108
	case "109":
		*e = Err109
	case "110":
		*e = Err110
	case "111":
		*e = Err111
	case "112":
		*e = Err112
	case "113":
		*e = Err113
	case "114":
		*e = Err114
	case "115":
		*e = Err115
	case "116":
		*e = Err116
	case "117":
		*e = Err117
	case "118":
		*e = Err118
	case "119":
		*e = Err119
	case "120":
		*e = Err120
	case "121":
		*e = Err121
	case "122":
		*e = Err122
	case "123":
		*e = Err123
	case "124":
		*e = Err124
	case "125":
		*e = Err125
	case "126":
		*e = Err126
	case "127":
		*e = Err127
	case "128":
		*e = Err128
	case "130":
		*e = Err130
	case "131":
		*e = Err131
	case "132":
		*e = Err132
	case "141":
		*e = Err141
	case "145":
		*e = Err145
	case "152":
		*e = Err152
	case "153":
		*e = Err153
	case "154":
		*e = Err154
	case "155":
		*e = Err155
	case "156":
		*e = Err156
	case "165":
		*e = Err165
	case "170":
		*e = Err170
	case "171":
		*e = Err171
	case "172":
		*e = Err172
	case "173":
		*e = Err173
	case "174":
		*e = Err174
	case "175":
		*e = Err175
	case "180":
		*e = Err180
	case "181":
		*e = Err181
	case "182":
		*e = Err182
	case "183":
		*e = Err183
	case "184":
		*e = Err184
	case "185":
		*e = Err185
	case "186":
		*e = Err186
	case "187":
		*e = Err187
	case "191":
		*e = Err191
	case "192":
		*e = Err192
	case "193":
		*e = Err193
	case "195":
		*e = Err195
	case "200":
		*e = Err200
	case "201":
		*e = Err201
	case "202":
		*e = Err202
	case "203":
		*e = Err203
	case "204":
		*e = Err204
	case "205":
		*e = Err205
	case "206":
		*e = Err206
	case "207":
		*e = Err207
	case "208":
		*e = Err208
	case "209":
		*e = Err209
	case "210":
		*e = Err210
	case "211":
		*e = Err211
	case "212":
		*e = Err212
	case "213":
		*e = Err213
	case "214":
		*e = Err214
	case "215":
		*e = Err215
	case "216":
		*e = Err216
	case "217":
		*e = Err217
	case "218":
		*e = Err218
	case "219":
		*e = Err219
	case "220":
		*e = Err220
	case "221":
		*e = Err221
	case "222":
		*e = Err222
	case "223":
		*e = Err223
	case "224":
		*e = Err224
	case "225":
		*e = Err225
	case "226":
		*e = Err226
	case "227":
		*e = Err227
	case "228":
		*e = Err228
	case "229":
		*e = Err229
	case "230":
		*e = Err230
	case "231":
		*e = Err231
	case "232":
		*e = Err232
	case "233":
		*e = Err233
	case "234":
		*e = Err234
	case "235":
		*e = Err235
	case "236":
		*e = Err236
	case "237":
		*e = Err237
	case "238":
		*e = Err238
	case "239":
		*e = Err239
	case "240":
		*e = Err240
	case "241":
		*e = Err241
	case "242":
		*e = Err242
	case "243":
		*e = Err243
	case "244":
		*e = Err244
	case "245":
		*e = Err245
	case "246":
		*e = Err246
	case "247":
		*e = Err247
	case "248":
		*e = Err248
	case "250":
		*e = Err250
	case "251":
		*e = Err251
	case "252":
		*e = Err252
	case "253":
		*e = Err253
	case "254":
		*e = Err254
	case "260":
		*e = Err260
	case "261":
		*e = Err261
	case "262":
		*e = Err262
	case "265":
		*e = Err265
	case "270":
		*e = Err270
	case "271":
		*e = Err271
	case "280":
		*e = Err280
	case "281":
		*e = Err281
	case "282":
		*e = Err282
	case "283":
		*e = Err283
	case "284":
		*e = Err284
	case "285":
		*e = Err285
	case "286":
		*e = Err286
	case "287":
		*e = Err287
	case "288":
		*e = Err288
	case "289":
		*e = Err289
	case "290":
		*e = Err290
	case "295":
		*e = Err295
	case "296":
		*e = Err296
	case "297":
		*e = Err297
	case "298":
		*e = Err298
	case "300":
		*e = Err300
	case "301":
		*e = Err301
	case "302":
		*e = Err302
	case "303":
		*e = Err303
	case "304":
		*e = Err304
	case "305":
		*e = Err305
	case "306":
		*e = Err306
	case "307":
		*e = Err307
	case "308":
		*e = Err308
	case "309":
		*e = Err309
	case "310":
		*e = Err310
	case "311":
		*e = Err311
	case "312":
		*e = Err312
	case "313":
		*e = Err313
	case "314":
		*e = Err314
	case "315":
		*e = Err315
	case "316":
		*e = Err316
	case "317":
		*e = Err317
	case "318":
		*e = Err318
	case "319":
		*e = Err319
	case "320":
		*e = Err320
	case "325":
		*e = Err325
	case "326":
		*e = Err326
	case "327":
		*e = Err327
	case "328":
		*e = Err328
	case "330":
		*e = Err330
	case "331":
		*e = Err331
	case "332":
		*e = Err332
	case "333":
		*e = Err333
	case "334":
		*e = Err334
	case "339":
		*e = Err339
	case "340":
		*e = Err340
	case "341":
		*e = Err341
	case "342":
		*e = Err342
	case "343":
		*e = Err343
	case "344":
		*e = Err344
	case "350":
		*e = Err350
	case "351":
		*e = Err351
	case "355":
		*e = Err355
	case "356":
		*e = Err356
	case "357":
		*e = Err357
	case "358":
		*e = Err358
	case "359":
		*e = Err359
	case "360":
		*e = Err360
	case "361":
		*e = Err361
	case "362":
		*e = Err362
	case "363":
		*e = Err363
	case "370":
		*e = Err370
	case "371":
		*e = Err371
	case "375":
		*e = Err375
	case "380":
		*e = Err380
	case "901":
		*e = Err901
	case "2000":
		*e = Err2000
	case "2001":
		*e = Err2001
	case "2002":
		*e = Err2002
	case "2003":
		*e = Err2003
	case "2004":
		*e = Err2004
	case "2005":
		*e = Err2005
	case "2006":
		*e = Err2006
	case "2007":
		*e = Err2007
	case "2008":
		*e = Err2008
	case "2009":
		*e = Err2009
	case "2100":
		*e = Err2100
	case "2101":
		*e = Err2101
	case "2102":
		*e = Err2102
	case "2103":
		*e = Err2103
	case "2104":
		*e = Err2104
	case "2105":
		*e = Err2105
	case "2106":
		*e = Err2106
	case "2107":
		*e = Err2107
	case "2108":
		*e = Err2108
	case "2109":
		*e = Err2109
	case "2200":
		*e = Err2200
	default:
		*e = ErrUnknown
	}
	return nil
}

func (e ErrCodeEnum) String() string {
	switch e {
	case 1:
		return EnUsI00001
	case 2:
		return EnUsI00002
	case 3:
		return EnUsI00003
	case 4:
		return EnUsI00004
	case 5:
		return EnUsI00005
	case 6:
		return EnUsI00006
	case 7:
		return EnUsI00007
	case 8:
		return EnUsI00008
	case 9:
		return EnUsI00009
	case 10:
		return EnUsI00010
	case 11:
		return EnUsI00011
	case 13:
		return EnUsE00001
	case 14:
		return EnUsE00002
	case 15:
		return EnUsE00003
	case 16:
		return EnUsE00004
	case 17:
		return EnUsE00005
	case 18:
		return EnUsE00006
	case 19:
		return EnUsE00007
	case 20:
		return EnUsE00008
	case 21:
		return EnUsE00009
	case 22:
		return EnUsE00010
	case 23:
		return EnUsE00011
	case 24:
		return EnUsE00012
	case 25:
		return EnUsE00013
	case 26:
		return EnUsE00014
	case 27:
		return EnUsE00015
	case 28:
		return EnUsE00016
	case 29:
		return EnUsE00017
	case 30:
		return EnUsE00018
	case 31:
		return EnUsE00019
	case 32:
		return EnUsE00020
	case 33:
		return EnUsE00021
	case 34:
		return EnUsE00022
	case 35:
		return EnUsE00023
	case 36:
		return EnUsE00024
	case 37:
		return EnUsE00025
	case 38:
		return EnUsE00026
	case 39:
		return EnUsE00027
	case 40:
		return EnUsE00028
	case 41:
		return EnUsE00029
	case 42:
		return EnUsE00030
	case 43:
		return EnUsE00031
	case 44:
		return EnUsE00032
	case 45:
		return EnUsE00033
	case 46:
		return EnUsE00034
	case 47:
		return EnUsE00035
	case 48:
		return EnUsE00036
	case 49:
		return EnUsE00037
	case 50:
		return EnUsE00038
	case 51:
		return EnUsE00039
	case 52:
		return EnUsE00040
	case 53:
		return EnUsE00041
	case 54:
		return EnUsE00042
	case 55:
		return EnUsE00043
	case 56:
		return EnUsE00044
	case 57:
		return EnUsE00045
	case 58:
		return EnUsE00046
	case 59:
		return EnUsE00047
	case 60:
		return EnUsE00048
	case 61:
		return EnUsE00049
	case 62:
		return EnUsE00050
	case 63:
		return EnUsE00051
	case 64:
		return EnUsE00052
	case 65:
		return EnUsE00053
	case 66:
		return EnUsE00054
	case 67:
		return EnUsE00055
	case 68:
		return EnUsE00056
	case 69:
		return EnUsE00057
	case 70:
		return EnUsE00058
	case 71:
		return EnUsE00059
	case 72:
		return EnUsE00060
	case 73:
		return EnUsE00061
	case 74:
		return EnUsE00062
	case 75:
		return EnUsE00063
	case 76:
		return EnUsE00064
	case 77:
		return EnUsE00065
	case 78:
		return EnUsE00066
	case 79:
		return EnUsE00067
	case 80:
		return EnUsE00068
	case 81:
		return EnUsE00069
	case 82:
		return EnUsE00070
	case 83:
		return EnUsE00071
	case 84:
		return EnUsE00072
	case 85:
		return EnUsE00073
	case 86:
		return EnUsE00074
	case 87:
		return EnUsE00075
	case 88:
		return EnUsE00076
	case 89:
		return EnUsE00077
	case 90:
		return EnUsE00078
	case 91:
		return EnUsE00079
	case 92:
		return EnUsE00080
	case 93:
		return EnUsE00081
	case 94:
		return EnUsE00082
	case 95:
		return EnUsE00083
	case 96:
		return EnUsE00084
	case 97:
		return EnUsE00085
	case 98:
		return EnUsE00086
	case 99:
		return EnUsE00087
	case 100:
		return EnUsE00088
	case 101:
		return EnUsE00089
	case 102:
		return EnUsE00090
	case 103:
		return EnUsE00091
	case 104:
		return EnUsE00092
	case 105:
		return EnUsE00093
	case 106:
		return EnUsE00094
	case 107:
		return EnUsE00095
	case 108:
		return EnUsE00096
	case 109:
		return EnUsE00097
	case 110:
		return EnUsE00098
	case 111:
		return EnUsE00099
	case 112:
		return EnUsE00100
	case 113:
		return EnUsE00101
	case 114:
		return EnUsE00102
	case 115:
		return EnUsE00103
	case 116:
		return EnUsE00104
	case 117:
		return EnUsE00105
	case 118:
		return EnUsE00106
	case 119:
		return EnUsE00107
	case 120:
		return EnUsE00108
	case 121:
		return EnUsE00109
	case 122:
		return EnUsE00110
	case 123:
		return EnUsE00111
	case 124:
		return EnUsE00112
	case 125:
		return EnUsE00113
	case 126:
		return EnUsE00114
	case 127:
		return EnUsE00115
	case 128:
		return EnUsE00116
	case 129:
		return EnUsE00117
	case 130:
		return EnUsE00118
	case 131:
		return EnUsE00119
	case 132:
		return EnUsE00120
	case 133:
		return EnUsE00121
	case 134:
		return EnUsE00122
	case 135:
		return EnUsE00123
	case 136:
		return EnUsE00124
	case 137:
		return EnUsE00125
	case 138:
		return EnUsE00126
	case 139:
		return EnUsE00127
	case 140:
		return EnUsE00128
	case 141:
		return EnUsE00129
	case 142:
		return EnUsE00130
	case 143:
		return EnUsE00131
	case 144:
		return EnUsE00132
	case 145:
		return EnUsE00133
	case 146:
		return EnUsE00134
	case 147:
		return EnUsE00135
	case 148:
		return EnUsE00136
	case 149:
		return EnUsE00137
	case 150:
		return EnUsE00138
	case 151:
		return EnUsE00139
	case 152:
		return EnUsE00140
	case 153:
		return EnUsE00141
	case 154:
		return EnUsE00142
	case 155:
		return EnUsE00143
	case 156:
		return EnUsE00144
	case 157:
		return EnUsE00145
	case 158:
		return EnUsErr0
	case 159:
		return EnUsErr1
	case 160:
		return EnUsErr2
	case 161:
		return EnUsErr3
	case 162:
		return EnUsErr4
	case 163:
		return EnUsErr5
	case 164:
		return EnUsErr6
	case 165:
		return EnUsErr7
	case 166:
		return EnUsErr8
	case 167:
		return EnUsErr9
	case 168:
		return EnUsErr10
	case 169:
		return EnUsErr11
	case 170:
		return EnUsErr12
	case 171:
		return EnUsErr13
	case 172:
		return EnUsErr14
	case 173:
		return EnUsErr15
	case 174:
		return EnUsErr16
	case 175:
		return EnUsErr17
	case 176:
		return EnUsErr18
	case 177:
		return EnUsErr19
	case 178:
		return EnUsErr20
	case 179:
		return EnUsErr21
	case 180:
		return EnUsErr22
	case 181:
		return EnUsErr23
	case 182:
		return EnUsErr24
	case 183:
		return EnUsErr25
	case 184:
		return EnUsErr26
	case 185:
		return EnUsErr27
	case 186:
		return EnUsErr28
	case 187:
		return EnUsErr29
	case 188:
		return EnUsErr30
	case 189:
		return EnUsErr31
	case 190:
		return EnUsErr32
	case 191:
		return EnUsErr33
	case 192:
		return EnUsErr34
	case 193:
		return EnUsErr35
	case 194:
		return EnUsErr36
	case 195:
		return EnUsErr37
	case 196:
		return EnUsErr38
	case 197:
		return EnUsErr39
	case 198:
		return EnUsErr40
	case 199:
		return EnUsErr41
	case 200:
		return EnUsErr42
	case 201:
		return EnUsErr43
	case 202:
		return EnUsErr44
	case 203:
		return EnUsErr45
	case 204:
		return EnUsErr46
	case 205:
		return EnUsErr47
	case 206:
		return EnUsErr48
	case 207:
		return EnUsErr49
	case 208:
		return EnUsErr50
	case 209:
		return EnUsErr51
	case 210:
		return EnUsErr52
	case 211:
		return EnUsErr53
	case 212:
		return EnUsErr54
	case 213:
		return EnUsErr55
	case 214:
		return EnUsErr56
	case 215:
		return EnUsErr57
	case 216:
		return EnUsErr58
	case 217:
		return EnUsErr59
	case 218:
		return EnUsErr60
	case 219:
		return EnUsErr61
	case 220:
		return EnUsErr62
	case 221:
		return EnUsErr63
	case 222:
		return EnUsErr64
	case 223:
		return EnUsErr65
	case 224:
		return EnUsErr66
	case 225:
		return EnUsErr67
	case 226:
		return EnUsErr68
	case 227:
		return EnUsErr69
	case 228:
		return EnUsErr70
	case 229:
		return EnUsErr71
	case 230:
		return EnUsErr72
	case 231:
		return EnUsErr73
	case 232:
		return EnUsErr74
	case 233:
		return EnUsErr75
	case 234:
		return EnUsErr76
	case 235:
		return EnUsErr77
	case 236:
		return EnUsErr78
	case 237:
		return EnUsErr79
	case 238:
		return EnUsErr80
	case 239:
		return EnUsErr81
	case 240:
		return EnUsErr82
	case 241:
		return EnUsErr83
	case 242:
		return EnUsErr84
	case 243:
		return EnUsErr85
	case 244:
		return EnUsErr86
	case 245:
		return EnUsErr87
	case 246:
		return EnUsErr88
	case 247:
		return EnUsErr89
	case 248:
		return EnUsErr90
	case 249:
		return EnUsErr91
	case 250:
		return EnUsErr92
	case 251:
		return EnUsErr93
	case 252:
		return EnUsErr94
	case 253:
		return EnUsErr95
	case 254:
		return EnUsErr96
	case 255:
		return EnUsErr97
	case 256:
		return EnUsErr98
	case 257:
		return EnUsErr99
	case 258:
		return EnUsErr100
	case 259:
		return EnUsErr101
	case 260:
		return EnUsErr102
	case 261:
		return EnUsErr103
	case 262:
		return EnUsErr104
	case 263:
		return EnUsErr105
	case 264:
		return EnUsErr106
	case 265:
		return EnUsErr107
	case 266:
		return EnUsErr108
	case 267:
		return EnUsErr109
	case 268:
		return EnUsErr110
	case 269:
		return EnUsErr111
	case 270:
		return EnUsErr112
	case 271:
		return EnUsErr113
	case 272:
		return EnUsErr114
	case 273:
		return EnUsErr115
	case 274:
		return EnUsErr116
	case 275:
		return EnUsErr117
	case 276:
		return EnUsErr118
	case 277:
		return EnUsErr119
	case 278:
		return EnUsErr120
	case 279:
		return EnUsErr121
	case 280:
		return EnUsErr122
	case 281:
		return EnUsErr123
	case 282:
		return EnUsErr124
	case 283:
		return EnUsErr125
	case 284:
		return EnUsErr126
	case 285:
		return EnUsErr127
	case 286:
		return EnUsErr128
	case 287:
		return EnUsErr130
	case 288:
		return EnUsErr131
	case 289:
		return EnUsErr132
	case 290:
		return EnUsErr141
	case 291:
		return EnUsErr145
	case 292:
		return EnUsErr152
	case 293:
		return EnUsErr153
	case 294:
		return EnUsErr154
	case 295:
		return EnUsErr155
	case 296:
		return EnUsErr156
	case 297:
		return EnUsErr165
	case 298:
		return EnUsErr170
	case 299:
		return EnUsErr171
	case 300:
		return EnUsErr172
	case 301:
		return EnUsErr173
	case 302:
		return EnUsErr174
	case 303:
		return EnUsErr175
	case 304:
		return EnUsErr180
	case 305:
		return EnUsErr181
	case 306:
		return EnUsErr182
	case 307:
		return EnUsErr183
	case 308:
		return EnUsErr184
	case 309:
		return EnUsErr185
	case 310:
		return EnUsErr186
	case 311:
		return EnUsErr187
	case 312:
		return EnUsErr191
	case 313:
		return EnUsErr192
	case 314:
		return EnUsErr193
	case 315:
		return EnUsErr195
	case 316:
		return EnUsErr200
	case 317:
		return EnUsErr201
	case 318:
		return EnUsErr202
	case 319:
		return EnUsErr203
	case 320:
		return EnUsErr204
	case 321:
		return EnUsErr205
	case 322:
		return EnUsErr206
	case 323:
		return EnUsErr207
	case 324:
		return EnUsErr208
	case 325:
		return EnUsErr209
	case 326:
		return EnUsErr210
	case 327:
		return EnUsErr211
	case 328:
		return EnUsErr212
	case 329:
		return EnUsErr213
	case 330:
		return EnUsErr214
	case 331:
		return EnUsErr215
	case 332:
		return EnUsErr216
	case 333:
		return EnUsErr217
	case 334:
		return EnUsErr218
	case 335:
		return EnUsErr219
	case 336:
		return EnUsErr220
	case 337:
		return EnUsErr221
	case 338:
		return EnUsErr222
	case 339:
		return EnUsErr223
	case 340:
		return EnUsErr224
	case 341:
		return EnUsErr225
	case 342:
		return EnUsErr226
	case 343:
		return EnUsErr227
	case 344:
		return EnUsErr228
	case 345:
		return EnUsErr229
	case 346:
		return EnUsErr230
	case 347:
		return EnUsErr231
	case 348:
		return EnUsErr232
	case 349:
		return EnUsErr233
	case 350:
		return EnUsErr234
	case 351:
		return EnUsErr235
	case 352:
		return EnUsErr236
	case 353:
		return EnUsErr237
	case 354:
		return EnUsErr238
	case 355:
		return EnUsErr239
	case 356:
		return EnUsErr240
	case 357:
		return EnUsErr241
	case 358:
		return EnUsErr242
	case 359:
		return EnUsErr243
	case 360:
		return EnUsErr244
	case 361:
		return EnUsErr245
	case 362:
		return EnUsErr246
	case 363:
		return EnUsErr247
	case 364:
		return EnUsErr248
	case 365:
		return EnUsErr250
	case 366:
		return EnUsErr251
	case 367:
		return EnUsErr252
	case 368:
		return EnUsErr253
	case 369:
		return EnUsErr254
	case 370:
		return EnUsErr260
	case 371:
		return EnUsErr261
	case 372:
		return EnUsErr262
	case 373:
		return EnUsErr265
	case 374:
		return EnUsErr270
	case 375:
		return EnUsErr271
	case 376:
		return EnUsErr280
	case 377:
		return EnUsErr281
	case 378:
		return EnUsErr282
	case 379:
		return EnUsErr283
	case 380:
		return EnUsErr284
	case 381:
		return EnUsErr285
	case 382:
		return EnUsErr286
	case 383:
		return EnUsErr287
	case 384:
		return EnUsErr288
	case 385:
		return EnUsErr289
	case 386:
		return EnUsErr290
	case 387:
		return EnUsErr295
	case 388:
		return EnUsErr296
	case 389:
		return EnUsErr297
	case 390:
		return EnUsErr298
	case 391:
		return EnUsErr300
	case 392:
		return EnUsErr301
	case 393:
		return EnUsErr302
	case 394:
		return EnUsErr303
	case 395:
		return EnUsErr304
	case 396:
		return EnUsErr305
	case 397:
		return EnUsErr306
	case 398:
		return EnUsErr307
	case 399:
		return EnUsErr308
	case 400:
		return EnUsErr309
	case 401:
		return EnUsErr310
	case 402:
		return EnUsErr311
	case 403:
		return EnUsErr312
	case 404:
		return EnUsErr313
	case 405:
		return EnUsErr314
	case 406:
		return EnUsErr315
	case 407:
		return EnUsErr316
	case 408:
		return EnUsErr317
	case 409:
		return EnUsErr318
	case 410:
		return EnUsErr319
	case 411:
		return EnUsErr320
	case 412:
		return EnUsErr325
	case 413:
		return EnUsErr326
	case 414:
		return EnUsErr327
	case 415:
		return EnUsErr328
	case 416:
		return EnUsErr330
	case 417:
		return EnUsErr331
	case 418:
		return EnUsErr332
	case 419:
		return EnUsErr333
	case 420:
		return EnUsErr334
	case 421:
		return EnUsErr339
	case 422:
		return EnUsErr340
	case 423:
		return EnUsErr341
	case 424:
		return EnUsErr342
	case 425:
		return EnUsErr343
	case 426:
		return EnUsErr344
	case 427:
		return EnUsErr350
	case 428:
		return EnUsErr351
	case 429:
		return EnUsErr355
	case 430:
		return EnUsErr356
	case 431:
		return EnUsErr357
	case 432:
		return EnUsErr358
	case 433:
		return EnUsErr359
	case 434:
		return EnUsErr360
	case 435:
		return EnUsErr361
	case 436:
		return EnUsErr362
	case 437:
		return EnUsErr363
	case 438:
		return EnUsErr370
	case 439:
		return EnUsErr371
	case 440:
		return EnUsErr375
	case 441:
		return EnUsErr380
	case 442:
		return EnUsErr901
	case 443:
		return EnUsErr2000
	case 444:
		return EnUsErr2001
	case 445:
		return EnUsErr2002
	case 446:
		return EnUsErr2003
	case 447:
		return EnUsErr2004
	case 448:
		return EnUsErr2005
	case 449:
		return EnUsErr2006
	case 450:
		return EnUsErr2007
	case 451:
		return EnUsErr2008
	case 452:
		return EnUsErr2009
	case 453:
		return EnUsErr2100
	case 454:
		return EnUsErr2101
	case 455:
		return EnUsErr2102
	case 456:
		return EnUsErr2103
	case 457:
		return EnUsErr2104
	case 458:
		return EnUsErr2105
	case 459:
		return EnUsErr2106
	case 460:
		return EnUsErr2107
	case 461:
		return EnUsErr2108
	case 462:
		return EnUsErr2109
	case 463:
		return EnUsErr2200
	default:
		return EnUsUnknown
	}
}
