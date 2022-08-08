package data

import (
	"time"
	cakes "xsis/test/features/cakes"
)

type Cake struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (data *Cake) toCore() cakes.Core {
	return cakes.Core{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func toCoreList(data []Cake) []cakes.Core {
	result := []cakes.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core cakes.Core) Cake {
	return Cake{
		Title:       core.Title,
		Description: core.Description,
		Rating:      core.Rating,
		Image:       core.Image,
	}
}
