package handler

type CheckRequest struct {
	FlightNumber string `json:"flight_number" example:"GA102"`
	Date         string `json:"date" example:"2025-07-12"`
}

type GenerateRequest struct {
	Name         string `json:"name" example:"Sarah"`
	ID           string `json:"id" example:"98123"`
	FlightNumber string `json:"flight_number" example:"ID102"`
	Date         string `json:"date" example:"2025-07-12"`
	Aircraft     string `json:"aircraft" example:"Airbus 320"`
}

type GenerateResponse struct {
	Success bool     `json:"success" example:"true"`
	Seats   []string `json:"seats" example:"[\"3B\", \"7C\", \"14D\"]"`
}

type CheckResponse struct {
	Success bool `json:"success" example:"true"`
	Exists  bool `json:"exists" example:"true"`
}
