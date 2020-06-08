package models

import "git/challenge-03/db"

type PaymentLink struct {
	ID 					  string
	ShortURL 			  string
	Type 				  string
	Name 				  string
	Description 		  string
	ShowDescription 	  bool
	Price				  int
	ExpirationDate	      string
	Weight				  int
	MaxInstallments  	  int
	Quantity			  int
	SKU					  string
	SoftDescriptor		  string
	ShippingName		  string
	ShippingPrice		  int
	ShippingOriginZipCode string
	ShippingType		  string
	RecurrentInterval	  string
	RecurrentEndDate	  string
}

func (link *PaymentLink) InsertLink() {
	db := db.Connect()
	defer db.Close()

	dbLinkInsert, err := db.Prepare("insert into payment_links(id, shortUrl, type, name, description, showDescription, price, expirationDate, weight, maxInstallments, quantity, sku, softDescriptor, shippingName, shippingPrice, shippingOriginZipCode, shippingType, recurrentInterval, recurrentEndDate) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)")
	if err != nil {
		panic(err.Error())
	}

	dbLinkInsert.Exec(link.ID, link.ShortURL, link.Type, link.Name, link.Description, link.ShowDescription, link.Price, link.ExpirationDate, link.Weight, link.MaxInstallments, link.Quantity, link.SKU, link.SoftDescriptor, link.ShippingName, link.ShippingPrice, link.ShippingOriginZipCode, link.ShippingType, link.RecurrentInterval, link.RecurrentEndDate)
}

func GetLinks() []PaymentLink {
	db := db.Connect()
	defer db.Close()

	dbLinks, err := db.Query("select * from payment_links")
	if err != nil {
		panic(err.Error())
	}

	link := PaymentLink{}
	links := []PaymentLink{}

	for dbLinks.Next() {
		err = dbLinks.Scan(&link.ID, &link.ShortURL, &link.Type, &link.Name, &link.Description, &link.ShowDescription, &link.Price, &link.ExpirationDate, &link.Weight, &link.MaxInstallments, &link.Quantity, &link.SKU, &link.SoftDescriptor, &link.ShippingName, &link.ShippingPrice, &link.ShippingOriginZipCode, &link.ShippingType, &link.RecurrentInterval, &link.RecurrentEndDate)
		if err != nil {
			panic(err.Error())
		}

		links := append(links, link)
	}

	return links
}

func GetLink(id string) PaymentLink {
	db := db.Connect()
	defer db.Close()

	dbLinks, err := db.Query("select * from payment_links where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	link := PaymentLink{}

	for dbLinks.Next() {
		err = dbLinks.Scan(&link.ID, &link.ShortURL, &link.Type, &link.Name, &link.Description, &link.ShowDescription, &link.Price, &link.ExpirationDate, &link.Weight, &link.MaxInstallments, &link.Quantity, &link.SKU, &link.SoftDescriptor, &link.ShippingName, &link.ShippingPrice, &link.ShippingOriginZipCode, &link.ShippingType, &link.RecurrentInterval, &link.RecurrentEndDate)
		if err != nil {
			panic(err.Error())
		}
	}

	return link
}

func (link *PaymentLink) UpdateLink() {
	db := db.Connect()
	defer db.Close()

	dbLinkUpdate, err := db.Prepare("update payment_links set shortUrl=$2, type=$3, name=$4, description=$5, showDescription=$6, price=$7, expirationDate=$8, weight=$9, maxInstallments=$10, quantity=$11, sku=$12, softDescriptor=$13, shippingName=$14, shippingPrice=$15, shippingOriginZipCode=$16, shippingType=$17, recurrentInterval=$18, recurrentEndDate=$19 where id=$1")
	if err != nil {
		panic(err.Error())
	}

	dbLinkUpdate.Exec(link.ID, link.ShortURL, link.Type, link.Name, link.Description, link.ShowDescription, link.Price, link.ExpirationDate, link.Weight, link.MaxInstallments, link.Quantity, link.SKU, link.SoftDescriptor, link.ShippingName, link.ShippingPrice, link.ShippingOriginZipCode, link.ShippingType, link.RecurrentInterval, link.RecurrentEndDate)

	return link
}

func (link *PaymentLink) DeleteLink() {
	db := db.Connect()
	defer db.Close()

	dbLinkUpdate, err := db.Prepare("delete from payment_links where id=$1")
	if err != nil {
		panic(err.Error())
	}

	dbLinkUpdate.Exec(link.ID)
}