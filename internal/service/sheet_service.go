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

	values := s.rowsToValues(rows)

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

// ReplaceSheet clears an entire sheet and writes header+rows.
func (s *SheetService) ReplaceSheet(ctx context.Context, sheetName string, rows []dto.Row) error {
	// Step 1: clear the sheet
	clearReq := &sheets.ClearValuesRequest{}
	_, err := s.svc.Spreadsheets.Values.Clear(s.spreadsheetID, sheetName, clearReq).Context(ctx).Do()
	if err != nil {
		return eris.Wrap(err, "error clearing sheet")
	}

	// Step 2: prepare the data with header first
	all := s.rowsToValues(rows)

	// Step 3: write back to sheet starting at A1
	vRange := sheetName + "!A1"
	valueRange := &sheets.ValueRange{
		Values: all,
	}

	_, err = s.svc.Spreadsheets.Values.Update(s.spreadsheetID, vRange, valueRange).
		ValueInputOption("RAW"). // or "USER_ENTERED"
		Context(ctx).
		Do()
	if err != nil {
		return eris.Wrap(err, "error updating sheet")
	}

	return nil
}

func (s *SheetService) rowsToValues(rows []dto.Row) [][]any {
	var values [][]any
	values = append(values, rows[0].ToHeader())
	for _, r := range rows {
		values = append(values, []any{
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

	return values
}
