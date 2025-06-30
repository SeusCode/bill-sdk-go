package receipt

type (
	Receipt struct {
		ReceiptNbr  string
		ReceiptDate string

		Subtotal float64
		Total    float64
	}

	ReceiptMetaData struct {
		ClientFiscalType         uint   `json:"client_fiscal_type,omitempty" uker:"required"`
		ClientPhone              string `json:"client_phone,omitempty"`
		ClientEmail              string `json:"client_email,omitempty"`
		ClientAddress            string `json:"client_address,omitempty"`
		ClientNameOrBusinessName string `json:"client_name,omitempty"`
		ClientDocument           int64
	}

	ReceiptItem struct {
		Id       string
		Desc     string
		Cantidad float64
		Price    float64
		Total    float64
	}

	GenerateReceiptPDFRequest struct {
		ReceiptData     Receipt         `json:"data" uker:"required"`
		ReceiptMetaData ReceiptMetaData `json:"metadata" uker:"required"`
		ReceiptProducts []ReceiptItem   `json:"products" uker:"required"`
	}
)
