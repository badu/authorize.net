package authorize_net

// -- request --

type GetAUJobSummaryRequest struct {
	Payload GetAUJobSummaryPayload `json:"getAUJobSummaryRequest"`
}
type GetAUJobSummaryPayload struct {
	ANetApiRequest
	Month string `json:"month"`
}

type GetAUJobDetailsRequest struct {
	Payload GetAUJobDetailsPayload `json:"getAUJobDetailsRequest"`
}
type GetAUJobDetailsPayload struct {
	ANetApiRequest
	Month              string  `json:"month"`
	ModifiedTypeFilter string  `json:"modifiedTypeFilter,omitempty"`
	Paging             *Paging `json:"paging,omitempty"`
}

// -- response --

type AuResponse struct {
	AuReasonCode      string `json:"auReasonCode"`
	ProfileCount      uint   `json:"profileCount"`
	ReasonDescription string `json:"reasonDescription"`
}
type GetAUJobSummaryResponse struct {
	ANetApiResponse
	AuSummary []AuResponse `json:"auSummary,omitempty"`
}

type ListOfAUDetails struct {
	AuUpdateOrAuDelete []string `json:"auUpdateOrAuDelete,omitempty"` //min=0
}

type GetAUJobDetailsResponse struct {
	ANetApiResponse
	TotalNumInResultSet int              `json:"totalNumInResultSet,omitempty"`
	AuDetails           *ListOfAUDetails `json:"auDetails,omitempty"`
}
