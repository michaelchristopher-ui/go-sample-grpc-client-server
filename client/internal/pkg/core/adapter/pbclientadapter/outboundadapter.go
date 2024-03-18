package pbclientadapter

type OutboundAdapter interface {
	Request(RequestRequest) (RequestResponse, error)
}

type RequestRequest struct {
	RequestString string
}

type RequestResponse struct {
	ResponseString string
}
