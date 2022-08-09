package business

import (
	"fmt"
	cakes "xsis/test/features/cakes"
)

type cakeUseCase struct {
	cakeData cakes.Data
}

func NewCakeBusiness(ckData cakes.Data) cakes.Business {
	return &cakeUseCase{
		cakeData: ckData,
	}
}

func (uc *cakeUseCase) GetCakeDetail(idCake int) (resp cakes.Core, err error) {
	resp, err = uc.cakeData.DetailOfCake(idCake)
	return resp, err
}

func (uc *cakeUseCase) GetAllCake() (data []cakes.Core, err error) {
	data, err = uc.cakeData.ListOfCake()
	return data, err
}

func (uc *cakeUseCase) PostNewCake(input cakes.Core) (row int, err error) {
	if input.Title == "" || input.Description == "" || input.Rating == 0 || input.Image == "" {
		return -1, fmt.Errorf("all input must be filled")
	}
	row, err = uc.cakeData.AddNewCake(input)
	return row, err
}

func (uc *cakeUseCase) PatchCake(idCake int, data cakes.Core) (row int, err error) {
	row, err = uc.cakeData.UpdateCake(idCake, data)
	return row, err
}

func (uc *cakeUseCase) DeleteCake(idCake int) (row int, err error) {
	row, err = uc.cakeData.DeleteCake(idCake)
	return row, err
}
