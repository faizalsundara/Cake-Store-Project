package business

import (
	"fmt"
	"testing"
	"time"
	cakes "xsis/test/features/cakes"

	"github.com/stretchr/testify/assert"
)

type mockCakeData struct{}

func (mock mockCakeData) DetailOfCake(id int) (cakes.Core, error) {
	return cakes.Core{
		ID: 1, Title: "Kue Lapis", Description: "Kue khas dari Indonesia", Rating: 7.9, Image: "https://kuelapis.png", CreatedAt: time.Time{}, UpdatedAt: time.Time{},
	}, nil
}

func (mock mockCakeData) ListOfCake() ([]cakes.Core, error) {
	return []cakes.Core{
		{
			ID: 1, Title: "Kue Lapis", Description: "Kue khas dari Indonesia", Rating: 7.9, Image: "https://kuelapis.png", CreatedAt: time.Time{}, UpdatedAt: time.Time{},
		},
	}, nil
}

func (mock mockCakeData) AddNewCake(input cakes.Core) (int, error) {
	return 1, nil
}

func (mock mockCakeData) UpdateCake(id int, input cakes.Core) (int, error) {
	return 1, nil
}

func (mock mockCakeData) DeleteCake(id int) (int, error) {
	return 1, nil
}

type mockCakeDataFailed struct{}

func (mock mockCakeDataFailed) DetailOfCake(id int) (cakes.Core, error) {
	return cakes.Core{}, fmt.Errorf("failed to get data")
}

func (mock mockCakeDataFailed) ListOfCake() ([]cakes.Core, error) {
	return nil, fmt.Errorf("failed to get data")
}

func (mock mockCakeDataFailed) AddNewCake(input cakes.Core) (int, error) {
	return 0, fmt.Errorf("failed to insert data")
}

func (mock mockCakeDataFailed) UpdateCake(id int, input cakes.Core) (int, error) {
	return 0, fmt.Errorf("failed to update data")
}

func (mock mockCakeDataFailed) DeleteCake(id int) (int, error) {
	return 0, fmt.Errorf("failed to delete data")
}

func TestGetCakeDetail(t *testing.T) {
	t.Run("Test Get Cake Detail Success", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeData{})
		result, err := cakeBusiness.GetCakeDetail(0)
		assert.Nil(t, err)
		assert.Equal(t, "Kue Lapis", result.Title)
	})
	t.Run("Test Get Cake Detail Failed", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeDataFailed{})
		result, err := cakeBusiness.GetCakeDetail(0)
		assert.NotNil(t, err)
		assert.Equal(t, cakes.Core{}, result)
	})
}

func TestGetAllCake(t *testing.T) {
	t.Run("Test Get All Cake Success", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeData{})
		result, err := cakeBusiness.GetAllCake()
		assert.Nil(t, err)
		assert.Equal(t, []cakes.Core{
			{
				ID:          1,
				Title:       "Kue Lapis",
				Description: "Kue khas dari Indonesia",
				Rating:      7.9,
				Image:       "https://kuelapis.png",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		}, result)
	})
	t.Run("Test Get All Cake Failed", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeDataFailed{})
		result, err := cakeBusiness.GetAllCake()
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestPostNewCake(t *testing.T) {
	t.Run("Test Post New Cake Success", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeData{})
		var input = cakes.Core{
			Title:       "Kue Lapis",
			Description: "Kue khas dari Indonesia",
			Rating:      7.9,
			Image:       "https://kuelapis.png",
		}
		row, err := cakeBusiness.PostNewCake(input)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Post New Cake Failed", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeDataFailed{})
		var input = cakes.Core{
			Title:       "Kue Lapis",
			Description: "Kue khas dari Indonesia",
			Rating:      7.9,
			Image:       "https://kuelapis.png",
		}
		row, err := cakeBusiness.PostNewCake(input)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
	t.Run("Test Post New Cake Failed When Title is Empty", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeDataFailed{})
		var input = cakes.Core{
			Description: "Kue khas dari Indonesia",
			Rating:      7.9,
			Image:       "https://kuelapis.png",
		}
		row, err := cakeBusiness.PostNewCake(input)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
	})
}

func TestPatchCake(t *testing.T) {
	t.Run("Test Patch Cake Success", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeData{})
		var update = cakes.Core{
			Title:       "Kue Lapis",
			Description: "Kue khas dari Indonesia",
			Rating:      7.9,
			Image:       "https://kuelapis.png",
		}
		row, err := cakeBusiness.PatchCake(1, update)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Patch Cake Failed", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeDataFailed{})
		var update = cakes.Core{
			Title:       "Kue Lapis",
			Description: "Kue khas dari Indonesia",
			Rating:      7.9,
			Image:       "https://kuelapis.png",
		}
		row, err := cakeBusiness.PatchCake(1, update)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
	t.Run("Test Patch Cake Failed When Title is Empty", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeDataFailed{})
		var update = cakes.Core{
			Description: "Kue khas dari Indonesia",
			Rating:      7.9,
			Image:       "https://kuelapis.png",
		}
		row, err := cakeBusiness.PatchCake(1, update)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
	})
}

func TestDeleteCake(t *testing.T) {
	t.Run("Test Delete Cake Success", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeData{})
		row, err := cakeBusiness.DeleteCake(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Delete Cake Failed", func(t *testing.T) {
		cakeBusiness := NewCakeBusiness(mockCakeDataFailed{})
		row, err := cakeBusiness.DeleteCake(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}
