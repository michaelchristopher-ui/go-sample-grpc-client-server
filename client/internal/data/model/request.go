package model

type RequestAPIReq struct {
	RequestString string `json:"request_string"`
}

type RequestAPIResp struct {
	ResponseString string `json:"response_string,omitempty"`
	Error          string `json:"error,omitempty"`
}
