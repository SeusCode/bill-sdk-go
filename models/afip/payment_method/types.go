package payment_method

type (
	SellConditionId int
	PaymentMethod   struct {
		Id          SellConditionId `json:"id"`
		Description string          `json:"desc"`
	}
)

const (
	Cash            SellConditionId = 1
	DebitCard       SellConditionId = 2
	CreditCard      SellConditionId = 3
	Check           SellConditionId = 4
	CheckingAccount SellConditionId = 5
	Other           SellConditionId = 8
)
