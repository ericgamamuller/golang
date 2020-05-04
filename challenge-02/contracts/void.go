package contracts

type VoidRequest struct {
	MerchantVoidId   string
	MerchantVoidDate string
	Card             VoidCard
}

type VoidResponse struct {
	VoidId             string
	CancellationStatus int
	ConfirmationStatus int
	Status             int
	ReasonCode         int
	ReturnCode         string
	ReturnMessage      string
}

type VoidCard struct {
	InputMode  string
	CardNumber string
}
