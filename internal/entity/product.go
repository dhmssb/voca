package entity

type Product struct {
	ID                 int64  `json:"id"`
	UserProduct        string `json:"user_product"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	Quantity           int32  `json:"quantity"`
	Price              int64  `json:"price"`
}
