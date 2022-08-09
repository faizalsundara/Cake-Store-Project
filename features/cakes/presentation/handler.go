package presentation

import (
	"net/http"
	"strconv"
	cakes "xsis/test/features/cakes"
	_requestCake "xsis/test/features/cakes/presentation/request"
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
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesWithData("success to get data", _responseCake.FromCore(result)))
}

func (h *CakeHandler) GetAllCake(c echo.Context) error {
	result, err := h.cakeBusiness.GetAllCake()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesWithData("success to get data", _responseCake.FromCoreList(result)))
}

func (h *CakeHandler) AddNewCake(c echo.Context) error {
	var dataCake _requestCake.Cake
	errBind := c.Bind(&dataCake)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}

	inputData := _requestCake.ToCore(dataCake)
	tx, err := h.cakeBusiness.PostNewCake(inputData)

	if tx == -1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to input data"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccesNoData("success to insert data"))
}

func (h *CakeHandler) UpdateCake(c echo.Context) error {
	var dataCake _requestCake.Cake
	idCake := c.Param("id")
	idCakeInt, _ := strconv.Atoi(idCake)
	errBind := c.Bind(&dataCake)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}

	inputData := _requestCake.ToCore(dataCake)
	tx, err := h.cakeBusiness.PatchCake(idCakeInt, inputData)
	if tx == -1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to update data"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccesNoData("success to update data"))
}

func (h *CakeHandler) DeleteCake(c echo.Context) error {
	idCake := c.Param("id")
	idCakeInt, _ := strconv.Atoi(idCake)

	row, _ := h.cakeBusiness.DeleteCake(idCakeInt)

	if row == 1 {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to delete data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesNoData("success to delete data"))
}
