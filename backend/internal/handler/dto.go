package handler

type CheckRequest struct {
	FlightNumber string `json:"flight_number" validate:"required" example:"GA102"`
	Date         string `json:"date" validate:"required" example:"2025-07-12"`
}

type GenerateRequest struct {
	Name         string `json:"name" validate:"required" example:"Sarah"`
	ID           string `json:"id" validate:"required" example:"98123"`
	FlightNumber string `json:"flight_number" validate:"required" example:"ID102"`
	Date         string `json:"date" validate:"required" example:"2025-07-12"`
	Aircraft     string `json:"aircraft" validate:"required" example:"Airbus 320"`
}

type GenerateResponse struct {
	Success bool     `json:"success" example:"true"`
	Seats   []string `json:"seats" example:"3B,7C,14D"`
}

type CheckResponse struct {
	Success bool `json:"success" example:"true"`
	Exists  bool `json:"exists" example:"true"`
}

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors" example:"date:date is required,flight_number:flight_number is required,aircraft:aircraft type not valid"`
}

type StatusConflictResponse struct {
	Error string `json:"error" example:"vouchers already generated"`
}
