package model

import (
	"database/sql"
	"errors"
	"log"
)

type Gallery struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GalleryModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var galleries = []Gallery{
	{
		Id:      "1",
		Title:   "Kuromi",
		Description: "Pictures of the best Sanrio Character.",
	},
	{
		Id:      "2",
		Title:   "Puppies",
		Description: "The cutest street dogs puppies I have found while walking in my neighborhood.",
	},
	{
		Id:      "3",
		Title:   "Cats",
		Description: "Collection of pictures of my favorite animal. Cats from Korea, Turkey and Kazakhstan!",
	},
	{
		Id:      "4",
		Title:   "PomPomPurin",
		Description: "My third favorite Sanrio character after Kuromi. His hat is so cute.",
	},
	{
		Id:      "5",
		Title:   "CinamonRoll",
		Description: "Second best Sanrio after Kuromi!. eventhough looks like a bunny he is actually a puppy.",
	},
}

func GetGalleries() []Gallery {
	return galleries
}

func GetGallery(id string) (*Gallery, error) {
	for _, g := range galleries {
		if g.Id == id {
			return &g, nil
		}
	}
	return nil, errors.New("Galleries not found")
}