package request

import (
	"time"
	cakes "xsis/test/features/cakes"
)

type Cake struct {
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Rating      float64   `json:"rating" form:"rating"`
	Image       string    `json:"image" form:"image"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

func ToCore(req Cake) cakes.Core {
	return cakes.Core{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		UpdatedAt:   req.UpdatedAt,
	}
}
