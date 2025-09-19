package main

import (
	"time"

	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/xml-to-gsheet/internal/config"
	"github.com/itsLeonB/xml-to-gsheet/internal/dto"
	"github.com/itsLeonB/xml-to-gsheet/internal/mapper"
	"github.com/itsLeonB/xml-to-gsheet/internal/service"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rotisserie/eris"
)

func main() {
	startTime := time.Now()
	logger := ezutil.NewSimpleLogger("XML-to-GSheet", true, 1)
	logger.Info("starting cronjob...")

	cfg := config.Load()
	url := cfg.Url

	logger.Infof("using spreadsheet ID: %s", cfg.SpreadsheetId)
	logger.Infof("using sheet name: %s", cfg.SheetName)

	scraper := service.NewScraperService[dto.Feed]()

	logger.Info("scraping url...")

	feed, err := scraper.ScrapeXML(url)
	if err != nil {
		logger.Fatalf(eris.ToString(err, true))
	}

	logger.Infof("obtained xml entries: %d", len(feed.Entries))

	rows := ezutil.MapSlice(feed.Entries, mapper.EntryToRow)

	logger.Info("start appending rows to sheet...")

	sheetService := service.NewSheetService(cfg)

	if err = sheetService.AppendRows(cfg.SheetName, rows); err != nil {
		logger.Fatalf(eris.ToString(err, true))
	}

	logger.Infof("finished job, elapsed time: %d miliseconds", time.Since(startTime).Milliseconds())
}
