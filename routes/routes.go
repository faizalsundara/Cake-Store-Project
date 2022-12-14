package routes

import (
	"xsis/test/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.GET("cakes/:id", presenter.CakePresenter.GetCakeDetail)
	e.GET("cakes", presenter.CakePresenter.GetAllCake)
	e.POST("cakes", presenter.CakePresenter.AddNewCake)
	e.PATCH("cakes/:id", presenter.CakePresenter.UpdateCake)
	e.DELETE("cakes/:id", presenter.CakePresenter.DeleteCake)

	return e
}
