package models

import "time"

type Books struct {
	ID            uint
	Name          string
	Description   string
	PublishedDate time.Time
	Edition int
	TotalPages int
	Language string
	isbn string
	imageUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Authors []Author
	Readers []Reader
}
