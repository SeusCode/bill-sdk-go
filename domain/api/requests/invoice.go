package requests

import (
	"github.com/seuscode/afip-sdk-go/domain/document"
	"github.com/seuscode/afip-sdk-go/domain/voucher"
)

/*
 * @param int $DocTipo Tipo de documento del comprador (REQUIRED, tipo 99 y numero de documento 0 para consumidor final, Facturas B y C hasta 95812 ó 191624 pesos, ver $limiteConsumidorFinal) (https://www.afip.gob.ar/fe/emision-autorizacion/datos-comprobantes.asp)
 *
 * //Ver descripcion de $items mas adelante.
 * @param array $Items Items del comprobante (REQUIRED. No Es estandar de AFIP, es propio del API y se usa para calcular otros importes que AFIP necesita, como ImpTotal, ImpNeto, ImpIVA, ImpOpEx, ImpTotConc)
 *
 * @param string $FchServDesde [NO SE USA EN CONCEPTO 1] Fecha de inicio del servicio (formato YYYYMMDD) (REQUIRED)
 * @param string $FchServHasta [NO SE USA EN CONCEPTO 1] Fecha de fin del servicio (formato YYYYMMDD) (REQUIRED)
 * @param string $FchVtoPago [NO SE USA EN CONCEPTO 1] Fecha de vencimiento del pago (formato YYYYMMDD)	(REQUIRED)
 *
 * @param array $Opcionales [OPCIONAL] Opcionales (Puede no enviarse, o enviarse null. Se usa para informacion complementaria/adicional en el comprobante)
 * @param array $Compradores [OPCIONAL] Compradores (Puede no enviarse, o enviarse null. Se usa para informacion de los compradores en caso de que haya mas de uno)
 */
type (
	CreateVoucher struct {
		PtoVta string `json:"PtoVta"`

		CbteTipo voucher.VoucherType    `json:"CbteTipo"`
		Concepto voucher.VoucherConcept `json:"Concepto"`

		DocTipo document.DocumentType `json:"DocTipo"`
		DocNro  int                   `json:"DocNro"`

		CbteFch *string `json:"CbteFch,omitempty"`

		FchServDesde *int `json:"FchServDesde,omitempty"`
		FchServHasta *int `json:"FchServHasta,omitempty"`
		FchVtoPago   *int `json:"FchVtoPago,omitempty"`

		Items []voucher.VoucherItems `json:"Items"`

		CbtesAsoc   []voucher.AsociatedVouchers `json:"CbtesAsoc,omitempty"`
		Iva         []voucher.VoucherShare      `json:"Iva,omitempty"`
		Tributos    []voucher.VoucherTributes   `json:"Tributos,omitempty"`
		Opcionales  []voucher.VoucherOptionals  `json:"Opcionales,omitempty"`
		Compradores []voucher.VoucherBuyers     `json:"Compradores,omitempty"`

		CompradorIvaExento *bool `json:"CompradorIvaExento,omitempty"`
		PagoContado        *bool `json:"pagoContado,omitempty"`
		GeneratePDF        *bool `json:"doPDF,omitempty"`

		MonId    *string  `json:"MonId,omitempty"`
		MonCotiz *float64 `json:"MonCotiz,omitempty"`
	}
)
