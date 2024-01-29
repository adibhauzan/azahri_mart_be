package request

type CreateProductCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=99"`
}
