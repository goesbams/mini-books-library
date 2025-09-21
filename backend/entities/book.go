package entities

import "github.com/go-playground/validator/v10"

type Book struct {
	ID              int    `json:"id" db:"id" form:"id"`
	Title           string `json:"title" db:"title" form:"title" validate:"required,min=2,max=255"`
	Author          string `json:"author" db:"author" form:"author" validate:"required,min=2,max=255"`
	CoverImageUrl   string `json:"cover_image_url" db:"cover_image_url" form:"cover_image_url" validate:"omitempty,url"`
	Description     string `json:"description" db:"description" form:"description" validate:"omitempty,max=1000"`
	PublicationDate string `json:"publication_date" db:"publication_date" form:"publication_date" validate:"required,datetime=2006-01-02"`
	NumberOfPages   int    `json:"number_of_pages" db:"number_of_pages" form:"number_of_pages" validate:"required,gt=0"`
	Isbn            string `json:"isbn" db:"isbn" form:"isbn" validate:"required,len=13,numeric"`
}

func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
