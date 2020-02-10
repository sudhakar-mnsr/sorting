package curlib

import (
   "encoding/csv"
   "io"
   "os"
   "strings"
)

type Currency struct {
   Code string `json:"currency_code"`
   Name string `json:"currency_name"`
   Number string `json:"currency_number"`
   Country string `json:"currency_country"`
}

type CurrencyRequest struct {
   Get string `json:"get"`
}

type CurrencyError struct {
   Error string `json:"currency_error"`
}

