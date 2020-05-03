package contracts

type PaymentRequest struct {
	MerchantOrderId string
	Payment         Payment
}

type PaymentResponse struct {
	MerchantOrderId string
	Payment         Payment
}

type Payment struct {
	Type               string
	SoftDescriptor     string
	PaymentDateTime    string
	Amount             int
	Installmentes      int
	Interest           string
	Capture            bool
	ProductId          int
	CreditCard         CreditCard
	PinPadInformation  PinPad
	PaymentId          string
	Status             int
	ConfirmationStatus int
	ReturnCode         string
	ReturnMessage      string
}

type CreditCard struct {
	CardNumber                     string
	ExpirationDate                 string
	SecurityCodeStatus             string
	SecurityCode                   string
	BrandId                        int
	IssuerId                       int
	InputMode                      string
	AuthenticationMethod           string
	TruncateCardNumberWhenPrinting bool
}

type PinPad struct {
	PhysicalCharacteristics string
	ReturnDataInfo          string
	SerialNumber            string
	TerminalId              string
}
