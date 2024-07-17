package concept

import "github.com/seuscode/bill-sdk-go/domain/voucher"

type (
	Concept struct {
		Id          voucher.VoucherConcept `json:"Id"`
		Description string                 `json:"Desc"`
		StartDate   string                 `json:"FchDesde"`
		EndDate     string                 `json:"FchHasta"`
	}
)
