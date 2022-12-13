package pets

type UseCase interface {
	CreatePets(params *CreatePetsParams) (*CreatePetsResponse, error)
	GetPets(params *GetPetsParams) (*GetPetsResponse, error)
	DeletePets(params *DeletePetsParams) (*DeletePetsResponse, error)
}
