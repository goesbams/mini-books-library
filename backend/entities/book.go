package entities

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	CoverImageUrl   string `json:"cover_image_url"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
	NumberOfPages   int    `json:"number_of_pages"`
	Isbn            string `json:"isbn"`
}
