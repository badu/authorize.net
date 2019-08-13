package authorize_net

// -- request --

type MobileDeviceLoginRequest struct {
	ANetApiRequest `json:"mobileDeviceLoginRequest"`
}

type MobileDevice struct {
	MobileDeviceId   string `json:"mobileDeviceId"`
	Description      string `json:"description,omitempty"`
	PhoneNumber      string `json:"phoneNumber,omitempty"`
	DevicePlatform   string `json:"devicePlatform,omitempty"`
	DeviceActivation string `json:"deviceActivation,omitempty"`
}

type MobileDeviceRegistrationRequest struct {
	Payload MobileDeviceRegistrationPayload `json:"mobileDeviceRegistrationRequest"`
}
type MobileDeviceRegistrationPayload struct {
	ANetApiRequest
	MobileDevice *MobileDevice `json:"mobileDevice"`
}

// -- response --

type MerchantContact struct {
	MerchantName    string `json:"merchantName,omitempty"`
	MerchantAddress string `json:"merchantAddress,omitempty"`
	MerchantCity    string `json:"merchantCity,omitempty"`
	MerchantState   string `json:"merchantState,omitempty"`
	MerchantZip     string `json:"merchantZip,omitempty"`
	MerchantPhone   string `json:"merchantPhone,omitempty"`
}

type Permission struct {
	PermissionName string `json:"permissionName,omitempty"`
}

type MobileDeviceLoginResponse struct {
	ANetApiResponse
	MerchantContact *MerchantContact `json:"merchantContact"`
	UserPermissions []Permission     `json:"userPermissions"`
	MerchantAccount *TransRetailInfo `json:"merchantAccount,omitempty"`
}

type MobileDeviceRegistrationResponse struct {
	ANetApiResponse
}
