package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/itsLeonB/xml-to-gsheet/internal/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.Load()
	url := cfg.Url

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status:", resp.Status)
		return
	}

	if err := scrapeUniqueTags(resp.Body); err != nil {
		fmt.Println("error:", err)
	}
}

func scrapeUniqueTags(r io.Reader) error {
	dec := xml.NewDecoder(r)

	seen := make(map[string]struct{})

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch se := tok.(type) {
		case xml.StartElement:
			name := se.Name.Local
			// include namespace if present
			if se.Name.Space != "" {
				name = se.Name.Space + ":" + se.Name.Local
			}
			if _, ok := seen[name]; !ok {
				seen[name] = struct{}{}
			}
		}
	}

	// print the unique tag names
	fmt.Println("Unique tags found:")
	for n := range seen {
		fmt.Println(" -", n)
	}
	return nil
}
