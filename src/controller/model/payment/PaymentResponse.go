package model

type PaymentResponse struct {
	//*model.RestResponse
	Url string `json:"url"`
}

func NewPaymentResponse(url string) *PaymentResponse {
	return &PaymentResponse{
		//RestResponse: model.NewRestResponse(status, time.Now()),
		Url: url,
	}
}
