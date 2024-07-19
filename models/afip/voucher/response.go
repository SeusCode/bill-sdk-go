package voucher

type (
	CreateVoucherResponse struct {
		CaeNumber         string      `json:"CAE"`
		VoucherNumber     int         `json:"voucher_number"`
		CaeExpirationDate string      `json:"CAEFchVto"`
		AfipInfo          interface{} `json:"info_to_afip"`
		QrString          string      `json:"string_qr_b64"`
		QrPresent         bool        `json:"qr_present"`
	}

	GetVoucherTypesResponse struct {
		Vouchers []Type `json:"voucher_types"`
	}
)
