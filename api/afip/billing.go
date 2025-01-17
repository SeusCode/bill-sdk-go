package afip

import (
	"errors"
	"fmt"

	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/concept"
	"github.com/seuscode/bill-sdk-go/models/afip/currency"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/afip/optionals"
	"github.com/seuscode/bill-sdk-go/models/afip/payment_method"
	"github.com/seuscode/bill-sdk-go/models/afip/pos"
	"github.com/seuscode/bill-sdk-go/models/afip/tax"
	"github.com/seuscode/bill-sdk-go/models/afip/voucher"
	"github.com/seuscode/bill-sdk-go/pkg/endpoints"
	"github.com/seuscode/bill-sdk-go/pkg/http"
)

type electronicBilling interface {
	/**
	 * Create PDF
	 *
	 * Send a request to Afip SDK server to create a PDF
	 *
	 * @param {object} data Data for PDF creation
	**/
	CreatePDF(data voucher.CreateVoucherPDFRequest, folderName, fileName string) (string, error)

	/*
	 Create a voucher from AFIP

	 Send to AFIP servers request for create a voucher and assign
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
	CreateVoucher(data *voucher.Voucher, response *voucher.CreateVoucherResponse) error

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
	GetVoucherInfo(voucherNumber, voucherPos int, voucherType voucher.VoucherType) (*voucher.GetVoucherInfoResponse, error)

	/**
	 * Asks to AFIP Servers for sales points availables {@see WS
	 * Specification item 4.11}
	 *
	 * @return array All sales points availables
	 **/
	GetSalesPoints() (*pos.GetPointsOfSaleResponse, error)

	/**
	 * Asks to AFIP Servers for voucher types availables {@see WS
	 * Specification item 4.4}
	 *
	 * @return array All voucher types availables
	 **/
	GetVoucherTypes() (*voucher.GetVoucherTypesResponse, error)

	/**
	 * Asks to AFIP Servers for voucher concepts availables {@see WS
	 * Specification item 4.5}
	 *
	 * @return array All voucher concepts availables
	 **/
	GetConceptTypes() (*concept.GetConceptTypesResponse, error)

	/**
	 * Asks to AFIP Servers for document types availables {@see WS
	 * Specification item 4.6}
	 *
	 * @return array All document types availables
	 **/
	GetDocumentTypes() (*document.GetDocumentTypesResponse, error)

	/**
	 * Asks to AFIP Servers for aliquot availables {@see WS
	 * Specification item 4.7}
	 *
	 * @return array All aliquot availables
	 **/
	GetAliquotTypes() (*aliquot.GetAliquotTypesResponse, error)

	/**
	 * Asks to AFIP Servers for currencies availables {@see WS
	 * Specification item 4.8}
	 *
	 * @return array All currencies availables
	 **/
	GetCurrenciesTypes() (*currency.GetCurrencyTypesResponse, error)

	/**
	 * Asks to AFIP Servers for tax availables {@see WS
	 * Specification item 4.10}
	 *
	 * @return array All tax availables
	 **/
	GetTaxTypes() (*tax.GetTaxTypesResponse, error)

	/**
	 * Asks to AFIP Servers for a currency cotization
	 *
	 * @return array All currency cotization i
	 **/
	GetCurrencyCotization(currencyId string) (*currency.GetCurrencyCotizationResponse, error)

	/**
	 * Asks to AFIP Servers for optional types
	 *
	 * @return array All optionals availables
	 **/
	GetOptionalTypes() (*optionals.GetOptionalTypesResponse, error)

	/**
	 * Asks to our API Servers for payment methods
	 *
	 * @return array All payment methods availables
	 **/
	GetPaymentMethods() (*payment_method.GetPaymentMethodsResponse, error)
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

func (e *eBilling) CreateVoucher(data *voucher.Voucher, response *voucher.CreateVoucherResponse) error {
	r := voucher.CreateVoucherRequest{
		CbteTipo: data.CbteTipo,
		Concepto: data.Concepto,

		DocTipo: data.DocTipo,
		DocNro:  data.DocNro,

		CbteFch: data.CbteFch,
		Items:   data.Items,

		CbtesAsoc:   data.CbtesAsoc,
		Tributos:    data.Tributos,
		Opcionales:  data.Opcionales,
		Compradores: data.Compradores,

		CompradorIvaExento: data.CompradorIvaExento,

		MonId:    data.MonId,
		MonCotiz: data.MonCotiz,
	}

	/*
		This fields are only required for vouchers
		wich concepts are not only products, validate
		that cases
	*/
	if r.Concepto != voucher.Productos {
		if data.FchServDesde == nil {
			return errors.New("missing required field for this voucher: FchServDesde")
		}

		if data.FchServHasta == nil {
			return errors.New("missing required field for this voucher: FchServDesde")
		}

		if data.FchVtoPago == nil {
			return errors.New("missing required field for this voucher: FchServDesde")
		}

		r.FchServDesde = data.FchServDesde
		r.FchServHasta = data.FchServHasta
		r.FchVtoPago = data.FchVtoPago
	}

	/*
		If document type is final consumer set document
		number as zero.
	*/
	if r.DocTipo == document.CF {
		r.DocNro = 0
	}

	f := false

	if r.CompradorIvaExento == nil {
		r.CompradorIvaExento = &f
	}

	if (r.MonId != nil && r.MonCotiz == nil) || (r.MonId == nil && r.MonCotiz != nil) {
		return errors.New("if you send one of this fields (MonId or MonCotiz) you must send the other too")
	}

	apiStatus, err := e.afip.HttpClient.Post(endpoints.INVOICE, r, response)
	if err != nil {
		return err
	}

	if apiStatus.Status.Type != http.SUCCESS {
		return fmt.Errorf("error (%s): %s", apiStatus.Status.Code, apiStatus.Status.Description)
	}

	return nil
}

func (e *eBilling) CreatePDF(data voucher.CreateVoucherPDFRequest, folderName, fileName string) (string, error) {
	fPath, err := e.afip.HttpClient.PostWithFileOnResponse(endpoints.GET_PDF, data, folderName, fileName)
	return fPath, err
}

func (e *eBilling) GetVoucherInfo(voucherNumber, voucherPos int, voucherType voucher.VoucherType) (*voucher.GetVoucherInfoResponse, error) {
	var res voucher.GetVoucherInfoResponse

	apiResponse, err := e.afip.HttpClient.Get(fmt.Sprintf("%s/%d/%d/%d", endpoints.INVOICE, voucherType, voucherPos, voucherNumber), &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetSalesPoints() (*pos.GetPointsOfSaleResponse, error) {
	var res pos.GetPointsOfSaleResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.SALES_POINTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetVoucherTypes() (*voucher.GetVoucherTypesResponse, error) {
	var res voucher.GetVoucherTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.VOUCHERS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetConceptTypes() (*concept.GetConceptTypesResponse, error) {
	var res concept.GetConceptTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.CONCEPTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetDocumentTypes() (*document.GetDocumentTypesResponse, error) {
	var res document.GetDocumentTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.DOCUMENTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetAliquotTypes() (*aliquot.GetAliquotTypesResponse, error) {
	var res aliquot.GetAliquotTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.ALIQUOTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetCurrenciesTypes() (*currency.GetCurrencyTypesResponse, error) {
	var res currency.GetCurrencyTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.CURRENCIES, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetTaxTypes() (*tax.GetTaxTypesResponse, error) {
	var res tax.GetTaxTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.TAXES, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetPaymentMethods() (*payment_method.GetPaymentMethodsResponse, error) {
	var res payment_method.GetPaymentMethodsResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.PAYMENT_METHODS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetCurrencyCotization(currencyId string) (*currency.GetCurrencyCotizationResponse, error) {
	var res currency.GetCurrencyCotizationResponse

	apiResponse, err := e.afip.HttpClient.Get(fmt.Sprintf("%s/%s", endpoints.COTIZATIONS, currencyId), &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetOptionalTypes() (*optionals.GetOptionalTypesResponse, error) {
	var res optionals.GetOptionalTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.OPTIONALS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}
