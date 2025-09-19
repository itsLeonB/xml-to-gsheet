package service

import (
	"context"

	"github.com/itsLeonB/xml-to-gsheet/internal/config"
	"github.com/itsLeonB/xml-to-gsheet/internal/dto"
	"github.com/rotisserie/eris"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetService struct {
	svc           *sheets.Service
	spreadsheetID string
}

func NewSheetService(cfg config.Config) *SheetService {
	svc, err := sheets.NewService(context.Background(), option.WithCredentialsJSON([]byte(cfg.GoogleServiceAccount)))
	if err != nil {
		panic(err)
	}

	return &SheetService{svc, cfg.SpreadsheetId}
}

func (s *SheetService) AppendRows(sheetName string, rows []dto.Row) error {
	if len(rows) < 1 {
		return eris.New("no rows to append")
	}

	var values [][]interface{}
	values = append(values, rows[0].ToHeader())
	for _, r := range rows {
		values = append(values, []interface{}{
			r.ID,
			r.Title,
			r.Description,
			r.ProductType,
			r.ItemGroupID,
			r.GoogleProductCategory,
			r.Link,
			r.ImageLink,
			r.Condition,
			r.Availability,
			r.Price,
			r.Brand,
			r.MPN,
			r.GTIN,
			r.Color,
			r.Size,
			r.ShippingCountry,
			r.ShippingService,
			r.ShippingPrice,
			r.AdditionalImageLink,
		})
	}

	valueRange := &sheets.ValueRange{Values: values}

	_, err := s.svc.Spreadsheets.Values.
		Append(s.spreadsheetID, sheetName, valueRange).
		ValueInputOption("RAW").
		Do()
	if err != nil {
		return eris.Wrap(err, "error appending to spreadsheet")
	}

	return nil
}
