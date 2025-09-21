package entities

type Book struct {
	ID              int    `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	Author          string `json:"author" db:"author"`
	CoverImageUrl   string `json:"cover_image_url" db:"cover_image_url"`
	Description     string `json:"description" db:"description"`
	PublicationDate string `json:"publication_date" db:"publication_date"`
	NumberOfPages   int    `json:"number_of_pages" db:"number_of_pages"`
	Isbn            string `json:"isbn" db:"isbn"`
}
