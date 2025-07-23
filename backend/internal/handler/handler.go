package handler

import (
	"net/http"

	"github.com/ibadi-id/backend-bc/internal/domain"
	"github.com/ibadi-id/backend-bc/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase *usecase.VoucherUsecase
}

func NewHandler(u *usecase.VoucherUsecase) *Handler {
	return &Handler{u}
}

// Check godoc
// @Summary Check if voucher already exists
// @Tags Voucher
// @Accept  json
// @Produce  json
// @Param request body handler.CheckRequest true "Flight info"
// @Success 200 {object} handler.CheckResponse
// @Failure 400 {object} map[string]string
// @Router /api/check [post]
func (h *Handler) Check(c echo.Context) error {
	var req CheckRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
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
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/generate [post]
func (h *Handler) Generate(c echo.Context) error {
	var req GenerateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate seats"})
	}
	if seats == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "vouchers already generated"})
	}

	return c.JSON(http.StatusOK, GenerateResponse{
		Success: true,
		Seats:   seats,
	})
}
