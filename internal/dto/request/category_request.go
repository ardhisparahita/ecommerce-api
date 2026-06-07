package request

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
