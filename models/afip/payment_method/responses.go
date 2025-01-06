package payment_method

type GetPaymentMethodsResponse struct {
	PaymentMethods []PaymentMethod `json:"payment_methods"`
}
