package afip

import (
	"fmt"
	"strings"

	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/concept"
	"github.com/seuscode/bill-sdk-go/models/afip/currency"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/afip/invoice"
	"github.com/seuscode/bill-sdk-go/models/afip/optionals"
	"github.com/seuscode/bill-sdk-go/models/afip/payment_method"
	"github.com/seuscode/bill-sdk-go/models/afip/pos"
	"github.com/seuscode/bill-sdk-go/models/afip/tribute"
	"github.com/seuscode/bill-sdk-go/models/afip/voucher"
	"github.com/seuscode/bill-sdk-go/pkg/http"
)

type electronicBilling interface {

	/**
	 * Get complete voucher information
	 *
	 * Asks to AFIP servers for complete information of voucher {@see WS
	 * Specification item 4.19}
	 *
	 * @param int number 		Number of voucher to get information
	 * @param int pos 	Point of sales of voucher to get information
	 * @param int type 			Type of voucher to get information
	 *
	 * @return array|null returns array with complete voucher information
	 * 	{@see WS Specification item 4.19} or null if there not exists
	 **/
	GetIssuedInvoiceData(Pos, invNbr int, invType invoice.InvoiceType) (*invoice.GetInvoiceDataResponse, *http.ApiErrorDetails)

	/*
	 Create a voucher from AFIP

	 Send to AFIP servers a request to issue an invoice and assign
	 CAE to them {@see WS Specification item 4.1}

	 @params array data Voucher parameters {@see WS Specification
	 	item 4.1.3}, some arrays were simplified for easy use {@example
	 	examples/createVoucher.js Example with all allowed
	 	 attributes}

	 @return array
	 	[CAE : CAE assigned to voucher, CAEFchVto : Expiration date
	 	for CAE (yyyy-mm-dd)] else returns complete response from
	 	AFIP {@see WS Specification item 4.1.3}
	**/
	IssueInvoice(data *invoice.IssueInvoiceRequest) (*invoice.IssueInvoiceResponse, *http.ApiErrorDetails)

	/**
	 * Create PDF
	 *
	 * Send a request to Afip SDK server to create a PDF
	 *
	 * @param {object} data Data for PDF creation
	**/
	GenerateInvoicePDF(data invoice.GenerateInvoicePDFRequest, folderName, fileName string) (string, error)

	/**
	 * Asks to AFIP Servers for sales points availables {@see WS
	 * Specification item 4.11}
	 *
	 * @return array All sales points availables
	 **/
	GetSalesPoints() (*pos.GetPointsOfSaleResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for voucher types availables {@see WS
	 * Specification item 4.4}
	 *
	 * @return array All voucher types availables
	 **/
	GetVouchers() (*voucher.GetVouchersResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for voucher concepts availables {@see WS
	 * Specification item 4.5}
	 *
	 * @return array All voucher concepts availables
	 **/
	GetConcepts() (*concept.GetConceptsResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for document types availables {@see WS
	 * Specification item 4.6}
	 *
	 * @return array All document types availables
	 **/
	GetDocuments() (*document.GetDocumentsResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for aliquot availables {@see WS
	 * Specification item 4.7}
	 *
	 * @return array All aliquot availables
	 **/
	GetAliquots() (*aliquot.GetAliquotsResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for currencies availables {@see WS
	 * Specification item 4.8}
	 *
	 * @return array All currencies availables
	 **/
	GetCurrencies() (*currency.GetCurrenciesResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for tax availables {@see WS
	 * Specification item 4.10}
	 *
	 * @return array All tax availables
	 **/
	GetTributes() (*tribute.GetTributesResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for a currency cotization
	 *
	 * @return array All currency cotization i
	 **/
	GetCurrencyExchangeRate(currencyId string) (*currency.GetCurrencyExchangeRateResponse, *http.ApiErrorDetails)

	/**
	 * Asks to AFIP Servers for optional types
	 *
	 * @return array All optionals availables
	 **/
	GetOptionals() (*optionals.GetOptionalsResponse, *http.ApiErrorDetails)

	/**
	 * Asks to our API Servers for payment methods
	 *
	 * @return array All payment methods availables
	 **/
	GetPaymentMethods() (*payment_method.GetPaymentMethodsResponse, *http.ApiErrorDetails)
}

type eBilling struct {
	afip *AfipData
}

func newElectronicBilling(afip *AfipData) electronicBilling {
	return &eBilling{
		afip: afip,
	}
}

/*
	-=============================-
	-== Functions Implementation =-
	-=============================-
*/

func (e *eBilling) IssueInvoice(invoiceData *invoice.IssueInvoiceRequest) (*invoice.IssueInvoiceResponse, *http.ApiErrorDetails) {
	var response invoice.IssueInvoiceResponse

	err := e.afip.HttpClient.Post(ENDPOINT_INVOICE, invoiceData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (e *eBilling) GenerateInvoicePDF(data invoice.GenerateInvoicePDFRequest, folderName, fileName string) (string, error) {
	fPath, err := e.afip.HttpClient.PostWithFileOnResponse(ENDPOINT_INVOICE_PDF, data, folderName, fileName)
	return fPath, err
}

func (e *eBilling) GetIssuedInvoiceData(pos, invNbr int, invType invoice.InvoiceType) (*invoice.GetInvoiceDataResponse, *http.ApiErrorDetails) {
	var res invoice.GetInvoiceDataResponse

	err := e.afip.HttpClient.Get(fmt.Sprintf("%s/%d/%d/%d", ENDPOINT_INVOICE, pos, invType, invNbr), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetSalesPoints() (*pos.GetPointsOfSaleResponse, *http.ApiErrorDetails) {
	var res pos.GetPointsOfSaleResponse

	err := e.afip.HttpClient.Get(ENDPOINT_POINTS_OF_SALES, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetVouchers() (*voucher.GetVouchersResponse, *http.ApiErrorDetails) {
	var res voucher.GetVouchersResponse

	err := e.afip.HttpClient.Get(ENDPOINT_VOUCHERS, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetConcepts() (*concept.GetConceptsResponse, *http.ApiErrorDetails) {
	var res concept.GetConceptsResponse

	err := e.afip.HttpClient.Get(ENDPOINT_CONCEPTS, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetDocuments() (*document.GetDocumentsResponse, *http.ApiErrorDetails) {
	var res document.GetDocumentsResponse

	err := e.afip.HttpClient.Get(ENDPOINT_DOCUMENTS, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetAliquots() (*aliquot.GetAliquotsResponse, *http.ApiErrorDetails) {
	var res aliquot.GetAliquotsResponse

	err := e.afip.HttpClient.Get(ENDPOINT_ALIQUOTS, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetCurrencies() (*currency.GetCurrenciesResponse, *http.ApiErrorDetails) {
	var res currency.GetCurrenciesResponse

	err := e.afip.HttpClient.Get(ENDPOINT_CURRENCIES, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetTributes() (*tribute.GetTributesResponse, *http.ApiErrorDetails) {
	var res tribute.GetTributesResponse

	err := e.afip.HttpClient.Get(ENDPOINT_TRIBUTES, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetPaymentMethods() (*payment_method.GetPaymentMethodsResponse, *http.ApiErrorDetails) {
	var res payment_method.GetPaymentMethodsResponse

	err := e.afip.HttpClient.Get(ENDPOINT_PAYMENT_METHODS, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetCurrencyExchangeRate(currencyId string) (*currency.GetCurrencyExchangeRateResponse, *http.ApiErrorDetails) {
	var res currency.GetCurrencyExchangeRateResponse

	err := e.afip.HttpClient.Get(strings.ReplaceAll(ENDPOINT_CURRENCY_EXCHANGE_RATE, "{currencyId}", currencyId), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (e *eBilling) GetOptionals() (*optionals.GetOptionalsResponse, *http.ApiErrorDetails) {
	var res optionals.GetOptionalsResponse

	err := e.afip.HttpClient.Get(ENDPOINT_OPTIONALS, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
