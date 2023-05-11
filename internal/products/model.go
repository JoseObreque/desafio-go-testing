package products

type Product struct {
	ID          string  `json:"id"`
	SellerID    string  `json:"seller_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
