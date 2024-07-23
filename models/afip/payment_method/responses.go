package payment_method

type GetPaymentMethodsResponse struct {
	PaymentMethods []PaymentMethod `json:"payments_types"`
}
