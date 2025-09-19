package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/itsLeonB/xml-to-gsheet/internal/config"
	"github.com/itsLeonB/xml-to-gsheet/internal/dto"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.Load()
	url := cfg.Url

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("bad status: %s", resp.Status))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var feed dto.Feed
	if err := xml.Unmarshal(data, &feed); err != nil {
		panic(err)
	}

	fmt.Println("Feed:", feed.Title, "updated:", feed.Updated)
	for _, e := range feed.Entries {
		fmt.Printf("\nTitle: %s\nPrice: %s\nSale Price: %s\nLink: %s\n", e.Title, e.Price, e.SalePrice, e.Link)
		fmt.Println("Shipping options:")
		for _, s := range e.Shipping {
			fmt.Printf(" - %s: %s (%s)\n", s.Service, s.Price, s.Country)
		}
		fmt.Println("Additional images:", e.AdditionalImageLinks)
	}
}
