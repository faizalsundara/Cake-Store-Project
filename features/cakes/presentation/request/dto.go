package request

import (
	cakes "xsis/test/features/cakes"
)

type Cake struct {
	Title       string  `json:"title" form:"title"`
	Description string  `json:"description" form:"description"`
	Rating      float64 `json:"rating" form:"rating"`
	Image       string  `json:"image" form:"image"`
}

func ToCore(req Cake) cakes.Core {
	return cakes.Core{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}
}
