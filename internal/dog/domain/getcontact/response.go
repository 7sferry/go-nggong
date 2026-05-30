package getcontact

type Response struct {
	DogID   int               `json:"dogId"`
	Contact map[string]string `json:"contact"`
}
