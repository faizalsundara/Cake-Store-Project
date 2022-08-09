package business

import (
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
