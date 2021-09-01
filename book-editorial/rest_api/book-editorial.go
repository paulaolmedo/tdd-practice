package rest_api

var Books = []string{"isbn1", "isbn2", "isbn3", "isbn4"}

type Editorial struct {
	Library []string
}

func NewEditorial() Editorial{
	return Editorial{
		Library: Books,
	}
}