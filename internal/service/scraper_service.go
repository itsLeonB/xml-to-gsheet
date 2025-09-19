package service

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/itsLeonB/ezutil/v2"
	"github.com/rotisserie/eris"
)

type ScraperService[T any] struct {
	logger ezutil.Logger
}

func NewScraperService[T any](logger ezutil.Logger) *ScraperService[T] {
	return &ScraperService[T]{logger}
}

func (ss *ScraperService[T]) ScrapeXML(url string) (T, error) {
	var zero T

	resp, err := http.Get(url)
	if err != nil {
		return zero, eris.Wrap(err, "error get request")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			ss.logger.Error(eris.ToString(eris.Wrap(err, "error closing body"), true))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return zero, eris.Errorf("response not OK: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, eris.Wrap(err, "error reading body")
	}

	var reqType T
	if err := xml.Unmarshal(data, &reqType); err != nil {
		return zero, eris.Wrap(err, "error unmarshaling xml body")
	}

	return reqType, nil
}
