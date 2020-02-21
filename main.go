package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"spreadsheet/sheet"
	"text/tabwriter"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"

)

func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := sheet.GetClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetID := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data!A2:F"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		// initialize tabwriter
		w := new(tabwriter.Writer)

		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 10, 8, 2, '\t', 0)

		defer w.Flush()

		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "Name", "Gender", "Class Level", "Home State", "Major")
		for _, row := range resp.Values {
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", row[0], row[1], row[2], row[3], row[4])
		}
	}

}
