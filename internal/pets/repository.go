package pets

type PetsRepository interface {
	GetPetById(id string) (*Pets, error)
	GetPets(params *GetPetsParams) (*[]Pets, error)
	DeletePet(id string) error
	CreatePet(pet *Pets) error
}
