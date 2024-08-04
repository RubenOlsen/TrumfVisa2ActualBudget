# Norsk Trumf Visa faktura til Actual Budget CSV konverterer

## English summary

This program will convert a Norwegian Trumf Visa invoice to a CSV file 
that can be imported into Actual Budget.

## Bakgrunn

Dette programmet konverterer en Trumf Visa PDF faktura til en CSV-fil 
som kan importeres til Actual Budget.

Trumf Visa lar deg kun hente ut en PDF fil i fra deres mobil app. 
Det er ingen mulighet for å hente ut en CSV fil eller annet maskinlesbart
format. 

Dette er elendig gjort av Trumf Visa og jeg håper at utviklerne der
virkelig skammer seg.


## Hvordan kompilere opp programmet

Dette programmet er skrevet i Golang og du må installer go på din maskin. 
Se https://go.dev/doc/install for mer informasjon.

Når du har installert go må du hente ned kildekoden til dette programmet.
Dette kan du enten gjøre ved å klone repoet eller laste ned zip fil.
Kjør deretter følgende kommando for å kompilere opp programmet:

```bash
go build
```

Du får nå en fil ved navn TrumVisa2ActualBudget som du kan kjøre.

## Kjøring av programmet

Denne filen tar en eller flere Trumf Visa PDF filer som argumenter og
konverterer disse til CSV filer som kan importeres til Actual Budget.

Filene lagres til samme katalog hvor PDF filene ligger.

### Eksempel på kjøring:

```bash
➜  TrumfVisa2ActualBudget git:(main) ✗ ./TrumVisa2ActualBudget testpdf/test.pdf
Wrote 67 transactions to testpdf/test.csv
➜  TrumfVisa2ActualBudget git:(main) ✗  wc -l testpdf/test.csv
      69 testpdf/test.csv
➜  TrumfVisa2ActualBudget git:(main) ✗
```

Grunnen til at det er to ekstra linjer i CSV filen kommer av at header linje 
samt en tom linje på slutten av filen.