package cake

import "time"

type Core struct {
	ID          int
	Title       string
	Description string
	Rating      float64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	// GetAllCake(limit, offset int) (data []Core, totalPage int, err error)
	GetCakeDetail(idCake int) (data Core, err error)
}

type Data interface {
	// ListOfCake(limit, offset int) (data []Core, err error)
	DetailOfCake(idCake int) (data Core, err error)
}
