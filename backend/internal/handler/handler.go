package handler

import (
	"errors"
	"net/http"

	"github.com/ibadi-id/airline-seat-assignment/backend/internal/domain"
	"github.com/ibadi-id/airline-seat-assignment/backend/internal/usecase"
	"github.com/ibadi-id/airline-seat-assignment/backend/pkg/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase usecase.VoucherUsecase
}

func NewHandler(u usecase.VoucherUsecase) *Handler {
	return &Handler{u}
}

// Check godoc
// @Summary Check if voucher already exists
// @Tags Voucher
// @Accept  json
// @Produce  json
// @Param request body handler.CheckRequest true "Flight info"
// @Success 200 {object} handler.CheckResponse
// @Failure 400 {object} handler.ValidationErrorResponse
// @Router /api/check [post]
func (h *Handler) Check(c echo.Context) error {
	var req CheckRequest
	errorMap := make(map[string]string)

	if err := c.Bind(&req); err != nil {
		errorMap["input"] = "invalid input"
		return c.JSON(http.StatusBadRequest, validator.ValidationError{Errors: errorMap})
	}

	if err := c.Validate(&req); err != nil {
		var ve *validator.ValidationError
		if errors.As(err, &ve) {
			return c.JSON(http.StatusBadRequest, ve)
		}
		errorMap["input"] = "invalid input"
		return c.JSON(http.StatusBadRequest, validator.ValidationError{Errors: errorMap})
	}

	exists, err := h.Usecase.Check(req.FlightNumber, req.Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to check assignment"})
	}

	return c.JSON(http.StatusOK, CheckResponse{Exists: exists, Success: true})
}

// Generate godoc
// @Summary Generate 3 random voucher seats
// @Tags Voucher
// @Accept  json
// @Produce  json
// @Param request body handler.GenerateRequest true "Voucher input"
// @Success 200 {object} handler.GenerateResponse
// @Failure 400 {object} handler.ValidationErrorResponse
// @Failure 409 {object} handler.StatusConflictResponse
// @Router /api/generate [post]
func (h *Handler) Generate(c echo.Context) error {
	var req GenerateRequest
	errorMap := make(map[string]string)

	if err := c.Bind(&req); err != nil {
		errorMap["input"] = "invalid input"
		return c.JSON(http.StatusBadRequest, validator.ValidationError{Errors: errorMap})
	}

	if err := c.Validate(&req); err != nil {
		var ve *validator.ValidationError
		if errors.As(err, &ve) {
			return c.JSON(http.StatusBadRequest, ve)
		}
		errorMap["input"] = "invalid input"
		return c.JSON(http.StatusBadRequest, validator.ValidationError{Errors: errorMap})
	}

	v := domain.Voucher{
		CrewName:     req.Name,
		CrewID:       req.ID,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.Date,
		AircraftType: req.Aircraft,
	}

	seats, err := h.Usecase.Generate(v)
	if err != nil {
		if err.Error() == "aircraft type not valid" {
			errorMap["aircraft"] = err.Error()
			return c.JSON(http.StatusBadRequest, validator.ValidationError{Errors: errorMap})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate seats"})
	}
	if seats == nil {
		return c.JSON(http.StatusConflict, StatusConflictResponse{Error: "vouchers already generated"})
	}

	return c.JSON(http.StatusOK, GenerateResponse{
		Success: true,
		Seats:   seats,
	})
}
