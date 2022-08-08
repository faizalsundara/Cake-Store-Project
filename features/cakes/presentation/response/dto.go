package response

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
}

func FromCore(data cakes.Core) Cake {
	return Cake{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func FromCoreList(data []cakes.Core) []Cake {
	result := []Cake{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
