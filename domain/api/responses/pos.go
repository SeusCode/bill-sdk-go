package responses

import "github.com/seuscode/bill-sdk-go/domain/pos"

type (
	GetPointsOfSaleResponse struct {
		POS []pos.PointOfSale `json:"sales_points"`
	}
)
