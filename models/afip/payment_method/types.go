package payment_method

type (
	PaymentMethod struct {
		Id          int    `json:"id"`
		Description string `json:"desc"`
	}
)
