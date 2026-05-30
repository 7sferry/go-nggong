package createcontact

type Response struct {
	ID        *int   `json:"id"`
	DogID     int    `json:"dogId"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Emergency string `json:"emergency"`
}
