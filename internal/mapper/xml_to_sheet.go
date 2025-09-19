package mapper

import "github.com/itsLeonB/xml-to-gsheet/internal/dto"

func EntryToRow(e dto.Entry) dto.Row {
	ship := dto.Shipping{}
	if len(e.Shipping) > 0 {
		ship = e.Shipping[0]
	}

	img := ""
	if len(e.AdditionalImageLinks) > 0 {
		img = e.AdditionalImageLinks[0]
	}

	return dto.Row{
		ID:                    e.ID,
		Title:                 e.Title,
		Description:           e.Description,
		ProductType:           e.ProductType,
		ItemGroupID:           e.ItemGroupID,
		GoogleProductCategory: e.GoogleProductCategory,
		Link:                  e.Link,
		ImageLink:             e.ImageLink,
		Condition:             e.Condition,
		Availability:          e.Availability,
		Price:                 e.Price,
		Brand:                 e.Brand,
		MPN:                   e.MPN,
		GTIN:                  e.GTIN,
		Color:                 e.Color,
		Size:                  e.Size,
		ShippingCountry:       ship.Country,
		ShippingService:       ship.Service,
		ShippingPrice:         ship.Price,
		AdditionalImageLink:   img,
	}
}
