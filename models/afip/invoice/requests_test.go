package invoice_test

import (
	"encoding/json"
	"testing"

	"github.com/seuscode/bill-sdk-go/v2/models/afip/invoice"
)

func TestGenerateInvoicePDFRequestOmitsEmptyPaperSize(t *testing.T) {
	body, err := json.Marshal(invoice.GenerateInvoicePDFRequest{})
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}

	if _, ok := payload["PaperSize"]; ok {
		t.Fatalf("expected PaperSize to be omitted, got payload %s", body)
	}
}

func TestGenerateInvoicePDFRequestIncludesPaperSize(t *testing.T) {
	body, err := json.Marshal(invoice.GenerateInvoicePDFRequest{
		PaperSize: invoice.PAPER_SIZE_80MM,
	})
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}

	if got := payload["PaperSize"]; got != string(invoice.PAPER_SIZE_80MM) {
		t.Fatalf("expected PaperSize %q, got %v in payload %s", invoice.PAPER_SIZE_80MM, got, body)
	}
}
