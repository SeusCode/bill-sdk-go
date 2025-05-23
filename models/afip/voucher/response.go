package voucher

type GetVouchersResponse struct {
	Vouchers []VoucherDetails `json:"vouchers"`
}
