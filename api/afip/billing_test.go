package afip

import (
	"testing"

	"github.com/seuscode/bill-sdk-go/v2/models/afip/invoice"
)

func TestInvoiceStatementFileNameKeepsPDFDefault(t *testing.T) {
	got := invoiceStatementFileName(invoice.GenerateInvoicePDFRequest{}, "invoice.pdf")
	if got != "invoice.pdf" {
		t.Fatalf("expected invoice.pdf, got %s", got)
	}
}

func TestInvoiceStatementFileNameKeepsPDFForA4(t *testing.T) {
	got := invoiceStatementFileName(invoice.GenerateInvoicePDFRequest{
		PaperSize: invoice.PAPER_SIZE_A4,
	}, "invoice.pdf")
	if got != "invoice.pdf" {
		t.Fatalf("expected invoice.pdf, got %s", got)
	}
}

func TestInvoiceStatementFileNameUsesHTMLFor80MM(t *testing.T) {
	got := invoiceStatementFileName(invoice.GenerateInvoicePDFRequest{
		PaperSize: invoice.PAPER_SIZE_80MM,
	}, "invoice.pdf")
	if got != "invoice.html" {
		t.Fatalf("expected invoice.html, got %s", got)
	}
}

func TestInvoiceStatementFileNameAddsHTMLExtensionFor80MM(t *testing.T) {
	got := invoiceStatementFileName(invoice.GenerateInvoicePDFRequest{
		PaperSize: invoice.PAPER_SIZE_80MM,
	}, "invoice")
	if got != "invoice.html" {
		t.Fatalf("expected invoice.html, got %s", got)
	}
}

func TestInvoiceStatementFileNameKeepsExplicitHTMLExtension(t *testing.T) {
	got := invoiceStatementFileName(invoice.GenerateInvoicePDFRequest{
		PaperSize: invoice.PAPER_SIZE_80MM,
	}, "invoice.html")
	if got != "invoice.html" {
		t.Fatalf("expected invoice.html, got %s", got)
	}
}
