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
	GetAllCake() (data []Core, err error)
	GetCakeDetail(idCake int) (data Core, err error)
	PostNewCake(data Core) (row int, err error)
	PatchCake(idCake int, data Core) (row int, err error)
}

type Data interface {
	ListOfCake() (data []Core, err error)
	DetailOfCake(idCake int) (data Core, err error)
	AddNewCake(data Core) (row int, err error)
	UpdateCake(idCake int, data Core) (row int, err error)
}
