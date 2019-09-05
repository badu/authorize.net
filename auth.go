package authorize_net

type AuthenticateTestRequest struct {
	ANetApiRequest `json:"authenticateTestRequest"`
}

type LogoutResponse struct {
	ANetApiResponse
}
