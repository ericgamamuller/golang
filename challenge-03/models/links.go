package models

import "git/challenge-03/db"

type PaymentLink struct {
	ID                    string
	ShortURL              string
	Type                  string
	Name                  string
	Description           string
	ShowDescription       bool
	Price                 int
	ExpirationDate        string
	Weight                int
	MaxInstallments       int
	Quantity              int
	SKU                   string
	SoftDescriptor        string
	ShippingName          string
	ShippingPrice         int
	ShippingOriginZipCode string
	ShippingType          string
	RecurrentInterval     string
	RecurrentEndDate      string
}

func (link *PaymentLink) InsertLink() {
	db := db.Connect()
	defer db.Close()

	dbLinkInsert, err := db.Prepare("insert into payment_links(id, short_url, type, name, price) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	dbLinkInsert.Exec(link.ID, link.ShortURL, link.Type, link.Name, link.Price)
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
		err = dbLinks.Scan(&link.ID, &link.ShortURL, &link.Type, &link.Name, &link.Price)
		if err != nil {
			panic(err.Error())
		}

		links = append(links, link)
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
		err = dbLinks.Scan(&link.ID, &link.ShortURL, &link.Type, &link.Name, &link.Price)
		if err != nil {
			panic(err.Error())
		}
	}

	return link
}

func (link *PaymentLink) UpdateLink() {
	db := db.Connect()
	defer db.Close()

	dbLinkUpdate, err := db.Prepare("update payment_links set short_url=$2, type=$3, name=$4, price=$5 where id=$1")
	if err != nil {
		panic(err.Error())
	}

	dbLinkUpdate.Exec(link.ID, link.ShortURL, link.Type, link.Name, link.Price)
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
