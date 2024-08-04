package main

import (
	"flag"
	"fmt"
	"github.com/RubenOlsen/pdf2txt/pdf"
	"os"
	"regexp"
	"strings"
)

/*
The original code was written by Shakeel Mahate and is licensed under the MIT License.

What I kept:
- PDF logic on when a new line is detected.

My modifications:
- Added the actual parsing of the Trumf Visa PDF file.
- Removed fatalf function as it's not needed for a short program like this.
- Where the CSV file is written.
- Usage function

*/

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of TrumfVisa2ActualBudget:\n\n")
	fmt.Fprintf(os.Stderr, "TrumfVisa2ActualBudget [flags] pdf-file ...\n\n")
	os.Exit(1)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}

	pattern := `^(\d{2}\.\d{2}\.\d{2})(\d{2}\.\d{2}\.\d{2})(.*?)([A-Z]{3})(-?\d?\.?\d+,\d{2})(-?\d?\.?\d+,\d{2})`
	rex := regexp.MustCompile(pattern)

	counter := 0

	for _, in := range flag.Args() {
		reader, err := pdf.Open(in)
		if err != nil {
			fmt.Printf("Could not open PDF file: %v\n", err)
			os.Exit(1)
		}

		out := strings.Replace(in, ".pdf", ".csv", 1)
		writer, err := os.Create(out)
		if err != nil {
			fmt.Printf("Could not write to CSV file %s: %v\n", out, err)
			os.Exit(1)
		}

		var b strings.Builder
		var ln string
		for i := 1; i <= reader.NumPage(); i++ {
			// Initialize y co-ordinate for the page
			y := 0.0
			for _, t := range reader.Page(i).Content().Text {
				// Check if we are on a new line
				if t.Y != y {
					y = t.Y
					matches := rex.FindStringSubmatch(ln)
					if len(matches) > 0 {
						date := matches[2]
						payee := matches[3]
						notes := ""
						outflow := strings.Replace(matches[6], ".", "", 1)
						inflow := "0"

						if !strings.Contains(outflow, "-") {
							inflow = outflow
							outflow = "0"
						}

						if matches[4] == "NOK" {
							notes = fmt.Sprintf("%s / %v", payee, matches[1])
						} else {
							notes = fmt.Sprintf("%s %s %s / %v", payee, matches[4], matches[5], matches[1])
						}

						b.WriteString(fmt.Sprintf("%s;%s;%s;;%s;%s\n", date, payee, notes, outflow, inflow))
						counter++
					}
					ln = ""
				}
				ln = ln + t.S
			}
		}

		fmt.Fprintf(writer, "Date;Payee;Notes;Category;Outflow;Inflow\n")
		fmt.Fprintf(writer, "%v\n", b.String())
		writer.Close()
		fmt.Fprintf(os.Stdout, "Wrote %v transactions to %s\n", counter, out)
	}
}
