package request

type CreateCategory struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateCategory struct {
	Name string `json:"name" form:"name" validate:"required"`
}
