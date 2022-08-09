package data

import (
	"database/sql"
	cakes "xsis/test/features/cakes"
)

type mysqlCakeRepository struct {
	db *sql.DB
}

func NewCakeRepository(conn *sql.DB) cakes.Data {
	return &mysqlCakeRepository{
		db: conn,
	}
}

func (repo *mysqlCakeRepository) DetailOfCake(idCake int) (cakes.Core, error) {
	result := repo.db.QueryRow("SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?", idCake)

	var Cake Cake
	err := result.Scan(&Cake.ID, &Cake.Title, &Cake.Description, &Cake.Rating, &Cake.Image, &Cake.CreatedAt, &Cake.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return Cake.toCore(), nil
}

func (repo *mysqlCakeRepository) ListOfCake() ([]cakes.Core, error) {
	results, err := repo.db.Query("SELECT id, title, description, rating, image, created_at, updated_at FROM cakes")
	if err != nil {
		return nil, err
	}

	var listCake []Cake
	for results.Next() {
		var cake Cake
		err := results.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			panic(err)
		}
		listCake = append(listCake, cake)
	}
	return toCoreList(listCake), nil
}
