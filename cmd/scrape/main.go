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

	resp, err := http.Get(url) // #nosec G107
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("error closing body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status:", resp.Status)
		return
	}

	if err := classifyTags(resp.Body); err != nil {
		fmt.Println("error:", err)
	}
}

func classifyTags(r io.Reader) error {
	dec := xml.NewDecoder(r)

	inEntry := false
	entryCounts := make(map[string]int)
	// track maximum count we ever saw for each tag in any single entry
	maxPerEntry := make(map[string]int)

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
			if se.Name.Space != "" {
				name = se.Name.Space + ":" + se.Name.Local
			}

			if se.Name.Local == "entry" && se.Name.Space == "http://www.w3.org/2005/Atom" {
				// new entry
				inEntry = true
				entryCounts = make(map[string]int)
				continue
			}

			if inEntry {
				entryCounts[name]++
				if entryCounts[name] > maxPerEntry[name] {
					maxPerEntry[name] = entryCounts[name]
				}
			}

		case xml.EndElement:
			if se.Name.Local == "entry" && se.Name.Space == "http://www.w3.org/2005/Atom" {
				inEntry = false
			}
		}
	}

	// classify tags
	var uniqueTags []string
	var singleValue []string
	var multiValue []string
	for tag, max := range maxPerEntry {
		uniqueTags = append(uniqueTags, tag)
		if max > 1 {
			multiValue = append(multiValue, tag)
		} else {
			singleValue = append(singleValue, tag)
		}
	}

	fmt.Println("=== All unique tags ===")
	for _, t := range uniqueTags {
		fmt.Println(" -", t)
	}

	fmt.Println("\n=== Tags with single value per entry (scalar fields) ===")
	for _, t := range singleValue {
		fmt.Println(" -", t)
	}

	fmt.Println("\n=== Tags with multiple values per entry (should be slices) ===")
	for _, t := range multiValue {
		fmt.Println(" -", t)
	}

	return nil
}
