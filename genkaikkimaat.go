package main

import (
	"astuart.co/goq"
	"encoding/csv"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
	"os"
)

const URL = "http://www.stat.fi/meta/luokitukset/valtio/001-2012/index.html"

type Table struct {
	Rows []TableRow `goquery:"table table tr"`
}

type TableRow struct {
	Cols []string `goquery:"td"`
}

func checkErr(str string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, str+":", err)
		os.Exit(1)
	}
}

func main() {
	res, err := http.Get(URL)
	checkErr("error getting "+URL, err)
	defer res.Body.Close()

	rdr, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	checkErr("error determining character set", err)

	var t Table

	err = goq.NewDecoder(rdr).Decode(&t)
	checkErr("error decoding", err)

	wr := csv.NewWriter(os.Stdout)

	for _, r := range t.Rows {
		err = wr.Write(r.Cols)
		checkErr("error writing", err)
	}
	wr.Flush()
	checkErr("error writing", wr.Error())
}
