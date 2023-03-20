package models


type BookModel struct {
	ID            uint `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PublishedDate time.Time `json:"published_date"`
	Edition int `json:"edition"`
	TotalPages int `json:"total_pages"`
	Language string `json:"language"`
	isbn string `json:"isbn"`
	imageUrl string `json:"image_url"`
	AuthorID int `json:"autor_id"`
	ReaderID int `json:"reader_id"`
}
