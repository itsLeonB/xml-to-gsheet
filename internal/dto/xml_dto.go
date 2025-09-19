package dto

import "encoding/xml"

type Feed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string   `xml:"title"`
	Link    Link     `xml:"link"`
	Updated string   `xml:"updated"`
	Entries []Entry  `xml:"entry"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Entry struct {
	ID                    string     `xml:"http://base.google.com/ns/1.0 id"`
	Title                 string     `xml:"http://base.google.com/ns/1.0 title"`
	Description           string     `xml:"http://base.google.com/ns/1.0 description"`
	ProductType           string     `xml:"http://base.google.com/ns/1.0 product_type"`
	ItemGroupID           string     `xml:"http://base.google.com/ns/1.0 item_group_id"`
	GoogleProductCategory string     `xml:"http://base.google.com/ns/1.0 google_product_category"`
	Link                  string     `xml:"http://base.google.com/ns/1.0 link"`
	ImageLink             string     `xml:"http://base.google.com/ns/1.0 image_link"`
	Condition             string     `xml:"http://base.google.com/ns/1.0 condition"`
	Availability          string     `xml:"http://base.google.com/ns/1.0 availability"`
	AvailabilityDate      string     `xml:"http://base.google.com/ns/1.0 availability_date"`
	Price                 string     `xml:"http://base.google.com/ns/1.0 price"`
	SalePrice             string     `xml:"http://base.google.com/ns/1.0 sale_price"`
	Brand                 string     `xml:"http://base.google.com/ns/1.0 brand"`
	MPN                   string     `xml:"http://base.google.com/ns/1.0 mpn"`
	GTIN                  string     `xml:"http://base.google.com/ns/1.0 gtin"`
	Color                 string     `xml:"http://base.google.com/ns/1.0 color"`
	HTMLCodeColor         string     `xml:"http://base.google.com/ns/1.0 html_code_color"`
	Size                  string     `xml:"http://base.google.com/ns/1.0 size"`
	Shipping              []Shipping `xml:"http://base.google.com/ns/1.0 shipping"`
	AdditionalImageLinks  []string   `xml:"http://base.google.com/ns/1.0 additional_image_link"`
}

type Shipping struct {
	Country string `xml:"http://base.google.com/ns/1.0 country"`
	Service string `xml:"http://base.google.com/ns/1.0 service"`
	Price   string `xml:"http://base.google.com/ns/1.0 price"`
}
