package createcontact

type Request struct {
	DogID     int    `json:"dogId"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Emergency string `json:"emergency"`
}
