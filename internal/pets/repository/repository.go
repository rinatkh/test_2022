package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/test_2022/internal/pets"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) pets.PetsRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetPetById(id string) (*pets.Pets, error) {
	var data []pets.Pets
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM pets WHERE id='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) CreatePet(pet *pets.Pets) error {
	res, err := p.db.Query("INSERT INTO pets(id, name, age, type, created_at) VALUES ($1, $2, $3, $4, $5)", pet.Id, pet.Name, pet.Age, pet.Type, pet.CreatedAt)

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeletePet(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM pets WHERE id='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) GetPets(params *pets.GetPetsParams) (*[]pets.Pets, error) {
	var result []pets.Pets
	queryStr := "SELECT * FROM pets WHERE id <> ''"

	if params.HasPhotos != nil {
		if *params.HasPhotos {
			queryStr += " AND id in (Select pets_id from photo)"
		} else {
			queryStr += " AND id not in (Select pets_id from photo)"
		}
	}

	queryStr += " ORDER BY created_at DESC"
	if params.Limit == 0 {
		queryStr += " LIMIT 1"
	} else {
		queryStr += fmt.Sprintf(" LIMIT %d", params.Limit)
	}
	queryStr += fmt.Sprintf(" OFFSET %d", params.Offset)

	err := p.db.Select(
		&result, queryStr)

	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}

	return &result, nil
}
