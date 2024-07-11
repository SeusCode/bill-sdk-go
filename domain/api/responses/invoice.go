package responses

type (
	CreateInvoiceResponse struct {
		CaeNumber         int `json:"CAE"`
		VoucherNumber     int `json:"voucher_number"`
		CaeExpirationDate int `json:"CAEFchVto"`
	}
)
