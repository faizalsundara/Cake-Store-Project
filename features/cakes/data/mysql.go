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

func (repo *mysqlCakeRepository) AddNewCake(input cakes.Core) (int, error) {
	var query = ("INSERT INTO cakes(title, description, rating, image) VALUES(?,?,?,?)")
	tx, errTx := repo.db.Prepare(query)
	if errTx != nil {
		return 0, errTx
	}

	result, err := tx.Exec(input.Title, input.Description, input.Rating, input.Image)

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func (repo *mysqlCakeRepository) UpdateCake(idCake int, input cakes.Core) (int, error) {
	var query = ("UPDATE cakes SET title = (?), description = (?), rating = (?), image = (?) WHERE id = ?")
	tx, err := repo.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(input.Title, input.Description, input.Rating, input.Image, idCake)

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func (repo *mysqlCakeRepository) DeleteCake(idCake int) (int, error) {
	var query = ("DELETE FROM cakes WHERE id = ?")
	tx, errTx := repo.db.Prepare(query)
	if errTx != nil {
		return 0, errTx
	}

	result := tx.QueryRow(idCake)

	var id int
	err := result.Scan(&id)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}
