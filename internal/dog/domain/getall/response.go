package getall

type Response struct {
	Dogs []DogWithPrice `json:"dogs"`
}

type DogWithPrice struct {
	ID        *int   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Price     any    `json:"price"`
	PriceType any    `json:"priceType"`
}
