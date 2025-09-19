package dto

import "reflect"

type Row struct {
	ID                    string
	Title                 string
	Description           string
	ProductType           string
	ItemGroupID           string
	GoogleProductCategory string
	Link                  string
	ImageLink             string
	Condition             string
	Availability          string
	Price                 string
	Brand                 string
	MPN                   string
	GTIN                  string
	Color                 string
	Size                  string
	ShippingCountry       string
	ShippingService       string
	ShippingPrice         string
	AdditionalImageLink   string
}

// Header returns the struct field names (column names).
func (r Row) ToHeader() []any {
	t := reflect.TypeOf(r)
	headers := make([]any, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		headers[i] = t.Field(i).Name
	}
	return headers
}
