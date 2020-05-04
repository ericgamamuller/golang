package contracts

type PaymentConfirmationRequest struct {
}

type PaymentConfirmationResponse struct {
	Status             int
	ConfirmationStatus int
	ReasonCode         int
	ReturnCode         string
	ReturnMessage      string
}
