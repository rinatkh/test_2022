package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/test_2022/internal/photo"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) photo.PhotoRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetPhotoById(id string) (*photo.Photo, error) {
	var data []photo.Photo
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM photo WHERE id='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) GetPhotoByPetsId(id string) (*[]photo.Photo, error) {
	var data []photo.Photo
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM photo WHERE pets_id='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data, nil
}

func (p postgresRepository) DeletePhoto(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM photo WHERE id='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeletePhotoByPetsId(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM photo WHERE pets_id='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) CreatePhoto(photo *photo.Photo) error {
	res, err := p.db.Query("INSERT INTO photo(id, url, created_at, pets_id) VALUES ($1, $2, $3, $4)", photo.Id, photo.Url, photo.CreatedAt, photo.PetsId)

	if res != nil {
		_ = res.Close()
	}
	return err
}
