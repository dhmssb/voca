package presentations

type CreateProductParams struct {
	UserProduct        string `json:"user_product"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	Quantity           int32  `json:"quantity"`
	Price              int64  `json:"price"`
}

type UpdateProductQuantityParams struct {
	ID       int64 `json:"id"`
	Quantity int32 `json:"quantity"`
}

type ListProductsParams struct {
	UserProduct string `json:"user_product"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

type GetProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
