package photo

import "time"

type Photo struct {
	Id        string    `json:"-" db:"id"`
	Url       string    `json:"url" db:"url"`
	PetsId    string    `json:"-" db:"pets_id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

type UploadPhotoParam struct {
	File string `json:"file"`
}

type UploadPhotoResponse struct {
	Photo
}
