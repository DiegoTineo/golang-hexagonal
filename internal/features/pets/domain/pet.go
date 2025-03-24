package domain

type Pet struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	PetType  string `json:"pet_type"`
	PetSound string `json:"pet_sound"`
}