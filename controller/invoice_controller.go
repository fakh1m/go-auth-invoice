package controller

import (
	"encoding/json"
	"net/http"
)

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoices := []map[string]interface{}{
		{"id": 1, "title": "Invoice A", "amount": 100},
	}
	json.NewEncoder(w).Encode(invoices)
}
