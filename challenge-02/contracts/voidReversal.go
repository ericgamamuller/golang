package contracts

type VoidReversalRequest struct {
}

type VoidReversalResponse struct {
	CancellationStatus int
	ConfirmationStatus int
	Status             int
	ReasonCode         int
	ReturnCode         string
	ReturnMessage      string
}
