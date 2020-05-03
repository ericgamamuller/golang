package contracts

type PaymentReversalRequest struct {
}

type PaymentReversalResponse struct {
	Status             int
	ConfirmationStatus int
	ReturnCode         string
	ReturnMessage      string
}
