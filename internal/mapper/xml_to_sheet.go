package mapper

import "github.com/itsLeonB/xml-to-gsheet/internal/dto"

func entryToRowWithPlaceholder(e dto.Entry, placeholders dto.Entry) dto.Row {
	ship := dto.Shipping{}
	if len(e.Shipping) > 0 {
		ship = e.Shipping[0]
	}

	img := ""
	if len(e.AdditionalImageLinks) > 0 {
		img = e.AdditionalImageLinks[0]
	}

	description := e.Description
	if description == "" {
		description = placeholders.Description
	}

	return dto.Row{
		ID:                    e.ID,
		Title:                 e.Title,
		Description:           description,
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

func GetSimpleMapper(placeholders dto.Entry) func(dto.Entry) dto.Row {
	return func(e dto.Entry) dto.Row {
		return entryToRowWithPlaceholder(e, placeholders)
	}
}
