package main

import (
	"context"
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

	scraper := service.NewScraperService[dto.Feed](logger)

	logger.Info("scraping url...")

	feed, err := scraper.ScrapeXML(url)
	if err != nil {
		logger.Fatalf(eris.ToString(err, true))
	}

	logger.Infof("obtained xml entries: %d", len(feed.Entries))

	logger.Infof("using placeholder description: %s", cfg.PlaceholderDescription)
	placeholders := dto.Entry{
		Description: cfg.PlaceholderDescription,
	}

	rows := ezutil.MapSlice(feed.Entries, mapper.GetSimpleMapper(placeholders))

	logger.Info("start updating rows to sheet...")

	sheetService := service.NewSheetService(cfg)

	if err = sheetService.ReplaceSheet(context.Background(), cfg.SheetName, rows); err != nil {
		logger.Fatalf(eris.ToString(err, true))
	}

	logger.Infof("finished job, elapsed time: %d miliseconds", time.Since(startTime).Milliseconds())
}
