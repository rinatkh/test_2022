package pets

import (
	"github.com/rinatkh/test_2022/internal/photo"
	"time"
)

type Pets struct {
	Id        string        `json:"id" db:"id"`
	Name      string        `json:"name" db:"name"`
	Age       int           `json:"age" db:"age"`
	Type      string        `json:"type" db:"type"`
	Photos    []photo.Photo `json:"photos" db:"-"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
}

type CreatePetsParams struct {
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Type   string        `json:"type"`
	Photos []photo.Photo `json:"photos"`
}

type CreatePetsResponse struct {
	Pets
}

type GetPetsParams struct {
	Limit     int64
	Offset    int64
	HasPhotos *bool
}

type GetPetsResponse struct {
	Count int64  `json:"count"`
	Items []Pets `json:"items"`
}

type DeletePetsParams struct {
	IDs []string `json:"ids"`
}

type Error struct {
	Id    string `json:"id"`
	Error string `json:"error"`
}

type DeletePetsResponse struct {
	Deleted int64   `json:"deleted"`
	Errors  []Error `json:"errors"`
}

type UploadPhotoParams struct {
}

type UploadPhotoResponse struct {
	photo.Photo
}
