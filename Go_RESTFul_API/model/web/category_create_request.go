package web

type CategoryCreateRequest struct {
	Name string `validate:"required,gte=1,lte=200"`
}
