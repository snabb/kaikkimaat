package main

import (
	"astuart.co/goq"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const URL = "http://publications.europa.eu/code/fi/fi-5000500.htm"

// Structured representation for github file name tableRows
type Table struct {
	Headings []string   `goquery:"table#listOfCountriesTable tr:nth-child(1) th"`
	Rows     []TableRow `goquery:"table#listOfCountriesTable tr:nth-child(n+2)"`
}

type TableRow struct {
	Cols []string `goquery:"td"`
}

func checkErr(str string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, str, err)
		os.Exit(1)
	}
}

func main() {
	// haetaan lähdedokumentti
	res, err := http.Get(URL)
	checkErr("error getting "+URL+":", err)
	defer res.Body.Close()

	var t Table

	err = goq.NewDecoder(res.Body).Decode(&t)
	checkErr("error decoding:", err)

	// taulukon ensimmäinen sarake on tyhjä jostain syystä, poista
	if len(t.Headings) > 0 && t.Headings[0] == "" {
		t.Headings = t.Headings[1:]
		for i := range t.Rows {
			if len(t.Rows[i].Cols) > 0 {
				t.Rows[i].Cols = t.Rows[i].Cols[1:]
			}
		}
	}
	// skipataan rivit joissa ei ole kaikkia tietoja
	var newRows []TableRow
	for _, r := range t.Rows {
		if len(r.Cols) != len(t.Headings) {
			continue
		}
		newRows = append(newRows, r)
	}
	t.Rows = newRows

	// poistetaan soft hyphenit, nbsp:t ja rivinvaihdot
	repl := strings.NewReplacer("\u00ad", "", "\u00a0", "", "-\n", "", "\n", " ")
	// sekä muu roska
	fixRe := regexp.MustCompile(`\s*\(.*\)$|^—$`)
	for i, str := range t.Headings {
		str = repl.Replace(str)
		t.Headings[i] = fixRe.ReplaceAllLiteralString(str, "")
	}
	for i, r := range t.Rows {
		for j, c := range r.Cols {
			c = repl.Replace(c)
			t.Rows[i].Cols[j] = fixRe.ReplaceAllLiteralString(c, "")
		}
	}

	// ulostetaan csv
	wr := csv.NewWriter(os.Stdout)

	err = wr.Write(t.Headings)
	checkErr("error writing:", err)

	for _, r := range t.Rows {
		err = wr.Write(r.Cols)
		checkErr("error writing:", err)
	}
	wr.Flush()
	checkErr("error writing:", wr.Error())
}
