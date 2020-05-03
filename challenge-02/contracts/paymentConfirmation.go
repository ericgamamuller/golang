package contracts

type PaymentConfirmationRequest struct {
}

type PaymentConfirmationResponse struct {
	Status             int
	ConfirmationStatus int
	ReturnCode         string
	ReturnMessage      string
}
