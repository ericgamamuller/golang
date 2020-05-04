package contracts

type PaymentReversalRequest struct {
}

type PaymentReversalResponse struct {
	Status             int
	ConfirmationStatus int
	ReasonCode         int
	ReturnCode         string
	ReturnMessage      string
}
