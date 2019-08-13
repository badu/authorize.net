package authorize_net

type Error struct {
	ErrorCode ErrCodeEnum `json:"errorCode,omitempty"`
	ErrorText string      `json:"errorText,omitempty"`
}

type ErrMessage struct {
	Code        ErrCodeEnum `json:"code"`
	Text        string      `json:"text"`
	Description string      `json:"description,omitempty"`
}

type Messages struct {
	ResultCode string       `json:"resultCode"` // "Ok" or "Error"
	Messages   []ErrMessage `json:"message"`    // min=0
}

type ANetApiResponse struct {
	RefId        string   `json:"refId,omitempty"` // Merchant-assigned reference ID for the request. If included in the request, this value will be included in the response. This feature might be especially useful for multi-threaded applications.
	Messages     Messages `json:"messages"`
	SessionToken string   `json:"sessionToken,omitempty"`
}
