package routes

import (
	"xsis/test/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.GET("cakes/:id", presenter.CakePresenter.GetCakeDetail)

	return e
}
