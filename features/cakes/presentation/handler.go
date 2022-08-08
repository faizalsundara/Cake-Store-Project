package presentation

import (
	"net/http"
	"strconv"
	cakes "xsis/test/features/cakes"
	_responseCake "xsis/test/features/cakes/presentation/response"
	"xsis/test/helper"

	"github.com/labstack/echo/v4"
)

type CakeHandler struct {
	cakeBusiness cakes.Business
}

func NewCakeHandler(business cakes.Business) *CakeHandler {
	return &CakeHandler{
		cakeBusiness: business,
	}
}

func (h *CakeHandler) GetCakeDetail(c echo.Context) error {
	idCake := c.Param("id")
	idCk, _ := strconv.Atoi(idCake)
	result, err := h.cakeBusiness.GetCakeDetail(idCk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesWithData("success to get data", _responseCake.FromCore(result)))
}
