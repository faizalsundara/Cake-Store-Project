package factory

import (
	"database/sql"
	_cakeBusiness "xsis/test/features/cakes/business"
	_cakeData "xsis/test/features/cakes/data"
	_cakePresentation "xsis/test/features/cakes/presentation"
)

type Presenter struct {
	CakePresenter *_cakePresentation.CakeHandler
}

func InitFactory(dbConn *sql.DB) Presenter {
	cakeData := _cakeData.NewCakeRepository(dbConn)
	cakeBusiness := _cakeBusiness.NewCakeBusiness(cakeData)
	cakePresentation := _cakePresentation.NewCakeHandler(cakeBusiness)

	return Presenter{
		CakePresenter: cakePresentation,
	}
}
