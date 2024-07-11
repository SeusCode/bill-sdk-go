package billing

import "github.com/seuscode/afip-sdk-go/domain/voucher"

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

	 @params bool returnResponse if is TRUE returns complete response
	 	from AFIP

	 @return array if returnResponse is set to false returns
	 	[CAE : CAE assigned to voucher, CAEFchVto : Expiration date
	 	for CAE (yyyy-mm-dd)] else returns complete response from
	 	AFIP {@see WS Specification item 4.1.3}
	**/
	CreateVoucher(data *voucher.Voucher, returnRes bool) error

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
	GetSalesPoints() (interface{}, error)

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
	GetConceptTypes() (interface{}, error)

	/**
	 * Asks to AFIP Servers for document types availables {@see WS
	 * Specification item 4.6}
	 *
	 * @return array All document types availables
	 **/
	GetDocumentTypes() (interface{}, error)

	/**
	 * Asks to AFIP Servers for aliquot availables {@see WS
	 * Specification item 4.7}
	 *
	 * @return array All aliquot availables
	 **/
	GetAliquotTypes() (interface{}, error)

	/**
	 * Asks to AFIP Servers for currencies availables {@see WS
	 * Specification item 4.8}
	 *
	 * @return array All currencies availables
	 **/
	GetCurrenciesTypes() (interface{}, error)

	/**
	 * Asks to AFIP Servers for voucher optional data available {@see WS
	 * Specification item 4.9}
	 *
	 * @return array All voucher optional data available
	 **/
	GetOptionsTypes() (interface{}, error)

	/**
	 * Asks to AFIP Servers for tax availables {@see WS
	 * Specification item 4.10}
	 *
	 * @return array All tax availables
	 **/
	GetTaxTypes() (interface{}, error)

	/**
	 * Asks to web service for servers status {@see WS
	 * Specification item 4.14}
	 *
	 * @return object { AppServer : Web Service status,
	 * DbServer : Database status, AuthServer : Autentication
	 * server status}
	 **/
	//GetServerStatus() (interface{}, error)
}

type eBilling struct {
}

type Options struct {
}

func NewElectronicBilling(config interface{}, opts Options) ElectronicBilling {
	return &eBilling{}
}

/*
	-=============================-
	-== Functions Implementation =-
	-=============================-
*/

func (e *eBilling) CreateVoucher(data *voucher.Voucher, returnRes bool) error {
	return nil
}

func (e *eBilling) CreatePDF(data interface{}) error {
	return nil
}

func (e *eBilling) GetVoucherInfo(voucherNumber, voucherPos int, voucherType voucher.VoucherType) error {
	return nil

}

func (e *eBilling) GetSalesPoints() (interface{}, error) {
	return nil, nil
}

func (e *eBilling) GetVoucherTypes() (interface{}, error) {
	return nil, nil
}

func (e *eBilling) GetConceptTypes() (interface{}, error) {
	return nil, nil
}

func (e *eBilling) GetDocumentTypes() (interface{}, error) {
	return nil, nil
}

func (e *eBilling) GetAliquotTypes() (interface{}, error) {
	return nil, nil
}

func (e *eBilling) GetCurrenciesTypes() (interface{}, error) {
	return nil, nil
}

func (e *eBilling) GetOptionsTypes() (interface{}, error) {
	return nil, nil
}

func (e *eBilling) GetTaxTypes() (interface{}, error) {
	return nil, nil
}
