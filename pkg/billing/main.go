package billing

import (
	"errors"
	"fmt"

	"github.com/seuscode/bill-sdk-go/domain/api/requests"
	"github.com/seuscode/bill-sdk-go/domain/api/responses"
	"github.com/seuscode/bill-sdk-go/domain/document"
	"github.com/seuscode/bill-sdk-go/domain/voucher"
	"github.com/seuscode/bill-sdk-go/endpoints"
	"github.com/seuscode/bill-sdk-go/gofip"
)

type ElectronicBilling interface {
	/**
	 * Create PDF
	 *
	 * Send a request to Afip SDK server to create a PDF
	 *
	 * @param {object} data Data for PDF creation
	**/
	CreatePDF(data interface{}) error

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
	CreateVoucher(data *voucher.Voucher, response *responses.CreateInvoiceResponse) error

	/**
	 * Create next voucher from AFIP
	 *
	 * This method combines Afip.getLastVoucher and Afip.createVoucher
	 * for create the next voucher
	 *
	 * @param array data Same to data in Afip.createVoucher except that
	 * 	don't need CbteDesde and CbteHasta attributes
	 *
	 * @return array [CAE : CAE assigned to voucher, CAEFchVto : Expiration
	 * 	date for CAE (yyyy-mm-dd), voucherNumber : Number assigned to
	 * 	voucher]
	 **/
	//CreateNextVoucher(data *Voucher) error

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
	GetVoucherInfo(voucherNumber, voucherPos int, voucherType voucher.VoucherType) error

	/**
	 * Create CAEA (Código de Autorización Electrónico Anticipado)
	 *
	 * Send a request to AFIP servers  to create a CAEA
	 *
	 * @param int period 	Time period
	 * @param periodFrequency period frequency
	 **/
	//CreateCAEA(period int, periodFrequency CaeaPeriodFrequency) error

	/**
	 * Asks to AFIP Servers for sales points availables {@see WS
	 * Specification item 4.11}
	 *
	 * @return array All sales points availables
	 **/
	GetSalesPoints() (*responses.GetPointsOfSaleResponse, error)

	/**
	 * Asks to AFIP Servers for voucher types availables {@see WS
	 * Specification item 4.4}
	 *
	 * @return array All voucher types availables
	 **/
	GetVoucherTypes() (interface{}, error)

	/**
	 * Asks to AFIP Servers for voucher concepts availables {@see WS
	 * Specification item 4.5}
	 *
	 * @return array All voucher concepts availables
	 **/
	GetConceptTypes() (*responses.GetConceptTypesResponse, error)

	/**
	 * Asks to AFIP Servers for document types availables {@see WS
	 * Specification item 4.6}
	 *
	 * @return array All document types availables
	 **/
	GetDocumentTypes() (*responses.GetDocumentTypesResponse, error)

	/**
	 * Asks to AFIP Servers for aliquot availables {@see WS
	 * Specification item 4.7}
	 *
	 * @return array All aliquot availables
	 **/
	GetAliquotTypes() (*responses.GetAliquotTypesResponse, error)

	/**
	 * Asks to AFIP Servers for currencies availables {@see WS
	 * Specification item 4.8}
	 *
	 * @return array All currencies availables
	 **/
	GetCurrenciesTypes() (*responses.GetCurrencyTypesResponse, error)

	/**
	 * Asks to AFIP Servers for tax availables {@see WS
	 * Specification item 4.10}
	 *
	 * @return array All tax availables
	 **/
	GetTaxTypes() (*responses.GetTaxTypesResponse, error)
}

type eBilling struct {
	afip *gofip.Gofip
}

type BillingOptions struct {
}

func NewElectronicBilling(afip *gofip.Gofip, opts BillingOptions) ElectronicBilling {
	return &eBilling{
		afip: afip,
	}
}

/*
	-=============================-
	-== Functions Implementation =-
	-=============================-
*/

func (e *eBilling) CreateVoucher(data *voucher.Voucher, response *responses.CreateInvoiceResponse) error {
	r := requests.CreateVoucher{
		PtoVta: e.afip.PointOfSale,

		CbteTipo: data.CbteTipo,
		Concepto: data.Concepto,

		DocTipo: data.DocTipo,
		DocNro:  data.DocNro,

		CbteFch: data.CbteFch,
		Items:   data.Items,

		CbtesAsoc:   data.CbtesAsoc,
		Iva:         data.Iva,
		Tributos:    data.Tributos,
		Opcionales:  data.Opcionales,
		Compradores: data.Compradores,

		CompradorIvaExento: data.CompradorIvaExento,
		PagoContado:        data.PagoContado,
		GeneratePDF:        data.GeneratePDF,

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

	if r.PagoContado == nil {
		r.PagoContado = &f
	}

	if r.GeneratePDF == nil {
		r.GeneratePDF = &f
	}

	if (r.MonId != nil && r.MonCotiz == nil) || (r.MonId == nil && r.MonCotiz != nil) {
		return errors.New("if you send one of this fields (MonId or MonCotiz) you must send the other too")
	}

	if r.MonId == nil {
		ars := "PES"
		arsCot := 1.0
		r.MonId = &ars
		r.MonCotiz = &arsCot
	}

	apiStatus, err := e.afip.HttpClient.Post(endpoints.INVOICE, r, response)
	if err != nil {
		return err
	}

	if apiStatus.Status.Type != gofip.SUCCESS {
		return fmt.Errorf("error (%s): %s", apiStatus.Status.Code, apiStatus.Status.Description)
	}

	return nil
}

func (e *eBilling) CreatePDF(data interface{}) error {
	return nil
}

func (e *eBilling) GetVoucherInfo(voucherNumber, voucherPos int, voucherType voucher.VoucherType) error {
	return nil

}

func (e *eBilling) GetSalesPoints() (*responses.GetPointsOfSaleResponse, error) {
	var res responses.GetPointsOfSaleResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.SALES_POINTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != gofip.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetVoucherTypes() (interface{}, error) {
	var res interface{}

	apiResponse, err := e.afip.HttpClient.Get(endpoints.VOUCHERS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != gofip.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return res, nil
}

func (e *eBilling) GetConceptTypes() (*responses.GetConceptTypesResponse, error) {
	var res responses.GetConceptTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.CONCEPTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != gofip.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetDocumentTypes() (*responses.GetDocumentTypesResponse, error) {
	var res responses.GetDocumentTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.DOCUMENTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != gofip.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetAliquotTypes() (*responses.GetAliquotTypesResponse, error) {
	var res responses.GetAliquotTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.ALIQUOTS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != gofip.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetCurrenciesTypes() (*responses.GetCurrencyTypesResponse, error) {
	var res responses.GetCurrencyTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.CURRENCIES, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != gofip.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (e *eBilling) GetTaxTypes() (*responses.GetTaxTypesResponse, error) {
	var res responses.GetTaxTypesResponse

	apiResponse, err := e.afip.HttpClient.Get(endpoints.TAXES, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != gofip.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}
