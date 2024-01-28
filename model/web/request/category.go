package request

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=99"`
}
