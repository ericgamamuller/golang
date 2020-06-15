package contracts

type CreateLinkRequest struct {
	Type            string
	Name            string
	Description     string
	ShowDescription bool
	Price           int
	ExpirationDate  string
	Weight          int
	MaxInstallments int
	Quantity        int
	SKU             string
	SoftDescriptor  string
	Shipping        Shipping
	Recurrent       Recurrent
}

type CreateLinkResponse struct {
	ID              string
	ShortURL        string
	Type            string
	Name            string
	Description     string
	ShowDescription bool
	Price           int
	ExpirationDate  string
	Weight          int
	MaxInstallments int
	Quantity        int
	SKU             string
	SoftDescriptor  string
	Shipping        Shipping
	Recurrent       Recurrent
}

type Shipping struct {
	Name          string
	Price         int
	OriginZipCode string
	Type          string
}

type Recurrent struct {
	RecurrentInterval string
	RecurrentEndDate  string
}
