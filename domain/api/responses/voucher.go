package responses

import "github.com/seuscode/bill-sdk-go/domain/voucher"

type (
	GetVoucherTypesResponse struct {
		Vouchers []voucher.Type `json:"voucher_types"`
	}
)
